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

import (
	"testing"

	"google.golang.org/genai"
)

func TestGenAIContentRoundTrip(t *testing.T) {
	in := &genai.Content{Role: genai.RoleModel, Parts: []*genai.Part{
		genai.NewPartFromText("hello"),
		{FunctionCall: &genai.FunctionCall{ID: "call_1", Name: "lookup", Args: map[string]any{"id": "42"}}},
	}}

	coreContent := FromGenAIContent(in)
	if coreContent.Role != RoleModel || len(coreContent.Parts) != 2 {
		t.Fatalf("core content mismatch: %#v", coreContent)
	}
	if coreContent.Parts[1].FunctionCall.Name != "lookup" || coreContent.Parts[1].FunctionCall.Args["id"] != "42" {
		t.Fatalf("function call mismatch: %#v", coreContent.Parts[1].FunctionCall)
	}

	out := ToGenAIContent(coreContent)
	if out.Role != genai.RoleModel || len(out.Parts) != 2 || out.Parts[0].Text != "hello" {
		t.Fatalf("round trip content mismatch: %#v", out)
	}
	if out.Parts[1].FunctionCall == nil || out.Parts[1].FunctionCall.Name != "lookup" {
		t.Fatalf("round trip function call mismatch: %#v", out.Parts[1])
	}
}

func TestFromGenAIConfigCopiesTools(t *testing.T) {
	cfg := &genai.GenerateContentConfig{Tools: []*genai.Tool{{FunctionDeclarations: []*genai.FunctionDeclaration{{
		Name:        "search",
		Description: "search docs",
		ParametersJsonSchema: map[string]any{
			"type": "object",
		},
	}}}}}

	got := FromGenAIConfig(cfg)
	if got == nil || len(got.Tools) != 1 || len(got.Tools[0].FunctionDeclarations) != 1 {
		t.Fatalf("tool config missing: %#v", got)
	}
	decl := got.Tools[0].FunctionDeclarations[0]
	if decl.Name != "search" || decl.Description != "search docs" {
		t.Fatalf("declaration mismatch: %#v", decl)
	}
}
