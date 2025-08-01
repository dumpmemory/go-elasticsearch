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

// DocumentRating type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/rank_eval/types.ts#L119-L126
type DocumentRating struct {
	// Id_ The document ID.
	Id_ string `json:"_id"`
	// Index_ The document’s index. For data streams, this should be the document’s backing
	// index.
	Index_ string `json:"_index"`
	// Rating The document’s relevance with regard to this search request.
	Rating int `json:"rating"`
}

func (s *DocumentRating) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "rating":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rating", err)
				}
				s.Rating = value
			case float64:
				f := int(v)
				s.Rating = f
			}

		}
	}
	return nil
}

// NewDocumentRating returns a DocumentRating.
func NewDocumentRating() *DocumentRating {
	r := &DocumentRating{}

	return r
}

type DocumentRatingVariant interface {
	DocumentRatingCaster() *DocumentRating
}

func (s *DocumentRating) DocumentRatingCaster() *DocumentRating {
	return s
}
