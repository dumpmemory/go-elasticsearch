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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

// IndexAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Actions.ts#L256-L265
type IndexAction struct {
	DocId              *string          `json:"doc_id,omitempty"`
	ExecutionTimeField *string          `json:"execution_time_field,omitempty"`
	Index              string           `json:"index"`
	OpType             *optype.OpType   `json:"op_type,omitempty"`
	Refresh            *refresh.Refresh `json:"refresh,omitempty"`
	Timeout            Duration         `json:"timeout,omitempty"`
}

func (s *IndexAction) UnmarshalJSON(data []byte) error {

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

		case "doc_id":
			if err := dec.Decode(&s.DocId); err != nil {
				return fmt.Errorf("%s | %w", "DocId", err)
			}

		case "execution_time_field":
			if err := dec.Decode(&s.ExecutionTimeField); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionTimeField", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "op_type":
			if err := dec.Decode(&s.OpType); err != nil {
				return fmt.Errorf("%s | %w", "OpType", err)
			}

		case "refresh":
			if err := dec.Decode(&s.Refresh); err != nil {
				return fmt.Errorf("%s | %w", "Refresh", err)
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return fmt.Errorf("%s | %w", "Timeout", err)
			}

		}
	}
	return nil
}

// NewIndexAction returns a IndexAction.
func NewIndexAction() *IndexAction {
	r := &IndexAction{}

	return r
}

type IndexActionVariant interface {
	IndexActionCaster() *IndexAction
}

func (s *IndexAction) IndexActionCaster() *IndexAction {
	return s
}
