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

// ClusterProcessOpenFileDescriptors type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cluster/stats/types.ts#L596-L612
type ClusterProcessOpenFileDescriptors struct {
	// Avg Average number of concurrently open file descriptors.
	// Returns `-1` if not supported.
	Avg int64 `json:"avg"`
	// Max Maximum number of concurrently open file descriptors allowed across all
	// selected nodes.
	// Returns `-1` if not supported.
	Max int64 `json:"max"`
	// Min Minimum number of concurrently open file descriptors across all selected
	// nodes.
	// Returns -1 if not supported.
	Min int64 `json:"min"`
}

func (s *ClusterProcessOpenFileDescriptors) UnmarshalJSON(data []byte) error {

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

		case "avg":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Avg", err)
				}
				s.Avg = value
			case float64:
				f := int64(v)
				s.Avg = f
			}

		case "max":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Max", err)
				}
				s.Max = value
			case float64:
				f := int64(v)
				s.Max = f
			}

		case "min":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Min", err)
				}
				s.Min = value
			case float64:
				f := int64(v)
				s.Min = f
			}

		}
	}
	return nil
}

// NewClusterProcessOpenFileDescriptors returns a ClusterProcessOpenFileDescriptors.
func NewClusterProcessOpenFileDescriptors() *ClusterProcessOpenFileDescriptors {
	r := &ClusterProcessOpenFileDescriptors{}

	return r
}
