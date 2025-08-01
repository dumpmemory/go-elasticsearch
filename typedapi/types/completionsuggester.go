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

// CompletionSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L164-L182
type CompletionSuggester struct {
	// Analyzer The analyzer to analyze the suggest text with.
	// Defaults to the search analyzer of the suggest field.
	Analyzer *string `json:"analyzer,omitempty"`
	// Contexts A value, geo point object, or a geo hash string to filter or boost the
	// suggestion on.
	Contexts map[string][]CompletionContext `json:"contexts,omitempty"`
	// Field The field to fetch the candidate suggestions from.
	// Needs to be set globally or per suggestion.
	Field string `json:"field"`
	// Fuzzy Enables fuzziness, meaning you can have a typo in your search and still get
	// results back.
	Fuzzy *SuggestFuzziness `json:"fuzzy,omitempty"`
	// Regex A regex query that expresses a prefix as a regular expression.
	Regex *RegexOptions `json:"regex,omitempty"`
	// Size The maximum corrections to be returned per suggest text token.
	Size *int `json:"size,omitempty"`
	// SkipDuplicates Whether duplicate suggestions should be filtered out.
	SkipDuplicates *bool `json:"skip_duplicates,omitempty"`
}

func (s *CompletionSuggester) UnmarshalJSON(data []byte) error {

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

		case "contexts":
			if s.Contexts == nil {
				s.Contexts = make(map[string][]CompletionContext, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := NewCompletionContext()
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "Contexts", err)
					}
					s.Contexts[key] = append(s.Contexts[key], *o)
				default:
					o := []CompletionContext{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "Contexts", err)
					}
					s.Contexts[key] = o
				}
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "fuzzy":
			if err := dec.Decode(&s.Fuzzy); err != nil {
				return fmt.Errorf("%s | %w", "Fuzzy", err)
			}

		case "regex":
			if err := dec.Decode(&s.Regex); err != nil {
				return fmt.Errorf("%s | %w", "Regex", err)
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "skip_duplicates":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipDuplicates", err)
				}
				s.SkipDuplicates = &value
			case bool:
				s.SkipDuplicates = &v
			}

		}
	}
	return nil
}

// NewCompletionSuggester returns a CompletionSuggester.
func NewCompletionSuggester() *CompletionSuggester {
	r := &CompletionSuggester{
		Contexts: make(map[string][]CompletionContext),
	}

	return r
}

type CompletionSuggesterVariant interface {
	CompletionSuggesterCaster() *CompletionSuggester
}

func (s *CompletionSuggester) CompletionSuggesterCaster() *CompletionSuggester {
	return s
}
