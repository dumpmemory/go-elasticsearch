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

// LifecycleExplainUnmanaged type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ilm/explain_lifecycle/types.ts#L60-L63
type LifecycleExplainUnmanaged struct {
	Index   string `json:"index"`
	Managed bool   `json:"managed,omitempty"`
}

func (s *LifecycleExplainUnmanaged) UnmarshalJSON(data []byte) error {

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

		case "managed":
			if err := dec.Decode(&s.Managed); err != nil {
				return fmt.Errorf("%s | %w", "Managed", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s LifecycleExplainUnmanaged) MarshalJSON() ([]byte, error) {
	type innerLifecycleExplainUnmanaged LifecycleExplainUnmanaged
	tmp := innerLifecycleExplainUnmanaged{
		Index:   s.Index,
		Managed: s.Managed,
	}

	tmp.Managed = false

	return json.Marshal(tmp)
}

// NewLifecycleExplainUnmanaged returns a LifecycleExplainUnmanaged.
func NewLifecycleExplainUnmanaged() *LifecycleExplainUnmanaged {
	r := &LifecycleExplainUnmanaged{}

	return r
}
