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

// This is provide all the types that are part of the union.
type _geoBounds struct {
	v types.GeoBounds
}

func NewGeoBounds() *_geoBounds {
	return &_geoBounds{v: nil}
}

func (u *_geoBounds) CoordsGeoBounds(coordsgeobounds types.CoordsGeoBoundsVariant) *_geoBounds {

	u.v = &coordsgeobounds

	return u
}

// Interface implementation for CoordsGeoBounds in GeoBounds union
func (u *_coordsGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	t := types.GeoBounds(u.v)
	return &t
}

func (u *_geoBounds) TopLeftBottomRightGeoBounds(topleftbottomrightgeobounds types.TopLeftBottomRightGeoBoundsVariant) *_geoBounds {

	u.v = &topleftbottomrightgeobounds

	return u
}

// Interface implementation for TopLeftBottomRightGeoBounds in GeoBounds union
func (u *_topLeftBottomRightGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	t := types.GeoBounds(u.v)
	return &t
}

func (u *_geoBounds) TopRightBottomLeftGeoBounds(toprightbottomleftgeobounds types.TopRightBottomLeftGeoBoundsVariant) *_geoBounds {

	u.v = &toprightbottomleftgeobounds

	return u
}

// Interface implementation for TopRightBottomLeftGeoBounds in GeoBounds union
func (u *_topRightBottomLeftGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	t := types.GeoBounds(u.v)
	return &t
}

func (u *_geoBounds) WktGeoBounds(wktgeobounds types.WktGeoBoundsVariant) *_geoBounds {

	u.v = &wktgeobounds

	return u
}

// Interface implementation for WktGeoBounds in GeoBounds union
func (u *_wktGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	t := types.GeoBounds(u.v)
	return &t
}

func (u *_geoBounds) GeoBoundsCaster() *types.GeoBounds {
	return &u.v
}
