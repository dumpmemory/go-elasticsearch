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

// Hit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/hits.ts#L40-L65
type Hit struct {
	Explanation_       *Explanation                 `json:"_explanation,omitempty"`
	Fields             map[string]json.RawMessage   `json:"fields,omitempty"`
	Highlight          map[string][]string          `json:"highlight,omitempty"`
	Id_                *string                      `json:"_id,omitempty"`
	IgnoredFieldValues map[string][]json.RawMessage `json:"ignored_field_values,omitempty"`
	Ignored_           []string                     `json:"_ignored,omitempty"`
	Index_             string                       `json:"_index"`
	InnerHits          map[string]InnerHitsResult   `json:"inner_hits,omitempty"`
	MatchedQueries     any                          `json:"matched_queries,omitempty"`
	Nested_            *NestedIdentity              `json:"_nested,omitempty"`
	Node_              *string                      `json:"_node,omitempty"`
	PrimaryTerm_       *int64                       `json:"_primary_term,omitempty"`
	Rank_              *int                         `json:"_rank,omitempty"`
	Routing_           *string                      `json:"_routing,omitempty"`
	Score_             *Float64                     `json:"_score,omitempty"`
	SeqNo_             *int64                       `json:"_seq_no,omitempty"`
	Shard_             *string                      `json:"_shard,omitempty"`
	Sort               []FieldValue                 `json:"sort,omitempty"`
	Source_            json.RawMessage              `json:"_source,omitempty"`
	Version_           *int64                       `json:"_version,omitempty"`
}

func (s *Hit) UnmarshalJSON(data []byte) error {

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

		case "_explanation":
			if err := dec.Decode(&s.Explanation_); err != nil {
				return fmt.Errorf("%s | %w", "Explanation_", err)
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "highlight":
			if s.Highlight == nil {
				s.Highlight = make(map[string][]string, 0)
			}
			if err := dec.Decode(&s.Highlight); err != nil {
				return fmt.Errorf("%s | %w", "Highlight", err)
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "ignored_field_values":
			if s.IgnoredFieldValues == nil {
				s.IgnoredFieldValues = make(map[string][]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.IgnoredFieldValues); err != nil {
				return fmt.Errorf("%s | %w", "IgnoredFieldValues", err)
			}

		case "_ignored":
			if err := dec.Decode(&s.Ignored_); err != nil {
				return fmt.Errorf("%s | %w", "Ignored_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "inner_hits":
			if s.InnerHits == nil {
				s.InnerHits = make(map[string]InnerHitsResult, 0)
			}
			if err := dec.Decode(&s.InnerHits); err != nil {
				return fmt.Errorf("%s | %w", "InnerHits", err)
			}

		case "matched_queries":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]Float64, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "MatchedQueries", err)
				}
				s.MatchedQueries = o
			case '[':
				o := []string{}
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "MatchedQueries", err)
				}
				s.MatchedQueries = o
			}

		case "_nested":
			if err := dec.Decode(&s.Nested_); err != nil {
				return fmt.Errorf("%s | %w", "Nested_", err)
			}

		case "_node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Node_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node_ = &o

		case "_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryTerm_", err)
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = &f
			}

		case "_rank":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rank_", err)
				}
				s.Rank_ = &value
			case float64:
				f := int(v)
				s.Rank_ = &f
			}

		case "_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Routing_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Routing_ = &o

		case "_score":
			if err := dec.Decode(&s.Score_); err != nil {
				return fmt.Errorf("%s | %w", "Score_", err)
			}

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return fmt.Errorf("%s | %w", "SeqNo_", err)
			}

		case "_shard":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Shard_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shard_ = &o

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return fmt.Errorf("%s | %w", "Sort", err)
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		}
	}
	return nil
}

// NewHit returns a Hit.
func NewHit() *Hit {
	r := &Hit{
		Fields:             make(map[string]json.RawMessage),
		Highlight:          make(map[string][]string),
		IgnoredFieldValues: make(map[string][]json.RawMessage),
		InnerHits:          make(map[string]InnerHitsResult),
	}

	return r
}
