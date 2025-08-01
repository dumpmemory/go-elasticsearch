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

// AllocationRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cat/allocation/types.ts#L25-L99
type AllocationRecord struct {
	// DiskAvail Free disk space available to Elasticsearch.
	// Elasticsearch retrieves this metric from the node’s operating system.
	// Disk-based shard allocation uses this metric to assign shards to nodes based
	// on available disk space.
	DiskAvail ByteSize `json:"disk.avail,omitempty"`
	// DiskIndices Disk space used by the node’s shards. Does not include disk space for the
	// translog or unassigned shards.
	// IMPORTANT: This metric double-counts disk space for hard-linked files, such
	// as those created when shrinking, splitting, or cloning an index.
	DiskIndices ByteSize `json:"disk.indices,omitempty"`
	// DiskIndicesForecast Sum of shard size forecasts
	DiskIndicesForecast ByteSize `json:"disk.indices.forecast,omitempty"`
	// DiskPercent Total percentage of disk space in use. Calculated as `disk.used /
	// disk.total`.
	DiskPercent Percentage `json:"disk.percent,omitempty"`
	// DiskTotal Total disk space for the node, including in-use and available space.
	DiskTotal ByteSize `json:"disk.total,omitempty"`
	// DiskUsed Total disk space in use.
	// Elasticsearch retrieves this metric from the node’s operating system (OS).
	// The metric includes disk space for: Elasticsearch, including the translog and
	// unassigned shards; the node’s operating system; any other applications or
	// files on the node.
	// Unlike `disk.indices`, this metric does not double-count disk space for
	// hard-linked files.
	DiskUsed ByteSize `json:"disk.used,omitempty"`
	// Host Network host for the node. Set using the `network.host` setting.
	Host *string `json:"host,omitempty"`
	// Ip IP address and port for the node.
	Ip *string `json:"ip,omitempty"`
	// Node Name for the node. Set using the `node.name` setting.
	Node *string `json:"node,omitempty"`
	// NodeRole Node roles
	NodeRole *string `json:"node.role,omitempty"`
	// Shards Number of primary and replica shards assigned to the node.
	Shards *string `json:"shards,omitempty"`
	// ShardsUndesired Amount of shards that are scheduled to be moved elsewhere in the cluster or
	// -1 other than desired balance allocator is used
	ShardsUndesired *string `json:"shards.undesired,omitempty"`
	// WriteLoadForecast Sum of index write load forecasts
	WriteLoadForecast Stringifieddouble `json:"write_load.forecast,omitempty"`
}

func (s *AllocationRecord) UnmarshalJSON(data []byte) error {

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

		case "disk.avail", "da", "diskAvail":
			if err := dec.Decode(&s.DiskAvail); err != nil {
				return fmt.Errorf("%s | %w", "DiskAvail", err)
			}

		case "disk.indices", "di", "diskIndices":
			if err := dec.Decode(&s.DiskIndices); err != nil {
				return fmt.Errorf("%s | %w", "DiskIndices", err)
			}

		case "disk.indices.forecast", "dif", "diskIndicesForecast":
			if err := dec.Decode(&s.DiskIndicesForecast); err != nil {
				return fmt.Errorf("%s | %w", "DiskIndicesForecast", err)
			}

		case "disk.percent", "dp", "diskPercent":
			if err := dec.Decode(&s.DiskPercent); err != nil {
				return fmt.Errorf("%s | %w", "DiskPercent", err)
			}

		case "disk.total", "dt", "diskTotal":
			if err := dec.Decode(&s.DiskTotal); err != nil {
				return fmt.Errorf("%s | %w", "DiskTotal", err)
			}

		case "disk.used", "du", "diskUsed":
			if err := dec.Decode(&s.DiskUsed); err != nil {
				return fmt.Errorf("%s | %w", "DiskUsed", err)
			}

		case "host", "h":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return fmt.Errorf("%s | %w", "Ip", err)
			}

		case "node", "n":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = &o

		case "node.role", "r", "role", "nodeRole":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NodeRole", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeRole = &o

		case "shards", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shards = &o

		case "shards.undesired":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ShardsUndesired", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ShardsUndesired = &o

		case "write_load.forecast", "wlf", "writeLoadForecast":
			if err := dec.Decode(&s.WriteLoadForecast); err != nil {
				return fmt.Errorf("%s | %w", "WriteLoadForecast", err)
			}

		}
	}
	return nil
}

// NewAllocationRecord returns a AllocationRecord.
func NewAllocationRecord() *AllocationRecord {
	r := &AllocationRecord{}

	return r
}
