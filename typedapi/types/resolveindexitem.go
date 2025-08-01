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
)

// ResolveIndexItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/resolve_index/ResolveIndexResponse.ts#L30-L35
type ResolveIndexItem struct {
	Aliases    []string `json:"aliases,omitempty"`
	Attributes []string `json:"attributes"`
	DataStream *string  `json:"data_stream,omitempty"`
	Name       string   `json:"name"`
}

func (s *ResolveIndexItem) UnmarshalJSON(data []byte) error {

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

		case "aliases":
			if err := dec.Decode(&s.Aliases); err != nil {
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "attributes":
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return fmt.Errorf("%s | %w", "DataStream", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		}
	}
	return nil
}

// NewResolveIndexItem returns a ResolveIndexItem.
func NewResolveIndexItem() *ResolveIndexItem {
	r := &ResolveIndexItem{}

	return r
}
