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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

// MTermVectorsOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/mtermvectors/types.ts#L35-L94
type MTermVectorsOperation struct {
	// Doc An artificial document (a document not present in the index) for which you
	// want to retrieve term vectors.
	Doc json.RawMessage `json:"doc,omitempty"`
	// FieldStatistics If `true`, the response includes the document count, sum of document
	// frequencies, and sum of total term frequencies.
	FieldStatistics *bool `json:"field_statistics,omitempty"`
	// Fields Comma-separated list or wildcard expressions of fields to include in the
	// statistics.
	// Used as the default list unless a specific field list is provided in the
	// `completion_fields` or `fielddata_fields` parameters.
	Fields []string `json:"fields,omitempty"`
	// Filter Filter terms based on their tf-idf scores.
	Filter *TermVectorsFilter `json:"filter,omitempty"`
	// Id_ The ID of the document.
	Id_ *string `json:"_id,omitempty"`
	// Index_ The index of the document.
	Index_ *string `json:"_index,omitempty"`
	// Offsets If `true`, the response includes term offsets.
	Offsets *bool `json:"offsets,omitempty"`
	// Payloads If `true`, the response includes term payloads.
	Payloads *bool `json:"payloads,omitempty"`
	// Positions If `true`, the response includes term positions.
	Positions *bool `json:"positions,omitempty"`
	// Routing Custom value used to route operations to a specific shard.
	Routing *string `json:"routing,omitempty"`
	// TermStatistics If true, the response includes term frequency and document frequency.
	TermStatistics *bool `json:"term_statistics,omitempty"`
	// Version If `true`, returns the document version as part of a hit.
	Version *int64 `json:"version,omitempty"`
	// VersionType Specific version type.
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MTermVectorsOperation) UnmarshalJSON(data []byte) error {

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

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return fmt.Errorf("%s | %w", "Doc", err)
			}

		case "field_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FieldStatistics", err)
				}
				s.FieldStatistics = &value
			case bool:
				s.FieldStatistics = &v
			}

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "offsets":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Offsets", err)
				}
				s.Offsets = &value
			case bool:
				s.Offsets = &v
			}

		case "payloads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Payloads", err)
				}
				s.Payloads = &value
			case bool:
				s.Payloads = &v
			}

		case "positions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Positions", err)
				}
				s.Positions = &value
			case bool:
				s.Positions = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "term_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TermStatistics", err)
				}
				s.TermStatistics = &value
			case bool:
				s.TermStatistics = &v
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
			}

		}
	}
	return nil
}

// NewMTermVectorsOperation returns a MTermVectorsOperation.
func NewMTermVectorsOperation() *MTermVectorsOperation {
	r := &MTermVectorsOperation{}

	return r
}

type MTermVectorsOperationVariant interface {
	MTermVectorsOperationCaster() *MTermVectorsOperation
}

func (s *MTermVectorsOperation) MTermVectorsOperationCaster() *MTermVectorsOperation {
	return s
}
