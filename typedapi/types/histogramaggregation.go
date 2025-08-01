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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

// HistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L519-L565
type HistogramAggregation struct {
	// ExtendedBounds Enables extending the bounds of the histogram beyond the data itself.
	ExtendedBounds *ExtendedBoundsdouble `json:"extended_bounds,omitempty"`
	// Field The name of the field to aggregate on.
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`
	// HardBounds Limits the range of buckets in the histogram.
	// It is particularly useful in the case of open data ranges that can result in
	// a very large number of buckets.
	HardBounds *ExtendedBoundsdouble `json:"hard_bounds,omitempty"`
	// Interval The interval for the buckets.
	// Must be a positive decimal.
	Interval *Float64 `json:"interval,omitempty"`
	// Keyed If `true`, returns buckets as a hash instead of an array, keyed by the bucket
	// keys.
	Keyed *bool `json:"keyed,omitempty"`
	// MinDocCount Only returns buckets that have `min_doc_count` number of documents.
	// By default, the response will fill gaps in the histogram with empty buckets.
	MinDocCount *int `json:"min_doc_count,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing *Float64 `json:"missing,omitempty"`
	// Offset By default, the bucket keys start with 0 and then continue in even spaced
	// steps of `interval`.
	// The bucket boundaries can be shifted by using the `offset` option.
	Offset *Float64 `json:"offset,omitempty"`
	// Order The sort order of the returned buckets.
	// By default, the returned buckets are sorted by their key ascending.
	Order  AggregateOrder `json:"order,omitempty"`
	Script *Script        `json:"script,omitempty"`
}

func (s *HistogramAggregation) UnmarshalJSON(data []byte) error {

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

		case "extended_bounds":
			if err := dec.Decode(&s.ExtendedBounds); err != nil {
				return fmt.Errorf("%s | %w", "ExtendedBounds", err)
			}

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

		case "hard_bounds":
			if err := dec.Decode(&s.HardBounds); err != nil {
				return fmt.Errorf("%s | %w", "HardBounds", err)
			}

		case "interval":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Interval", err)
				}
				f := Float64(value)
				s.Interval = &f
			case float64:
				f := Float64(v)
				s.Interval = &f
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

		case "min_doc_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int(v)
				s.MinDocCount = &f
			}

		case "missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Missing", err)
				}
				f := Float64(value)
				s.Missing = &f
			case float64:
				f := Float64(v)
				s.Missing = &f
			}

		case "offset":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Offset", err)
				}
				f := Float64(value)
				s.Offset = &f
			case float64:
				f := Float64(v)
				s.Offset = &f
			}

		case "order":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = o
			case '[':
				o := make([]map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = o
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewHistogramAggregation returns a HistogramAggregation.
func NewHistogramAggregation() *HistogramAggregation {
	r := &HistogramAggregation{}

	return r
}

type HistogramAggregationVariant interface {
	HistogramAggregationCaster() *HistogramAggregation
}

func (s *HistogramAggregation) HistogramAggregationCaster() *HistogramAggregation {
	return s
}
