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

// RemoveIndexAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/update_aliases/types.ts#L124-L139
type RemoveIndexAction struct {
	// Index Data stream or index for the action.
	// Supports wildcards (`*`).
	Index *string `json:"index,omitempty"`
	// Indices Data streams or indices for the action.
	// Supports wildcards (`*`).
	Indices []string `json:"indices,omitempty"`
	// MustExist If `true`, the alias must exist to perform the action.
	MustExist *bool `json:"must_exist,omitempty"`
}

func (s *RemoveIndexAction) UnmarshalJSON(data []byte) error {

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

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}
			}

		case "must_exist":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MustExist", err)
				}
				s.MustExist = &value
			case bool:
				s.MustExist = &v
			}

		}
	}
	return nil
}

// NewRemoveIndexAction returns a RemoveIndexAction.
func NewRemoveIndexAction() *RemoveIndexAction {
	r := &RemoveIndexAction{}

	return r
}

type RemoveIndexActionVariant interface {
	RemoveIndexActionCaster() *RemoveIndexAction
}

func (s *RemoveIndexAction) RemoveIndexActionCaster() *RemoveIndexAction {
	return s
}
