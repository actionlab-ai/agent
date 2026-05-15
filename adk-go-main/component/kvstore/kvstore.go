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

// Package kvstore defines a cache/session-memory port for Redis-compatible
// implementations while keeping vendor SDKs outside the ADK core.
package kvstore

import (
	"context"
	"time"
)

type Entry struct {
	Key       string            `json:"key"`
	Value     []byte            `json:"value"`
	ExpiresAt time.Time         `json:"expires_at,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type Store interface {
	Name() string
	Get(ctx context.Context, key string) (*Entry, error)
	Set(ctx context.Context, entry Entry, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Scan(ctx context.Context, prefix string, limit int) ([]string, error)
}

type AdapterFactory interface {
	NewKVStore(ctx context.Context, cfg map[string]any) (Store, error)
}
