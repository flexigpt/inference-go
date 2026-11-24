# Gemini Interactions API

## Overview

The Interactions API is the recommended standard API for new Gemini projects and applications. It supports agentic workflows, server-side state management, and real-time conversations.

Gemini models support multimodal input and output. Supported use cases include reasoning across text and images, content generation, dialogue agents, summarization, classification, audio, video, and code workflows.

## Methods

| Method                | HTTP Request                        | Description                                |
| --------------------- | ----------------------------------- | ------------------------------------------ |
| Create an interaction | `POST /v1/interactions`             | Creates a new interaction.                 |
| Cancel an interaction | `POST /v1/interactions/{id}/cancel` | Cancels an interaction by ID.              |
| Get an interaction    | `GET /v1/interactions/{id}`         | Retrieves full details for an interaction. |
| Delete an interaction | `DELETE /v1/interactions/{id}`      | Deletes an interaction by ID.              |

## Create an Interaction

`POST https://generativelanguage.googleapis.com/v1/interactions`

Creates a new interaction.

### Request Body

| Field                     | Type                                          |                Required | Description                                                                              |
| ------------------------- | --------------------------------------------- | ----------------------: | ---------------------------------------------------------------------------------------- |
| `model`                   | `ModelOption`                                 | When `agent` is not set | Model used to generate the interaction.                                                  |
| `agent`                   | `AgentOption`                                 | When `model` is not set | Managed agent used to generate the interaction.                                          |
| `input`                   | `Content`, `Content[]`, `Step[]`, or `string` |                     Yes | Input for the interaction.                                                               |
| `generation_config`       | `GenerationConfig`                            |                      No | Model configuration. Valid only when `model` is set. Alternative to `agent_config`.      |
| `agent_config`            | `DynamicAgentConfig`                          |                      No | Agent configuration. Valid only when `agent` is set. Alternative to `generation_config`. |
| `system_instruction`      | `string`                                      |                      No | System instruction for the interaction.                                                  |
| `tools`                   | `Tool[]`                                      |                      No | Tools that the model can call.                                                           |
| `response_format`         | `ResponseFormat` or `ResponseFormat[]`        |                      No | Required format for generated output.                                                    |
| `previous_interaction_id` | `string`                                      |                      No | ID of a previous interaction.                                                            |
| `service_tier`            | `ServiceTier`                                 |                      No | Service tier for the interaction.                                                        |
| `safety_settings`         | `SafetySetting[]`                             |                      No | Safety settings.                                                                         |
| `continuation_token`      | `string`                                      |                      No | Token used to resume long decoding.                                                      |
| `stream`                  | `boolean`                                     |                      No | Enables server-sent event streaming.                                                     |
| `store`                   | `boolean`                                     |                      No | Stores the request and response for later retrieval.                                     |
| `background`              | `boolean`                                     |                      No | Runs the interaction in the background.                                                  |

### Model Options

| Model                         | Description                                                                                                       |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `gemma-4-26b-a4b-it`          | Gemma 4 26B A4B IT.                                                                                               |
| `gemma-4-31b-it`              | Gemma 4 31B IT.                                                                                                   |
| `gemini-2.5-flash`            | Hybrid reasoning model with a 1M-token context window and thinking budgets.                                       |
| `gemini-2.5-pro`              | Multipurpose model for coding and complex reasoning tasks.                                                        |
| `gemini-2.5-flash-lite`       | Small and cost-effective model for high-volume workloads.                                                         |
| `gemini-3.5-flash-lite`       | Fast and cost-effective model for high-throughput execution.                                                      |
| `gemini-2.5-flash-image`      | Native image-generation model optimized for speed, flexibility, and contextual understanding.                     |
| `gemini-3.1-flash-lite`       | Cost-efficient model for high-volume agentic tasks, translation, and simple data processing.                      |
| `gemini-3.1-flash-lite-image` | Gemini 3.1 Flash Lite Image.                                                                                      |
| `gemini-3-pro-image`          | Gemini 3 Pro Image.                                                                                               |
| `gemini-3.1-flash-image`      | Gemini 3.1 Flash Image.                                                                                           |
| `gemini-nano-banana-2.1`      | Gemini Nano Banana 2.1.                                                                                           |
| `gemini-3.5-flash`            | Earlier Flash model for routine, high-throughput workloads.                                                       |
| `gemini-3.6-flash`            | Previous-generation Flash model for general agentic and multimodal tasks.                                         |
| `gemini-3.7-flash`            | High-speed Flash model for coding, agentic tool use, and multi-step execution.                                    |
| `gemini-3.8-flash`            | Flash model for long-horizon software engineering, autonomous agents, and complex enterprise workflows.           |
| `gemini-4-argon`              | Frontier model for complex software engineering and knowledge work.                                               |
| `gemini-4-carbon`             | Reasoning model for agentic coding, multimodal knowledge work, scientific discovery, and cybersecurity defense.   |
| `gemini-4-hydrogen`           | Frontier model for fast and dependable execution across coding, agentic workflows, and multimodal knowledge work. |
| `gemini-4-neon`               | Frontier model for complex software engineering, analytical, and science workflows.                               |

### Agent Options

| Agent                               | Description                                                                               |
| ----------------------------------- | ----------------------------------------------------------------------------------------- |
| `deep-research-preview-04-2026`     | Gemini Deep Research Agent.                                                               |
| `deep-research-max-preview-04-2026` | Gemini Deep Research Max Agent.                                                           |
| `antigravity-preview-05-2026`       | Managed agent for multi-step tasks that require reasoning, file operations, and tool use. |
| `antigravity-preview-09-2026`       | Managed agent for multi-step tasks that require reasoning, file operations, and tool use. |

### Generation Configuration

| Field                | Type                                   | Description                                      |
| -------------------- | -------------------------------------- | ------------------------------------------------ |
| `seed`               | `integer`                              | Seed used during decoding for reproducibility.   |
| `stop_sequences`     | `string[]`                             | Character sequences that stop output generation. |
| `tool_choice`        | `ToolChoiceConfig` or `ToolChoiceMode` | Tool-choice configuration.                       |
| `thinking_level`     | `ThinkingLevel`                        | Thought-token level generated by the model.      |
| `thinking_summaries` | `ThinkingSummaries`                    | Controls whether thought summaries are returned. |
| `max_output_tokens`  | `integer`                              | Maximum number of output tokens.                 |
| `video_config`       | `VideoConfig`                          | Video-generation configuration.                  |

#### Tool Choice Modes

- `auto`: The model decides whether to call a tool.
- `any`: The model must choose from available tools.
- `none`: The model cannot call tools.
- `validated`: The model uses validated tool choice.

#### Thinking Levels

- `minimal`: Little or no thinking.
- `low`: Low thinking level.
- `medium`: Medium thinking level.
- `high`: High thinking level.
- `xhigh`: Extra-high thinking level.

#### Thinking Summaries

- `auto`: Automatically include thought summaries when available.
- `none`: Do not include thought summaries.

#### Video Configuration

| Field  | Type        | Description                                                                                                        |
| ------ | ----------- | ------------------------------------------------------------------------------------------------------------------ |
| `task` | `VideoTask` | Optional video-generation task. If omitted, the model selects an appropriate task from the prompt and input media. |

#### Video Tasks

- `text_to_video`: Generates video from a text prompt.
- `image_to_video`: Generates video from one or two source images. The first image is the starting frame. The optional second image is the ending frame.
- `reference_to_video`: Generates video from reference media such as images, audio, or video.
- `edit`: Modifies an input video.
- `extend`: Extends an input video.

### Agent Configuration

`agent_config` is valid only when `agent` is set.

| Field  | Type     | Required | Description        |
| ------ | -------- | -------: | ------------------ |
| `type` | `object` |      Yes | Must be `dynamic`. |

### Service Tiers

- `flex`: Flex service tier.
- `standard`: Standard service tier.
- `priority`: Priority service tier.

### Safety Settings

| Field       | Type              | Required | Description                                             |
| ----------- | ----------------- | -------: | ------------------------------------------------------- |
| `type`      | `HarmCategory`    |      Yes | Harm category to block.                                 |
| `threshold` | `SafetyThreshold` |      Yes | Harm threshold that determines when content is blocked. |
| `method`    | `SafetyMethod`    |       No | Blocking method. Defaults to probability scoring.       |

#### Harm Categories

- `hate_speech`: Content that promotes violence or incites hatred against people or groups based on certain attributes.
- `dangerous_content`: Content that promotes, facilitates, or enables dangerous activities.
- `harassment`: Abusive, threatening, bullying, tormenting, or ridiculing content.
- `sexually_explicit`: Sexually explicit content.
- `jailbreak`: Prompts intended to bypass safety filters.

#### Safety Thresholds

- `block_low_and_above`: Block content with low, medium, or high harm probability.
- `block_medium_and_above`: Block content with medium or high harm probability.
- `block_only_high`: Block content with high harm probability.
- `off`: Disable the safety filter.

#### Safety Methods

- `severity`: Uses probability and severity scores.
- `probability`: Uses the probability score.

### Simple Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput("Hello, how are you?"),
      ),
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Multi-Turn Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  history := []gaos_interactions.Step{
    gaos_interactions.NewStep(gaos_interactions.UserInputStep{
      Content: []gaos_interactions.Content{
        gaos_interactions.NewContent(gaos_interactions.TextContent{
          Text: "Hello!",
        }),
      },
    }),
    gaos_interactions.NewStep(gaos_interactions.ModelOutputStep{
      Content: []gaos_interactions.Content{
        gaos_interactions.NewContent(gaos_interactions.TextContent{
          Text: "Hi there! How can I help you today?",
        }),
      },
    }),
    gaos_interactions.NewStep(gaos_interactions.UserInputStep{
      Content: []gaos_interactions.Content{
        gaos_interactions.NewContent(gaos_interactions.TextContent{
          Text: "What is the capital of France?",
        }),
      },
    }),
  }

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(gaos_interactions.NewInteractionsInput(history)),
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Image Input Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  contents := []gaos_interactions.Content{
    gaos_interactions.NewContent(gaos_interactions.TextContent{
      Text: "What is in this picture?",
    }),
    gaos_interactions.NewContent(gaos_interactions.ImageContent{
      Data:     genai.Ptr("BASE64_ENCODED_IMAGE"),
      MimeType: gaos_interactions.ImageContentMimeTypeImagePng.ToPointer(),
    }),
  }

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(gaos_interactions.NewInteractionsInput(contents)),
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Document Input Go Example

```go
package main

import (
  "context"
  "encoding/base64"
  "fmt"
  "log"
  "os"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  data, err := os.ReadFile("report.pdf")
  if err != nil {
    log.Fatal(err)
  }

  contents := []gaos_interactions.Content{
    gaos_interactions.NewContent(gaos_interactions.TextContent{
      Text: "Summarize this document.",
    }),
    gaos_interactions.NewContent(gaos_interactions.DocumentContent{
      Data: genai.Ptr(base64.StdEncoding.EncodeToString(data)),
      MimeType: gaos_interactions.
        DocumentContentMimeTypeApplicationPdf.
        ToPointer(),
    }),
  }

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-4-hydrogen"),
      Input: genai.Ptr(gaos_interactions.NewInteractionsInput(contents)),
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Function Calling Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  function := gaos_interactions.Function{
    Name:        genai.Ptr("get_weather"),
    Description: genai.Ptr("Get the current weather in a given location."),
    Parameters: map[string]any{
      "type": "object",
      "properties": map[string]any{
        "location": map[string]any{
          "type":        "string",
          "description": "The city and state, for example San Francisco, CA.",
        },
      },
      "required": []string{"location"},
    },
  }

  tool := gaos_interactions.NewTool(function)

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "What is the weather like in Boston, MA?",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction == nil {
    return
  }

  for _, step := range res.Interaction.Steps {
    if step.FunctionCallStep != nil {
      fmt.Println("Function:", step.FunctionCallStep.Name)
      fmt.Println("Arguments:", step.FunctionCallStep.Arguments)
    }
  }
}
```

### Deep Research Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateAgentInteraction{
      Agent: gaos_interactions.AgentOption(
        "deep-research-preview-04-2026",
      ),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "Find current cancer treatment research.",
        ),
      ),
      Background: genai.Ptr(true),
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil {
    fmt.Println(res.Interaction.Status)
  }
}
```

## Interaction Resource

An `Interaction` is the resource returned by the API.

| Field                     | Type                                          | Access           | Description                                                  |
| ------------------------- | --------------------------------------------- | ---------------- | ------------------------------------------------------------ |
| `id`                      | `string`                                      | Output only      | Unique interaction completion ID.                            |
| `status`                  | `InteractionStatus`                           | Output only      | Current interaction status.                                  |
| `created`                 | `string`                                      | Output only      | Creation time in ISO 8601 format: `YYYY-MM-DDThh:mm:ssZ`.    |
| `updated`                 | `string`                                      | Output only      | Last update time in ISO 8601 format: `YYYY-MM-DDThh:mm:ssZ`. |
| `steps`                   | `Step[]`                                      | Output only      | Steps that make up the interaction.                          |
| `model`                   | `ModelOption`                                 | Input and output | Model used for the interaction.                              |
| `agent`                   | `AgentOption`                                 | Input and output | Agent used for the interaction.                              |
| `input`                   | `Content`, `Content[]`, `Step[]`, or `string` | Input            | Interaction input.                                           |
| `generation_config`       | `GenerationConfig`                            | Input            | Model configuration.                                         |
| `agent_config`            | `DynamicAgentConfig`                          | Input            | Agent configuration.                                         |
| `system_instruction`      | `string`                                      | Input            | System instruction.                                          |
| `tools`                   | `Tool[]`                                      | Input            | Tool declarations.                                           |
| `response_format`         | `ResponseFormat` or `ResponseFormat[]`        | Input            | Required response format.                                    |
| `usage`                   | `Usage`                                       | Output only      | Token-use statistics.                                        |
| `previous_interaction_id` | `string`                                      | Input            | Previous interaction ID.                                     |
| `service_tier`            | `ServiceTier`                                 | Input and output | Service tier.                                                |
| `safety_settings`         | `SafetySetting[]`                             | Input            | Safety settings.                                             |
| `errors`                  | `Error[]`                                     | Output only      | Diagnostic errors recorded for the interaction.              |
| `continuation_token`      | `string`                                      | Input and output | Token used to resume long decoding.                          |
| `stream`                  | `boolean`                                     | Input only       | Enables streaming.                                           |
| `store`                   | `boolean`                                     | Input only       | Stores the request and response.                             |
| `background`              | `boolean`                                     | Input only       | Runs the interaction in the background.                      |
| `max_total_tokens`        | `string`                                      | Input            | Maximum token count for an agent run.                        |

### Interaction Status

- `in_progress`: The interaction is in progress.
- `requires_action`: The interaction requires user action or input.
- `completed`: The interaction completed.
- `failed`: The interaction failed.
- `cancelled`: The interaction was cancelled.
- `incomplete`: The interaction completed with incomplete results, for example because it reached the output-token limit.
- `requires_continuation`: The interaction requires another request. Send the returned `continuation_token` in the next request.

## Usage

| Field                         | Type                   | Description                                                       |
| ----------------------------- | ---------------------- | ----------------------------------------------------------------- |
| `total_input_tokens`          | `integer`              | Number of prompt tokens.                                          |
| `initial_input_tokens`        | `integer`              | Number of input tokens in the first turn.                         |
| `input_tokens_by_modality`    | `ModalityTokens[]`     | Input-token breakdown by modality.                                |
| `total_cached_tokens`         | `integer`              | Number of cached prompt tokens.                                   |
| `cached_tokens_by_modality`   | `ModalityTokens[]`     | Cached-token breakdown by modality.                               |
| `total_output_tokens`         | `integer`              | Total generated output tokens.                                    |
| `output_tokens_by_modality`   | `ModalityTokens[]`     | Output-token breakdown by modality.                               |
| `total_tool_use_tokens`       | `integer`              | Tokens used in tool-use prompts.                                  |
| `tool_use_tokens_by_modality` | `ModalityTokens[]`     | Tool-use token breakdown by modality.                             |
| `total_thought_tokens`        | `integer`              | Thought tokens generated by thinking models.                      |
| `total_tokens`                | `integer`              | Total token count, including prompt, output, and internal tokens. |
| `grounding_tool_count`        | `GroundingToolCount[]` | Grounding tool-use counts.                                        |

### Modality Tokens

| Field      | Type               | Description                               |
| ---------- | ------------------ | ----------------------------------------- |
| `modality` | `ResponseModality` | Modality associated with the token count. |
| `tokens`   | `integer`          | Number of tokens.                         |

### Response Modalities

- `text`
- `image`
- `audio`
- `video`
- `document`

### Grounding Tool Count

| Field   | Type                | Description                |
| ------- | ------------------- | -------------------------- |
| `type`  | `GroundingToolType` | Grounding tool type.       |
| `count` | `integer`           | Number of grounding calls. |

#### Grounding Tool Types

- `google_search`: Google Web Search, Image Search, or Web Grounding for Enterprise.
- `google_maps`: Google Maps.

## Content

A `Content` value is one of these types:

- `AudioContent`
- `DocumentContent`
- `ImageContent`
- `TextContent`
- `VideoContent`

### Audio Content

| Field         | Type            | Required | Description                |
| ------------- | --------------- | -------: | -------------------------- |
| `type`        | `string`        |      Yes | Must be `audio`.           |
| `data`        | `string`        |       No | Base64-encoded audio data. |
| `uri`         | `string`        |       No | URI of the audio.          |
| `channels`    | `integer`       |       No | Number of audio channels.  |
| `sample_rate` | `integer`       |       No | Audio sample rate.         |
| `mime_type`   | `AudioMimeType` |       No | Audio MIME type.           |

#### Audio MIME Types

- `audio/wav`
- `audio/mp3`
- `audio/aiff`
- `audio/aac`
- `audio/ogg`
- `audio/flac`
- `audio/mpeg`
- `audio/m4a`
- `audio/l16`
- `audio/opus`
- `audio/alaw`
- `audio/mulaw`
- `audio/webm`

### Document Content

| Field       | Type               | Required | Description                   |
| ----------- | ------------------ | -------: | ----------------------------- |
| `type`      | `string`           |      Yes | Must be `document`.           |
| `data`      | `string`           |       No | Base64-encoded document data. |
| `uri`       | `string`           |       No | URI of the document.          |
| `mime_type` | `DocumentMimeType` |       No | Document MIME type.           |

#### Document MIME Types

- `application/pdf`
- `text/csv`
- `application/vnd.openxmlformats-officedocument.presentationml.presentation`
- `application/vnd.ms-powerpoint`
- `application/vnd.oasis.opendocument.presentation`
- `application/vnd.openxmlformats-officedocument.wordprocessingml.document`
- `application/msword`
- `application/vnd.openxmlformats-officedocument.wordprocessingml.template`
- `application/vnd.oasis.opendocument.text`
- `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- `application/vnd.ms-excel`
- `application/vnd.oasis.opendocument.spreadsheet`

### Image Content

| Field        | Type              | Required | Description                 |
| ------------ | ----------------- | -------: | --------------------------- |
| `type`       | `string`          |      Yes | Must be `image`.            |
| `data`       | `string`          |       No | Base64-encoded image data.  |
| `uri`        | `string`          |       No | URI of the image.           |
| `resolution` | `MediaResolution` |       No | Requested media resolution. |
| `mime_type`  | `ImageMimeType`   |       No | Image MIME type.            |

#### Image MIME Types

- `image/png`
- `image/jpeg`
- `image/webp`
- `image/heic`
- `image/heif`
- `image/gif`
- `image/bmp`
- `image/tiff`

### Text Content

| Field  | Type     | Required | Description     |
| ------ | -------- | -------: | --------------- |
| `type` | `string` |      Yes | Must be `text`. |
| `text` | `string` |      Yes | Text content.   |

### Video Content

| Field        | Type                                       | Required | Description                                                                        |
| ------------ | ------------------------------------------ | -------: | ---------------------------------------------------------------------------------- |
| `type`       | `string`                                   |      Yes | Must be `video`.                                                                   |
| `data`       | `string`                                   |       No | Base64-encoded video data.                                                         |
| `uri`        | `string`                                   |       No | URI of the video.                                                                  |
| `resolution` | `MediaResolution`                          |       No | Requested media resolution.                                                        |
| `processing` | `MediaProcessing` or `VideoProcessingMode` |       No | Video processing method.                                                           |
| `name`       | `string`                                   |       No | User-defined content name. The model can reference the name in its final response. |
| `mime_type`  | `VideoMimeType`                            |       No | Video MIME type.                                                                   |

#### Media Resolutions

- `low`
- `medium`
- `high`
- `ultra_high`

#### Media Processing

| Field          | Type     | Required | Description                                                                                                              |
| -------------- | -------- | -------: | ------------------------------------------------------------------------------------------------------------------------ |
| `type`         | `string` |      Yes | Must be `static`.                                                                                                        |
| `start_offset` | `string` |       No | Segment start time in seconds with an `s` suffix, for example `10.5s`. Must be non-negative.                             |
| `end_offset`   | `string` |       No | Segment end time in seconds with an `s` suffix, for example `30s`. Must be non-negative and greater than `start_offset`. |
| `fps`          | `number` |       No | Video frame-rate sampling density.                                                                                       |

#### Video Processing Modes

- `static`: Fixed-rate frame extraction. All selected frames are placed in context.
- `agentic`: Model-driven dynamic navigation.

#### Video MIME Types

- `video/mp4`
- `video/mpeg`
- `video/mpg`
- `video/mov`
- `video/avi`
- `video/x-flv`
- `video/webm`
- `video/wmv`
- `video/3gpp`
- `video/jpeg2000`

## Tools

A `Tool` is one of these types:

- `CodeExecution`
- `ComputerUse`
- `FileSearch`
- `Function`
- `GoogleMaps`
- `GoogleSearch`
- `McpServer`
- `UrlContext`

### Code Execution

Enables model-managed code execution.

| Field  | Type     | Required | Description               |
| ------ | -------- | -------: | ------------------------- |
| `type` | `string` |      Yes | Must be `code_execution`. |

#### Code Execution Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.CodeExecution{})

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "Calculate the first 10 Fibonacci numbers.",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Computer Use

Enables model-controlled interaction with a browser, mobile environment, or desktop environment.

| Field                               | Type                     | Required | Description                                        |
| ----------------------------------- | ------------------------ | -------: | -------------------------------------------------- |
| `type`                              | `string`                 |      Yes | Must be `computer_use`.                            |
| `environment`                       | `ComputerEnvironment`    |       No | Operating environment.                             |
| `excluded_predefined_functions`     | `string[]`               |       No | Predefined functions excluded from the model call. |
| `enable_prompt_injection_detection` | `boolean`                |       No | Enables prompt-injection detection.                |
| `disabled_safety_policies`          | `ComputerSafetyPolicy[]` |       No | Disabled computer-use safety policies.             |

#### Computer Environments

- `browser`
- `mobile`
- `desktop`

#### Computer Safety Policies

- `financial_transactions`
- `sensitive_data_modification`
- `communication_tool`
- `account_creation`
- `data_modification`
- `user_consent_management`
- `legal_terms_and_agreements`

#### Computer Use Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.ComputerUse{})

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput("Find a flight to Tokyo."),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### File Search

Searches one or more file-search stores.

| Field                     | Type       | Required | Description                                                         |
| ------------------------- | ---------- | -------: | ------------------------------------------------------------------- |
| `type`                    | `string`   |      Yes | Must be `file_search`.                                              |
| `file_search_store_names` | `string[]` |       No | File-search store names to search.                                  |
| `top_k`                   | `integer`  |       No | Number of semantic retrieval chunks to return.                      |
| `metadata_filter`         | `string`   |       No | Metadata filter applied to semantic retrieval documents and chunks. |

#### File Search Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.FileSearch{
    FileSearchStoreNames: []string{
      "fileSearchStores/your-store-id",
    },
  })

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "What documents are available?",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Function

Declares an application function that the model can call.

| Field         | Type     | Required | Description                          |
| ------------- | -------- | -------: | ------------------------------------ |
| `type`        | `string` |      Yes | Must be `function`.                  |
| `name`        | `string` |       No | Function name.                       |
| `description` | `string` |       No | Function description.                |
| `parameters`  | `object` |       No | JSON Schema for function parameters. |

The API supports a subset of JSON Schema. The following keywords are ignored and are not enforced:

- `minLength`
- `maxLength`
- `minProperties`
- `maxProperties`
- `uniqueItems`
- `contains`
- `minContains`
- `maxContains`

### Google Maps

Enables the model to call Google Maps.

| Field           | Type      | Required | Description                                        |
| --------------- | --------- | -------: | -------------------------------------------------- |
| `type`          | `string`  |      Yes | Must be `google_maps`.                             |
| `enable_widget` | `boolean` |       No | Returns a widget context token in the tool result. |
| `latitude`      | `number`  |       No | User latitude.                                     |
| `longitude`     | `number`  |       No | User longitude.                                    |

#### Google Maps Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.GoogleMaps{
    Latitude:  genai.Ptr(37.7749),
    Longitude: genai.Ptr(-122.4194),
  })

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "What is the best food near me?",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### Google Search

Enables Google Search grounding.

| Field          | Type           | Required | Description                       |
| -------------- | -------------- | -------: | --------------------------------- |
| `type`         | `string`       |      Yes | Must be `google_search`.          |
| `search_types` | `SearchType[]` |       No | Search-grounding types to enable. |

#### Search Types

- `web_search`: Returns text web-search results.
- `image_search`: Returns image-search results and image bytes.

#### Google Search Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.GoogleSearch{})

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "Who is the current president of France?",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### MCP Server

Declares an MCP server that the model can use to perform actions.

| Field           | Type             | Required | Description                                                   |
| --------------- | ---------------- | -------: | ------------------------------------------------------------- |
| `type`          | `string`         |      Yes | Must be `mcp_server`.                                         |
| `name`          | `string`         |       No | MCP server name.                                              |
| `url`           | `string`         |       No | Full MCP server endpoint URL.                                 |
| `headers`       | `object`         |       No | Headers for authentication, timeouts, or other configuration. |
| `allowed_tools` | `AllowedTools[]` |       No | Allowed MCP tools.                                            |

#### Allowed Tools

| Field   | Type             | Description             |
| ------- | ---------------- | ----------------------- |
| `mode`  | `ToolChoiceMode` | Tool selection mode.    |
| `tools` | `string[]`       | Names of allowed tools. |

#### MCP Server Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  mcpServer := gaos_interactions.MCPServer{
    Name: genai.Ptr("weather_service"),
    URL:  genai.Ptr("https://gemini-api-demos.uc.r.appspot.com/mcp"),
  }

  tool := gaos_interactions.NewTool(mcpServer)

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "Today is 12-05-2025. What is the temperature in London?",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

### URL Context

Enables the model to fetch context from URLs.

| Field  | Type     | Required | Description            |
| ------ | -------- | -------: | ---------------------- |
| `type` | `string` |      Yes | Must be `url_context`. |

#### URL Context Go Example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "google.golang.org/genai"
  gaos_interactions "google.golang.org/genai/interactions/models/interactions"
  "google.golang.org/genai/interactions/models/operations"
)

func main() {
  ctx := context.Background()

  client, err := genai.NewClient(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  tool := gaos_interactions.NewTool(gaos_interactions.URLContext{})

  body := operations.NewCreateInteractionRequestBody(
    gaos_interactions.CreateModelInteraction{
      Model: gaos_interactions.Model("gemini-3.8-flash"),
      Input: genai.Ptr(
        gaos_interactions.NewInteractionsInput(
          "Summarize https://www.example.com.",
        ),
      ),
      Tools: []gaos_interactions.Tool{tool},
    },
  )

  res, err := client.Interactions.Create(
    ctx,
    operations.CreateInteractionRequest{Body: body},
  )
  if err != nil {
    log.Fatal(err)
  }

  if res.Interaction != nil && res.Interaction.OutputText != nil {
    fmt.Println(*res.Interaction.OutputText)
  }
}
```

## Response Formats

A `ResponseFormat` is one of these types:

- `AudioResponseFormat`
- `ImageResponseFormat`
- `TextResponseFormat`
- `VideoResponseFormat`

### Audio Response Format

| Field         | Type                  | Required | Description                                                 |
| ------------- | --------------------- | -------: | ----------------------------------------------------------- |
| `type`        | `string`              |      Yes | Must be `audio`.                                            |
| `mime_type`   | `AudioOutputMimeType` |       No | Audio output MIME type.                                     |
| `delivery`    | `DeliveryMode`        |       No | Audio delivery mode.                                        |
| `sample_rate` | `integer`             |       No | Audio sample rate in Hz.                                    |
| `bit_rate`    | `integer`             |       No | Bit rate in bits per second. Applies to compressed formats. |

#### Audio Output MIME Types

- `audio/mp3`
- `audio/ogg_opus`
- `audio/l16`
- `audio/wav`
- `audio/alaw`
- `audio/mulaw`

### Image Response Format

| Field          | Type               | Required | Description                                             |
| -------------- | ------------------ | -------: | ------------------------------------------------------- |
| `type`         | `string`           |      Yes | Must be `image`.                                        |
| `mime_type`    | `string`           |       No | Image output MIME type. Only `image/jpeg` is supported. |
| `delivery`     | `DeliveryMode`     |       No | Image delivery mode.                                    |
| `aspect_ratio` | `ImageAspectRatio` |       No | Image aspect ratio.                                     |
| `image_size`   | `ImageSize`        |       No | Image output size.                                      |

#### Image Aspect Ratios

- `1:1`
- `2:3`
- `3:2`
- `3:4`
- `4:3`
- `4:5`
- `5:4`
- `9:16`
- `16:9`
- `21:9`
- `1:8`
- `8:1`
- `1:4`
- `4:1`

#### Image Sizes

- `512`
- `1K`
- `2K`
- `4K`

### Text Response Format

| Field       | Type                 | Required | Description                           |
| ----------- | -------------------- | -------: | ------------------------------------- |
| `type`      | `string`             |      Yes | Must be `text`.                       |
| `mime_type` | `TextOutputMimeType` |       No | Text output MIME type.                |
| `schema`    | `object`             |       No | JSON Schema required for JSON output. |

#### Text Output MIME Types

- `application/json`
- `text/plain`

For `application/json`, the `schema` field defines the required JSON output structure.

The API ignores these JSON Schema keywords:

- `minLength`
- `maxLength`
- `minProperties`
- `maxProperties`
- `uniqueItems`
- `contains`
- `minContains`
- `maxContains`

The API enforces `minItems` and `maxItems`. Schemas that are too complex to enforce return `400 INVALID_ARGUMENT`.

### Video Response Format

| Field          | Type                    | Required | Description                                  |
| -------------- | ----------------------- | -------: | -------------------------------------------- |
| `type`         | `string`                |      Yes | Must be `video`.                             |
| `delivery`     | `DeliveryMode`          |       No | Video delivery mode.                         |
| `aspect_ratio` | `VideoAspectRatio`      |       No | Video aspect ratio.                          |
| `duration`     | `string`                |       No | Video duration.                              |
| `resolution`   | `VideoOutputResolution` |       No | Video output resolution. Defaults to `720p`. |

#### Delivery Modes

- `inline`: Data is included directly in the response.
- `uri`: Data is returned as a URI.

#### Video Aspect Ratios

- `16:9`
- `9:16`

#### Video Output Resolutions

- `360p`
- `720p`
- `1080p`
- `4k`

## Steps

A `Step` is one of these types:

- `CodeExecutionCallStep`
- `CodeExecutionResultStep`
- `FileSearchCallStep`
- `FileSearchResultStep`
- `FunctionCallStep`
- `FunctionResultStep`
- `GoogleMapsCallStep`
- `GoogleMapsResultStep`
- `GoogleSearchCallStep`
- `GoogleSearchResultStep`
- `McpServerToolCallStep`
- `McpServerToolResultStep`
- `ModelOutputStep`
- `ProcessingCallStep`
- `ProcessingResultStep`
- `ThoughtStep`
- `UrlContextCallStep`
- `UrlContextResultStep`
- `UserInputStep`

### Code Execution Call Step

| Field       | Type                     | Required | Description                    |
| ----------- | ------------------------ | -------: | ------------------------------ |
| `type`      | `string`                 |      Yes | Must be `code_execution_call`. |
| `id`        | `string`                 |      Yes | Unique tool-call ID.           |
| `arguments` | `CodeExecutionArguments` |      Yes | Code-execution arguments.      |
| `signature` | `string`                 |       No | Backend validation signature.  |

#### Code Execution Arguments

| Field      | Type           | Description      |
| ---------- | -------------- | ---------------- |
| `language` | `CodeLanguage` | Code language.   |
| `code`     | `string`       | Code to execute. |

#### Code Languages

- `python`: Python 3.10 or later, with NumPy and SimPy available.
- `bash`: Bash. Available only for Gemini 3.0 models and later.

### Code Execution Result Step

| Field       | Type                    | Required | Description                            |
| ----------- | ----------------------- | -------: | -------------------------------------- |
| `type`      | `string`                |      Yes | Must be `code_execution_result`.       |
| `call_id`   | `string`                |      Yes | ID of the related code-execution call. |
| `result`    | `Content[]` or `string` |      Yes | Execution output.                      |
| `is_error`  | `boolean`               |       No | Indicates an execution error.          |
| `signature` | `string`                |       No | Backend validation signature.          |

### File Search Call Step

| Field       | Type     | Required | Description                   |
| ----------- | -------- | -------: | ----------------------------- |
| `type`      | `string` |      Yes | Must be `file_search_call`.   |
| `id`        | `string` |      Yes | Unique tool-call ID.          |
| `signature` | `string` |       No | Backend validation signature. |

### File Search Result Step

| Field       | Type     | Required | Description                         |
| ----------- | -------- | -------: | ----------------------------------- |
| `type`      | `string` |      Yes | Must be `file_search_result`.       |
| `call_id`   | `string` |      Yes | ID of the related file-search call. |
| `signature` | `string` |       No | Backend validation signature.       |

### Function Call Step

| Field       | Type     | Required | Description              |
| ----------- | -------- | -------: | ------------------------ |
| `type`      | `string` |      Yes | Must be `function_call`. |
| `name`      | `string` |      Yes | Function name.           |
| `arguments` | `object` |      Yes | Function arguments.      |
| `id`        | `string` |      Yes | Unique tool-call ID.     |

### Function Result Step

| Field      | Type                               | Required | Description                      |
| ---------- | ---------------------------------- | -------: | -------------------------------- |
| `type`     | `string`                           |      Yes | Must be `function_result`.       |
| `call_id`  | `string`                           |      Yes | ID of the related function call. |
| `result`   | `Content[]`, `object`, or `string` |      Yes | Function result.                 |
| `name`     | `string`                           |       No | Function name.                   |
| `is_error` | `boolean`                          |       No | Indicates a function error.      |

### Google Maps Call Step

| Field       | Type                      | Required | Description                   |
| ----------- | ------------------------- | -------: | ----------------------------- |
| `type`      | `string`                  |      Yes | Must be `google_maps_call`.   |
| `id`        | `string`                  |      Yes | Unique tool-call ID.          |
| `arguments` | `GoogleMapsCallArguments` |       No | Google Maps call arguments.   |
| `signature` | `string`                  |       No | Backend validation signature. |

#### Google Maps Call Arguments

| Field     | Type       | Description         |
| --------- | ---------- | ------------------- |
| `queries` | `string[]` | Queries to execute. |

### Google Maps Result Step

| Field       | Type                     | Required | Description                         |
| ----------- | ------------------------ | -------: | ----------------------------------- |
| `type`      | `string`                 |      Yes | Must be `google_maps_result`.       |
| `call_id`   | `string`                 |      Yes | ID of the related Google Maps call. |
| `result`    | `GoogleMapsResultItem[]` |      Yes | Google Maps results.                |
| `signature` | `string`                 |       No | Backend validation signature.       |

#### Google Maps Result Item

| Field                  | Type                | Description                       |
| ---------------------- | ------------------- | --------------------------------- |
| `places`               | `GoogleMapsPlace[]` | Places found by the query.        |
| `widget_context_token` | `string`            | Google Maps widget context token. |

#### Google Maps Place

| Field             | Type              | Description                                    |
| ----------------- | ----------------- | ---------------------------------------------- |
| `place_id`        | `string`          | Place ID in `places/{place_id}` format.        |
| `name`            | `string`          | Place name.                                    |
| `url`             | `string`          | Place URL.                                     |
| `review_snippets` | `ReviewSnippet[]` | Review snippets used to generate the response. |

#### Review Snippet

| Field       | Type     | Description             |
| ----------- | -------- | ----------------------- |
| `title`     | `string` | Review title.           |
| `url`       | `string` | Google Maps review URL. |
| `review_id` | `string` | Review snippet ID.      |

### Google Search Call Step

| Field         | Type                        | Required | Description                    |
| ------------- | --------------------------- | -------: | ------------------------------ |
| `type`        | `string`                    |      Yes | Must be `google_search_call`.  |
| `id`          | `string`                    |      Yes | Unique tool-call ID.           |
| `arguments`   | `GoogleSearchCallArguments` |      Yes | Google Search call arguments.  |
| `search_type` | `SearchType`                |       No | Enabled grounding search type. |
| `signature`   | `string`                    |       No | Backend validation signature.  |

#### Google Search Call Arguments

| Field     | Type       | Description         |
| --------- | ---------- | ------------------- |
| `queries` | `string[]` | Web-search queries. |

### Google Search Result Step

| Field       | Type                       | Required | Description                           |
| ----------- | -------------------------- | -------: | ------------------------------------- |
| `type`      | `string`                   |      Yes | Must be `google_search_result`.       |
| `call_id`   | `string`                   |      Yes | ID of the related Google Search call. |
| `result`    | `GoogleSearchResultItem[]` |      Yes | Search results.                       |
| `is_error`  | `boolean`                  |       No | Indicates a Google Search error.      |
| `signature` | `string`                   |       No | Backend validation signature.         |

#### Google Search Result Item

| Field                | Type     | Description                                                             |
| -------------------- | -------- | ----------------------------------------------------------------------- |
| `search_suggestions` | `string` | Web content snippet suitable for display in an application or web view. |

### MCP Server Tool Call Step

| Field         | Type     | Required | Description                     |
| ------------- | -------- | -------: | ------------------------------- |
| `type`        | `string` |      Yes | Must be `mcp_server_tool_call`. |
| `name`        | `string` |      Yes | Name of the called tool.        |
| `server_name` | `string` |      Yes | Name of the MCP server.         |
| `arguments`   | `object` |      Yes | Tool arguments.                 |
| `id`          | `string` |      Yes | Unique tool-call ID.            |

### MCP Server Tool Result Step

| Field         | Type                               | Required | Description                       |
| ------------- | ---------------------------------- | -------: | --------------------------------- |
| `type`        | `string`                           |      Yes | Must be `mcp_server_tool_result`. |
| `call_id`     | `string`                           |      Yes | ID of the related MCP tool call.  |
| `result`      | `Content[]`, `object`, or `string` |      Yes | MCP tool result.                  |
| `name`        | `string`                           |       No | MCP tool name.                    |
| `server_name` | `string`                           |       No | MCP server name.                  |

### Model Output Step

| Field     | Type        | Required | Description                     |
| --------- | ----------- | -------: | ------------------------------- |
| `type`    | `string`    |      Yes | Must be `model_output`.         |
| `content` | `Content[]` |       No | Content generated by the model. |

### Processing Call Step

A server-initiated media-processing step, for example for video understanding.

| Field  | Type     | Required | Description                |
| ------ | -------- | -------: | -------------------------- |
| `type` | `string` |      Yes | Must be `processing_call`. |
| `id`   | `string` |      Yes | Unique processing-call ID. |

### Processing Result Step

| Field     | Type     | Required | Description                        |
| --------- | -------- | -------: | ---------------------------------- |
| `type`    | `string` |      Yes | Must be `processing_result`.       |
| `call_id` | `string` |      Yes | ID of the related processing call. |

### Thought Step

| Field       | Type        | Required | Description                   |
| ----------- | ----------- | -------: | ----------------------------- |
| `type`      | `string`    |      Yes | Must be `thought`.            |
| `signature` | `string`    |       No | Backend validation signature. |
| `summary`   | `Content[]` |       No | Summary of the thought.       |

### URL Context Call Step

| Field       | Type                      | Required | Description                   |
| ----------- | ------------------------- | -------: | ----------------------------- |
| `type`      | `string`                  |      Yes | Must be `url_context_call`.   |
| `id`        | `string`                  |      Yes | Unique tool-call ID.          |
| `arguments` | `UrlContextCallArguments` |      Yes | URL context call arguments.   |
| `signature` | `string`                  |       No | Backend validation signature. |

#### URL Context Call Arguments

| Field  | Type       | Description    |
| ------ | ---------- | -------------- |
| `urls` | `string[]` | URLs to fetch. |

### URL Context Result Step

| Field       | Type                 | Required | Description                         |
| ----------- | -------------------- | -------: | ----------------------------------- |
| `type`      | `string`             |      Yes | Must be `url_context_result`.       |
| `call_id`   | `string`             |      Yes | ID of the related URL context call. |
| `result`    | `UrlContextResult[]` |      Yes | URL retrieval results.              |
| `is_error`  | `boolean`            |       No | Indicates a URL context error.      |
| `signature` | `string`             |       No | Backend validation signature.       |

#### URL Context Result

| Field    | Type                 | Description       |
| -------- | -------------------- | ----------------- |
| `url`    | `string`             | Fetched URL.      |
| `status` | `UrlRetrievalStatus` | Retrieval status. |

#### URL Retrieval Statuses

- `success`: URL retrieval succeeded.
- `error`: URL retrieval failed because of an error.
- `paywall`: URL retrieval failed because the content is behind a paywall.
- `unsafe`: URL retrieval failed because the content is unsafe.

### User Input Step

| Field     | Type        | Required | Description                   |
| --------- | ----------- | -------: | ----------------------------- |
| `type`    | `string`    |      Yes | Must be `user_input`.         |
| `content` | `Content[]` |       No | Content provided by the user. |

## Tool Choice Configuration

| Field           | Type           | Description                         |
| --------------- | -------------- | ----------------------------------- |
| `allowed_tools` | `AllowedTools` | Allowed tools and tool-choice mode. |

## Streaming Events

When `stream` is `true`, the API returns server-sent events with content type `text/event-stream`.

An `InteractionSseEvent` is one of these types:

- `ErrorEvent`
- `InteractionCreatedEvent`
- `InteractionCompletedEvent`
- `InteractionStatusUpdate`
- `StepStart`
- `StepDelta`
- `StepStop`

### Error Event

| Field        | Type     | Required | Description                      |
| ------------ | -------- | -------: | -------------------------------- |
| `event_type` | `string` |      Yes | Must be `error`.                 |
| `error`      | `Error`  |       No | Error details.                   |
| `event_id`   | `string` |       No | Token used to resume the stream. |

### Interaction Created Event

| Field         | Type                             | Required | Description                      |
| ------------- | -------------------------------- | -------: | -------------------------------- |
| `event_type`  | `string`                         |      Yes | Must be `interaction.created`.   |
| `interaction` | `InteractionSseEventInteraction` |      Yes | Partial interaction resource.    |
| `event_id`    | `string`                         |       No | Token used to resume the stream. |

### Interaction Completed Event

An `interaction.completed` event signals that the interaction ended. The interaction can end after completion, cancellation, or natural termination. No further input can be sent after this event.

| Field         | Type                             | Required | Description                      |
| ------------- | -------------------------------- | -------: | -------------------------------- |
| `event_type`  | `string`                         |      Yes | Must be `interaction.completed`. |
| `interaction` | `InteractionSseEventInteraction` |      Yes | Completed interaction resource.  |
| `event_id`    | `string`                         |       No | Token used to resume the stream. |

### Interaction SSE Event Interaction

| Field          | Type                | Description                          |
| -------------- | ------------------- | ------------------------------------ |
| `id`           | `string`            | Unique interaction completion ID.    |
| `status`       | `InteractionStatus` | Interaction status.                  |
| `model`        | `string`            | Model that completes the prompt.     |
| `created`      | `string`            | Creation time in ISO 8601 format.    |
| `updated`      | `string`            | Last update time in ISO 8601 format. |
| `usage`        | `Usage`             | Token-use statistics.                |
| `service_tier` | `ServiceTier`       | Service tier.                        |
| `steps`        | `Step[]`            | Interaction steps, when included.    |
| `agent`        | `string`            | Agent used by the interaction.       |
| `object`       | `string`            | Resource type.                       |

### Interaction Status Update

| Field                | Type                | Required | Description                                               |
| -------------------- | ------------------- | -------: | --------------------------------------------------------- |
| `event_type`         | `string`            |      Yes | Must be `interaction.status_update`.                      |
| `interaction_id`     | `string`            |      Yes | Interaction ID.                                           |
| `status`             | `InteractionStatus` |      Yes | Updated interaction status.                               |
| `continuation_token` | `string`            |       No | Token used to resume decoding from the latest checkpoint. |
| `usage`              | `Usage`             |       No | Model-usage statistics at the checkpoint.                 |
| `event_id`           | `string`            |       No | Token used to resume the stream.                          |

### Step Start

| Field        | Type      | Required | Description                      |
| ------------ | --------- | -------: | -------------------------------- |
| `event_type` | `string`  |      Yes | Must be `step.start`.            |
| `index`      | `integer` |      Yes | Step index.                      |
| `step`       | `Step`    |      Yes | Started step.                    |
| `event_id`   | `string`  |       No | Token used to resume the stream. |

### Step Delta

| Field        | Type                | Required | Description                      |
| ------------ | ------------------- | -------: | -------------------------------- |
| `event_type` | `string`            |      Yes | Must be `step.delta`.            |
| `index`      | `integer`           |      Yes | Step index.                      |
| `delta`      | `StepDeltaData`     |      Yes | Incremental step data.           |
| `metadata`   | `StepDeltaMetadata` |       No | Step metadata.                   |
| `event_id`   | `string`            |       No | Token used to resume the stream. |

#### Step Delta Metadata

| Field         | Type    | Description                             |
| ------------- | ------- | --------------------------------------- |
| `total_usage` | `Usage` | Total interaction token-use statistics. |

### Step Stop

| Field        | Type      | Required | Description                                        |
| ------------ | --------- | -------: | -------------------------------------------------- |
| `event_type` | `string`  |      Yes | Must be `step.stop`.                               |
| `index`      | `integer` |      Yes | Step index.                                        |
| `usage`      | `Usage`   |       No | Cumulative model usage from the interaction start. |
| `step_usage` | `Usage`   |       No | Usage for this specific step.                      |
| `event_id`   | `string`  |       No | Token used to resume the stream.                   |

## Step Delta Data

A `StepDeltaData` value is one of these types:

- `ArgumentsDelta`
- `AudioDelta`
- `CodeExecutionCallDelta`
- `CodeExecutionResultDelta`
- `DocumentDelta`
- `FileSearchCallDelta`
- `FileSearchResultDelta`
- `FunctionResultDelta`
- `GoogleMapsCallDelta`
- `GoogleMapsResultDelta`
- `GoogleSearchCallDelta`
- `GoogleSearchResultDelta`
- `ImageDelta`
- `McpServerToolCallDelta`
- `McpServerToolResultDelta`
- `TextDelta`
- `ThoughtSignatureDelta`
- `ThoughtSummaryDelta`
- `UrlContextCallDelta`
- `UrlContextResultDelta`
- `VideoDelta`

### Arguments Delta

| Field       | Type     | Required | Description                       |
| ----------- | -------- | -------: | --------------------------------- |
| `type`      | `string` |      Yes | Must be `arguments_delta`.        |
| `arguments` | `string` |       No | Incremental serialized arguments. |

### Text Delta

| Field  | Type     | Required | Description               |
| ------ | -------- | -------: | ------------------------- |
| `type` | `string` |      Yes | Must be `text`.           |
| `text` | `string` |      Yes | Incremental text content. |

### Audio Delta

| Field         | Type            | Required | Description                |
| ------------- | --------------- | -------: | -------------------------- |
| `type`        | `string`        |      Yes | Must be `audio`.           |
| `data`        | `string`        |       No | Base64-encoded audio data. |
| `uri`         | `string`        |       No | Audio URI.                 |
| `mime_type`   | `AudioMimeType` |       No | Audio MIME type.           |
| `sample_rate` | `integer`       |       No | Audio sample rate.         |
| `channels`    | `integer`       |       No | Number of audio channels.  |

### Document Delta

| Field       | Type               | Required | Description                   |
| ----------- | ------------------ | -------: | ----------------------------- |
| `type`      | `string`           |      Yes | Must be `document`.           |
| `data`      | `string`           |       No | Base64-encoded document data. |
| `uri`       | `string`           |       No | Document URI.                 |
| `mime_type` | `DocumentMimeType` |       No | Document MIME type.           |

### Image Delta

| Field        | Type              | Required | Description                |
| ------------ | ----------------- | -------: | -------------------------- |
| `type`       | `string`          |      Yes | Must be `image`.           |
| `data`       | `string`          |       No | Base64-encoded image data. |
| `uri`        | `string`          |       No | Image URI.                 |
| `mime_type`  | `ImageMimeType`   |       No | Image MIME type.           |
| `resolution` | `MediaResolution` |       No | Image resolution.          |

### Video Delta

| Field        | Type              | Required | Description                |
| ------------ | ----------------- | -------: | -------------------------- |
| `type`       | `string`          |      Yes | Must be `video`.           |
| `data`       | `string`          |       No | Base64-encoded video data. |
| `uri`        | `string`          |       No | Video URI.                 |
| `mime_type`  | `VideoMimeType`   |       No | Video MIME type.           |
| `resolution` | `MediaResolution` |       No | Video resolution.          |

### Thought Signature Delta

| Field       | Type     | Required | Description                  |
| ----------- | -------- | -------: | ---------------------------- |
| `type`      | `string` |      Yes | Must be `thought_signature`. |
| `signature` | `string` |       No | Backend thought signature.   |

### Thought Summary Delta

| Field     | Type      | Required | Description                   |
| --------- | --------- | -------: | ----------------------------- |
| `type`    | `string`  |      Yes | Must be `thought_summary`.    |
| `content` | `Content` |       No | Thought-summary content item. |

### Code Execution Call Delta

| Field       | Type                     | Required | Description                    |
| ----------- | ------------------------ | -------: | ------------------------------ |
| `type`      | `string`                 |      Yes | Must be `code_execution_call`. |
| `arguments` | `CodeExecutionArguments` |      Yes | Code-execution arguments.      |
| `signature` | `string`                 |       No | Backend validation signature.  |

### Code Execution Result Delta

| Field       | Type      | Required | Description                      |
| ----------- | --------- | -------: | -------------------------------- |
| `type`      | `string`  |      Yes | Must be `code_execution_result`. |
| `result`    | `string`  |      Yes | Execution result.                |
| `is_error`  | `boolean` |       No | Indicates an execution error.    |
| `signature` | `string`  |       No | Backend validation signature.    |

### File Search Call Delta

| Field       | Type     | Required | Description                   |
| ----------- | -------- | -------: | ----------------------------- |
| `type`      | `string` |      Yes | Must be `file_search_call`.   |
| `signature` | `string` |       No | Backend validation signature. |

### File Search Result Delta

| Field       | Type                 | Required | Description                   |
| ----------- | -------------------- | -------: | ----------------------------- |
| `type`      | `string`             |      Yes | Must be `file_search_result`. |
| `result`    | `FileSearchResult[]` |      Yes | File-search results.          |
| `signature` | `string`             |       No | Backend validation signature. |

### Function Result Delta

| Field      | Type                               | Required | Description                 |
| ---------- | ---------------------------------- | -------: | --------------------------- |
| `type`     | `string`                           |      Yes | Must be `function_result`.  |
| `result`   | `Content[]`, `object`, or `string` |      Yes | Function result.            |
| `name`     | `string`                           |       No | Function name.              |
| `is_error` | `boolean`                          |       No | Indicates a function error. |

### Google Maps Call Delta

| Field       | Type                      | Required | Description                   |
| ----------- | ------------------------- | -------: | ----------------------------- |
| `type`      | `string`                  |      Yes | Must be `google_maps_call`.   |
| `arguments` | `GoogleMapsCallArguments` |       No | Google Maps query arguments.  |
| `signature` | `string`                  |       No | Backend validation signature. |

### Google Maps Result Delta

| Field       | Type                     | Required | Description                   |
| ----------- | ------------------------ | -------: | ----------------------------- |
| `type`      | `string`                 |      Yes | Must be `google_maps_result`. |
| `result`    | `GoogleMapsResultItem[]` |       No | Google Maps results.          |
| `signature` | `string`                 |       No | Backend validation signature. |

### Google Search Call Delta

| Field       | Type                        | Required | Description                   |
| ----------- | --------------------------- | -------: | ----------------------------- |
| `type`      | `string`                    |      Yes | Must be `google_search_call`. |
| `arguments` | `GoogleSearchCallArguments` |      Yes | Search-query arguments.       |
| `signature` | `string`                    |       No | Backend validation signature. |

### Google Search Result Delta

| Field       | Type                       | Required | Description                     |
| ----------- | -------------------------- | -------: | ------------------------------- |
| `type`      | `string`                   |      Yes | Must be `google_search_result`. |
| `result`    | `GoogleSearchResultItem[]` |      Yes | Google Search results.          |
| `is_error`  | `boolean`                  |       No | Indicates a search error.       |
| `signature` | `string`                   |       No | Backend validation signature.   |

### MCP Server Tool Call Delta

| Field         | Type     | Required | Description                     |
| ------------- | -------- | -------: | ------------------------------- |
| `type`        | `string` |      Yes | Must be `mcp_server_tool_call`. |
| `name`        | `string` |      Yes | MCP tool name.                  |
| `server_name` | `string` |      Yes | MCP server name.                |
| `arguments`   | `object` |      Yes | Tool arguments.                 |

### MCP Server Tool Result Delta

| Field         | Type                               | Required | Description                       |
| ------------- | ---------------------------------- | -------: | --------------------------------- |
| `type`        | `string`                           |      Yes | Must be `mcp_server_tool_result`. |
| `result`      | `Content[]`, `object`, or `string` |      Yes | MCP tool result.                  |
| `name`        | `string`                           |       No | MCP tool name.                    |
| `server_name` | `string`                           |       No | MCP server name.                  |

### URL Context Call Delta

| Field       | Type                      | Required | Description                   |
| ----------- | ------------------------- | -------: | ----------------------------- |
| `type`      | `string`                  |      Yes | Must be `url_context_call`.   |
| `arguments` | `UrlContextCallArguments` |      Yes | URL context arguments.        |
| `signature` | `string`                  |       No | Backend validation signature. |

### URL Context Result Delta

| Field       | Type                 | Required | Description                      |
| ----------- | -------------------- | -------: | -------------------------------- |
| `type`      | `string`             |      Yes | Must be `url_context_result`.    |
| `result`    | `UrlContextResult[]` |      Yes | URL retrieval results.           |
| `is_error`  | `boolean`            |       No | Indicates a URL retrieval error. |
| `signature` | `string`             |       No | Backend validation signature.    |

## Error

| Field     | Type     | Description                         |
| --------- | -------- | ----------------------------------- |
| `code`    | `string` | URI that identifies the error type. |
| `message` | `string` | Human-readable error message.       |
