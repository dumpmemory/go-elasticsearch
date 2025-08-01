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

// RankEvalRequestItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/rank_eval/types.ts#L98-L109
type RankEvalRequestItem struct {
	// Id The search request’s ID, used to group result details later.
	Id string `json:"id"`
	// Params The search template parameters.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Ratings List of document ratings
	Ratings []DocumentRating `json:"ratings"`
	// Request The query being evaluated.
	Request *RankEvalQuery `json:"request,omitempty"`
	// TemplateId The search template Id
	TemplateId *string `json:"template_id,omitempty"`
}

func (s *RankEvalRequestItem) UnmarshalJSON(data []byte) error {

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "ratings":
			if err := dec.Decode(&s.Ratings); err != nil {
				return fmt.Errorf("%s | %w", "Ratings", err)
			}

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		case "template_id":
			if err := dec.Decode(&s.TemplateId); err != nil {
				return fmt.Errorf("%s | %w", "TemplateId", err)
			}

		}
	}
	return nil
}

// NewRankEvalRequestItem returns a RankEvalRequestItem.
func NewRankEvalRequestItem() *RankEvalRequestItem {
	r := &RankEvalRequestItem{
		Params: make(map[string]json.RawMessage),
	}

	return r
}

type RankEvalRequestItemVariant interface {
	RankEvalRequestItemCaster() *RankEvalRequestItem
}

func (s *RankEvalRequestItem) RankEvalRequestItemCaster() *RankEvalRequestItem {
	return s
}
