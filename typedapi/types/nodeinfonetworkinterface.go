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
// https://github.com/elastic/elasticsearch-specification/tree/33e8a1c9cad22a5946ac735c4fba31af2da2cec2

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// NodeInfoNetworkInterface type.
//
// https://github.com/elastic/elasticsearch-specification/blob/33e8a1c9cad22a5946ac735c4fba31af2da2cec2/specification/nodes/info/types.ts#L328-L332
type NodeInfoNetworkInterface struct {
	Address    string `json:"address"`
	MacAddress string `json:"mac_address"`
	Name       string `json:"name"`
}

func (s *NodeInfoNetworkInterface) UnmarshalJSON(data []byte) error {

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

		case "address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Address = o

		case "mac_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MacAddress = o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodeInfoNetworkInterface returns a NodeInfoNetworkInterface.
func NewNodeInfoNetworkInterface() *NodeInfoNetworkInterface {
	r := &NodeInfoNetworkInterface{}

	return r
}
