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

// Package functionscoremode
package functionscoremode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/compound.ts#L268-L293
type FunctionScoreMode struct {
	Name string
}

var (
	Multiply = FunctionScoreMode{"multiply"}

	Sum = FunctionScoreMode{"sum"}

	Avg = FunctionScoreMode{"avg"}

	First = FunctionScoreMode{"first"}

	Max = FunctionScoreMode{"max"}

	Min = FunctionScoreMode{"min"}
)

func (f FunctionScoreMode) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FunctionScoreMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "multiply":
		*f = Multiply
	case "sum":
		*f = Sum
	case "avg":
		*f = Avg
	case "first":
		*f = First
	case "max":
		*f = Max
	case "min":
		*f = Min
	default:
		*f = FunctionScoreMode{string(text)}
	}

	return nil
}

func (f FunctionScoreMode) String() string {
	return f.Name
}
