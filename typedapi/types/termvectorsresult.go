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

// TermVectorsResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/mtermvectors/types.ts#L96-L104
type TermVectorsResult struct {
	Error       *ErrorCause           `json:"error,omitempty"`
	Found       *bool                 `json:"found,omitempty"`
	Id_         *string               `json:"_id,omitempty"`
	Index_      string                `json:"_index"`
	TermVectors map[string]TermVector `json:"term_vectors,omitempty"`
	Took        *int64                `json:"took,omitempty"`
	Version_    *int64                `json:"_version,omitempty"`
}

func (s *TermVectorsResult) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "found":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Found", err)
				}
				s.Found = &value
			case bool:
				s.Found = &v
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "term_vectors":
			if s.TermVectors == nil {
				s.TermVectors = make(map[string]TermVector, 0)
			}
			if err := dec.Decode(&s.TermVectors); err != nil {
				return fmt.Errorf("%s | %w", "TermVectors", err)
			}

		case "took":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Took", err)
				}
				s.Took = &value
			case float64:
				f := int64(v)
				s.Took = &f
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		}
	}
	return nil
}

// NewTermVectorsResult returns a TermVectorsResult.
func NewTermVectorsResult() *TermVectorsResult {
	r := &TermVectorsResult{
		TermVectors: make(map[string]TermVector),
	}

	return r
}
