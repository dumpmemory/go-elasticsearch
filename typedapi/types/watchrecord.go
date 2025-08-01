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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionstatus"
)

// WatchRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/execute_watch/types.ts#L27-L39
type WatchRecord struct {
	Condition    WatcherCondition                `json:"condition"`
	Input        WatcherInput                    `json:"input"`
	Messages     []string                        `json:"messages"`
	Metadata     Metadata                        `json:"metadata,omitempty"`
	Node         string                          `json:"node"`
	Result       ExecutionResult                 `json:"result"`
	State        executionstatus.ExecutionStatus `json:"state"`
	Status       *WatchStatus                    `json:"status,omitempty"`
	TriggerEvent TriggerEventResult              `json:"trigger_event"`
	User         string                          `json:"user"`
	WatchId      string                          `json:"watch_id"`
}

func (s *WatchRecord) UnmarshalJSON(data []byte) error {

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

		case "condition":
			if err := dec.Decode(&s.Condition); err != nil {
				return fmt.Errorf("%s | %w", "Condition", err)
			}

		case "input":
			if err := dec.Decode(&s.Input); err != nil {
				return fmt.Errorf("%s | %w", "Input", err)
			}

		case "messages":
			if err := dec.Decode(&s.Messages); err != nil {
				return fmt.Errorf("%s | %w", "Messages", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
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

		case "result":
			if err := dec.Decode(&s.Result); err != nil {
				return fmt.Errorf("%s | %w", "Result", err)
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "trigger_event":
			if err := dec.Decode(&s.TriggerEvent); err != nil {
				return fmt.Errorf("%s | %w", "TriggerEvent", err)
			}

		case "user":
			if err := dec.Decode(&s.User); err != nil {
				return fmt.Errorf("%s | %w", "User", err)
			}

		case "watch_id":
			if err := dec.Decode(&s.WatchId); err != nil {
				return fmt.Errorf("%s | %w", "WatchId", err)
			}

		}
	}
	return nil
}

// NewWatchRecord returns a WatchRecord.
func NewWatchRecord() *WatchRecord {
	r := &WatchRecord{}

	return r
}
