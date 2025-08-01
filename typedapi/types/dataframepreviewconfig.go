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

// DataframePreviewConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/preview_data_frame_analytics/types.ts#L27-L33
type DataframePreviewConfig struct {
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
}

func (s *DataframePreviewConfig) UnmarshalJSON(data []byte) error {

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

		case "analysis":
			if err := dec.Decode(&s.Analysis); err != nil {
				return fmt.Errorf("%s | %w", "Analysis", err)
			}

		case "analyzed_fields":
			if err := dec.Decode(&s.AnalyzedFields); err != nil {
				return fmt.Errorf("%s | %w", "AnalyzedFields", err)
			}

		case "max_num_threads":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNumThreads", err)
				}
				s.MaxNumThreads = &value
			case float64:
				f := int(v)
				s.MaxNumThreads = &f
			}

		case "model_memory_limit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelMemoryLimit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		}
	}
	return nil
}

// NewDataframePreviewConfig returns a DataframePreviewConfig.
func NewDataframePreviewConfig() *DataframePreviewConfig {
	r := &DataframePreviewConfig{}

	return r
}

type DataframePreviewConfigVariant interface {
	DataframePreviewConfigCaster() *DataframePreviewConfig
}

func (s *DataframePreviewConfig) DataframePreviewConfigCaster() *DataframePreviewConfig {
	return s
}
