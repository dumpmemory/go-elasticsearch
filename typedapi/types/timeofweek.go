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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/day"
)

// TimeOfWeek type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Schedule.ts#L116-L119
type TimeOfWeek struct {
	At []string  `json:"at"`
	On []day.Day `json:"on"`
}

// NewTimeOfWeek returns a TimeOfWeek.
func NewTimeOfWeek() *TimeOfWeek {
	r := &TimeOfWeek{}

	return r
}

type TimeOfWeekVariant interface {
	TimeOfWeekCaster() *TimeOfWeek
}

func (s *TimeOfWeek) TimeOfWeekCaster() *TimeOfWeek {
	return s
}
