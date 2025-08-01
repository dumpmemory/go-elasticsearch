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

// SoftDeletes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L51-L64
type SoftDeletes struct {
	// Enabled Indicates whether soft deletes are enabled on the index.
	Enabled *bool `json:"enabled,omitempty"`
	// RetentionLease The maximum period to retain a shard history retention lease before it is
	// considered expired.
	// Shard history retention leases ensure that soft deletes are retained during
	// merges on the Lucene
	// index. If a soft delete is merged away before it can be replicated to a
	// follower the following
	// process will fail due to incomplete history on the leader.
	RetentionLease *RetentionLease `json:"retention_lease,omitempty"`
}

func (s *SoftDeletes) UnmarshalJSON(data []byte) error {

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

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "retention_lease":
			if err := dec.Decode(&s.RetentionLease); err != nil {
				return fmt.Errorf("%s | %w", "RetentionLease", err)
			}

		}
	}
	return nil
}

// NewSoftDeletes returns a SoftDeletes.
func NewSoftDeletes() *SoftDeletes {
	r := &SoftDeletes{}

	return r
}

type SoftDeletesVariant interface {
	SoftDeletesCaster() *SoftDeletes
}

func (s *SoftDeletes) SoftDeletesCaster() *SoftDeletes {
	return s
}
