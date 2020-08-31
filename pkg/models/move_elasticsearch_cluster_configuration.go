// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveElasticsearchClusterConfiguration The configuration object for moving Elasticsearch clusters.
//
// swagger:model MoveElasticsearchClusterConfiguration
type MoveElasticsearchClusterConfiguration struct {

	// Identifiers for the Elasticsearch clusters.
	// Required: true
	ClusterIds []string `json:"cluster_ids"`

	// Plan override to apply to the Elasticsearch clusters being moved.
	PlanOverride *TransientElasticsearchPlanConfiguration `json:"plan_override,omitempty"`
}

// Validate validates this move elasticsearch cluster configuration
func (m *MoveElasticsearchClusterConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanOverride(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveElasticsearchClusterConfiguration) validateClusterIds(formats strfmt.Registry) error {

	if err := validate.Required("cluster_ids", "body", m.ClusterIds); err != nil {
		return err
	}

	return nil
}

func (m *MoveElasticsearchClusterConfiguration) validatePlanOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanOverride) { // not required
		return nil
	}

	if m.PlanOverride != nil {
		if err := m.PlanOverride.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_override")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveElasticsearchClusterConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveElasticsearchClusterConfiguration) UnmarshalBinary(b []byte) error {
	var res MoveElasticsearchClusterConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
