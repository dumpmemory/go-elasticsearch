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

// TrainedModelTreeNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/put_trained_model/types.ts#L81-L91
type TrainedModelTreeNode struct {
	DecisionType *string  `json:"decision_type,omitempty"`
	DefaultLeft  *bool    `json:"default_left,omitempty"`
	LeafValue    *Float64 `json:"leaf_value,omitempty"`
	LeftChild    *int     `json:"left_child,omitempty"`
	NodeIndex    int      `json:"node_index"`
	RightChild   *int     `json:"right_child,omitempty"`
	SplitFeature *int     `json:"split_feature,omitempty"`
	SplitGain    *int     `json:"split_gain,omitempty"`
	Threshold    *Float64 `json:"threshold,omitempty"`
}

func (s *TrainedModelTreeNode) UnmarshalJSON(data []byte) error {

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

		case "decision_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DecisionType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DecisionType = &o

		case "default_left":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DefaultLeft", err)
				}
				s.DefaultLeft = &value
			case bool:
				s.DefaultLeft = &v
			}

		case "leaf_value":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LeafValue", err)
				}
				f := Float64(value)
				s.LeafValue = &f
			case float64:
				f := Float64(v)
				s.LeafValue = &f
			}

		case "left_child":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LeftChild", err)
				}
				s.LeftChild = &value
			case float64:
				f := int(v)
				s.LeftChild = &f
			}

		case "node_index":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NodeIndex", err)
				}
				s.NodeIndex = value
			case float64:
				f := int(v)
				s.NodeIndex = f
			}

		case "right_child":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RightChild", err)
				}
				s.RightChild = &value
			case float64:
				f := int(v)
				s.RightChild = &f
			}

		case "split_feature":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SplitFeature", err)
				}
				s.SplitFeature = &value
			case float64:
				f := int(v)
				s.SplitFeature = &f
			}

		case "split_gain":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SplitGain", err)
				}
				s.SplitGain = &value
			case float64:
				f := int(v)
				s.SplitGain = &f
			}

		case "threshold":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Threshold", err)
				}
				f := Float64(value)
				s.Threshold = &f
			case float64:
				f := Float64(v)
				s.Threshold = &f
			}

		}
	}
	return nil
}

// NewTrainedModelTreeNode returns a TrainedModelTreeNode.
func NewTrainedModelTreeNode() *TrainedModelTreeNode {
	r := &TrainedModelTreeNode{}

	return r
}

type TrainedModelTreeNodeVariant interface {
	TrainedModelTreeNodeCaster() *TrainedModelTreeNode
}

func (s *TrainedModelTreeNode) TrainedModelTreeNodeCaster() *TrainedModelTreeNode {
	return s
}
