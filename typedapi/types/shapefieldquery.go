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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

// ShapeFieldQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/specialized.ts#L383-L396
type ShapeFieldQuery struct {
	// IndexedShape Queries using a pre-indexed shape.
	IndexedShape *FieldLookup `json:"indexed_shape,omitempty"`
	// Relation Spatial relation between the query shape and the document shape.
	Relation *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	// Shape Queries using an inline shape definition in GeoJSON or Well Known Text (WKT)
	// format.
	Shape json.RawMessage `json:"shape,omitempty"`
}

func (s *ShapeFieldQuery) UnmarshalJSON(data []byte) error {

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

		case "indexed_shape":
			if err := dec.Decode(&s.IndexedShape); err != nil {
				return fmt.Errorf("%s | %w", "IndexedShape", err)
			}

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return fmt.Errorf("%s | %w", "Relation", err)
			}

		case "shape":
			if err := dec.Decode(&s.Shape); err != nil {
				return fmt.Errorf("%s | %w", "Shape", err)
			}

		}
	}
	return nil
}

// NewShapeFieldQuery returns a ShapeFieldQuery.
func NewShapeFieldQuery() *ShapeFieldQuery {
	r := &ShapeFieldQuery{}

	return r
}

type ShapeFieldQueryVariant interface {
	ShapeFieldQueryCaster() *ShapeFieldQuery
}

func (s *ShapeFieldQuery) ShapeFieldQueryCaster() *ShapeFieldQuery {
	return s
}
