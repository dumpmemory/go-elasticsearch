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

// SourceOnlyRepository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotRepository.ts#L104-L114
type SourceOnlyRepository struct {
	// Settings The repository settings.
	Settings SourceOnlyRepositorySettings `json:"settings"`
	// Type The source-only repository type.
	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *SourceOnlyRepository) UnmarshalJSON(data []byte) error {

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

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SourceOnlyRepository) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepository SourceOnlyRepository
	tmp := innerSourceOnlyRepository{
		Settings: s.Settings,
		Type:     s.Type,
		Uuid:     s.Uuid,
	}

	tmp.Type = "source"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepository returns a SourceOnlyRepository.
func NewSourceOnlyRepository() *SourceOnlyRepository {
	r := &SourceOnlyRepository{}

	return r
}

type SourceOnlyRepositoryVariant interface {
	SourceOnlyRepositoryCaster() *SourceOnlyRepository
}

func (s *SourceOnlyRepository) SourceOnlyRepositoryCaster() *SourceOnlyRepository {
	return s
}

func (s *SourceOnlyRepository) RepositoryCaster() *Repository {
	o := Repository(s)
	return &o
}
