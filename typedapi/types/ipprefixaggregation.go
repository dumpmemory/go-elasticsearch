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

// IpPrefixAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L1196-L1225
type IpPrefixAggregation struct {
	// AppendPrefixLength Defines whether the prefix length is appended to IP address keys in the
	// response.
	AppendPrefixLength *bool `json:"append_prefix_length,omitempty"`
	// Field The IP address field to aggregation on. The field mapping type must be `ip`.
	Field string `json:"field"`
	// IsIpv6 Defines whether the prefix applies to IPv6 addresses.
	IsIpv6 *bool `json:"is_ipv6,omitempty"`
	// Keyed Defines whether buckets are returned as a hash rather than an array in the
	// response.
	Keyed *bool `json:"keyed,omitempty"`
	// MinDocCount Minimum number of documents in a bucket for it to be included in the
	// response.
	MinDocCount *int64 `json:"min_doc_count,omitempty"`
	// PrefixLength Length of the network prefix. For IPv4 addresses the accepted range is [0,
	// 32].
	// For IPv6 addresses the accepted range is [0, 128].
	PrefixLength int `json:"prefix_length"`
}

func (s *IpPrefixAggregation) UnmarshalJSON(data []byte) error {

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

		case "append_prefix_length":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AppendPrefixLength", err)
				}
				s.AppendPrefixLength = &value
			case bool:
				s.AppendPrefixLength = &v
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "is_ipv6":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsIpv6", err)
				}
				s.IsIpv6 = &value
			case bool:
				s.IsIpv6 = &v
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
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int64(v)
				s.MinDocCount = &f
			}

		case "prefix_length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrefixLength", err)
				}
				s.PrefixLength = value
			case float64:
				f := int(v)
				s.PrefixLength = f
			}

		}
	}
	return nil
}

// NewIpPrefixAggregation returns a IpPrefixAggregation.
func NewIpPrefixAggregation() *IpPrefixAggregation {
	r := &IpPrefixAggregation{}

	return r
}

type IpPrefixAggregationVariant interface {
	IpPrefixAggregationCaster() *IpPrefixAggregation
}

func (s *IpPrefixAggregation) IpPrefixAggregationCaster() *IpPrefixAggregation {
	return s
}
