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

// Package vectorstore defines a vendor-neutral vector retrieval port.
//
// Milvus, Qdrant, pgvector, Redis Vector Set, and in-memory implementations
// should live behind this interface so agents and tools never depend on a
// product SDK directly.
package vectorstore

import "context"

type Document struct {
	ID        string         `json:"id"`
	Text      string         `json:"text,omitempty"`
	Vector    []float32      `json:"vector,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Namespace string         `json:"namespace,omitempty"`
}

type SearchRequest struct {
	Namespace string         `json:"namespace,omitempty"`
	Vector    []float32      `json:"vector,omitempty"`
	Text      string         `json:"text,omitempty"`
	TopK      int            `json:"top_k,omitempty"`
	Filter    map[string]any `json:"filter,omitempty"`
}

type SearchResult struct {
	Document Document `json:"document"`
	Score    float64  `json:"score"`
}

type Store interface {
	Name() string
	Upsert(ctx context.Context, docs ...Document) error
	Search(ctx context.Context, req SearchRequest) ([]SearchResult, error)
	Delete(ctx context.Context, namespace string, ids ...string) error
}

type AdapterFactory interface {
	NewVectorStore(ctx context.Context, cfg map[string]any) (Store, error)
}
