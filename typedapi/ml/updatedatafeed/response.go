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

package updatedatafeed

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package updatedatafeed
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/update_datafeed/MlUpdateDatafeedResponse.ts#L31-L49
type Response struct {
	Aggregations           map[string]types.Aggregations `json:"aggregations,omitempty"`
	Authorization          *types.DatafeedAuthorization  `json:"authorization,omitempty"`
	ChunkingConfig         types.ChunkingConfig          `json:"chunking_config"`
	DatafeedId             string                        `json:"datafeed_id"`
	DelayedDataCheckConfig *types.DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`
	Frequency              types.Duration                `json:"frequency,omitempty"`
	Indices                []string                      `json:"indices"`
	IndicesOptions         *types.IndicesOptions         `json:"indices_options,omitempty"`
	JobId                  string                        `json:"job_id"`
	MaxEmptySearches       *int                          `json:"max_empty_searches,omitempty"`
	Query                  types.Query                   `json:"query"`
	QueryDelay             types.Duration                `json:"query_delay"`
	RuntimeMappings        types.RuntimeFields           `json:"runtime_mappings,omitempty"`
	ScriptFields           map[string]types.ScriptField  `json:"script_fields,omitempty"`
	ScrollSize             int                           `json:"scroll_size"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Aggregations: make(map[string]types.Aggregations, 0),
		ScriptFields: make(map[string]types.ScriptField, 0),
	}
	return r
}
