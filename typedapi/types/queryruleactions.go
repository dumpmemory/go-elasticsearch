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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package types

// QueryRuleActions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/query_rules/_types/QueryRuleset.ts#L110-L126
type QueryRuleActions struct {
	// Docs The documents to apply the rule to.
	// Only one of `ids` or `docs` may be specified and at least one must be
	// specified.
	// There is a maximum value of 100 documents in a rule.
	// You can specify the following attributes for each document:
	//
	// * `_index`: The index of the document to pin.
	// * `_id`: The unique document ID.
	Docs []PinnedDoc `json:"docs,omitempty"`
	// Ids The unique document IDs of the documents to apply the rule to.
	// Only one of `ids` or `docs` may be specified and at least one must be
	// specified.
	Ids []string `json:"ids,omitempty"`
}

// NewQueryRuleActions returns a QueryRuleActions.
func NewQueryRuleActions() *QueryRuleActions {
	r := &QueryRuleActions{}

	return r
}

// true

type QueryRuleActionsVariant interface {
	QueryRuleActionsCaster() *QueryRuleActions
}

func (s *QueryRuleActions) QueryRuleActionsCaster() *QueryRuleActions {
	return s
}
