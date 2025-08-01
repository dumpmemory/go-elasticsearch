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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainingpriority"
)

// TrainedModelAssignmentTaskParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/TrainedModel.ts#L370-L406
type TrainedModelAssignmentTaskParameters struct {
	// CacheSize The size of the trained model cache.
	CacheSize ByteSize `json:"cache_size,omitempty"`
	// DeploymentId The unique identifier for the trained model deployment.
	DeploymentId string `json:"deployment_id"`
	// ModelBytes The size of the trained model in bytes.
	ModelBytes ByteSize `json:"model_bytes"`
	// ModelId The unique identifier for the trained model.
	ModelId string `json:"model_id"`
	// NumberOfAllocations The total number of allocations this model is assigned across ML nodes.
	NumberOfAllocations      int                               `json:"number_of_allocations"`
	PerAllocationMemoryBytes ByteSize                          `json:"per_allocation_memory_bytes"`
	PerDeploymentMemoryBytes ByteSize                          `json:"per_deployment_memory_bytes"`
	Priority                 trainingpriority.TrainingPriority `json:"priority"`
	// QueueCapacity Number of inference requests are allowed in the queue at a time.
	QueueCapacity int `json:"queue_capacity"`
	// ThreadsPerAllocation Number of threads per allocation.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
}

func (s *TrainedModelAssignmentTaskParameters) UnmarshalJSON(data []byte) error {

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

		case "cache_size":
			if err := dec.Decode(&s.CacheSize); err != nil {
				return fmt.Errorf("%s | %w", "CacheSize", err)
			}

		case "deployment_id":
			if err := dec.Decode(&s.DeploymentId); err != nil {
				return fmt.Errorf("%s | %w", "DeploymentId", err)
			}

		case "model_bytes":
			if err := dec.Decode(&s.ModelBytes); err != nil {
				return fmt.Errorf("%s | %w", "ModelBytes", err)
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return fmt.Errorf("%s | %w", "ModelId", err)
			}

		case "number_of_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfAllocations", err)
				}
				s.NumberOfAllocations = value
			case float64:
				f := int(v)
				s.NumberOfAllocations = f
			}

		case "per_allocation_memory_bytes":
			if err := dec.Decode(&s.PerAllocationMemoryBytes); err != nil {
				return fmt.Errorf("%s | %w", "PerAllocationMemoryBytes", err)
			}

		case "per_deployment_memory_bytes":
			if err := dec.Decode(&s.PerDeploymentMemoryBytes); err != nil {
				return fmt.Errorf("%s | %w", "PerDeploymentMemoryBytes", err)
			}

		case "priority":
			if err := dec.Decode(&s.Priority); err != nil {
				return fmt.Errorf("%s | %w", "Priority", err)
			}

		case "queue_capacity":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueueCapacity", err)
				}
				s.QueueCapacity = value
			case float64:
				f := int(v)
				s.QueueCapacity = f
			}

		case "threads_per_allocation":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ThreadsPerAllocation", err)
				}
				s.ThreadsPerAllocation = value
			case float64:
				f := int(v)
				s.ThreadsPerAllocation = f
			}

		}
	}
	return nil
}

// NewTrainedModelAssignmentTaskParameters returns a TrainedModelAssignmentTaskParameters.
func NewTrainedModelAssignmentTaskParameters() *TrainedModelAssignmentTaskParameters {
	r := &TrainedModelAssignmentTaskParameters{}

	return r
}
