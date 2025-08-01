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

// IndexingSlowlogSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L603-L608
type IndexingSlowlogSettings struct {
	Level     *string                   `json:"level,omitempty"`
	Reformat  *bool                     `json:"reformat,omitempty"`
	Source    *int                      `json:"source,omitempty"`
	Threshold *IndexingSlowlogTresholds `json:"threshold,omitempty"`
}

func (s *IndexingSlowlogSettings) UnmarshalJSON(data []byte) error {

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

		case "level":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Level", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Level = &o

		case "reformat":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Reformat", err)
				}
				s.Reformat = &value
			case bool:
				s.Reformat = &v
			}

		case "source":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Source", err)
				}
				s.Source = &value
			case float64:
				f := int(v)
				s.Source = &f
			}

		case "threshold":
			if err := dec.Decode(&s.Threshold); err != nil {
				return fmt.Errorf("%s | %w", "Threshold", err)
			}

		}
	}
	return nil
}

// NewIndexingSlowlogSettings returns a IndexingSlowlogSettings.
func NewIndexingSlowlogSettings() *IndexingSlowlogSettings {
	r := &IndexingSlowlogSettings{}

	return r
}

type IndexingSlowlogSettingsVariant interface {
	IndexingSlowlogSettingsCaster() *IndexingSlowlogSettings
}

func (s *IndexingSlowlogSettings) IndexingSlowlogSettingsCaster() *IndexingSlowlogSettings {
	return s
}
