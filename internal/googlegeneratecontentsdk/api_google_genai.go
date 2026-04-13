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

	req, warns, err := sdkutil.NormalizeRequestForSDK(
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
	applyGoogleGenerateContentThinkingPolicy(config, &req.ModelParam)

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

	emitText := func(chunk string) error {
		if strings.TrimSpace(chunk) == "" {
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
		if strings.TrimSpace(chunk) == "" {
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

	for chunkResp, chunkErr := range client.Models.GenerateContentStream(ctx, string(modelName), contents, config) {
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
			accGrounding = cand.GroundingMetadata
		}
		if cand.Content == nil {
			continue
		}

		for _, part := range cand.Content.Parts {
			if part == nil {
				continue
			}
			accParts = append(accParts, part)

			if part.Text == "" {
				continue
			}
			if part.Thought {
				streamWriteErr = writeThinking(part.Text)
			} else {
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

	var outs []spec.OutputUnion

	// Rolling text / thinking buffers – flushed on type-switch or function call.
	var textBuf strings.Builder
	var thinkBuf strings.Builder
	var thinkSig []byte

	flushText := func() {
		if textBuf.Len() == 0 {
			return
		}
		text := textBuf.String()
		textBuf.Reset()
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindOutputMessage,
			OutputMessage: &spec.InputOutputContent{
				ID:     respID,
				Role:   spec.RoleAssistant,
				Status: status,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				}},
			},
		})
	}

	flushThinking := func() {
		if thinkBuf.Len() == 0 {
			return
		}
		thinking := thinkBuf.String()
		sig := thoughtSignatureToString(thinkSig)
		thinkBuf.Reset()
		thinkSig = nil
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindReasoningMessage,
			ReasoningMessage: &spec.ReasoningContent{
				ID:        respID,
				Role:      spec.RoleAssistant,
				Status:    status,
				Signature: sig,
				Thinking:  []string{thinking},
			},
		})
	}

	for i, part := range cand.Content.Parts {
		if part == nil {
			continue
		}
		switch {
		case part.Text != "" && part.Thought:
			// Thinking / reasoning part – flush any pending regular text first.
			flushText()
			thinkBuf.WriteString(part.Text)
			if len(part.ThoughtSignature) > 0 {
				thinkSig = part.ThoughtSignature
			}

		case part.Text != "" && !part.Thought:
			// Regular text part – flush any pending thinking first.
			flushThinking()
			textBuf.WriteString(part.Text)

		case part.FunctionCall != nil:
			// Function call – flush pending text/thinking before emitting.
			flushText()
			flushThinking()

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
		}
	}

	// Flush any remaining buffered text / thinking.
	flushText()
	flushThinking()

	// Grounding metadata → WebSearchToolCall + WebSearchToolOutput.
	if webSearchChoiceID != "" && cand.GroundingMetadata != nil {
		outs = append(outs, groundingToWebSearchOutputs(cand.GroundingMetadata, webSearchChoiceID)...)
	}

	if len(outs) == 0 {
		return nil
	}
	return outs
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

	for qi, query := range gm.WebSearchQueries {
		if strings.TrimSpace(query) == "" {
			continue
		}
		callID := fmt.Sprintf("ws_call_%d", qi)
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindWebSearchToolCall,
			WebSearchToolCall: &spec.ToolCall{
				ChoiceID: choiceID,
				Type:     spec.ToolTypeWebSearch,
				Role:     spec.RoleAssistant,
				ID:       callID,
				CallID:   callID,
				Name:     spec.DefaultWebSearchToolName,
				Status:   spec.StatusCompleted,
				WebSearchToolCallItems: []spec.WebSearchToolCallItemUnion{{
					Kind:       spec.WebSearchToolCallKindSearch,
					SearchItem: &spec.WebSearchToolCallSearch{Query: query},
				}},
			},
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
	if len(wsItems) > 0 {
		outs = append(outs, spec.OutputUnion{
			Kind: spec.OutputKindWebSearchToolOutput,
			WebSearchToolOutput: &spec.ToolOutput{
				ChoiceID:                 choiceID,
				Type:                     spec.ToolTypeWebSearch,
				Role:                     spec.RoleAssistant,
				ID:                       "ws_output",
				CallID:                   "ws_output",
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
			if err != nil || len(resolved) == 0 {
				return nil, errors.New(
					"googleGenerateContent: toolPolicy=any requires allowedTools with at least one resolvable function/custom tool",
				)
			}
			for _, t := range resolved {
				if t.Type == spec.ToolTypeWebSearch {
					continue
				}
				fcc.AllowedFunctionNames = append(fcc.AllowedFunctionNames, t.Name)
			}
		}

	case spec.ToolPolicyModeTool:
		fcc.Mode = genai.FunctionCallingConfigModeAny
		resolved, err := sdkutil.ResolveAllowedTools(policy.AllowedTools, toolChoiceNameMap)
		if err != nil || len(resolved) == 0 {
			return nil, errors.New(
				"googleGenerateContent: toolPolicy=tool requires exactly one resolvable function/custom tool in allowedTools",
			)
		}
		if len(resolved) != 1 {
			return nil, errors.New(
				"googleGenerateContent: toolPolicy=tool requires exactly one callable (function/custom) tool; webSearch cannot be forced on Gemini GenerateContent",
			)
		}
		fcc.AllowedFunctionNames = []string{resolved[0].Name}

	default:
		return nil, fmt.Errorf("googleGenerateContent: unknown toolPolicy.mode %q", policy.Mode)
	}

	return &genai.ToolConfig{FunctionCallingConfig: fcc}, nil
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
		genai.FinishReasonMalformedFunctionCall,
		genai.FinishReasonUnexpectedToolCall,
		genai.FinishReasonOther,
		genai.FinishReasonUnspecified:
		return spec.StatusCompleted
	case genai.FinishReasonMaxTokens:
		return spec.StatusIncomplete
	case genai.FinishReasonSafety,
		genai.FinishReasonRecitation,
		genai.FinishReasonProhibitedContent,
		genai.FinishReasonSPII,
		genai.FinishReasonBlocklist:
		return spec.StatusFailed
	default:
		return spec.StatusCompleted
	}
}
