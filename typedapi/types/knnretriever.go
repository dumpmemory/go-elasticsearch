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

// KnnRetriever type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Retriever.ts#L115-L133
type KnnRetriever struct {
	// Field The name of the vector field to search against.
	Field string `json:"field"`
	// Filter Query to filter the documents that can match.
	Filter []Query `json:"filter,omitempty"`
	// K Number of nearest neighbors to return as top hits.
	K int `json:"k"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are not
	// included in the top documents.
	MinScore *float32 `json:"min_score,omitempty"`
	// Name_ Retriever name.
	Name_ *string `json:"_name,omitempty"`
	// NumCandidates Number of nearest neighbor candidates to consider per shard.
	NumCandidates int `json:"num_candidates"`
	// QueryVector Query vector. Must have the same number of dimensions as the vector field you
	// are searching against. You must provide a query_vector_builder or
	// query_vector, but not both.
	QueryVector []float32 `json:"query_vector,omitempty"`
	// QueryVectorBuilder Defines a model to build a query vector.
	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`
	// RescoreVector Apply oversampling and rescoring to quantized vectors
	RescoreVector *RescoreVector `json:"rescore_vector,omitempty"`
	// Similarity The minimum similarity required for a document to be considered a match.
	Similarity *float32 `json:"similarity,omitempty"`
}

func (s *KnnRetriever) UnmarshalJSON(data []byte) error {

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

		case "field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Field = o

		case "filter":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewQuery()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}

				s.Filter = append(s.Filter, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Filter); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}
			}

		case "k":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "K", err)
				}
				s.K = value
			case float64:
				f := int(v)
				s.K = f
			}

		case "min_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinScore", err)
				}
				f := float32(value)
				s.MinScore = &f
			case float64:
				f := float32(v)
				s.MinScore = &f
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name_ = &o

		case "num_candidates":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumCandidates", err)
				}
				s.NumCandidates = value
			case float64:
				f := int(v)
				s.NumCandidates = f
			}

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		case "query_vector_builder":
			if err := dec.Decode(&s.QueryVectorBuilder); err != nil {
				return fmt.Errorf("%s | %w", "QueryVectorBuilder", err)
			}

		case "rescore_vector":
			if err := dec.Decode(&s.RescoreVector); err != nil {
				return fmt.Errorf("%s | %w", "RescoreVector", err)
			}

		case "similarity":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Similarity", err)
				}
				f := float32(value)
				s.Similarity = &f
			case float64:
				f := float32(v)
				s.Similarity = &f
			}

		}
	}
	return nil
}

// NewKnnRetriever returns a KnnRetriever.
func NewKnnRetriever() *KnnRetriever {
	r := &KnnRetriever{}

	return r
}

type KnnRetrieverVariant interface {
	KnnRetrieverCaster() *KnnRetriever
}

func (s *KnnRetriever) KnnRetrieverCaster() *KnnRetriever {
	return s
}
