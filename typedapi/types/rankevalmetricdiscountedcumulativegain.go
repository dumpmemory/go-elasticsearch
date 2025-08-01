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

// RankEvalMetricDiscountedCumulativeGain type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/rank_eval/types.ts#L66-L77
type RankEvalMetricDiscountedCumulativeGain struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// Normalize If set to true, this metric will calculate the Normalized DCG.
	Normalize *bool `json:"normalize,omitempty"`
}

func (s *RankEvalMetricDiscountedCumulativeGain) UnmarshalJSON(data []byte) error {

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

		case "k":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "K", err)
				}
				s.K = &value
			case float64:
				f := int(v)
				s.K = &f
			}

		case "normalize":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Normalize", err)
				}
				s.Normalize = &value
			case bool:
				s.Normalize = &v
			}

		}
	}
	return nil
}

// NewRankEvalMetricDiscountedCumulativeGain returns a RankEvalMetricDiscountedCumulativeGain.
func NewRankEvalMetricDiscountedCumulativeGain() *RankEvalMetricDiscountedCumulativeGain {
	r := &RankEvalMetricDiscountedCumulativeGain{}

	return r
}

type RankEvalMetricDiscountedCumulativeGainVariant interface {
	RankEvalMetricDiscountedCumulativeGainCaster() *RankEvalMetricDiscountedCumulativeGain
}

func (s *RankEvalMetricDiscountedCumulativeGain) RankEvalMetricDiscountedCumulativeGainCaster() *RankEvalMetricDiscountedCumulativeGain {
	return s
}
