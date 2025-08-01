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

// ExtendedBoundsdouble type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L508-L517
type ExtendedBoundsdouble struct {
	// Max Maximum value for the bound.
	Max *Float64 `json:"max,omitempty"`
	// Min Minimum value for the bound.
	Min *Float64 `json:"min,omitempty"`
}

func (s *ExtendedBoundsdouble) UnmarshalJSON(data []byte) error {

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

		case "max":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Max", err)
				}
				f := Float64(value)
				s.Max = &f
			case float64:
				f := Float64(v)
				s.Max = &f
			}

		case "min":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Min", err)
				}
				f := Float64(value)
				s.Min = &f
			case float64:
				f := Float64(v)
				s.Min = &f
			}

		}
	}
	return nil
}

// NewExtendedBoundsdouble returns a ExtendedBoundsdouble.
func NewExtendedBoundsdouble() *ExtendedBoundsdouble {
	r := &ExtendedBoundsdouble{}

	return r
}

type ExtendedBoundsdoubleVariant interface {
	ExtendedBoundsdoubleCaster() *ExtendedBoundsdouble
}

func (s *ExtendedBoundsdouble) ExtendedBoundsdoubleCaster() *ExtendedBoundsdouble {
	return s
}
