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

// DataTierPhaseStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L83-L94
type DataTierPhaseStatistics struct {
	DocCount                    int64 `json:"doc_count"`
	IndexCount                  int64 `json:"index_count"`
	NodeCount                   int64 `json:"node_count"`
	PrimaryShardCount           int64 `json:"primary_shard_count"`
	PrimaryShardSizeAvgBytes    int64 `json:"primary_shard_size_avg_bytes"`
	PrimaryShardSizeMadBytes    int64 `json:"primary_shard_size_mad_bytes"`
	PrimaryShardSizeMedianBytes int64 `json:"primary_shard_size_median_bytes"`
	PrimarySizeBytes            int64 `json:"primary_size_bytes"`
	TotalShardCount             int64 `json:"total_shard_count"`
	TotalSizeBytes              int64 `json:"total_size_bytes"`
}

func (s *DataTierPhaseStatistics) UnmarshalJSON(data []byte) error {

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

		case "doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocCount", err)
				}
				s.DocCount = value
			case float64:
				f := int64(v)
				s.DocCount = f
			}

		case "index_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexCount", err)
				}
				s.IndexCount = value
			case float64:
				f := int64(v)
				s.IndexCount = f
			}

		case "node_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NodeCount", err)
				}
				s.NodeCount = value
			case float64:
				f := int64(v)
				s.NodeCount = f
			}

		case "primary_shard_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryShardCount", err)
				}
				s.PrimaryShardCount = value
			case float64:
				f := int64(v)
				s.PrimaryShardCount = f
			}

		case "primary_shard_size_avg_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryShardSizeAvgBytes", err)
				}
				s.PrimaryShardSizeAvgBytes = value
			case float64:
				f := int64(v)
				s.PrimaryShardSizeAvgBytes = f
			}

		case "primary_shard_size_mad_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryShardSizeMadBytes", err)
				}
				s.PrimaryShardSizeMadBytes = value
			case float64:
				f := int64(v)
				s.PrimaryShardSizeMadBytes = f
			}

		case "primary_shard_size_median_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryShardSizeMedianBytes", err)
				}
				s.PrimaryShardSizeMedianBytes = value
			case float64:
				f := int64(v)
				s.PrimaryShardSizeMedianBytes = f
			}

		case "primary_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimarySizeBytes", err)
				}
				s.PrimarySizeBytes = value
			case float64:
				f := int64(v)
				s.PrimarySizeBytes = f
			}

		case "total_shard_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalShardCount", err)
				}
				s.TotalShardCount = value
			case float64:
				f := int64(v)
				s.TotalShardCount = f
			}

		case "total_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSizeBytes", err)
				}
				s.TotalSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeBytes = f
			}

		}
	}
	return nil
}

// NewDataTierPhaseStatistics returns a DataTierPhaseStatistics.
func NewDataTierPhaseStatistics() *DataTierPhaseStatistics {
	r := &DataTierPhaseStatistics{}

	return r
}
