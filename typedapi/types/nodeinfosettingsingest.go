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

// NodeInfoSettingsIngest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/info/types.ts#L97-L132
type NodeInfoSettingsIngest struct {
	Append          *NodeInfoIngestInfo `json:"append,omitempty"`
	Attachment      *NodeInfoIngestInfo `json:"attachment,omitempty"`
	Bytes           *NodeInfoIngestInfo `json:"bytes,omitempty"`
	Circle          *NodeInfoIngestInfo `json:"circle,omitempty"`
	Convert         *NodeInfoIngestInfo `json:"convert,omitempty"`
	Csv             *NodeInfoIngestInfo `json:"csv,omitempty"`
	Date            *NodeInfoIngestInfo `json:"date,omitempty"`
	DateIndexName   *NodeInfoIngestInfo `json:"date_index_name,omitempty"`
	Dissect         *NodeInfoIngestInfo `json:"dissect,omitempty"`
	DotExpander     *NodeInfoIngestInfo `json:"dot_expander,omitempty"`
	Drop            *NodeInfoIngestInfo `json:"drop,omitempty"`
	Enrich          *NodeInfoIngestInfo `json:"enrich,omitempty"`
	Fail            *NodeInfoIngestInfo `json:"fail,omitempty"`
	Foreach         *NodeInfoIngestInfo `json:"foreach,omitempty"`
	Geoip           *NodeInfoIngestInfo `json:"geoip,omitempty"`
	Grok            *NodeInfoIngestInfo `json:"grok,omitempty"`
	Gsub            *NodeInfoIngestInfo `json:"gsub,omitempty"`
	Inference       *NodeInfoIngestInfo `json:"inference,omitempty"`
	Join            *NodeInfoIngestInfo `json:"join,omitempty"`
	Json            *NodeInfoIngestInfo `json:"json,omitempty"`
	Kv              *NodeInfoIngestInfo `json:"kv,omitempty"`
	Lowercase       *NodeInfoIngestInfo `json:"lowercase,omitempty"`
	Pipeline        *NodeInfoIngestInfo `json:"pipeline,omitempty"`
	Remove          *NodeInfoIngestInfo `json:"remove,omitempty"`
	Rename          *NodeInfoIngestInfo `json:"rename,omitempty"`
	Script          *NodeInfoIngestInfo `json:"script,omitempty"`
	Set             *NodeInfoIngestInfo `json:"set,omitempty"`
	SetSecurityUser *NodeInfoIngestInfo `json:"set_security_user,omitempty"`
	Sort            *NodeInfoIngestInfo `json:"sort,omitempty"`
	Split           *NodeInfoIngestInfo `json:"split,omitempty"`
	Trim            *NodeInfoIngestInfo `json:"trim,omitempty"`
	Uppercase       *NodeInfoIngestInfo `json:"uppercase,omitempty"`
	Urldecode       *NodeInfoIngestInfo `json:"urldecode,omitempty"`
	UserAgent       *NodeInfoIngestInfo `json:"user_agent,omitempty"`
}

// NewNodeInfoSettingsIngest returns a NodeInfoSettingsIngest.
func NewNodeInfoSettingsIngest() *NodeInfoSettingsIngest {
	r := &NodeInfoSettingsIngest{}

	return r
}
