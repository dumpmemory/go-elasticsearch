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

// CommandCancelAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cluster/reroute/types.ts#L45-L50
type CommandCancelAction struct {
	AllowPrimary *bool  `json:"allow_primary,omitempty"`
	Index        string `json:"index"`
	Node         string `json:"node"`
	Shard        int    `json:"shard"`
}

func (s *CommandCancelAction) UnmarshalJSON(data []byte) error {

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

		case "allow_primary":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowPrimary", err)
				}
				s.AllowPrimary = &value
			case bool:
				s.AllowPrimary = &v
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = o

		case "shard":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Shard", err)
				}
				s.Shard = value
			case float64:
				f := int(v)
				s.Shard = f
			}

		}
	}
	return nil
}

// NewCommandCancelAction returns a CommandCancelAction.
func NewCommandCancelAction() *CommandCancelAction {
	r := &CommandCancelAction{}

	return r
}

type CommandCancelActionVariant interface {
	CommandCancelActionCaster() *CommandCancelAction
}

func (s *CommandCancelAction) CommandCancelActionCaster() *CommandCancelAction {
	return s
}
