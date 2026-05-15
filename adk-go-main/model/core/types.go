// Copyright 2026 Archinfra / yuanyp8
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package core defines Archinfra ADK's provider-neutral business model.
//
// These DTOs are intentionally small, JSON-stable, and independent from any
// vendor SDK. Runtime code may still convert to/from legacy SDK structs at
// package boundaries while new front-end, REST, and model adapters can exchange
// this package directly.
package core

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleModel     = "model"
	RoleTool      = "tool"
)

const (
	FinishReasonStop      = "stop"
	FinishReasonToolCalls = "tool_calls"
	FinishReasonLength    = "length"
	FinishReasonError     = "error"
)

// Content is a provider-neutral chat turn.
type Content struct {
	Role  string  `json:"role,omitempty"`
	Parts []*Part `json:"parts,omitempty"`
}

// Part is one item in a Content message. Exactly one payload should be set.
type Part struct {
	Text             string            `json:"text,omitempty"`
	FunctionCall     *FunctionCall     `json:"function_call,omitempty"`
	FunctionResponse *FunctionResponse `json:"function_response,omitempty"`
	InlineData       *Blob             `json:"inline_data,omitempty"`
	FileData         *FileData         `json:"file_data,omitempty"`
	Metadata         map[string]any    `json:"metadata,omitempty"`
}

type FunctionCall struct {
	ID   string         `json:"id,omitempty"`
	Name string         `json:"name,omitempty"`
	Args map[string]any `json:"args,omitempty"`
}

type FunctionResponse struct {
	ID       string         `json:"id,omitempty"`
	Name     string         `json:"name,omitempty"`
	Response map[string]any `json:"response,omitempty"`
}

type Blob struct {
	MIMEType string `json:"mime_type,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

type FileData struct {
	MIMEType string `json:"mime_type,omitempty"`
	URI      string `json:"uri,omitempty"`
}

// GenerateContentConfig contains model-generation options used by all adapters.
type GenerateContentConfig struct {
	SystemInstruction *Content       `json:"system_instruction,omitempty"`
	Temperature       *float32       `json:"temperature,omitempty"`
	TopP              *float32       `json:"top_p,omitempty"`
	TopK              *float32       `json:"top_k,omitempty"`
	MaxOutputTokens   int32          `json:"max_output_tokens,omitempty"`
	StopSequences     []string       `json:"stop_sequences,omitempty"`
	Tools             []*Tool        `json:"tools,omitempty"`
	ResponseSchema    any            `json:"response_schema,omitempty"`
	Extra             map[string]any `json:"extra,omitempty"`
}

type Tool struct {
	FunctionDeclarations []*FunctionDeclaration `json:"function_declarations,omitempty"`
	Extensions           map[string]any         `json:"extensions,omitempty"`
}

type FunctionDeclaration struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters,omitempty"`
}

// LLMRequest is the canonical request shape for provider adapters.
type LLMRequest struct {
	Model    string                 `json:"model,omitempty"`
	Contents []*Content             `json:"contents,omitempty"`
	Config   *GenerateContentConfig `json:"config,omitempty"`
	Tools    map[string]any         `json:"-"`
}

// LLMResponse is the canonical response shape for provider adapters and APIs.
type LLMResponse struct {
	Content        *Content       `json:"content,omitempty"`
	Usage          *UsageMetadata `json:"usage,omitempty"`
	CustomMetadata map[string]any `json:"custom_metadata,omitempty"`
	ModelVersion   string         `json:"model_version,omitempty"`
	Partial        bool           `json:"partial,omitempty"`
	TurnComplete   bool           `json:"turn_complete,omitempty"`
	Interrupted    bool           `json:"interrupted,omitempty"`
	ErrorCode      string         `json:"error_code,omitempty"`
	ErrorMessage   string         `json:"error_message,omitempty"`
	FinishReason   string         `json:"finish_reason,omitempty"`
	AvgLogprobs    float64        `json:"avg_logprobs,omitempty"`
}

type UsageMetadata struct {
	PromptTokenCount     int32 `json:"prompt_token_count,omitempty"`
	CandidatesTokenCount int32 `json:"candidates_token_count,omitempty"`
	TotalTokenCount      int32 `json:"total_token_count,omitempty"`
}

func NewTextContent(text, role string) *Content {
	return &Content{Role: role, Parts: []*Part{{Text: text}}}
}
