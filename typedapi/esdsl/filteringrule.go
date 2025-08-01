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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringpolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringrulerule"
)

type _filteringRule struct {
	v *types.FilteringRule
}

func NewFilteringRule(order int, policy filteringpolicy.FilteringPolicy, rule filteringrulerule.FilteringRuleRule, value string) *_filteringRule {

	tmp := &_filteringRule{v: types.NewFilteringRule()}

	tmp.Order(order)

	tmp.Policy(policy)

	tmp.Rule(rule)

	tmp.Value(value)

	return tmp

}

func (s *_filteringRule) CreatedAt(datetime types.DateTimeVariant) *_filteringRule {

	s.v.CreatedAt = *datetime.DateTimeCaster()

	return s
}

func (s *_filteringRule) Field(field string) *_filteringRule {

	s.v.Field = field

	return s
}

func (s *_filteringRule) Id(id string) *_filteringRule {

	s.v.Id = id

	return s
}

func (s *_filteringRule) Order(order int) *_filteringRule {

	s.v.Order = order

	return s
}

func (s *_filteringRule) Policy(policy filteringpolicy.FilteringPolicy) *_filteringRule {

	s.v.Policy = policy
	return s
}

func (s *_filteringRule) Rule(rule filteringrulerule.FilteringRuleRule) *_filteringRule {

	s.v.Rule = rule
	return s
}

func (s *_filteringRule) UpdatedAt(datetime types.DateTimeVariant) *_filteringRule {

	s.v.UpdatedAt = *datetime.DateTimeCaster()

	return s
}

func (s *_filteringRule) Value(value string) *_filteringRule {

	s.v.Value = value

	return s
}

func (s *_filteringRule) FilteringRuleCaster() *types.FilteringRule {
	return s.v
}
