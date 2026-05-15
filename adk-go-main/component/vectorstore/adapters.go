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

package vectorstore

import (
	"context"
	"errors"
	"fmt"
)

var ErrAdapterNotLinked = errors.New("vectorstore adapter implementation is not linked")

// MilvusConfig captures the stable options expected from a Milvus adapter.
type MilvusConfig struct {
	Address    string `json:"address"`
	Database   string `json:"database,omitempty"`
	Collection string `json:"collection"`
	Dimension  int    `json:"dimension"`
	MetricType string `json:"metric_type,omitempty"`
}

// NewMilvusStore is an integration seam. Product builds can replace this with a
// real adapter package without leaking Milvus SDK types into agent/tool code.
func NewMilvusStore(ctx context.Context, cfg MilvusConfig) (Store, error) {
	_ = ctx
	return nil, fmt.Errorf("%w: milvus address=%q collection=%q", ErrAdapterNotLinked, cfg.Address, cfg.Collection)
}
