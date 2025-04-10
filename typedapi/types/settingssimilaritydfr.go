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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfraftereffect"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfrbasicmodel"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

// SettingsSimilarityDfr type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/indices/_types/IndexSettings.ts#L208-L213
type SettingsSimilarityDfr struct {
	AfterEffect   dfraftereffect.DFRAfterEffect `json:"after_effect"`
	BasicModel    dfrbasicmodel.DFRBasicModel   `json:"basic_model"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityDfr) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityDfr SettingsSimilarityDfr
	tmp := innerSettingsSimilarityDfr{
		AfterEffect:   s.AfterEffect,
		BasicModel:    s.BasicModel,
		Normalization: s.Normalization,
		Type:          s.Type,
	}

	tmp.Type = "DFR"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityDfr returns a SettingsSimilarityDfr.
func NewSettingsSimilarityDfr() *SettingsSimilarityDfr {
	r := &SettingsSimilarityDfr{}

	return r
}

// true

type SettingsSimilarityDfrVariant interface {
	SettingsSimilarityDfrCaster() *SettingsSimilarityDfr
}

func (s *SettingsSimilarityDfr) SettingsSimilarityDfrCaster() *SettingsSimilarityDfr {
	return s
}
