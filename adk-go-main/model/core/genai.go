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

package core

import "google.golang.org/genai"

func FromGenAIContent(in *genai.Content) *Content {
	if in == nil {
		return nil
	}
	out := &Content{Role: normalizeRoleFromGenAI(in.Role), Parts: make([]*Part, 0, len(in.Parts))}
	for _, p := range in.Parts {
		out.Parts = append(out.Parts, FromGenAIPart(p))
	}
	return out
}

func FromGenAIPart(in *genai.Part) *Part {
	if in == nil {
		return nil
	}
	out := &Part{Text: in.Text}
	if in.FunctionCall != nil {
		out.FunctionCall = &FunctionCall{ID: in.FunctionCall.ID, Name: in.FunctionCall.Name, Args: cloneMap(in.FunctionCall.Args)}
	}
	if in.FunctionResponse != nil {
		out.FunctionResponse = &FunctionResponse{ID: in.FunctionResponse.ID, Name: in.FunctionResponse.Name, Response: cloneMap(in.FunctionResponse.Response)}
	}
	if in.InlineData != nil {
		out.InlineData = &Blob{MIMEType: in.InlineData.MIMEType, Data: append([]byte(nil), in.InlineData.Data...)}
	}
	if in.FileData != nil {
		out.FileData = &FileData{MIMEType: in.FileData.MIMEType, URI: in.FileData.FileURI}
	}
	return out
}

func ToGenAIContent(in *Content) *genai.Content {
	if in == nil {
		return nil
	}
	out := &genai.Content{Role: normalizeRoleToGenAI(in.Role), Parts: make([]*genai.Part, 0, len(in.Parts))}
	for _, p := range in.Parts {
		out.Parts = append(out.Parts, ToGenAIPart(p))
	}
	return out
}

func ToGenAIPart(in *Part) *genai.Part {
	if in == nil {
		return nil
	}
	out := &genai.Part{Text: in.Text}
	if in.FunctionCall != nil {
		out.FunctionCall = &genai.FunctionCall{ID: in.FunctionCall.ID, Name: in.FunctionCall.Name, Args: cloneMap(in.FunctionCall.Args)}
	}
	if in.FunctionResponse != nil {
		out.FunctionResponse = &genai.FunctionResponse{ID: in.FunctionResponse.ID, Name: in.FunctionResponse.Name, Response: cloneMap(in.FunctionResponse.Response)}
	}
	if in.InlineData != nil {
		out.InlineData = &genai.Blob{MIMEType: in.InlineData.MIMEType, Data: append([]byte(nil), in.InlineData.Data...)}
	}
	if in.FileData != nil {
		out.FileData = &genai.FileData{MIMEType: in.FileData.MIMEType, FileURI: in.FileData.URI}
	}
	return out
}

func FromGenAIConfig(in *genai.GenerateContentConfig) *GenerateContentConfig {
	if in == nil {
		return nil
	}
	out := &GenerateContentConfig{
		SystemInstruction: FromGenAIContent(in.SystemInstruction),
		Temperature:       in.Temperature,
		TopP:              in.TopP,
		TopK:              in.TopK,
		MaxOutputTokens:   in.MaxOutputTokens,
		StopSequences:     append([]string(nil), in.StopSequences...),
		ResponseSchema:    in.ResponseJsonSchema,
	}
	for _, t := range in.Tools {
		if t == nil {
			continue
		}
		tool := &Tool{}
		for _, decl := range t.FunctionDeclarations {
			if decl == nil {
				continue
			}
			tool.FunctionDeclarations = append(tool.FunctionDeclarations, &FunctionDeclaration{Name: decl.Name, Description: decl.Description, Parameters: decl.ParametersJsonSchema})
		}
		out.Tools = append(out.Tools, tool)
	}
	return out
}

func FromADKLLMRequest(modelName string, contents []*genai.Content, cfg *genai.GenerateContentConfig, tools map[string]any) *LLMRequest {
	out := &LLMRequest{Model: modelName, Config: FromGenAIConfig(cfg), Tools: tools}
	for _, c := range contents {
		out.Contents = append(out.Contents, FromGenAIContent(c))
	}
	return out
}

func ToADKLLMResponse(in *LLMResponse) *genai.Content {
	if in == nil {
		return nil
	}
	return ToGenAIContent(in.Content)
}

func normalizeRoleFromGenAI(role string) string {
	switch role {
	case genai.RoleModel:
		return RoleModel
	case genai.RoleUser:
		return RoleUser
	default:
		return role
	}
}

func normalizeRoleToGenAI(role string) string {
	switch role {
	case RoleAssistant, RoleModel:
		return genai.RoleModel
	case RoleSystem, RoleUser, RoleTool:
		return genai.RoleUser
	default:
		return role
	}
}

func cloneMap(in map[string]any) map[string]any {
	if in == nil {
		return nil
	}
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
