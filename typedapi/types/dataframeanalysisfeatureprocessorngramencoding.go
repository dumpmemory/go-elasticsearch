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

// DataframeAnalysisFeatureProcessorNGramEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/DataframeAnalytics.ts#L274-L286
type DataframeAnalysisFeatureProcessorNGramEncoding struct {
	Custom *bool `json:"custom,omitempty"`
	// FeaturePrefix The feature name prefix. Defaults to ngram_<start>_<length>.
	FeaturePrefix *string `json:"feature_prefix,omitempty"`
	// Field The name of the text field to encode.
	Field string `json:"field"`
	// Length Specifies the length of the n-gram substring. Defaults to 50. Must be greater
	// than 0.
	Length *int `json:"length,omitempty"`
	// NGrams Specifies which n-grams to gather. It’s an array of integer values where the
	// minimum value is 1, and a maximum value is 5.
	NGrams []int `json:"n_grams"`
	// Start Specifies the zero-indexed start of the n-gram substring. Negative values are
	// allowed for encoding n-grams of string suffixes. Defaults to 0.
	Start *int `json:"start,omitempty"`
}

func (s *DataframeAnalysisFeatureProcessorNGramEncoding) UnmarshalJSON(data []byte) error {

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

		case "custom":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Custom", err)
				}
				s.Custom = &value
			case bool:
				s.Custom = &v
			}

		case "feature_prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FeaturePrefix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FeaturePrefix = &o

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Length", err)
				}
				s.Length = &value
			case float64:
				f := int(v)
				s.Length = &f
			}

		case "n_grams":
			if err := dec.Decode(&s.NGrams); err != nil {
				return fmt.Errorf("%s | %w", "NGrams", err)
			}

		case "start":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Start", err)
				}
				s.Start = &value
			case float64:
				f := int(v)
				s.Start = &f
			}

		}
	}
	return nil
}

// NewDataframeAnalysisFeatureProcessorNGramEncoding returns a DataframeAnalysisFeatureProcessorNGramEncoding.
func NewDataframeAnalysisFeatureProcessorNGramEncoding() *DataframeAnalysisFeatureProcessorNGramEncoding {
	r := &DataframeAnalysisFeatureProcessorNGramEncoding{}

	return r
}

type DataframeAnalysisFeatureProcessorNGramEncodingVariant interface {
	DataframeAnalysisFeatureProcessorNGramEncodingCaster() *DataframeAnalysisFeatureProcessorNGramEncoding
}

func (s *DataframeAnalysisFeatureProcessorNGramEncoding) DataframeAnalysisFeatureProcessorNGramEncodingCaster() *DataframeAnalysisFeatureProcessorNGramEncoding {
	return s
}
