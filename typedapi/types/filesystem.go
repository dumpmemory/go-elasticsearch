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

// FileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L771-L789
type FileSystem struct {
	// Data List of all file stores.
	Data []DataPathStats `json:"data,omitempty"`
	// IoStats Contains I/O statistics for the node.
	IoStats *IoStats `json:"io_stats,omitempty"`
	// Timestamp Last time the file stores statistics were refreshed.
	// Recorded in milliseconds since the Unix Epoch.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Total Contains statistics for all file stores of the node.
	Total *FileSystemTotal `json:"total,omitempty"`
}

func (s *FileSystem) UnmarshalJSON(data []byte) error {

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

		case "data":
			if err := dec.Decode(&s.Data); err != nil {
				return fmt.Errorf("%s | %w", "Data", err)
			}

		case "io_stats":
			if err := dec.Decode(&s.IoStats); err != nil {
				return fmt.Errorf("%s | %w", "IoStats", err)
			}

		case "timestamp":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Timestamp", err)
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

		}
	}
	return nil
}

// NewFileSystem returns a FileSystem.
func NewFileSystem() *FileSystem {
	r := &FileSystem{}

	return r
}
