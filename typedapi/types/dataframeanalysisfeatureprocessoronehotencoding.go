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

// DataframeAnalysisFeatureProcessorOneHotEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/DataframeAnalytics.ts#L288-L293
type DataframeAnalysisFeatureProcessorOneHotEncoding struct {
	// Field The name of the field to encode.
	Field string `json:"field"`
	// HotMap The one hot map mapping the field value with the column name.
	HotMap string `json:"hot_map"`
}

func (s *DataframeAnalysisFeatureProcessorOneHotEncoding) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "hot_map":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "HotMap", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HotMap = o

		}
	}
	return nil
}

// NewDataframeAnalysisFeatureProcessorOneHotEncoding returns a DataframeAnalysisFeatureProcessorOneHotEncoding.
func NewDataframeAnalysisFeatureProcessorOneHotEncoding() *DataframeAnalysisFeatureProcessorOneHotEncoding {
	r := &DataframeAnalysisFeatureProcessorOneHotEncoding{}

	return r
}

type DataframeAnalysisFeatureProcessorOneHotEncodingVariant interface {
	DataframeAnalysisFeatureProcessorOneHotEncodingCaster() *DataframeAnalysisFeatureProcessorOneHotEncoding
}

func (s *DataframeAnalysisFeatureProcessorOneHotEncoding) DataframeAnalysisFeatureProcessorOneHotEncodingCaster() *DataframeAnalysisFeatureProcessorOneHotEncoding {
	return s
}
