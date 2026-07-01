curl https://openrouter.ai/api/v1/models\?sort=most-popular \
     -H "Authorization: Bearer <token>" &> models.json

python3 - models.json > top-20-filtered.json <<'PY'
import sys
import json

EXCLUDED_PROVIDERS = {"openai", "anthropic", "google"}

with open(sys.argv[1], "r", encoding="utf-8") as f:
    raw = json.load(f)

filtered = []

for model in raw.get("data", []):
    provider = model.get("id", "").split("/", 1)[0].lower()

    if provider in EXCLUDED_PROVIDERS:
        continue

    filtered.append(model)

    if len(filtered) == 20:
        break

print(json.dumps({"data": filtered}, indent=2, ensure_ascii=False))
PY
