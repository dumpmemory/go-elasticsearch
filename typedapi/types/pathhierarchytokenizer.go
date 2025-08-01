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

// PathHierarchyTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/tokenizers.ts#L97-L104
type PathHierarchyTokenizer struct {
	BufferSize  Stringifiedinteger `json:"buffer_size,omitempty"`
	Delimiter   *string            `json:"delimiter,omitempty"`
	Replacement *string            `json:"replacement,omitempty"`
	Reverse     Stringifiedboolean `json:"reverse,omitempty"`
	Skip        Stringifiedinteger `json:"skip,omitempty"`
	Type        string             `json:"type,omitempty"`
	Version     *string            `json:"version,omitempty"`
}

func (s *PathHierarchyTokenizer) UnmarshalJSON(data []byte) error {

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

		case "buffer_size":
			if err := dec.Decode(&s.BufferSize); err != nil {
				return fmt.Errorf("%s | %w", "BufferSize", err)
			}

		case "delimiter":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Delimiter", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Delimiter = &o

		case "replacement":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Replacement", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Replacement = &o

		case "reverse":
			if err := dec.Decode(&s.Reverse); err != nil {
				return fmt.Errorf("%s | %w", "Reverse", err)
			}

		case "skip":
			if err := dec.Decode(&s.Skip); err != nil {
				return fmt.Errorf("%s | %w", "Skip", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s PathHierarchyTokenizer) MarshalJSON() ([]byte, error) {
	type innerPathHierarchyTokenizer PathHierarchyTokenizer
	tmp := innerPathHierarchyTokenizer{
		BufferSize:  s.BufferSize,
		Delimiter:   s.Delimiter,
		Replacement: s.Replacement,
		Reverse:     s.Reverse,
		Skip:        s.Skip,
		Type:        s.Type,
		Version:     s.Version,
	}

	tmp.Type = "path_hierarchy"

	return json.Marshal(tmp)
}

// NewPathHierarchyTokenizer returns a PathHierarchyTokenizer.
func NewPathHierarchyTokenizer() *PathHierarchyTokenizer {
	r := &PathHierarchyTokenizer{}

	return r
}

type PathHierarchyTokenizerVariant interface {
	PathHierarchyTokenizerCaster() *PathHierarchyTokenizer
}

func (s *PathHierarchyTokenizer) PathHierarchyTokenizerCaster() *PathHierarchyTokenizer {
	return s
}

func (s *PathHierarchyTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	o := TokenizerDefinition(s)
	return &o
}
