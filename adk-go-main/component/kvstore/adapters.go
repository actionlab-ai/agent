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

package kvstore

import (
	"context"
	"errors"
	"fmt"
)

var ErrAdapterNotLinked = errors.New("kvstore adapter implementation is not linked")

type RedisConfig struct {
	Address  string `json:"address"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DB       int    `json:"db,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

// NewRedisStore is an integration seam. Product builds can replace this with a
// real adapter package without leaking Redis SDK types into session/memory code.
func NewRedisStore(ctx context.Context, cfg RedisConfig) (Store, error) {
	_ = ctx
	return nil, fmt.Errorf("%w: redis address=%q db=%d", ErrAdapterNotLinked, cfg.Address, cfg.DB)
}
