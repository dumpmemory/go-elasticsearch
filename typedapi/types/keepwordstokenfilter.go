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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// KeepWordsTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/_types/analysis/token_filters.ts#L225-L230
type KeepWordsTokenFilter struct {
	KeepWords     []string `json:"keep_words,omitempty"`
	KeepWordsCase *bool    `json:"keep_words_case,omitempty"`
	KeepWordsPath *string  `json:"keep_words_path,omitempty"`
	Type          string   `json:"type,omitempty"`
	Version       *string  `json:"version,omitempty"`
}

func (s *KeepWordsTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "keep_words":
			if err := dec.Decode(&s.KeepWords); err != nil {
				return fmt.Errorf("%s | %w", "KeepWords", err)
			}

		case "keep_words_case":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "KeepWordsCase", err)
				}
				s.KeepWordsCase = &value
			case bool:
				s.KeepWordsCase = &v
			}

		case "keep_words_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KeepWordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeepWordsPath = &o

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
func (s KeepWordsTokenFilter) MarshalJSON() ([]byte, error) {
	type innerKeepWordsTokenFilter KeepWordsTokenFilter
	tmp := innerKeepWordsTokenFilter{
		KeepWords:     s.KeepWords,
		KeepWordsCase: s.KeepWordsCase,
		KeepWordsPath: s.KeepWordsPath,
		Type:          s.Type,
		Version:       s.Version,
	}

	tmp.Type = "keep"

	return json.Marshal(tmp)
}

// NewKeepWordsTokenFilter returns a KeepWordsTokenFilter.
func NewKeepWordsTokenFilter() *KeepWordsTokenFilter {
	r := &KeepWordsTokenFilter{}

	return r
}

// true

type KeepWordsTokenFilterVariant interface {
	KeepWordsTokenFilterCaster() *KeepWordsTokenFilter
}

func (s *KeepWordsTokenFilter) KeepWordsTokenFilterCaster() *KeepWordsTokenFilter {
	return s
}
