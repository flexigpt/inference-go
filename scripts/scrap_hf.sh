
curl https://huggingface.co/api/models\?pipeline_tag\=text-generation\&num_parameters\=min:32B\&inference_provider\=fireworks-ai,cerebras,novita,groq,together,nscale,fal-ai,featherless-ai,replicate,zai-org,cohere,scaleway,hf-inference,publicai,ovhcloud,deepinfra,wavespeed\&base_model_relation\=base\&full\=true &> hf_models.json

cutoff=$(date -u -d '6 months ago' +%s)

jq --argjson cutoff "$cutoff" '
  map(
    select(
      ((try (.lastModified | sub("\\.[0-9]+Z$"; "Z") | fromdateiso8601) catch 0) >= $cutoff)
      and
      ((.downloads // 0) > 10000)
    )
    | del(.siblings, .sha, .tags)
  )
' hf_models.json > hf_models_last_6_months_10k_downloads_clean.json

HF_TOKEN=hf_your_token python3 - hf_models_last_6_months_10k_downloads_clean.json > hf_models-with-providers-and-params.json <<'PY'
import os
import sys
import json
import time
import urllib.request
import urllib.parse
import urllib.error

HF_TOKEN = os.environ.get("HF_TOKEN")
HF_SLEEP = float(os.environ.get("HF_SLEEP", "0"))
ROUTER_BASE_URL = os.environ.get("HF_ROUTER_BASE", "https://router.huggingface.co/v1")

MAX_REASONABLE_CONTEXT = 10_000_000

CONFIG_KEYS = [
    "model_type",
    "architectures",
    "max_position_embeddings",
    "max_sequence_length",
    "max_seq_len",
    "seq_length",
    "n_positions",
    "n_ctx",
    "sliding_window",
    "rope_scaling",
    "torch_dtype",
    "vocab_size",
    "hidden_size",
    "num_hidden_layers",
    "num_attention_heads",
    "num_key_value_heads",
]

GENERATION_KEYS = [
    "max_new_tokens",
    "max_length",
    "min_new_tokens",
    "temperature",
    "top_p",
    "top_k",
    "do_sample",
    "repetition_penalty",
    "bos_token_id",
    "eos_token_id",
    "pad_token_id",
]

def get_models_container(raw):
    if isinstance(raw, dict) and isinstance(raw.get("data"), list):
        return raw["data"], "data"
    if isinstance(raw, list):
        return raw, None
    raise SystemExit("Input JSON must be an array or an object with a data array")

def http_get_json(url):
    headers = {
        "Accept": "application/json",
        "User-Agent": "hf-json-enricher/1.0",
    }

    if HF_TOKEN:
        headers["Authorization"] = "Bearer " + HF_TOKEN

    req = urllib.request.Request(url, headers=headers)

    try:
        with urllib.request.urlopen(req, timeout=30) as res:
            status = getattr(res, "status", 200)
            body = res.read().decode("utf-8")
            return json.loads(body), None, status
    except urllib.error.HTTPError as e:
        body = e.read().decode("utf-8", errors="replace")
        return None, {
            "statusCode": e.code,
            "error": body[:5000],
        }, e.code
    except Exception as e:
        return None, {
            "error": str(e),
        }, None

def fetch_hf_model_api(model_id):
    expands = [
        "inferenceProviderMapping",
        "pipeline_tag",
        "gated",
        "private",
    ]

    query = urllib.parse.urlencode(
        [("expand[]", x) for x in expands]
    )

    url = (
        "https://huggingface.co/api/models/"
        + urllib.parse.quote(model_id, safe="/")
        + "?"
        + query
    )

    return http_get_json(url)

def fetch_repo_json_file(model_id, filename):
    url = (
        "https://huggingface.co/"
        + urllib.parse.quote(model_id, safe="/")
        + "/resolve/main/"
        + filename
    )

    return http_get_json(url)

def int_or_none(value):
    if isinstance(value, bool):
        return None

    if isinstance(value, int):
        n = value
    elif isinstance(value, str) and value.isdigit():
        n = int(value)
    else:
        return None

    if n <= 0:
        return None

    # Ignore huge tokenizer sentinel values.
    if n > MAX_REASONABLE_CONTEXT:
        return None

    return n

def pick_fields(obj, keys):
    if not isinstance(obj, dict):
        return {}

    return {
        key: obj[key]
        for key in keys
        if key in obj
    }

def add_context_candidate(candidates, source, key, value):
    n = int_or_none(value)
    if n is not None:
        candidates.append({
            "source": source,
            "key": key,
            "value": n,
        })

def infer_reasoning(model_id, tokenizer_config):
    text = model_id.lower()
    signals = []

    name_keywords = [
        "thinking",
        "reasoning",
        "reasoner",
        "deepseek-r1",
        "qwq",
        "glm-z1",
        "o1",
    ]

    for kw in name_keywords:
        if kw in text:
            signals.append("model_id_contains:" + kw)

    chat_template = ""
    if isinstance(tokenizer_config, dict):
        chat_template = str(tokenizer_config.get("chat_template") or "").lower()

    template_keywords = [
        "<think",
        "reasoning",
        "thinking",
        "enable_thinking",
        "thinking_budget",
        "reasoning_content",
    ]

    for kw in template_keywords:
        if kw in chat_template:
            signals.append("chat_template_contains:" + kw)

    return {
        "known": False,
        "supported": None,
        "likely": bool(signals),
        "signals": signals,
        "note": "HF metadata usually does not expose a universal reasoning-support flag. This is a heuristic.",
    }

def infer_tools(tokenizer_config):
    chat_template = ""
    if isinstance(tokenizer_config, dict):
        chat_template = str(tokenizer_config.get("chat_template") or "").lower()

    signals = []

    template_keywords = [
        "tool_calls",
        "tool_call",
        "tools",
        "function_call",
        "functions",
    ]

    for kw in template_keywords:
        if kw in chat_template:
            signals.append("chat_template_contains:" + kw)

    return {
        "known": False,
        "supported": None,
        "likely": bool(signals),
        "signals": signals,
        "note": "HF metadata does not consistently expose tool/function-calling support. This is a heuristic.",
    }

def enrich_providers(model, model_id, hf_model, err, status):
    if err:
        model["hfProviderLookup"] = {
            "ok": False,
            "statusCode": status,
            **err,
        }

        model["routerApi"] = {
            "baseUrl": ROUTER_BASE_URL,
            "chatCompletionsUrl": ROUTER_BASE_URL + "/chat/completions",
            "modelNames": [],
            "liveModelNames": [],
            "preferredLiveModelName": None,
        }

        return

    mapping = hf_model.get("inferenceProviderMapping") or {}

    providers = []
    for provider_name, provider_info in mapping.items():
        provider_obj = {
            "provider": provider_name,
            "routerModel": f"{model_id}:{provider_name}",
            "routerBaseUrl": ROUTER_BASE_URL,
            "chatCompletionsUrl": ROUTER_BASE_URL + "/chat/completions",
        }

        if isinstance(provider_info, dict):
            provider_obj.update(provider_info)
        else:
            provider_obj["raw"] = provider_info

        providers.append(provider_obj)

    live_providers = [
        p for p in providers
        if str(p.get("status", "")).lower() == "live"
    ]

    live_model_names = [
        p["routerModel"]
        for p in live_providers
    ]

    model["hfProviderLookup"] = {
        "ok": True,
        "statusCode": status,
        "pipeline_tag": hf_model.get("pipeline_tag"),
        "gated": hf_model.get("gated"),
        "private": hf_model.get("private"),
        "inferenceProviderMapping": mapping,
        "inferenceProviders": providers,
        "liveInferenceProviders": live_providers,
    }

    model["routerApi"] = {
        "baseUrl": ROUTER_BASE_URL,
        "chatCompletionsUrl": ROUTER_BASE_URL + "/chat/completions",
        "modelNames": [
            p["routerModel"]
            for p in providers
        ],
        "liveModelNames": live_model_names,
        "preferredLiveModelName": live_model_names[0] if live_model_names else None,
    }

def enrich_params(model, model_id):
    config, config_err, _ = fetch_repo_json_file(model_id, "config.json")
    tokenizer_config, tok_err, _ = fetch_repo_json_file(model_id, "tokenizer_config.json")
    generation_config, gen_err, _ = fetch_repo_json_file(model_id, "generation_config.json")

    context_candidates = []

    if isinstance(tokenizer_config, dict):
        add_context_candidate(
            context_candidates,
            "tokenizer_config.json",
            "model_max_length",
            tokenizer_config.get("model_max_length"),
        )

    if isinstance(config, dict):
        for key in [
            "max_position_embeddings",
            "max_sequence_length",
            "max_seq_len",
            "seq_length",
            "n_positions",
            "n_ctx",
        ]:
            add_context_candidate(
                context_candidates,
                "config.json",
                key,
                config.get(key),
            )

    context_window = context_candidates[0]["value"] if context_candidates else None

    chat_template = None
    if isinstance(tokenizer_config, dict):
        chat_template = tokenizer_config.get("chat_template")

    generation_selected = pick_fields(generation_config, GENERATION_KEYS)

    model["modelParamsBestEffort"] = {
        "ok": True,
        "files": {
            "config.json": {
                "ok": isinstance(config, dict),
                "error": config_err,
            },
            "tokenizer_config.json": {
                "ok": isinstance(tokenizer_config, dict),
                "error": tok_err,
            },
            "generation_config.json": {
                "ok": isinstance(generation_config, dict),
                "error": gen_err,
            },
        },
        "contextWindowTokens": context_window,
        "contextWindowCandidates": context_candidates,
        "maxOutputTokens": {
            "hardLimit": None,
            "defaultMaxNewTokens": generation_selected.get("max_new_tokens"),
            "defaultMaxLength": generation_selected.get("max_length"),
            "note": "Hard output-token limits are usually provider-specific and are not reliably exposed in HF repo metadata.",
        },
        "promptTokens": {
            "availableFromMetadata": False,
            "note": "Prompt token count requires the actual prompt and the model tokenizer.",
        },
        "reasoning": infer_reasoning(model_id, tokenizer_config),
        "tools": infer_tools(tokenizer_config),
        "chatTemplate": {
            "available": bool(chat_template),
            "supportsSystemMessageLikely": (
                "system" in str(chat_template).lower()
                if chat_template
                else None
            ),
        },
        "selectedConfig": pick_fields(config, CONFIG_KEYS),
        "selectedTokenizerConfig": {
            "model_max_length": (
                tokenizer_config.get("model_max_length")
                if isinstance(tokenizer_config, dict)
                else None
            ),
            "bos_token": (
                tokenizer_config.get("bos_token")
                if isinstance(tokenizer_config, dict)
                else None
            ),
            "eos_token": (
                tokenizer_config.get("eos_token")
                if isinstance(tokenizer_config, dict)
                else None
            ),
            "pad_token": (
                tokenizer_config.get("pad_token")
                if isinstance(tokenizer_config, dict)
                else None
            ),
        },
        "selectedGenerationConfig": generation_selected,
    }

if len(sys.argv) < 2:
    raise SystemExit("Usage: python3 - models.json > models-enriched.json")

with open(sys.argv[1], "r", encoding="utf-8") as f:
    raw = json.load(f)

models, data_key = get_models_container(raw)

for i, model in enumerate(models, start=1):
    if not isinstance(model, dict):
        continue

    model_id = model.get("modelId") or model.get("id")

    if not model_id:
        model["hfProviderLookup"] = {
            "ok": False,
            "error": "Missing modelId or id",
        }
        model["modelParamsBestEffort"] = {
            "ok": False,
            "error": "Missing modelId or id",
        }
        continue

    print(f"[{i}/{len(models)}] {model_id}", file=sys.stderr)

    hf_model, err, status = fetch_hf_model_api(model_id)
    enrich_providers(model, model_id, hf_model or {}, err, status)

    enrich_params(model, model_id)

    if HF_SLEEP:
        time.sleep(HF_SLEEP)

if data_key:
    raw[data_key] = models
    print(json.dumps(raw, indent=2, ensure_ascii=False))
else:
    print(json.dumps(models, indent=2, ensure_ascii=False))
PY
