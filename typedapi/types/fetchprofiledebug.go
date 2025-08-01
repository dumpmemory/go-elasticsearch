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

// FetchProfileDebug type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/profile.ts#L250-L253
type FetchProfileDebug struct {
	FastPath     *int     `json:"fast_path,omitempty"`
	StoredFields []string `json:"stored_fields,omitempty"`
}

func (s *FetchProfileDebug) UnmarshalJSON(data []byte) error {

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

		case "fast_path":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FastPath", err)
				}
				s.FastPath = &value
			case float64:
				f := int(v)
				s.FastPath = &f
			}

		case "stored_fields":
			if err := dec.Decode(&s.StoredFields); err != nil {
				return fmt.Errorf("%s | %w", "StoredFields", err)
			}

		}
	}
	return nil
}

// NewFetchProfileDebug returns a FetchProfileDebug.
func NewFetchProfileDebug() *FetchProfileDebug {
	r := &FetchProfileDebug{}

	return r
}
