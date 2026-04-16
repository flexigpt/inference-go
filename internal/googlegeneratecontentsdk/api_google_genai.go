package googlegeneratecontentsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"strings"
	"sync"
	"time"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

// GoogleGenerateContentAPI implements CompletionProvider for Google's Generative AI API.
type GoogleGenerateContentAPI struct {
	ProviderParam *spec.ProviderParam
	debugger      spec.CompletionDebugger
	client        *genai.Client
	mu            sync.RWMutex
}

// NewGoogleGenerateContentAPI creates a new instance of the Google GenAI provider.
func NewGoogleGenerateContentAPI(
	pi spec.ProviderParam,
	debugger spec.CompletionDebugger,
) (*GoogleGenerateContentAPI, error) {
	if pi.Name == "" {
		return nil, errors.New("google genai api LLM: invalid args")
	}
	return &GoogleGenerateContentAPI{
		ProviderParam: &pi,
		debugger:      debugger,
	}, nil
}

func (api *GoogleGenerateContentAPI) InitLLM(ctx context.Context) error {
	api.mu.Lock()
	defer api.mu.Unlock()

	if api.ProviderParam == nil {
		api.client = nil
		return errors.New("google genai api LLM: no ProviderParam found")
	}

	if strings.TrimSpace(api.ProviderParam.APIKey) == "" {
		logutil.Debug(
			string(api.ProviderParam.Name) + ": No API key given. Not initializing Google GenAI client",
		)
		api.client = nil
		return nil
	}

	pi := *api.ProviderParam // snapshot under lock

	cc := &genai.ClientConfig{
		APIKey:  pi.APIKey,
		Backend: genai.BackendGeminiAPI,
	}

	httpOpts := genai.HTTPOptions{}

	// Custom base URL / path prefix (optional).
	baseURL := spec.DefaultGoogleGenerateContentOrigin
	if pi.Origin != "" || strings.TrimSpace(pi.ChatCompletionPathPrefix) != "" {
		baseURL = strings.TrimSuffix(pi.Origin, "/")
		if prefix := strings.Trim(strings.TrimSpace(pi.ChatCompletionPathPrefix), "/"); prefix != "" {
			baseURL = strings.TrimRight(baseURL+"/"+prefix, "/")
		}
		if baseURL != "" {
			httpOpts.BaseURL = baseURL + "/"
		}
	}

	// Custom default headers (optional).
	if len(pi.DefaultHeaders) > 0 || strings.TrimSpace(pi.APIKeyHeaderKey) != "" {
		httpOpts.Headers = make(http.Header)

		for k, v := range pi.DefaultHeaders {
			httpOpts.Headers.Set(strings.TrimSpace(k), strings.TrimSpace(v))
		}
		if hdr := strings.TrimSpace(pi.APIKeyHeaderKey); hdr != "" {
			httpOpts.Headers.Set(hdr, pi.APIKey)
		}
	}

	cc.HTTPOptions = httpOpts

	// Debugger HTTP client (optional).
	if api.debugger != nil {
		if httpClient := api.debugger.HTTPClient(nil); httpClient != nil {
			cc.HTTPClient = httpClient
		}
	}

	client, err := genai.NewClient(ctx, cc)
	if err != nil {
		return fmt.Errorf("google genai api LLM: failed to create client: %w", err)
	}
	api.client = client

	logutil.Info(
		"google genai api LLM provider initialized",
		"name", string(pi.Name),
		"URL", baseURL,
	)
	return nil
}

func (api *GoogleGenerateContentAPI) DeInitLLM(ctx context.Context) error {
	api.mu.Lock()
	var name spec.ProviderName
	if api.ProviderParam != nil {
		name = api.ProviderParam.Name
	}
	api.client = nil
	api.mu.Unlock()
	logutil.Info(
		"google genai api LLM: provider de initialized",
		"name", string(name),
	)
	return nil
}

func (api *GoogleGenerateContentAPI) GetProviderInfo(ctx context.Context) *spec.ProviderParam {
	api.mu.RLock()
	defer api.mu.RUnlock()
	if api.ProviderParam == nil {
		return nil
	}
	cp := *api.ProviderParam
	cp.DefaultHeaders = sdkutil.CloneStringMap(cp.DefaultHeaders)
	return &cp
}

func (api *GoogleGenerateContentAPI) IsConfigured(ctx context.Context) bool {
	api.mu.RLock()
	defer api.mu.RUnlock()
	return api.ProviderParam != nil && strings.TrimSpace(api.ProviderParam.APIKey) != ""
}

func (api *GoogleGenerateContentAPI) SetProviderAPIKey(ctx context.Context, apiKey string) error {
	api.mu.Lock()
	defer api.mu.Unlock()

	if api.ProviderParam == nil {
		return errors.New("google genai api LLM: no ProviderParam found")
	}
	// Allow empty to clear.
	api.ProviderParam.APIKey = strings.TrimSpace(apiKey)
	return nil
}

func (api *GoogleGenerateContentAPI) GetProviderCapability(ctx context.Context) (spec.ModelCapabilities, error) {
	return googleGenerateContentSDKCapability, nil
}

func (api *GoogleGenerateContentAPI) FetchCompletion(
	ctx context.Context,
	inReq *spec.FetchCompletionRequest,
	opts *spec.FetchCompletionOptions,
) (*spec.FetchCompletionResponse, error) {
	api.mu.RLock()
	client := api.client
	var pi spec.ProviderParam
	if api.ProviderParam != nil {
		pi = *api.ProviderParam
	}
	api.mu.RUnlock()

	if client == nil {
		return nil, errors.New("google genai api LLM: client not initialized")
	}
	if inReq == nil || len(inReq.Inputs) == 0 || inReq.ModelParam.Name == "" {
		return nil, errors.New("google genai api LLM: empty completion data")
	}

	req, _, warns, err := sdkutil.NormalizeRequestForSDK(
		ctx, inReq, opts, spec.ProviderSDKTypeGoogleGenerateContent, googleGenerateContentSDKCapability,
	)
	if err != nil {
		return nil, err
	}

	// Sanitize reasoning inputs: keep only Google-native signed thoughts.
	req.Inputs = sanitizeGoogleGenerateContentReasoningInputs(req.Inputs)

	// Build genai contents + system instruction.
	contents, sysInstruction, err := toGoogleGenerateContentContents(ctx, req.ModelParam.SystemPrompt, req.Inputs)
	if err != nil {
		return nil, err
	}

	// Resolve timeout.
	timeout := spec.DefaultAPITimeout
	if req.ModelParam.Timeout > 0 {
		timeout = time.Duration(req.ModelParam.Timeout) * time.Second
	}

	// Build GenerateContentConfig.
	config := &genai.GenerateContentConfig{
		HTTPOptions: &genai.HTTPOptions{Timeout: &timeout},
	}

	if sysInstruction != nil {
		config.SystemInstruction = sysInstruction
	}
	if req.ModelParam.MaxOutputLength > 0 {
		config.MaxOutputTokens = sdkutil.ClampIntToInt32(req.ModelParam.MaxOutputLength)
	}
	if t := req.ModelParam.Temperature; t != nil {
		v := float32(*t)
		config.Temperature = &v
	}
	if len(req.ModelParam.StopSequences) > 0 {
		config.StopSequences = req.ModelParam.StopSequences
	}

	// Thinking / reasoning.
	reasoningCaps := googleGenerateContentSDKCapability.ReasoningCapabilities
	if resolved := resolveGoogleGenerateContentReasoningCapabilities(ctx, opts, req.ModelParam.Name); resolved != nil {
		reasoningCaps = resolved
	}
	if err := applyGoogleGenerateContentThinkingPolicy(config, &req.ModelParam, reasoningCaps); err != nil {
		return nil, err
	}

	// Output format.
	if req.ModelParam.OutputParam != nil {
		if err := applyGoogleGenerateContentOutputParam(config, req.ModelParam.OutputParam); err != nil {
			return nil, err
		}
	}

	// Tools + tool policy.
	var toolChoiceNameMap map[string]spec.ToolChoice
	var webSearchChoiceID string

	if len(req.ToolChoices) > 0 {
		toolChoicesForProvider := req.ToolChoices
		if req.ToolPolicy != nil {
			switch req.ToolPolicy.Mode {
			case spec.ToolPolicyModeNone:
				// Gemini web search is not governed by FunctionCallingConfig, so
				// the only reliable way to enforce "none" is to omit all tools.
				toolChoicesForProvider = nil
			case spec.ToolPolicyModeAny, spec.ToolPolicyModeTool:
				// Gemini can force callable function/custom tools, but not
				// GoogleSearch grounding. Drop web search tools here so the
				// strict policy semantics stay correct.
				toolChoicesForProvider = filterGoogleGenerateContentCallableToolChoices(req.ToolChoices)
				if len(toolChoicesForProvider) == 0 {
					return nil, fmt.Errorf(
						"googleGenerateContent: toolPolicy.mode=%q cannot be satisfied because Gemini GenerateContent cannot force webSearch; provide at least one function/custom tool or use toolPolicy=auto",
						req.ToolPolicy.Mode,
					)
				}
			default:
				// Fall.
			}
		}

		tools, nameMap, wsChoiceID, buildErr := buildGoogleGenerateContentTools(toolChoicesForProvider)
		if buildErr != nil {
			return nil, buildErr
		}
		if len(tools) > 0 {
			config.Tools = tools
			toolChoiceNameMap = nameMap
			webSearchChoiceID = wsChoiceID
			if req.ToolPolicy != nil {
				tc, policyErr := buildGoogleGenerateContentToolConfig(req.ToolPolicy, toolChoiceNameMap)
				if policyErr != nil {
					return nil, policyErr
				}
				config.ToolConfig = tc
			}

		}
	}

	// Debug span.
	var span spec.CompletionSpan
	if api.debugger != nil {
		ctx, span = api.debugger.StartSpan(ctx, &spec.CompletionSpanStart{
			Provider: pi.Name,
			Model:    req.ModelParam.Name,
			Request:  req,
			Options:  opts,
		})
	}

	var (
		normalizedResp *spec.FetchCompletionResponse
		rawResp        *genai.GenerateContentResponse
		apiErr         error
	)

	useStream := req.ModelParam.Stream && opts != nil && opts.StreamHandler != nil
	if useStream {
		normalizedResp, rawResp, apiErr = api.doStreaming(
			ctx, client, pi.Name, req.ModelParam.Name,
			contents, config, opts, toolChoiceNameMap, webSearchChoiceID,
		)
	} else {
		normalizedResp, rawResp, apiErr = api.doNonStreaming(
			ctx, client, req.ModelParam.Name,
			contents, config, toolChoiceNameMap, webSearchChoiceID,
		)
	}

	if normalizedResp != nil && len(warns) > 0 {
		normalizedResp.Warnings = append(normalizedResp.Warnings, warns...)
	}

	if span != nil {
		end := spec.CompletionSpanEnd{
			ProviderResponse: rawResp,
			Response:         normalizedResp,
			Err:              apiErr,
		}
		if normalizedResp != nil {
			if dd := span.End(&end); dd != nil && normalizedResp.DebugDetails == nil {
				normalizedResp.DebugDetails = dd
			}
		} else {
			_ = span.End(&end)
		}
	}

	return normalizedResp, apiErr
}

func (api *GoogleGenerateContentAPI) doNonStreaming(
	ctx context.Context,
	client *genai.Client,
	modelName spec.ModelName,
	contents []*genai.Content,
	config *genai.GenerateContentConfig,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, *genai.GenerateContentResponse, error) {
	resp := &spec.FetchCompletionResponse{}

	genResp, err := client.Models.GenerateContent(ctx, string(modelName), contents, config)

	resp.Usage = usageFromGenAIResponse(genResp)
	if err != nil {
		resp.Error = &spec.Error{Message: err.Error()}
		return resp, genResp, err
	}
	resp.Outputs = outputsFromGenAIResponse(genResp, toolChoiceNameMap, webSearchChoiceID)
	return resp, genResp, nil
}

func (api *GoogleGenerateContentAPI) doStreaming(
	ctx context.Context,
	client *genai.Client,
	providerName spec.ProviderName,
	modelName spec.ModelName,
	contents []*genai.Content,
	config *genai.GenerateContentConfig,
	opts *spec.FetchCompletionOptions,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, *genai.GenerateContentResponse, error) {
	resp := &spec.FetchCompletionResponse{}
	streamCfg := sdkutil.ResolveStreamConfig(opts)
	// Work around a Google GenAI SDK streaming timeout lifecycle bug:
	// when HTTPOptions.Timeout is set, the SDK may create a derived request
	// context inside sendStreamRequest() and defer cancel() there, even though
	// the response body is consumed later by the returned iterator.
	//
	// To avoid premature cancellation, apply the timeout on our outer context
	// for streaming and clear config.HTTPOptions.Timeout only in the cases
	// where the SDK would otherwise derive its own timeout context.
	streamCtx, streamCancel, streamConfig := prepareGoogleGenerateContentStreamCall(ctx, config)
	if streamCancel != nil {
		defer streamCancel()
	}

	emitText := func(chunk string) error {
		if chunk == "" {
			return nil
		}
		return sdkutil.SafeCallStreamHandler(opts.StreamHandler, spec.StreamEvent{
			Kind:          spec.StreamContentKindText,
			Provider:      providerName,
			Model:         modelName,
			CompletionKey: opts.CompletionKey,
			Text:          &spec.StreamTextChunk{Text: chunk},
		})
	}

	emitThinking := func(chunk string) error {
		if chunk == "" {
			return nil
		}
		return sdkutil.SafeCallStreamHandler(opts.StreamHandler, spec.StreamEvent{
			Kind:          spec.StreamContentKindThinking,
			Provider:      providerName,
			Model:         modelName,
			CompletionKey: opts.CompletionKey,
			Thinking:      &spec.StreamThinkingChunk{Text: chunk},
		})
	}

	// Flush calls are for end cleaning not for continuation.
	writeText, flushText := sdkutil.NewBufferedStreamer(emitText, streamCfg.FlushInterval, streamCfg.FlushChunkSize)
	writeThinking, flushThinking := sdkutil.NewBufferedStreamer(
		emitThinking,
		streamCfg.FlushInterval,
		streamCfg.FlushChunkSize,
	)

	// Accumulated state across all stream chunks.
	var (
		accUsage       *genai.GenerateContentResponseUsageMetadata
		accFinish      genai.FinishReason
		accResponseID  string
		accParts       []*genai.Part
		accGrounding   *genai.GroundingMetadata
		streamWriteErr error
		streamErr      error
	)
	stream := client.Models.GenerateContentStream(streamCtx, string(modelName), contents, streamConfig)
	for chunkResp, chunkErr := range stream {
		if chunkErr != nil {
			streamErr = chunkErr
			break
		}
		if chunkResp == nil {
			continue
		}

		if chunkResp.UsageMetadata != nil {
			accUsage = chunkResp.UsageMetadata
		}
		if chunkResp.ResponseID != "" {
			accResponseID = chunkResp.ResponseID
		}

		if len(chunkResp.Candidates) == 0 || chunkResp.Candidates[0] == nil {
			continue
		}
		cand := chunkResp.Candidates[0]
		if cand.FinishReason != "" {
			accFinish = cand.FinishReason
		}
		if cand.GroundingMetadata != nil {
			accGrounding = mergeGoogleGenerateContentGroundingMetadata(accGrounding, cand.GroundingMetadata)
		}
		if cand.Content == nil {
			continue
		}

		for _, part := range cand.Content.Parts {
			if part == nil {
				continue
			}
			accParts = append(accParts, part)

			switch {
			case part.Thought:
				if part.Text == "" {
					continue
				}
				streamWriteErr = writeThinking(part.Text)

			case part.FunctionCall != nil:
				// We dont stream function calls.
			case part.Text != "":
				if streamWriteErr != nil {
					break
				}
				streamWriteErr = writeText(part.Text)
			}
			if streamWriteErr != nil {
				break
			}
		}
		if streamWriteErr != nil {
			break
		}
	}

	if flushText != nil {
		flushText()
	}
	if flushThinking != nil {
		flushThinking()
	}

	// Consolidate raw per-chunk stream parts into the canonical single-part-per-kind
	// form that GenerateContent (non-streaming) returns: adjacent thought fragments are
	// merged into one thought part and adjacent text fragments into one text part, with
	// the last non-empty ThoughtSignature winning (signatures arrive on the final chunk).
	// Without this, outputsFromGenAIResponse produces many tiny, mostly-unsigned
	// ReasoningMessage / OutputMessage entries instead of the single signed entries that
	// callers and the multi-turn round-trip path expect.
	accParts = consolidateGoogleGenerateContentStreamParts(accParts)

	combinedErr := errors.Join(streamErr, streamWriteErr)

	// Build a synthetic GenerateContentResponse from all accumulated state so
	// the shared outputsFromGenAIResponse path can process it uniformly.
	synthResp := &genai.GenerateContentResponse{
		ResponseID:    accResponseID,
		UsageMetadata: accUsage,
	}
	if len(accParts) > 0 || accGrounding != nil || accFinish != "" {
		synthResp.Candidates = []*genai.Candidate{{
			Content: &genai.Content{
				Role:  genai.RoleModel,
				Parts: accParts,
			},
			FinishReason:      accFinish,
			GroundingMetadata: accGrounding,
		}}
	}

	resp.Usage = usageFromGenAIResponse(synthResp)
	if combinedErr != nil {
		resp.Error = &spec.Error{Message: combinedErr.Error()}
	}
	resp.Outputs = outputsFromGenAIResponse(synthResp, toolChoiceNameMap, webSearchChoiceID)

	return resp, synthResp, combinedErr
}

func mergeGoogleGenerateContentGroundingMetadata(
	dst *genai.GroundingMetadata,
	src *genai.GroundingMetadata,
) *genai.GroundingMetadata {
	if dst == nil {
		return src
	}
	if src == nil {
		return dst
	}
	out := *dst
	out.WebSearchQueries = append(out.WebSearchQueries, src.WebSearchQueries...)
	out.GroundingChunks = append(out.GroundingChunks, src.GroundingChunks...)
	return &out
}

// prepareGoogleGenerateContentStreamCall applies the configured timeout on the
// outer context for streaming and returns a config clone with HTTP timeout
// cleared only when needed.
//
// This avoids a Google GenAI SDK issue where the SDK creates a derived timeout
// context inside the request setup helper and defers cancel() there, even
// though the returned iterator continues consuming the response body after that
// helper has already returned.
func prepareGoogleGenerateContentStreamCall(
	ctx context.Context,
	config *genai.GenerateContentConfig,
) (streamCtx context.Context, cancel context.CancelFunc, streamConfig *genai.GenerateContentConfig) {
	streamCtx = ctx
	streamConfig = config

	timeout, ok := googleGenerateContentStreamingTimeoutToApply(ctx, config)
	if !ok {
		return streamCtx, nil, streamConfig
	}

	streamCtx, cancel = context.WithTimeout(ctx, timeout)
	streamConfig = cloneGoogleGenerateContentStreamConfigWithoutTimeout(config)
	return streamCtx, cancel, streamConfig
}

// googleGenerateContentStreamingTimeoutToApply reports whether the stream call
// should apply config.HTTPOptions.Timeout on the outer context instead of
// passing it through to the SDK.
//
// We only do this when the SDK would otherwise derive a child timeout context:
//   - timeout is set and > 0
//   - and the parent context has no earlier/equal deadline
func googleGenerateContentStreamingTimeoutToApply(
	ctx context.Context,
	config *genai.GenerateContentConfig,
) (time.Duration, bool) {
	if config == nil || config.HTTPOptions == nil || config.HTTPOptions.Timeout == nil {
		return 0, false
	}

	timeout := *config.HTTPOptions.Timeout
	if timeout <= 0 {
		return 0, false
	}

	if deadline, ok := ctx.Deadline(); ok {
		remaining := time.Until(deadline)
		if remaining <= timeout {
			return 0, false
		}
	}

	return timeout, true
}

// cloneGoogleGenerateContentStreamConfigWithoutTimeout returns a shallow clone
// of config with HTTPOptions.Timeout cleared. This is used only for streaming
// after the timeout has been moved to the outer context.
func cloneGoogleGenerateContentStreamConfigWithoutTimeout(
	config *genai.GenerateContentConfig,
) *genai.GenerateContentConfig {
	if config == nil {
		return nil
	}

	out := *config
	if config.HTTPOptions != nil {
		httpOpts := *config.HTTPOptions
		httpOpts.Timeout = nil
		out.HTTPOptions = &httpOpts
	}
	return &out
}

// consolidateGoogleGenerateContentStreamParts merges the raw per-chunk parts
// accumulated during streaming into the canonical single-part-per-kind form that
// GenerateContent (non-streaming) returns.
//
// Rules:
//   - Adjacent thought parts are merged into one: texts are concatenated and the
//     last non-empty ThoughtSignature wins (signatures arrive on the final chunk).
//   - Adjacent plain-text (answer) parts are merged the same way.
//   - FunctionCall parts are kept verbatim and act as group boundaries: a new
//     text/thought accumulation starts after each function call.
//
// This makes resp.Outputs from the streaming path structurally identical to the
// non-streaming path so that callers and the multi-turn signature round-trip work
// correctly regardless of which path was used.
func consolidateGoogleGenerateContentStreamParts(parts []*genai.Part) []*genai.Part {
	if len(parts) == 0 {
		return nil
	}

	type partKind uint8
	const (
		kindThought  partKind = iota // Thought == true, FunctionCall == nil
		kindText                     // Thought == false, FunctionCall == nil
		kindFuncCall                 // FunctionCall != nil
	)

	classOf := func(p *genai.Part) partKind {
		if p.FunctionCall != nil {
			return kindFuncCall
		}
		if p.Thought {
			return kindThought
		}
		return kindText
	}

	type mergeGroup struct {
		kind     partKind
		textBuf  strings.Builder
		sig      []byte
		funcCall *genai.FunctionCall
	}

	var groups []*mergeGroup

	for _, p := range parts {
		if p == nil {
			continue
		}

		k := classOf(p)

		// Function calls always become their own group (never merged).
		if k == kindFuncCall {
			g := &mergeGroup{kind: kindFuncCall, funcCall: p.FunctionCall}
			if len(p.ThoughtSignature) > 0 {
				g.sig = p.ThoughtSignature
			}
			groups = append(groups, g)
			continue
		}

		// Merge into the immediately preceding same-kind group when present.
		if n := len(groups); n > 0 {
			if last := groups[n-1]; last.kind == k {
				if p.Text != "" {
					last.textBuf.WriteString(p.Text)
				}
				if len(p.ThoughtSignature) > 0 {
					last.sig = p.ThoughtSignature
				}
				continue
			}
		}

		// Start a fresh group.
		g := &mergeGroup{kind: k}
		if p.Text != "" {
			g.textBuf.WriteString(p.Text)
		}
		if len(p.ThoughtSignature) > 0 {
			g.sig = p.ThoughtSignature
		}
		groups = append(groups, g)
	}

	if len(groups) == 0 {
		return nil
	}

	out := make([]*genai.Part, 0, len(groups))
	for _, g := range groups {
		var p *genai.Part
		switch g.kind {
		case kindFuncCall:
			p = &genai.Part{
				FunctionCall:     g.funcCall,
				ThoughtSignature: g.sig,
			}
		case kindThought:
			p = &genai.Part{
				Thought:          true,
				Text:             g.textBuf.String(),
				ThoughtSignature: g.sig,
			}
		default: // kindText
			p = &genai.Part{
				Text:             g.textBuf.String(),
				ThoughtSignature: g.sig,
			}
		}
		out = append(out, p)
	}
	return out
}

// outputsFromGenAIResponse converts a GenerateContentResponse to a slice of
// spec.OutputUnion entries, preserving the natural ordering of the model's
// response parts (thinking before text when the model emits them in that order).
func outputsFromGenAIResponse(
	genResp *genai.GenerateContentResponse,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) []spec.OutputUnion {
	if genResp == nil || len(genResp.Candidates) == 0 {
		return nil
	}

	cand := genResp.Candidates[0]
	if cand == nil || cand.Content == nil {
		return nil
	}

	status := mapGenAIFinishReasonToStatus(cand.FinishReason)
	respID := genResp.ResponseID

	outs := make([]spec.OutputUnion, 0, len(cand.Content.Parts)+2)

	for i, part := range cand.Content.Parts {
		if part == nil {
			continue
		}
		sig := thoughtSignatureToString(part.ThoughtSignature)

		switch {
		case part.FunctionCall != nil:

			fc := part.FunctionCall
			name := strings.TrimSpace(fc.Name)
			if name == "" {
				continue
			}

			tc, ok := toolChoiceNameMap[name]
			if !ok || tc.ID == "" {
				logutil.Debug("googleGenerateContent: received unknown function call in response", "name", name)
				continue
			}

			id := strings.TrimSpace(fc.ID)
			if id == "" {
				id = fmt.Sprintf("%s_%d", name, i)
			}

			argsJSON := "{}"
			if fc.Args != nil {
				if b, marshalErr := json.Marshal(fc.Args); marshalErr == nil {
					argsJSON = string(b)
				}
			}

			call := spec.ToolCall{
				ChoiceID:  tc.ID,
				Type:      tc.Type,
				Role:      spec.RoleAssistant,
				ID:        id,
				CallID:    id,
				Name:      name,
				Arguments: argsJSON,
				Signature: sig,
				Status:    status,
			}

			kind := spec.OutputKindFunctionToolCall
			if tc.Type == spec.ToolTypeCustom {
				kind = spec.OutputKindCustomToolCall
			}

			out := spec.OutputUnion{Kind: kind}
			switch kind {
			case spec.OutputKindCustomToolCall:
				out.CustomToolCall = &call
			default:
				out.FunctionToolCall = &call
			}
			outs = append(outs, out)
		case part.Text != "" || len(part.ThoughtSignature) > 0:
			if part.Thought {
				msg := &spec.ReasoningContent{
					ID:        respID,
					Role:      spec.RoleAssistant,
					Status:    status,
					Signature: sig,
				}
				if part.Text != "" {
					msg.Thinking = []string{part.Text}
				}
				outs = append(outs, spec.OutputUnion{
					Kind:             spec.OutputKindReasoningMessage,
					ReasoningMessage: msg,
				})
				continue
			}

			outs = append(outs, spec.OutputUnion{
				Kind: spec.OutputKindOutputMessage,
				OutputMessage: &spec.InputOutputContent{
					ID:     respID,
					Role:   spec.RoleAssistant,
					Status: status,
					Contents: []spec.InputOutputContentItemUnion{{
						Kind: spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{
							Text:      part.Text,
							Signature: sig,
						},
					}},
				},
			})
		}
	}

	// Grounding metadata → WebSearchToolCall + WebSearchToolOutput.
	if webSearchChoiceID != "" && cand.GroundingMetadata != nil {
		outs = append(outs, groundingToWebSearchOutputs(cand.GroundingMetadata, webSearchChoiceID)...)
	}

	if len(outs) == 0 {
		return nil
	}
	return outs
}

// groundingToWebSearchOutputs converts Google grounding metadata to spec
// WebSearchToolCall (one per query) and a single WebSearchToolOutput (all
// grounding chunks).
func groundingToWebSearchOutputs(
	gm *genai.GroundingMetadata,
	choiceID string,
) []spec.OutputUnion {
	if gm == nil || choiceID == "" {
		return nil
	}

	var outs []spec.OutputUnion
	const callID = "ws_call"
	var callItems []spec.WebSearchToolCallItemUnion

	for _, query := range gm.WebSearchQueries {

		if strings.TrimSpace(query) == "" {
			continue
		}
		callItems = append(callItems, spec.WebSearchToolCallItemUnion{
			Kind:       spec.WebSearchToolCallKindSearch,
			SearchItem: &spec.WebSearchToolCallSearch{Query: query},
		})
	}

	var wsItems []spec.WebSearchToolOutputItemUnion
	for _, chunk := range gm.GroundingChunks {
		if chunk.Web == nil {
			continue
		}
		wsItems = append(wsItems, spec.WebSearchToolOutputItemUnion{
			Kind: spec.WebSearchToolOutputKindSearch,
			SearchItem: &spec.WebSearchToolOutputSearch{
				URL:   chunk.Web.URI,
				Title: chunk.Web.Title,
			},
		})
	}
	if len(callItems) > 0 || len(wsItems) > 0 {
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindWebSearchToolCall,
			WebSearchToolCall: &spec.ToolCall{
				ChoiceID:               choiceID,
				Type:                   spec.ToolTypeWebSearch,
				Role:                   spec.RoleAssistant,
				ID:                     callID,
				CallID:                 callID,
				Name:                   spec.DefaultWebSearchToolName,
				Status:                 spec.StatusCompleted,
				WebSearchToolCallItems: callItems,
			},
		})
	}

	if len(wsItems) > 0 {
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindWebSearchToolOutput,
			WebSearchToolOutput: &spec.ToolOutput{
				ChoiceID:                 choiceID,
				Type:                     spec.ToolTypeWebSearch,
				Role:                     spec.RoleAssistant,
				ID:                       callID,
				CallID:                   callID,
				Name:                     spec.DefaultWebSearchToolName,
				Status:                   spec.StatusCompleted,
				WebSearchToolOutputItems: wsItems,
			},
		})
	}

	return outs
}

func usageFromGenAIResponse(genResp *genai.GenerateContentResponse) *spec.Usage {
	uOut := &spec.Usage{}
	if genResp == nil || genResp.UsageMetadata == nil {
		return uOut
	}
	u := genResp.UsageMetadata

	uOut.InputTokensTotal = int64(u.PromptTokenCount)
	uOut.InputTokensCached = int64(u.CachedContentTokenCount)
	uOut.InputTokensUncached = int64(max(int(u.PromptTokenCount)-int(u.CachedContentTokenCount), 0))
	uOut.OutputTokens = int64(u.CandidatesTokenCount)
	uOut.ReasoningTokens = int64(u.ThoughtsTokenCount)
	return uOut
}

// buildGoogleGenerateContentTools converts spec ToolChoices to a slice of *genai.Tool.
// Function/custom tools go into one Tool{FunctionDeclarations: [...]};
// the web-search tool goes into a separate Tool{GoogleSearch: ...}.
// Returns (tools, functionNameMap, webSearchChoiceID, error).
func buildGoogleGenerateContentTools(
	toolChoices []spec.ToolChoice,
) (tools []*genai.Tool, nameMap map[string]spec.ToolChoice, webSearchChoiceID string, err error) {
	if len(toolChoices) == 0 {
		return nil, nil, "", nil
	}

	ordered, nameMap := sdkutil.BuildToolChoiceNameMapping(toolChoices)

	var funcDecls []*genai.FunctionDeclaration

	webSearchAdded := false

	for _, tw := range ordered {
		tc := tw.Choice
		name := tw.Name

		switch tc.Type {
		case spec.ToolTypeFunction, spec.ToolTypeCustom:
			if name == "" {
				continue
			}
			schema := tc.Arguments
			if schema == nil {
				schema = sdkutil.EmptyJSONArgs
			}
			// Copy so we do not mutate the caller's map.
			schemaCopy := make(map[string]any, len(schema))
			maps.Copy(schemaCopy, schema)

			fd := &genai.FunctionDeclaration{
				Name:                 name,
				ParametersJsonSchema: schemaCopy,
			}
			if desc := sdkutil.ToolDescription(tc); desc != "" {
				fd.Description = desc
			}
			funcDecls = append(funcDecls, fd)

		case spec.ToolTypeWebSearch:
			if webSearchAdded {
				continue
			}
			webSearchChoiceID = tc.ID
			webSearchAdded = true
		}
	}

	if len(funcDecls) > 0 {
		tools = append(tools, &genai.Tool{FunctionDeclarations: funcDecls})
	}

	if webSearchAdded {
		gsTool := &genai.GoogleSearch{}
		// ExcludeDomains is Vertex AI only – set it anyway so it is available
		// when callers target Vertex AI via a custom origin.
		if wsTC, ok := firstWebSearchToolChoice(toolChoices); ok &&
			wsTC.WebSearchArguments != nil &&
			len(wsTC.WebSearchArguments.BlockedDomains) > 0 &&
			len(wsTC.WebSearchArguments.AllowedDomains) == 0 {
			gsTool.ExcludeDomains = wsTC.WebSearchArguments.BlockedDomains
		}
		tools = append(tools, &genai.Tool{GoogleSearch: gsTool})
	}

	return tools, nameMap, webSearchChoiceID, nil
}

// firstWebSearchToolChoice returns the first ToolChoice of type webSearch.
func firstWebSearchToolChoice(tools []spec.ToolChoice) (spec.ToolChoice, bool) {
	for _, tc := range tools {
		if tc.Type == spec.ToolTypeWebSearch {
			return tc, true
		}
	}
	return spec.ToolChoice{}, false
}

// buildGoogleGenerateContentToolConfig converts a spec ToolPolicy to *genai.ToolConfig.
//
//   - auto  → FunctionCallingConfigModeAuto
//   - none  → FunctionCallingConfigModeNone
//   - any   → FunctionCallingConfigModeAny  (+ optional AllowedFunctionNames)
//   - tool  → FunctionCallingConfigModeAny  + specific AllowedFunctionNames
func buildGoogleGenerateContentToolConfig(
	policy *spec.ToolPolicy,
	toolChoiceNameMap map[string]spec.ToolChoice,
) (*genai.ToolConfig, error) {
	if policy == nil {
		return nil, errors.New("invalid policy input")
	}

	fcc := &genai.FunctionCallingConfig{}

	switch policy.Mode {
	case spec.ToolPolicyModeAuto:
		fcc.Mode = genai.FunctionCallingConfigModeAuto

	case spec.ToolPolicyModeNone:
		fcc.Mode = genai.FunctionCallingConfigModeNone

	case spec.ToolPolicyModeAny:
		fcc.Mode = genai.FunctionCallingConfigModeAny
		// Optionally restrict to a named subset.
		if len(policy.AllowedTools) > 0 {
			resolved, err := sdkutil.ResolveAllowedTools(policy.AllowedTools, toolChoiceNameMap)
			if err != nil {
				return nil, err
			}
			callable := filterGoogleGenerateContentCallableResolvedAllowedTool(resolved)
			if len(callable) == 0 {
				return nil, errors.New(
					"googleGenerateContent: toolPolicy=any cannot be satisfied with empty or webSearch-only allowedTools; provide at least one function/custom tool or use auto",
				)
			}
			for _, t := range callable {
				fcc.AllowedFunctionNames = append(fcc.AllowedFunctionNames, t.Name)
			}
		}

	case spec.ToolPolicyModeTool:
		fcc.Mode = genai.FunctionCallingConfigModeAny
		resolved, err := sdkutil.ResolveAllowedTools(policy.AllowedTools, toolChoiceNameMap)
		if err != nil {
			return nil, err
		}
		callable := filterGoogleGenerateContentCallableResolvedAllowedTool(resolved)
		if len(callable) != 1 {
			return nil, errors.New(
				"googleGenerateContent: toolPolicy=tool requires exactly one callable (function/custom) tool; webSearch cannot be forced on Gemini GenerateContent",
			)
		}
		fcc.AllowedFunctionNames = []string{callable[0].Name}

	default:
		return nil, fmt.Errorf("googleGenerateContent: unknown toolPolicy.mode %q", policy.Mode)
	}

	return &genai.ToolConfig{FunctionCallingConfig: fcc}, nil
}

func filterGoogleGenerateContentCallableResolvedAllowedTool(
	tools []sdkutil.ResolvedAllowedTool,
) []sdkutil.ResolvedAllowedTool {
	if len(tools) == 0 {
		return nil
	}

	out := make([]sdkutil.ResolvedAllowedTool, 0, len(tools))
	for _, tc := range tools {
		switch tc.Type {
		case spec.ToolTypeFunction, spec.ToolTypeCustom:
			out = append(out, tc)
		default:
			// Fall.
		}
	}
	return out
}

func filterGoogleGenerateContentCallableToolChoices(toolChoices []spec.ToolChoice) []spec.ToolChoice {
	if len(toolChoices) == 0 {
		return nil
	}

	out := make([]spec.ToolChoice, 0, len(toolChoices))
	for _, tc := range toolChoices {
		switch tc.Type {
		case spec.ToolTypeFunction, spec.ToolTypeCustom:
			out = append(out, tc)
		default:
			// Fall.
		}
	}
	return out
}

func applyGoogleGenerateContentOutputParam(
	config *genai.GenerateContentConfig,
	op *spec.OutputParam,
) error {
	if config == nil || op == nil || op.Format == nil {
		return nil
	}

	switch op.Format.Kind {
	case spec.OutputFormatKindText:
		config.ResponseMIMEType = "text/plain"
		return nil

	case spec.OutputFormatKindJSONSchema:
		if op.Format.JSONSchemaParam == nil || len(op.Format.JSONSchemaParam.Schema) == 0 {
			return errors.New("googleGenerateContent: outputParam.format=jsonSchema requires jsonSchemaParam.schema")
		}
		config.ResponseMIMEType = "application/json"
		// Use ResponseJsonSchema (accepts any) with the raw map to avoid
		// the complexity of converting map[string]any → *genai.Schema.
		config.ResponseJsonSchema = op.Format.JSONSchemaParam.Schema
		return nil

	default:
		return fmt.Errorf("googleGenerateContent: unknown output format kind %q", op.Format.Kind)
	}
}

func mapGenAIFinishReasonToStatus(reason genai.FinishReason) spec.Status {
	switch reason {
	case genai.FinishReasonStop,
		genai.FinishReasonOther,
		genai.FinishReasonUnspecified:
		return spec.StatusCompleted
	case genai.FinishReasonMaxTokens:
		return spec.StatusIncomplete
	case genai.FinishReasonMalformedFunctionCall,
		genai.FinishReasonUnexpectedToolCall,
		genai.FinishReasonSafety,
		genai.FinishReasonRecitation,
		genai.FinishReasonProhibitedContent,
		genai.FinishReasonSPII,
		genai.FinishReasonBlocklist:
		return spec.StatusFailed
	default:
		return spec.StatusCompleted
	}
}
