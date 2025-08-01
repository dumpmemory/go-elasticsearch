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

// AllField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/mapping/meta-fields.ts#L29-L40
type AllField struct {
	Analyzer                 string `json:"analyzer"`
	Enabled                  bool   `json:"enabled"`
	OmitNorms                bool   `json:"omit_norms"`
	SearchAnalyzer           string `json:"search_analyzer"`
	Similarity               string `json:"similarity"`
	Store                    bool   `json:"store"`
	StoreTermVectorOffsets   bool   `json:"store_term_vector_offsets"`
	StoreTermVectorPayloads  bool   `json:"store_term_vector_payloads"`
	StoreTermVectorPositions bool   `json:"store_term_vector_positions"`
	StoreTermVectors         bool   `json:"store_term_vectors"`
}

func (s *AllField) UnmarshalJSON(data []byte) error {

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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Analyzer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = o

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "omit_norms":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OmitNorms", err)
				}
				s.OmitNorms = value
			case bool:
				s.OmitNorms = v
			}

		case "search_analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SearchAnalyzer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchAnalyzer = o

		case "similarity":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Similarity", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Similarity = o

		case "store":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Store", err)
				}
				s.Store = value
			case bool:
				s.Store = v
			}

		case "store_term_vector_offsets":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StoreTermVectorOffsets", err)
				}
				s.StoreTermVectorOffsets = value
			case bool:
				s.StoreTermVectorOffsets = v
			}

		case "store_term_vector_payloads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StoreTermVectorPayloads", err)
				}
				s.StoreTermVectorPayloads = value
			case bool:
				s.StoreTermVectorPayloads = v
			}

		case "store_term_vector_positions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StoreTermVectorPositions", err)
				}
				s.StoreTermVectorPositions = value
			case bool:
				s.StoreTermVectorPositions = v
			}

		case "store_term_vectors":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StoreTermVectors", err)
				}
				s.StoreTermVectors = value
			case bool:
				s.StoreTermVectors = v
			}

		}
	}
	return nil
}

// NewAllField returns a AllField.
func NewAllField() *AllField {
	r := &AllField{}

	return r
}

type AllFieldVariant interface {
	AllFieldCaster() *AllField
}

func (s *AllField) AllFieldCaster() *AllField {
	return s
}
