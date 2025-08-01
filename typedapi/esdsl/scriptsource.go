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

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _scriptSource struct {
	v types.ScriptSource
}

func NewScriptSource() *_scriptSource {
	return &_scriptSource{v: nil}
}

func (u *_scriptSource) String(string string) *_scriptSource {

	u.v = &string

	return u
}

func (u *_scriptSource) SearchRequestBody(searchrequestbody types.SearchRequestBodyVariant) *_scriptSource {

	u.v = searchrequestbody.SearchRequestBodyCaster()

	return u
}

// Interface implementation for SearchRequestBody in ScriptSource union
func (u *_searchRequestBody) ScriptSourceCaster() *types.ScriptSource {
	t := types.ScriptSource(u.v)
	return &t
}

func (u *_scriptSource) ScriptSourceCaster() *types.ScriptSource {
	return &u.v
}
