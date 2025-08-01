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

// LinearInterpolationSmoothingModel type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L438-L442
type LinearInterpolationSmoothingModel struct {
	BigramLambda  Float64 `json:"bigram_lambda"`
	TrigramLambda Float64 `json:"trigram_lambda"`
	UnigramLambda Float64 `json:"unigram_lambda"`
}

func (s *LinearInterpolationSmoothingModel) UnmarshalJSON(data []byte) error {

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

		case "bigram_lambda":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BigramLambda", err)
				}
				f := Float64(value)
				s.BigramLambda = f
			case float64:
				f := Float64(v)
				s.BigramLambda = f
			}

		case "trigram_lambda":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TrigramLambda", err)
				}
				f := Float64(value)
				s.TrigramLambda = f
			case float64:
				f := Float64(v)
				s.TrigramLambda = f
			}

		case "unigram_lambda":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "UnigramLambda", err)
				}
				f := Float64(value)
				s.UnigramLambda = f
			case float64:
				f := Float64(v)
				s.UnigramLambda = f
			}

		}
	}
	return nil
}

// NewLinearInterpolationSmoothingModel returns a LinearInterpolationSmoothingModel.
func NewLinearInterpolationSmoothingModel() *LinearInterpolationSmoothingModel {
	r := &LinearInterpolationSmoothingModel{}

	return r
}

type LinearInterpolationSmoothingModelVariant interface {
	LinearInterpolationSmoothingModelCaster() *LinearInterpolationSmoothingModel
}

func (s *LinearInterpolationSmoothingModel) LinearInterpolationSmoothingModelCaster() *LinearInterpolationSmoothingModel {
	return s
}
