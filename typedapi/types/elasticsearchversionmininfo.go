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

// ElasticsearchVersionMinInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Base.ts#L120-L128
type ElasticsearchVersionMinInfo struct {
	BuildFlavor                      string `json:"build_flavor"`
	Int                              string `json:"number"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
}

func (s *ElasticsearchVersionMinInfo) UnmarshalJSON(data []byte) error {

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

		case "build_flavor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildFlavor", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildFlavor = o

		case "number":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Int", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Int = o

		case "minimum_index_compatibility_version":
			if err := dec.Decode(&s.MinimumIndexCompatibilityVersion); err != nil {
				return fmt.Errorf("%s | %w", "MinimumIndexCompatibilityVersion", err)
			}

		case "minimum_wire_compatibility_version":
			if err := dec.Decode(&s.MinimumWireCompatibilityVersion); err != nil {
				return fmt.Errorf("%s | %w", "MinimumWireCompatibilityVersion", err)
			}

		}
	}
	return nil
}

// NewElasticsearchVersionMinInfo returns a ElasticsearchVersionMinInfo.
func NewElasticsearchVersionMinInfo() *ElasticsearchVersionMinInfo {
	r := &ElasticsearchVersionMinInfo{}

	return r
}
