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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveElasticsearchClusterDetails Information about the Elasticsearch cluster that you want to move.
//
// swagger:model MoveElasticsearchClusterDetails
type MoveElasticsearchClusterDetails struct {

	// If only validating the move, then the plan configuration that would be applied to the cluster.
	CalculatedPlan *TransientElasticsearchPlanConfiguration `json:"calculated_plan,omitempty"`

	// Identifier for the Elasticsearch cluster.
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// A list of errors that occurred if the attempt to move the cluster failed.
	Errors []*BasicFailedReplyElement `json:"errors"`
}

// Validate validates this move elasticsearch cluster details
func (m *MoveElasticsearchClusterDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveElasticsearchClusterDetails) validateCalculatedPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CalculatedPlan) { // not required
		return nil
	}

	if m.CalculatedPlan != nil {
		if err := m.CalculatedPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_plan")
			}
			return err
		}
	}

	return nil
}

func (m *MoveElasticsearchClusterDetails) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *MoveElasticsearchClusterDetails) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveElasticsearchClusterDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveElasticsearchClusterDetails) UnmarshalBinary(b []byte) error {
	var res MoveElasticsearchClusterDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
