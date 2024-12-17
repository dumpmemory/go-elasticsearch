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
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

// Package tasktype
package tasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/inference/_types/TaskType.ts#L20-L28
type TaskType struct {
	Name string
}

var (
	Sparseembedding = TaskType{"sparse_embedding"}

	Textembedding = TaskType{"text_embedding"}

	Rerank = TaskType{"rerank"}

	Completion = TaskType{"completion"}
)

func (t TaskType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "sparse_embedding":
		*t = Sparseembedding
	case "text_embedding":
		*t = Textembedding
	case "rerank":
		*t = Rerank
	case "completion":
		*t = Completion
	default:
		*t = TaskType{string(text)}
	}

	return nil
}

func (t TaskType) String() string {
	return t.Name
}