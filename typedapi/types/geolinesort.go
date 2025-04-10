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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// GeoLineSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/_types/aggregations/metric.ts#L151-L156
type GeoLineSort struct {
	// Field The name of the numeric field to use as the sort key for ordering the points.
	Field string `json:"field"`
}

func (s *GeoLineSort) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// NewGeoLineSort returns a GeoLineSort.
func NewGeoLineSort() *GeoLineSort {
	r := &GeoLineSort{}

	return r
}

// true

type GeoLineSortVariant interface {
	GeoLineSortCaster() *GeoLineSort
}

func (s *GeoLineSort) GeoLineSortCaster() *GeoLineSort {
	return s
}
