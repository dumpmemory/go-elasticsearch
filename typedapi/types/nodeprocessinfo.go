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

// NodeProcessInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/info/types.ts#L402-L409
type NodeProcessInfo struct {
	// Id Process identifier (PID)
	Id int64 `json:"id"`
	// Mlockall Indicates if the process address space has been successfully locked in memory
	Mlockall bool `json:"mlockall"`
	// RefreshIntervalInMillis Refresh interval for the process statistics
	RefreshIntervalInMillis int64 `json:"refresh_interval_in_millis"`
}

func (s *NodeProcessInfo) UnmarshalJSON(data []byte) error {

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

		case "id":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Id", err)
				}
				s.Id = value
			case float64:
				f := int64(v)
				s.Id = f
			}

		case "mlockall":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Mlockall", err)
				}
				s.Mlockall = value
			case bool:
				s.Mlockall = v
			}

		case "refresh_interval_in_millis":
			if err := dec.Decode(&s.RefreshIntervalInMillis); err != nil {
				return fmt.Errorf("%s | %w", "RefreshIntervalInMillis", err)
			}

		}
	}
	return nil
}

// NewNodeProcessInfo returns a NodeProcessInfo.
func NewNodeProcessInfo() *NodeProcessInfo {
	r := &NodeProcessInfo{}

	return r
}
