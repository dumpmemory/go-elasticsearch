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

// CgroupCpu type.
//
// https://github.com/elastic/elasticsearch-specification/blob/33e8a1c9cad22a5946ac735c4fba31af2da2cec2/specification/nodes/_types/Stats.ts#L199-L204
type CgroupCpu struct {
	CfsPeriodMicros *int           `json:"cfs_period_micros,omitempty"`
	CfsQuotaMicros  *int           `json:"cfs_quota_micros,omitempty"`
	ControlGroup    *string        `json:"control_group,omitempty"`
	Stat            *CgroupCpuStat `json:"stat,omitempty"`
}

func (s *CgroupCpu) UnmarshalJSON(data []byte) error {

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

		case "cfs_period_micros":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CfsPeriodMicros = &value
			case float64:
				f := int(v)
				s.CfsPeriodMicros = &f
			}

		case "cfs_quota_micros":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CfsQuotaMicros = &value
			case float64:
				f := int(v)
				s.CfsQuotaMicros = &f
			}

		case "control_group":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ControlGroup = &o

		case "stat":
			if err := dec.Decode(&s.Stat); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCgroupCpu returns a CgroupCpu.
func NewCgroupCpu() *CgroupCpu {
	r := &CgroupCpu{}

	return r
}
