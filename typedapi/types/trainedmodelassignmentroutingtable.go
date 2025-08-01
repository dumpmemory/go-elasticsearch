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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/routingstate"
)

// TrainedModelAssignmentRoutingTable type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/TrainedModel.ts#L443-L461
type TrainedModelAssignmentRoutingTable struct {
	// CurrentAllocations Current number of allocations.
	CurrentAllocations int `json:"current_allocations"`
	// Reason The reason for the current state. It is usually populated only when the
	// `routing_state` is `failed`.
	Reason *string `json:"reason,omitempty"`
	// RoutingState The current routing state.
	RoutingState routingstate.RoutingState `json:"routing_state"`
	// TargetAllocations Target number of allocations.
	TargetAllocations int `json:"target_allocations"`
}

func (s *TrainedModelAssignmentRoutingTable) UnmarshalJSON(data []byte) error {

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

		case "current_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentAllocations", err)
				}
				s.CurrentAllocations = value
			case float64:
				f := int(v)
				s.CurrentAllocations = f
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "routing_state":
			if err := dec.Decode(&s.RoutingState); err != nil {
				return fmt.Errorf("%s | %w", "RoutingState", err)
			}

		case "target_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TargetAllocations", err)
				}
				s.TargetAllocations = value
			case float64:
				f := int(v)
				s.TargetAllocations = f
			}

		}
	}
	return nil
}

// NewTrainedModelAssignmentRoutingTable returns a TrainedModelAssignmentRoutingTable.
func NewTrainedModelAssignmentRoutingTable() *TrainedModelAssignmentRoutingTable {
	r := &TrainedModelAssignmentRoutingTable{}

	return r
}
