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

// RankEvalMetricExpectedReciprocalRank type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/_global/rank_eval/types.ts#L79-L88
type RankEvalMetricExpectedReciprocalRank struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// MaximumRelevance The highest relevance grade used in the user-supplied relevance judgments.
	MaximumRelevance int `json:"maximum_relevance"`
}

func (s *RankEvalMetricExpectedReciprocalRank) UnmarshalJSON(data []byte) error {

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

		case "maximum_relevance":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaximumRelevance", err)
				}
				s.MaximumRelevance = value
			case float64:
				f := int(v)
				s.MaximumRelevance = f
			}

		}
	}
	return nil
}

// NewRankEvalMetricExpectedReciprocalRank returns a RankEvalMetricExpectedReciprocalRank.
func NewRankEvalMetricExpectedReciprocalRank() *RankEvalMetricExpectedReciprocalRank {
	r := &RankEvalMetricExpectedReciprocalRank{}

	return r
}

// true

type RankEvalMetricExpectedReciprocalRankVariant interface {
	RankEvalMetricExpectedReciprocalRankCaster() *RankEvalMetricExpectedReciprocalRank
}

func (s *RankEvalMetricExpectedReciprocalRank) RankEvalMetricExpectedReciprocalRankCaster() *RankEvalMetricExpectedReciprocalRank {
	return s
}
