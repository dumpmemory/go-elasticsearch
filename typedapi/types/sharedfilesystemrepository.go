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

// SharedFileSystemRepository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotRepository.ts#L80-L90
type SharedFileSystemRepository struct {
	// Settings The repository settings.
	Settings SharedFileSystemRepositorySettings `json:"settings"`
	// Type The shared file system repository type.
	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *SharedFileSystemRepository) UnmarshalJSON(data []byte) error {

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
func (s SharedFileSystemRepository) MarshalJSON() ([]byte, error) {
	type innerSharedFileSystemRepository SharedFileSystemRepository
	tmp := innerSharedFileSystemRepository{
		Settings: s.Settings,
		Type:     s.Type,
		Uuid:     s.Uuid,
	}

	tmp.Type = "fs"

	return json.Marshal(tmp)
}

// NewSharedFileSystemRepository returns a SharedFileSystemRepository.
func NewSharedFileSystemRepository() *SharedFileSystemRepository {
	r := &SharedFileSystemRepository{}

	return r
}

type SharedFileSystemRepositoryVariant interface {
	SharedFileSystemRepositoryCaster() *SharedFileSystemRepository
}

func (s *SharedFileSystemRepository) SharedFileSystemRepositoryCaster() *SharedFileSystemRepository {
	return s
}

func (s *SharedFileSystemRepository) RepositoryCaster() *Repository {
	o := Repository(s)
	return &o
}
