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

// PercentileRanksAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/metric.ts#L192-L214
type PercentileRanksAggregation struct {
	// Field The field on which to run the aggregation.
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`
	// Hdr Uses the alternative High Dynamic Range Histogram algorithm to calculate
	// percentile ranks.
	Hdr *HdrMethod `json:"hdr,omitempty"`
	// Keyed By default, the aggregation associates a unique string key with each bucket
	// and returns the ranges as a hash rather than an array.
	// Set to `false` to disable this behavior.
	Keyed *bool `json:"keyed,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
	// Tdigest Sets parameters for the default TDigest algorithm used to calculate
	// percentile ranks.
	Tdigest *TDigest `json:"tdigest,omitempty"`
	// Values An array of values for which to calculate the percentile ranks.
	Values *[]Float64 `json:"values,omitempty"`
}

func (s *PercentileRanksAggregation) UnmarshalJSON(data []byte) error {

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

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "hdr":
			if err := dec.Decode(&s.Hdr); err != nil {
				return fmt.Errorf("%s | %w", "Hdr", err)
			}

		case "keyed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Keyed", err)
				}
				s.Keyed = &value
			case bool:
				s.Keyed = &v
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "tdigest":
			if err := dec.Decode(&s.Tdigest); err != nil {
				return fmt.Errorf("%s | %w", "Tdigest", err)
			}

		case "values":
			if err := dec.Decode(&s.Values); err != nil {
				return fmt.Errorf("%s | %w", "Values", err)
			}

		}
	}
	return nil
}

// NewPercentileRanksAggregation returns a PercentileRanksAggregation.
func NewPercentileRanksAggregation() *PercentileRanksAggregation {
	r := &PercentileRanksAggregation{}

	return r
}

type PercentileRanksAggregationVariant interface {
	PercentileRanksAggregationCaster() *PercentileRanksAggregation
}

func (s *PercentileRanksAggregation) PercentileRanksAggregationCaster() *PercentileRanksAggregation {
	return s
}
