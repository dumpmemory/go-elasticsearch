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
)

// IndexSettingsTimeSeries type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L354-L357
type IndexSettingsTimeSeries struct {
	EndTime   DateTime `json:"end_time,omitempty"`
	StartTime DateTime `json:"start_time,omitempty"`
}

func (s *IndexSettingsTimeSeries) UnmarshalJSON(data []byte) error {

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

		case "end_time":
			if err := dec.Decode(&s.EndTime); err != nil {
				return fmt.Errorf("%s | %w", "EndTime", err)
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		}
	}
	return nil
}

// NewIndexSettingsTimeSeries returns a IndexSettingsTimeSeries.
func NewIndexSettingsTimeSeries() *IndexSettingsTimeSeries {
	r := &IndexSettingsTimeSeries{}

	return r
}

type IndexSettingsTimeSeriesVariant interface {
	IndexSettingsTimeSeriesCaster() *IndexSettingsTimeSeries
}

func (s *IndexSettingsTimeSeries) IndexSettingsTimeSeriesCaster() *IndexSettingsTimeSeries {
	return s
}
