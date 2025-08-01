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

// Sql type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L392-L395
type Sql struct {
	Available bool                  `json:"available"`
	Enabled   bool                  `json:"enabled"`
	Features  map[string]int        `json:"features"`
	Queries   map[string]XpackQuery `json:"queries"`
}

func (s *Sql) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "features":
			if s.Features == nil {
				s.Features = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.Features); err != nil {
				return fmt.Errorf("%s | %w", "Features", err)
			}

		case "queries":
			if s.Queries == nil {
				s.Queries = make(map[string]XpackQuery, 0)
			}
			if err := dec.Decode(&s.Queries); err != nil {
				return fmt.Errorf("%s | %w", "Queries", err)
			}

		}
	}
	return nil
}

// NewSql returns a Sql.
func NewSql() *Sql {
	r := &Sql{
		Features: make(map[string]int),
		Queries:  make(map[string]XpackQuery),
	}

	return r
}
