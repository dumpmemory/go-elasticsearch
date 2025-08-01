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

// Hop type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/graph/_types/Hop.ts#L23-L36
type Hop struct {
	// Connections Specifies one or more fields from which you want to extract terms that are
	// associated with the specified vertices.
	Connections *Hop `json:"connections,omitempty"`
	// Query An optional guiding query that constrains the Graph API as it explores
	// connected terms.
	Query *Query `json:"query,omitempty"`
	// Vertices Contains the fields you are interested in.
	Vertices []VertexDefinition `json:"vertices"`
}

// NewHop returns a Hop.
func NewHop() *Hop {
	r := &Hop{}

	return r
}

type HopVariant interface {
	HopCaster() *Hop
}

func (s *Hop) HopCaster() *Hop {
	return s
}
