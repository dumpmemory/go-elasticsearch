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

// TextEmbeddingInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/inference.ts#L252-L262
type TextEmbeddingInferenceOptions struct {
	// EmbeddingSize The number of dimensions in the embedding output
	EmbeddingSize *int `json:"embedding_size,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   Vocabulary                   `json:"vocabulary"`
}

func (s *TextEmbeddingInferenceOptions) UnmarshalJSON(data []byte) error {

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

		case "embedding_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "EmbeddingSize", err)
				}
				s.EmbeddingSize = &value
			case float64:
				f := int(v)
				s.EmbeddingSize = &f
			}

		case "results_field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResultsField", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultsField = &o

		case "tokenization":
			if err := dec.Decode(&s.Tokenization); err != nil {
				return fmt.Errorf("%s | %w", "Tokenization", err)
			}

		case "vocabulary":
			if err := dec.Decode(&s.Vocabulary); err != nil {
				return fmt.Errorf("%s | %w", "Vocabulary", err)
			}

		}
	}
	return nil
}

// NewTextEmbeddingInferenceOptions returns a TextEmbeddingInferenceOptions.
func NewTextEmbeddingInferenceOptions() *TextEmbeddingInferenceOptions {
	r := &TextEmbeddingInferenceOptions{}

	return r
}

type TextEmbeddingInferenceOptionsVariant interface {
	TextEmbeddingInferenceOptionsCaster() *TextEmbeddingInferenceOptions
}

func (s *TextEmbeddingInferenceOptions) TextEmbeddingInferenceOptionsCaster() *TextEmbeddingInferenceOptions {
	return s
}
