package googleinteractionssdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"google.golang.org/genai"
	"google.golang.org/genai/interactions/models/apierrors"
	"google.golang.org/genai/interactions/models/interactions"
	"google.golang.org/genai/interactions/models/operations"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

// GoogleInteractionsAPI implements stateless model interactions through the
// official Google GenAI Interactions SDK.
//
// This adapter deliberately supports only Create requests. It never enables
// server-side storage, background execution, previous interaction IDs, agents,
// or SDK automatic continuation.
type GoogleInteractionsAPI struct {
	ProviderParam *spec.ProviderParam
	debugger      spec.CompletionDebugger
	client        *genai.Client
	mu            sync.RWMutex
}

func NewGoogleInteractionsAPI(
	pi spec.ProviderParam,
	debugger spec.CompletionDebugger,
) (*GoogleInteractionsAPI, error) {
	if pi.Name == "" {
		return nil, errors.New("google interactions api LLM: invalid args")
	}

	if pi.SDKType == "" {
		pi.SDKType = spec.ProviderSDKTypeGoogleInteractions
	}
	if pi.SDKType != spec.ProviderSDKTypeGoogleInteractions {
		return nil, fmt.Errorf(
			"google interactions api LLM: invalid sdk type %q",
			pi.SDKType,
		)
	}
	if strings.TrimSpace(pi.Origin) == "" {
		pi.Origin = spec.DefaultGoogleInteractionsOrigin
	}
	if strings.TrimSpace(pi.ChatCompletionPathPrefix) == "" {
		pi.ChatCompletionPathPrefix = spec.DefaultGoogleInteractionsPrefix
	}
	if strings.TrimSpace(pi.APIKeyHeaderKey) == "" {
		pi.APIKeyHeaderKey = spec.DefaultGoogleInteractionsAPIKeyHeaderKey
	}

	pi.DefaultHeaders = sdkutil.CloneStringMap(pi.DefaultHeaders)

	return &GoogleInteractionsAPI{
		ProviderParam: &pi,
		debugger:      debugger,
	}, nil
}

func (api *GoogleInteractionsAPI) InitLLM(ctx context.Context) error {
	api.mu.Lock()
	defer api.mu.Unlock()

	if api.ProviderParam == nil {
		api.client = nil
		return errors.New("google interactions api LLM: no ProviderParam found")
	}
	if strings.TrimSpace(api.ProviderParam.APIKey) == "" {
		api.client = nil
		logutil.Debug(
			string(api.ProviderParam.Name) +
				": No API key given. Not initializing Google Interactions client",
		)
		return nil
	}

	pi := *api.ProviderParam
	httpOptions := genai.HTTPOptions{}

	baseURL := strings.TrimRight(strings.TrimSpace(pi.Origin), "/")
	if baseURL == "" {
		baseURL = spec.DefaultGoogleInteractionsOrigin
	}
	if prefix := strings.Trim(
		strings.TrimSpace(pi.ChatCompletionPathPrefix),
		"/",
	); prefix != "" {
		baseURL = strings.TrimRight(baseURL+"/"+prefix, "/")
	}
	httpOptions.BaseURL = strings.TrimRight(baseURL, "/") + "/"

	if len(pi.DefaultHeaders) > 0 || strings.TrimSpace(pi.APIKeyHeaderKey) != "" {
		httpOptions.Headers = make(http.Header)
		for key, value := range pi.DefaultHeaders {
			httpOptions.Headers.Set(strings.TrimSpace(key), strings.TrimSpace(value))
		}
		if header := strings.TrimSpace(pi.APIKeyHeaderKey); header != "" {
			httpOptions.Headers.Set(header, pi.APIKey)
		}
	}

	clientConfig := &genai.ClientConfig{
		APIKey:      pi.APIKey,
		Backend:     genai.BackendGeminiAPI,
		HTTPOptions: httpOptions,
	}
	if api.debugger != nil {
		if client := api.debugger.HTTPClient(nil); client != nil {
			clientConfig.HTTPClient = client
		}
	}

	client, err := genai.NewClient(ctx, clientConfig)
	if err != nil {
		return fmt.Errorf(
			"google interactions api LLM: failed to create client: %w",
			err,
		)
	}

	api.client = client
	logutil.Info(
		"google interactions api LLM provider initialized",
		"name", string(pi.Name),
		"URL", baseURL,
	)
	return nil
}

func (api *GoogleInteractionsAPI) DeInitLLM(_ context.Context) error {
	api.mu.Lock()
	var name spec.ProviderName
	if api.ProviderParam != nil {
		name = api.ProviderParam.Name
	}
	api.client = nil
	api.mu.Unlock()

	logutil.Info(
		"google interactions api LLM: provider de initialized",
		"name", string(name),
	)
	return nil
}

func (api *GoogleInteractionsAPI) GetProviderInfo(
	_ context.Context,
) *spec.ProviderParam {
	api.mu.RLock()
	defer api.mu.RUnlock()

	if api.ProviderParam == nil {
		return nil
	}

	out := *api.ProviderParam
	out.DefaultHeaders = sdkutil.CloneStringMap(out.DefaultHeaders)
	return &out
}

func (api *GoogleInteractionsAPI) IsConfigured(_ context.Context) bool {
	api.mu.RLock()
	defer api.mu.RUnlock()

	return api.ProviderParam != nil &&
		strings.TrimSpace(api.ProviderParam.APIKey) != ""
}

func (api *GoogleInteractionsAPI) SetProviderAPIKey(
	_ context.Context,
	apiKey string,
) error {
	api.mu.Lock()
	defer api.mu.Unlock()

	if api.ProviderParam == nil {
		return errors.New("google interactions api LLM: no ProviderParam found")
	}

	api.ProviderParam.APIKey = strings.TrimSpace(apiKey)
	api.client = nil
	return nil
}

func (api *GoogleInteractionsAPI) GetProviderCapability(
	_ context.Context,
) (spec.ModelCapabilities, error) {
	return googleInteractionsSDKCapability(), nil
}

func (api *GoogleInteractionsAPI) FetchCompletion(
	ctx context.Context,
	inReq *spec.FetchCompletionRequest,
	opts *spec.FetchCompletionOptions,
) (*spec.FetchCompletionResponse, error) {
	api.mu.RLock()
	client := api.client
	var providerInfo spec.ProviderParam
	if api.ProviderParam != nil {
		providerInfo = *api.ProviderParam
	}
	api.mu.RUnlock()

	if client == nil {
		return nil, errors.New("google interactions api LLM: client not initialized")
	}
	if inReq == nil || len(inReq.Inputs) == 0 || inReq.ModelParam.Name == "" {
		return nil, errors.New("google interactions api LLM: invalid data")
	}

	req, _, warnings, err := sdkutil.NormalizeRequestForSDK(
		ctx,
		inReq,
		opts,
		spec.ProviderSDKTypeGoogleInteractions,
		googleInteractionsSDKCapability(),
	)
	if err != nil {
		return nil, err
	}

	req.Inputs = sanitizeGoogleInteractionsReasoningInputs(req.Inputs)
	useStream := req.ModelParam.Stream && opts != nil && opts.StreamHandler != nil

	requestBuild, err := buildGoogleInteractionsCreateInteraction(req, useStream)
	if err != nil {
		return nil, err
	}
	warnings = append(warnings, requestBuild.Warnings...)

	request := operations.CreateInteractionRequest{
		Body: operations.NewCreateInteractionRequestBody(requestBuild.Interaction),
	}

	if req.ModelParam.Timeout > 0 {
		timeout := time.Duration(req.ModelParam.Timeout) * time.Second
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	var span spec.CompletionSpan
	if api.debugger != nil {
		ctx, span = api.debugger.StartSpan(ctx, &spec.CompletionSpanStart{
			Provider: providerInfo.Name,
			Model:    req.ModelParam.Name,
			Request:  req,
			Options:  opts,
		})
	}

	var (
		response *spec.FetchCompletionResponse
		raw      any
		callErr  error
	)

	if useStream {
		response, raw, callErr = api.doStreaming(
			ctx,
			client,
			providerInfo.Name,
			req.ModelParam.Name,
			request,
			opts,
			requestBuild.ToolChoiceNameMap,
			requestBuild.WebSearchChoiceID,
		)
	} else {
		response, raw, callErr = api.doNonStreaming(
			ctx,
			client,
			request,
			requestBuild.ToolChoiceNameMap,
			requestBuild.WebSearchChoiceID,
		)
	}

	if response != nil && len(warnings) > 0 {
		response.Warnings = append(response.Warnings, warnings...)
	}

	if span != nil {
		end := spec.CompletionSpanEnd{
			ProviderResponse: raw,
			Response:         response,
			Err:              callErr,
		}
		if response != nil {
			if details := span.End(&end); details != nil && response.DebugDetails == nil {
				response.DebugDetails = details
			}
		} else {
			_ = span.End(&end)
		}
	}

	return response, callErr
}

func (api *GoogleInteractionsAPI) doNonStreaming(
	ctx context.Context,
	client *genai.Client,
	request operations.CreateInteractionRequest,
	toolChoiceNameMap map[string]spec.ToolChoice,
	webSearchChoiceID string,
) (*spec.FetchCompletionResponse, any, error) {
	result, sdkErr := client.Interactions.Create(
		ctx,
		request,
		operations.WithAutomaticContinuation(false),
	)

	var interaction *interactions.Interaction
	if result != nil {
		interaction = result.Interaction
	}
	if interaction == nil {
		interaction = partialGoogleInteractionFromError(sdkErr)
	}

	response, responseErr := responseFromGoogleInteraction(
		interaction,
		toolChoiceNameMap,
		webSearchChoiceID,
	)
	callErr := errors.Join(sdkErr, responseErr)

	if response != nil && callErr != nil && response.Error == nil {
		response.Error = &spec.Error{Message: callErr.Error()}
	}

	return response, result, callErr
}

func partialGoogleInteractionFromError(err error) *interactions.Interaction {
	if err == nil {
		return nil
	}

	if interrupted, ok := errors.AsType[*apierrors.ContinuationInterruptedError](err); ok {
		return interrupted.PartialInteraction
	}

	return nil
}
