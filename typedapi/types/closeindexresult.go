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

// CloseIndexResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/close/CloseIndexResponse.ts#L32-L35
type CloseIndexResult struct {
	Closed bool                        `json:"closed"`
	Shards map[string]CloseShardResult `json:"shards,omitempty"`
}

func (s *CloseIndexResult) UnmarshalJSON(data []byte) error {

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

		case "closed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Closed", err)
				}
				s.Closed = value
			case bool:
				s.Closed = v
			}

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string]CloseShardResult, 0)
			}
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		}
	}
	return nil
}

// NewCloseIndexResult returns a CloseIndexResult.
func NewCloseIndexResult() *CloseIndexResult {
	r := &CloseIndexResult{
		Shards: make(map[string]CloseShardResult),
	}

	return r
}
