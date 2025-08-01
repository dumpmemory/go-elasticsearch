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

// AggregationBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/profile.ts#L26-L39
type AggregationBreakdown struct {
	BuildAggregation        int64  `json:"build_aggregation"`
	BuildAggregationCount   int64  `json:"build_aggregation_count"`
	BuildLeafCollector      int64  `json:"build_leaf_collector"`
	BuildLeafCollectorCount int64  `json:"build_leaf_collector_count"`
	Collect                 int64  `json:"collect"`
	CollectCount            int64  `json:"collect_count"`
	Initialize              int64  `json:"initialize"`
	InitializeCount         int64  `json:"initialize_count"`
	PostCollection          *int64 `json:"post_collection,omitempty"`
	PostCollectionCount     *int64 `json:"post_collection_count,omitempty"`
	Reduce                  int64  `json:"reduce"`
	ReduceCount             int64  `json:"reduce_count"`
}

func (s *AggregationBreakdown) UnmarshalJSON(data []byte) error {

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

		case "build_aggregation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildAggregation", err)
				}
				s.BuildAggregation = value
			case float64:
				f := int64(v)
				s.BuildAggregation = f
			}

		case "build_aggregation_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildAggregationCount", err)
				}
				s.BuildAggregationCount = value
			case float64:
				f := int64(v)
				s.BuildAggregationCount = f
			}

		case "build_leaf_collector":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildLeafCollector", err)
				}
				s.BuildLeafCollector = value
			case float64:
				f := int64(v)
				s.BuildLeafCollector = f
			}

		case "build_leaf_collector_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildLeafCollectorCount", err)
				}
				s.BuildLeafCollectorCount = value
			case float64:
				f := int64(v)
				s.BuildLeafCollectorCount = f
			}

		case "collect":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Collect", err)
				}
				s.Collect = value
			case float64:
				f := int64(v)
				s.Collect = f
			}

		case "collect_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CollectCount", err)
				}
				s.CollectCount = value
			case float64:
				f := int64(v)
				s.CollectCount = f
			}

		case "initialize":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Initialize", err)
				}
				s.Initialize = value
			case float64:
				f := int64(v)
				s.Initialize = f
			}

		case "initialize_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "InitializeCount", err)
				}
				s.InitializeCount = value
			case float64:
				f := int64(v)
				s.InitializeCount = f
			}

		case "post_collection":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PostCollection", err)
				}
				s.PostCollection = &value
			case float64:
				f := int64(v)
				s.PostCollection = &f
			}

		case "post_collection_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PostCollectionCount", err)
				}
				s.PostCollectionCount = &value
			case float64:
				f := int64(v)
				s.PostCollectionCount = &f
			}

		case "reduce":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Reduce", err)
				}
				s.Reduce = value
			case float64:
				f := int64(v)
				s.Reduce = f
			}

		case "reduce_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReduceCount", err)
				}
				s.ReduceCount = value
			case float64:
				f := int64(v)
				s.ReduceCount = f
			}

		}
	}
	return nil
}

// NewAggregationBreakdown returns a AggregationBreakdown.
func NewAggregationBreakdown() *AggregationBreakdown {
	r := &AggregationBreakdown{}

	return r
}
