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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cardinalityexecutionmode"
)

// CardinalityAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/metric.ts#L93-L105
type CardinalityAggregation struct {
	// ExecutionHint Mechanism by which cardinality aggregations is run.
	ExecutionHint *cardinalityexecutionmode.CardinalityExecutionMode `json:"execution_hint,omitempty"`
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	// PrecisionThreshold A unique count below which counts are expected to be close to accurate.
	// This allows to trade memory for accuracy.
	PrecisionThreshold *int    `json:"precision_threshold,omitempty"`
	Rehash             *bool   `json:"rehash,omitempty"`
	Script             *Script `json:"script,omitempty"`
}

func (s *CardinalityAggregation) UnmarshalJSON(data []byte) error {

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

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "precision_threshold":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrecisionThreshold", err)
				}
				s.PrecisionThreshold = &value
			case float64:
				f := int(v)
				s.PrecisionThreshold = &f
			}

		case "rehash":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rehash", err)
				}
				s.Rehash = &value
			case bool:
				s.Rehash = &v
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewCardinalityAggregation returns a CardinalityAggregation.
func NewCardinalityAggregation() *CardinalityAggregation {
	r := &CardinalityAggregation{}

	return r
}

type CardinalityAggregationVariant interface {
	CardinalityAggregationCaster() *CardinalityAggregation
}

func (s *CardinalityAggregation) CardinalityAggregationCaster() *CardinalityAggregation {
	return s
}
