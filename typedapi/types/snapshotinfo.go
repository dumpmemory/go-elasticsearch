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

// SnapshotInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotInfo.ts#L41-L71
type SnapshotInfo struct {
	DataStreams        []string                `json:"data_streams"`
	Duration           Duration                `json:"duration,omitempty"`
	DurationInMillis   *int64                  `json:"duration_in_millis,omitempty"`
	EndTime            DateTime                `json:"end_time,omitempty"`
	EndTimeInMillis    *int64                  `json:"end_time_in_millis,omitempty"`
	Failures           []SnapshotShardFailure  `json:"failures,omitempty"`
	FeatureStates      []InfoFeatureState      `json:"feature_states,omitempty"`
	IncludeGlobalState *bool                   `json:"include_global_state,omitempty"`
	IndexDetails       map[string]IndexDetails `json:"index_details,omitempty"`
	Indices            []string                `json:"indices,omitempty"`
	Metadata           Metadata                `json:"metadata,omitempty"`
	Reason             *string                 `json:"reason,omitempty"`
	Repository         *string                 `json:"repository,omitempty"`
	Shards             *ShardStatistics        `json:"shards,omitempty"`
	Snapshot           string                  `json:"snapshot"`
	StartTime          DateTime                `json:"start_time,omitempty"`
	StartTimeInMillis  *int64                  `json:"start_time_in_millis,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Uuid               string                  `json:"uuid"`
	Version            *string                 `json:"version,omitempty"`
	VersionId          *int64                  `json:"version_id,omitempty"`
}

func (s *SnapshotInfo) UnmarshalJSON(data []byte) error {

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

		case "data_streams":
			if err := dec.Decode(&s.DataStreams); err != nil {
				return fmt.Errorf("%s | %w", "DataStreams", err)
			}

		case "duration":
			if err := dec.Decode(&s.Duration); err != nil {
				return fmt.Errorf("%s | %w", "Duration", err)
			}

		case "duration_in_millis":
			if err := dec.Decode(&s.DurationInMillis); err != nil {
				return fmt.Errorf("%s | %w", "DurationInMillis", err)
			}

		case "end_time":
			if err := dec.Decode(&s.EndTime); err != nil {
				return fmt.Errorf("%s | %w", "EndTime", err)
			}

		case "end_time_in_millis":
			if err := dec.Decode(&s.EndTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "EndTimeInMillis", err)
			}

		case "failures":
			if err := dec.Decode(&s.Failures); err != nil {
				return fmt.Errorf("%s | %w", "Failures", err)
			}

		case "feature_states":
			if err := dec.Decode(&s.FeatureStates); err != nil {
				return fmt.Errorf("%s | %w", "FeatureStates", err)
			}

		case "include_global_state":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeGlobalState", err)
				}
				s.IncludeGlobalState = &value
			case bool:
				s.IncludeGlobalState = &v
			}

		case "index_details":
			if s.IndexDetails == nil {
				s.IndexDetails = make(map[string]IndexDetails, 0)
			}
			if err := dec.Decode(&s.IndexDetails); err != nil {
				return fmt.Errorf("%s | %w", "IndexDetails", err)
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "repository":
			if err := dec.Decode(&s.Repository); err != nil {
				return fmt.Errorf("%s | %w", "Repository", err)
			}

		case "shards":
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		case "snapshot":
			if err := dec.Decode(&s.Snapshot); err != nil {
				return fmt.Errorf("%s | %w", "Snapshot", err)
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		case "state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = &o

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_id":
			if err := dec.Decode(&s.VersionId); err != nil {
				return fmt.Errorf("%s | %w", "VersionId", err)
			}

		}
	}
	return nil
}

// NewSnapshotInfo returns a SnapshotInfo.
func NewSnapshotInfo() *SnapshotInfo {
	r := &SnapshotInfo{
		IndexDetails: make(map[string]IndexDetails),
	}

	return r
}
