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

// DataframeClassificationSummaryRecall type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/evaluate_data_frame/types.ts#L106-L109
type DataframeClassificationSummaryRecall struct {
	AvgRecall Float64                    `json:"avg_recall"`
	Classes   []DataframeEvaluationClass `json:"classes"`
}

func (s *DataframeClassificationSummaryRecall) UnmarshalJSON(data []byte) error {

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

		case "avg_recall":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AvgRecall", err)
				}
				f := Float64(value)
				s.AvgRecall = f
			case float64:
				f := Float64(v)
				s.AvgRecall = f
			}

		case "classes":
			if err := dec.Decode(&s.Classes); err != nil {
				return fmt.Errorf("%s | %w", "Classes", err)
			}

		}
	}
	return nil
}

// NewDataframeClassificationSummaryRecall returns a DataframeClassificationSummaryRecall.
func NewDataframeClassificationSummaryRecall() *DataframeClassificationSummaryRecall {
	r := &DataframeClassificationSummaryRecall{}

	return r
}
