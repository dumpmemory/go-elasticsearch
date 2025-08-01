// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AliasDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/AliasDefinition.ts#L22-L54
type AliasDefinition struct {
	// Filter Query used to limit documents the alias can access.
	Filter *Query `json:"filter,omitempty"`
	// IndexRouting Value used to route indexing operations to a specific shard.
	// If specified, this overwrites the `routing` value for indexing operations.
	IndexRouting *string `json:"index_routing,omitempty"`
	// IsHidden If `true`, the alias is hidden.
	// All indices for the alias must have the same `is_hidden` value.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// IsWriteIndex If `true`, the index is the write index for the alias.
	IsWriteIndex *bool `json:"is_write_index,omitempty"`
	// Routing Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty"`
	// SearchRouting Value used to route search operations to a specific shard.
	// If specified, this overwrites the `routing` value for search operations.
	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *AliasDefinition) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "index_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndexRouting", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexRouting = &o

		case "is_hidden":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsHidden", err)
				}
				s.IsHidden = &value
			case bool:
				s.IsHidden = &v
			}

		case "is_write_index":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsWriteIndex", err)
				}
				s.IsWriteIndex = &value
			case bool:
				s.IsWriteIndex = &v
			}

		case "routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Routing = &o

		case "search_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SearchRouting", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchRouting = &o

		}
	}
	return nil
}

// NewAliasDefinition returns a AliasDefinition.
func NewAliasDefinition() *AliasDefinition {
	r := &AliasDefinition{}

	return r
}

type AliasDefinitionVariant interface {
	AliasDefinitionCaster() *AliasDefinition
}

func (s *AliasDefinition) AliasDefinitionCaster() *AliasDefinition {
	return s
}
