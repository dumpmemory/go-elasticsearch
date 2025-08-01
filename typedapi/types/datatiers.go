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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataTiers type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L349-L359
type DataTiers struct {
	Available   bool                     `json:"available"`
	DataCold    DataTierPhaseStatistics  `json:"data_cold"`
	DataContent DataTierPhaseStatistics  `json:"data_content"`
	DataFrozen  *DataTierPhaseStatistics `json:"data_frozen,omitempty"`
	DataHot     DataTierPhaseStatistics  `json:"data_hot"`
	DataWarm    DataTierPhaseStatistics  `json:"data_warm"`
	Enabled     bool                     `json:"enabled"`
}

func (s *DataTiers) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "data_cold":
			if err := dec.Decode(&s.DataCold); err != nil {
				return fmt.Errorf("%s | %w", "DataCold", err)
			}

		case "data_content":
			if err := dec.Decode(&s.DataContent); err != nil {
				return fmt.Errorf("%s | %w", "DataContent", err)
			}

		case "data_frozen":
			if err := dec.Decode(&s.DataFrozen); err != nil {
				return fmt.Errorf("%s | %w", "DataFrozen", err)
			}

		case "data_hot":
			if err := dec.Decode(&s.DataHot); err != nil {
				return fmt.Errorf("%s | %w", "DataHot", err)
			}

		case "data_warm":
			if err := dec.Decode(&s.DataWarm); err != nil {
				return fmt.Errorf("%s | %w", "DataWarm", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		}
	}
	return nil
}

// NewDataTiers returns a DataTiers.
func NewDataTiers() *DataTiers {
	r := &DataTiers{}

	return r
}
