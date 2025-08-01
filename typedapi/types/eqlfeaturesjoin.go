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

// EqlFeaturesJoin type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L106-L112
type EqlFeaturesJoin struct {
	JoinQueriesFiveOrMore uint `json:"join_queries_five_or_more"`
	JoinQueriesFour       uint `json:"join_queries_four"`
	JoinQueriesThree      uint `json:"join_queries_three"`
	JoinQueriesTwo        uint `json:"join_queries_two"`
	JoinUntil             uint `json:"join_until"`
}

// NewEqlFeaturesJoin returns a EqlFeaturesJoin.
func NewEqlFeaturesJoin() *EqlFeaturesJoin {
	r := &EqlFeaturesJoin{}

	return r
}
