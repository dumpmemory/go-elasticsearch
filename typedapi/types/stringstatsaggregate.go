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

// StringStatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/Aggregate.ts#L795-L806
type StringStatsAggregate struct {
	AvgLength         *Float64            `json:"avg_length,omitempty"`
	AvgLengthAsString *string             `json:"avg_length_as_string,omitempty"`
	Count             int64               `json:"count"`
	Distribution      *map[string]Float64 `json:"distribution,omitempty"`
	Entropy           *Float64            `json:"entropy,omitempty"`
	MaxLength         *int                `json:"max_length,omitempty"`
	MaxLengthAsString *string             `json:"max_length_as_string,omitempty"`
	Meta              Metadata            `json:"meta,omitempty"`
	MinLength         *int                `json:"min_length,omitempty"`
	MinLengthAsString *string             `json:"min_length_as_string,omitempty"`
}

func (s *StringStatsAggregate) UnmarshalJSON(data []byte) error {

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

		case "avg_length":
			if err := dec.Decode(&s.AvgLength); err != nil {
				return fmt.Errorf("%s | %w", "AvgLength", err)
			}

		case "avg_length_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AvgLengthAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AvgLengthAsString = &o

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "distribution":
			if err := dec.Decode(&s.Distribution); err != nil {
				return fmt.Errorf("%s | %w", "Distribution", err)
			}

		case "entropy":
			if err := dec.Decode(&s.Entropy); err != nil {
				return fmt.Errorf("%s | %w", "Entropy", err)
			}

		case "max_length":
			if err := dec.Decode(&s.MaxLength); err != nil {
				return fmt.Errorf("%s | %w", "MaxLength", err)
			}

		case "max_length_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxLengthAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxLengthAsString = &o

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "min_length":
			if err := dec.Decode(&s.MinLength); err != nil {
				return fmt.Errorf("%s | %w", "MinLength", err)
			}

		case "min_length_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MinLengthAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinLengthAsString = &o

		}
	}
	return nil
}

// NewStringStatsAggregate returns a StringStatsAggregate.
func NewStringStatsAggregate() *StringStatsAggregate {
	r := &StringStatsAggregate{}

	return r
}
