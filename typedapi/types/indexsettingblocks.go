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

// IndexSettingBlocks type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L264-L270
type IndexSettingBlocks struct {
	Metadata            Stringifiedboolean `json:"metadata,omitempty"`
	Read                Stringifiedboolean `json:"read,omitempty"`
	ReadOnly            Stringifiedboolean `json:"read_only,omitempty"`
	ReadOnlyAllowDelete Stringifiedboolean `json:"read_only_allow_delete,omitempty"`
	Write               Stringifiedboolean `json:"write,omitempty"`
}

func (s *IndexSettingBlocks) UnmarshalJSON(data []byte) error {

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

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "read":
			if err := dec.Decode(&s.Read); err != nil {
				return fmt.Errorf("%s | %w", "Read", err)
			}

		case "read_only":
			if err := dec.Decode(&s.ReadOnly); err != nil {
				return fmt.Errorf("%s | %w", "ReadOnly", err)
			}

		case "read_only_allow_delete":
			if err := dec.Decode(&s.ReadOnlyAllowDelete); err != nil {
				return fmt.Errorf("%s | %w", "ReadOnlyAllowDelete", err)
			}

		case "write":
			if err := dec.Decode(&s.Write); err != nil {
				return fmt.Errorf("%s | %w", "Write", err)
			}

		}
	}
	return nil
}

// NewIndexSettingBlocks returns a IndexSettingBlocks.
func NewIndexSettingBlocks() *IndexSettingBlocks {
	r := &IndexSettingBlocks{}

	return r
}

type IndexSettingBlocksVariant interface {
	IndexSettingBlocksCaster() *IndexSettingBlocks
}

func (s *IndexSettingBlocks) IndexSettingBlocksCaster() *IndexSettingBlocks {
	return s
}
