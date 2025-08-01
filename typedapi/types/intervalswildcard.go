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

// IntervalsWildcard type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/fulltext.ts#L319-L334
type IntervalsWildcard struct {
	// Analyzer Analyzer used to analyze the `pattern`.
	// Defaults to the top-level field's analyzer.
	Analyzer *string `json:"analyzer,omitempty"`
	// Pattern Wildcard pattern used to find matching terms.
	Pattern string `json:"pattern"`
	// UseField If specified, match intervals from this field rather than the top-level
	// field.
	// The `pattern` is normalized using the search analyzer from this field, unless
	// `analyzer` is specified separately.
	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsWildcard) UnmarshalJSON(data []byte) error {

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
			s.Analyzer = &o

		case "pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pattern = o

		case "use_field":
			if err := dec.Decode(&s.UseField); err != nil {
				return fmt.Errorf("%s | %w", "UseField", err)
			}

		}
	}
	return nil
}

// NewIntervalsWildcard returns a IntervalsWildcard.
func NewIntervalsWildcard() *IntervalsWildcard {
	r := &IntervalsWildcard{}

	return r
}

type IntervalsWildcardVariant interface {
	IntervalsWildcardCaster() *IntervalsWildcard
}

func (s *IntervalsWildcard) IntervalsWildcardCaster() *IntervalsWildcard {
	return s
}
