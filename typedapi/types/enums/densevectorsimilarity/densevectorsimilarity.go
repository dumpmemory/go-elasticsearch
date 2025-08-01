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

// Package densevectorsimilarity
package densevectorsimilarity

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/mapping/DenseVectorProperty.ts#L82-L127
type DenseVectorSimilarity struct {
	Name string
}

var (
	Cosine = DenseVectorSimilarity{"cosine"}

	Dotproduct = DenseVectorSimilarity{"dot_product"}

	L2norm = DenseVectorSimilarity{"l2_norm"}

	Maxinnerproduct = DenseVectorSimilarity{"max_inner_product"}
)

func (d DenseVectorSimilarity) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DenseVectorSimilarity) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "cosine":
		*d = Cosine
	case "dot_product":
		*d = Dotproduct
	case "l2_norm":
		*d = L2norm
	case "max_inner_product":
		*d = Maxinnerproduct
	default:
		*d = DenseVectorSimilarity{string(text)}
	}

	return nil
}

func (d DenseVectorSimilarity) String() string {
	return d.Name
}
