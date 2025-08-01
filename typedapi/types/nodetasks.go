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

// NodeTasks type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/tasks/_types/TaskListResponseBase.ts#L49-L57
type NodeTasks struct {
	Attributes       map[string]string   `json:"attributes,omitempty"`
	Host             *string             `json:"host,omitempty"`
	Ip               *string             `json:"ip,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Roles            []string            `json:"roles,omitempty"`
	Tasks            map[string]TaskInfo `json:"tasks"`
	TransportAddress *string             `json:"transport_address,omitempty"`
}

func (s *NodeTasks) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return fmt.Errorf("%s | %w", "Ip", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "tasks":
			if s.Tasks == nil {
				s.Tasks = make(map[string]TaskInfo, 0)
			}
			if err := dec.Decode(&s.Tasks); err != nil {
				return fmt.Errorf("%s | %w", "Tasks", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		}
	}
	return nil
}

// NewNodeTasks returns a NodeTasks.
func NewNodeTasks() *NodeTasks {
	r := &NodeTasks{
		Attributes: make(map[string]string),
		Tasks:      make(map[string]TaskInfo),
	}

	return r
}
