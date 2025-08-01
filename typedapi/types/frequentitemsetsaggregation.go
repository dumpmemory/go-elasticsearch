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

// FrequentItemSetsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L1241-L1268
type FrequentItemSetsAggregation struct {
	// Fields Fields to analyze.
	Fields []FrequentItemSetsField `json:"fields"`
	// Filter Query that filters documents from analysis.
	Filter *Query `json:"filter,omitempty"`
	// MinimumSetSize The minimum size of one item set.
	MinimumSetSize *int `json:"minimum_set_size,omitempty"`
	// MinimumSupport The minimum support of one item set.
	MinimumSupport *Float64 `json:"minimum_support,omitempty"`
	// Size The number of top item sets to return.
	Size *int `json:"size,omitempty"`
}

func (s *FrequentItemSetsAggregation) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "minimum_set_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinimumSetSize", err)
				}
				s.MinimumSetSize = &value
			case float64:
				f := int(v)
				s.MinimumSetSize = &f
			}

		case "minimum_support":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinimumSupport", err)
				}
				f := Float64(value)
				s.MinimumSupport = &f
			case float64:
				f := Float64(v)
				s.MinimumSupport = &f
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

		}
	}
	return nil
}

// NewFrequentItemSetsAggregation returns a FrequentItemSetsAggregation.
func NewFrequentItemSetsAggregation() *FrequentItemSetsAggregation {
	r := &FrequentItemSetsAggregation{}

	return r
}

type FrequentItemSetsAggregationVariant interface {
	FrequentItemSetsAggregationCaster() *FrequentItemSetsAggregation
}

func (s *FrequentItemSetsAggregation) FrequentItemSetsAggregationCaster() *FrequentItemSetsAggregation {
	return s
}
