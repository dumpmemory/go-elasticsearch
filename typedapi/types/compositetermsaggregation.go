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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

// CompositeTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L181-L181
type CompositeTermsAggregation struct {
	// Field Either `field` or `script` must be present
	Field         *string                    `json:"field,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`
	// Script Either `field` or `script` must be present
	Script    *Script              `json:"script,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeTermsAggregation) UnmarshalJSON(data []byte) error {

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

		case "missing_bucket":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MissingBucket", err)
				}
				s.MissingBucket = &value
			case bool:
				s.MissingBucket = &v
			}

		case "missing_order":
			if err := dec.Decode(&s.MissingOrder); err != nil {
				return fmt.Errorf("%s | %w", "MissingOrder", err)
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return fmt.Errorf("%s | %w", "Order", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "value_type":
			if err := dec.Decode(&s.ValueType); err != nil {
				return fmt.Errorf("%s | %w", "ValueType", err)
			}

		}
	}
	return nil
}

// NewCompositeTermsAggregation returns a CompositeTermsAggregation.
func NewCompositeTermsAggregation() *CompositeTermsAggregation {
	r := &CompositeTermsAggregation{}

	return r
}

type CompositeTermsAggregationVariant interface {
	CompositeTermsAggregationCaster() *CompositeTermsAggregation
}

func (s *CompositeTermsAggregation) CompositeTermsAggregationCaster() *CompositeTermsAggregation {
	return s
}
