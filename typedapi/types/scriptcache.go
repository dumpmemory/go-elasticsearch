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

// ScriptCache type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L1109-L1123
type ScriptCache struct {
	// CacheEvictions Total number of times the script cache has evicted old data.
	CacheEvictions *int64 `json:"cache_evictions,omitempty"`
	// CompilationLimitTriggered Total number of times the script compilation circuit breaker has limited
	// inline script compilations.
	CompilationLimitTriggered *int64 `json:"compilation_limit_triggered,omitempty"`
	// Compilations Total number of inline script compilations performed by the node.
	Compilations *int64  `json:"compilations,omitempty"`
	Context      *string `json:"context,omitempty"`
}

func (s *ScriptCache) UnmarshalJSON(data []byte) error {

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

		case "cache_evictions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CacheEvictions", err)
				}
				s.CacheEvictions = &value
			case float64:
				f := int64(v)
				s.CacheEvictions = &f
			}

		case "compilation_limit_triggered":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CompilationLimitTriggered", err)
				}
				s.CompilationLimitTriggered = &value
			case float64:
				f := int64(v)
				s.CompilationLimitTriggered = &f
			}

		case "compilations":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compilations", err)
				}
				s.Compilations = &value
			case float64:
				f := int64(v)
				s.Compilations = &f
			}

		case "context":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Context", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Context = &o

		}
	}
	return nil
}

// NewScriptCache returns a ScriptCache.
func NewScriptCache() *ScriptCache {
	r := &ScriptCache{}

	return r
}
