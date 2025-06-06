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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _shingleTokenFilter struct {
	v *types.ShingleTokenFilter
}

func NewShingleTokenFilter() *_shingleTokenFilter {

	return &_shingleTokenFilter{v: types.NewShingleTokenFilter()}

}

func (s *_shingleTokenFilter) FillerToken(fillertoken string) *_shingleTokenFilter {

	s.v.FillerToken = &fillertoken

	return s
}

func (s *_shingleTokenFilter) MaxShingleSize(maxshinglesize string) *_shingleTokenFilter {

	s.v.MaxShingleSize = maxshinglesize

	return s
}

func (s *_shingleTokenFilter) MinShingleSize(minshinglesize string) *_shingleTokenFilter {

	s.v.MinShingleSize = minshinglesize

	return s
}

func (s *_shingleTokenFilter) OutputUnigrams(outputunigrams bool) *_shingleTokenFilter {

	s.v.OutputUnigrams = &outputunigrams

	return s
}

func (s *_shingleTokenFilter) OutputUnigramsIfNoShingles(outputunigramsifnoshingles bool) *_shingleTokenFilter {

	s.v.OutputUnigramsIfNoShingles = &outputunigramsifnoshingles

	return s
}

func (s *_shingleTokenFilter) TokenSeparator(tokenseparator string) *_shingleTokenFilter {

	s.v.TokenSeparator = &tokenseparator

	return s
}

func (s *_shingleTokenFilter) Version(versionstring string) *_shingleTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_shingleTokenFilter) ShingleTokenFilterCaster() *types.ShingleTokenFilter {
	return s.v
}
