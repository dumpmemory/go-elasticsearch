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

// Package jinaaitextembeddingtask
package jinaaitextembeddingtask

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/inference/_types/CommonTypes.ts#L1508-L1513
type JinaAITextEmbeddingTask struct {
	Name string
}

var (
	Classification = JinaAITextEmbeddingTask{"classification"}

	Clustering = JinaAITextEmbeddingTask{"clustering"}

	Ingest = JinaAITextEmbeddingTask{"ingest"}

	Search = JinaAITextEmbeddingTask{"search"}
)

func (j JinaAITextEmbeddingTask) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JinaAITextEmbeddingTask) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "classification":
		*j = Classification
	case "clustering":
		*j = Clustering
	case "ingest":
		*j = Ingest
	case "search":
		*j = Search
	default:
		*j = JinaAITextEmbeddingTask{string(text)}
	}

	return nil
}

func (j JinaAITextEmbeddingTask) String() string {
	return j.Name
}
