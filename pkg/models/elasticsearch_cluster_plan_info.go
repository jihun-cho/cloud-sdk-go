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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchClusterPlanInfo Information about the Elasticsearch cluster plan.
// swagger:model ElasticsearchClusterPlanInfo
type ElasticsearchClusterPlanInfo struct {

	// If this plan completed or failed (ie is not pending), when the attempt ended (ISO format in UTC)
	// Format: date-time
	AttemptEndTime strfmt.DateTime `json:"attempt_end_time,omitempty"`

	// When this plan attempt (ie to apply the plan to the cluster) started (ISO format in UTC)
	// Required: true
	// Format: date-time
	AttemptStartTime *strfmt.DateTime `json:"attempt_start_time"`

	// Either the plan ended successfully, or is not yet completed (and no errors have occurred)
	// Required: true
	Healthy *bool `json:"healthy"`

	// plan
	Plan *ElasticsearchClusterPlan `json:"plan,omitempty"`

	// A UUID for each plan attempt
	PlanAttemptID string `json:"plan_attempt_id,omitempty"`

	// plan attempt log
	// Required: true
	PlanAttemptLog []*ClusterPlanStepInfo `json:"plan_attempt_log"`

	// A human readable name for each plan attempt, only populated when retrieving plan histories
	PlanAttemptName string `json:"plan_attempt_name,omitempty"`

	// If this plan is not current or pending, when the plan was no longer active (ISO format in UTC)
	// Format: date-time
	PlanEndTime strfmt.DateTime `json:"plan_end_time,omitempty"`

	// Information describing the source that facilitated the plans current state
	Source *ChangeSourceInfo `json:"source,omitempty"`
}

// Validate validates this elasticsearch cluster plan info
func (m *ElasticsearchClusterPlanInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttemptEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttemptStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanAttemptLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterPlanInfo) validateAttemptEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AttemptEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("attempt_end_time", "body", "date-time", m.AttemptEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validateAttemptStartTime(formats strfmt.Registry) error {

	if err := validate.Required("attempt_start_time", "body", m.AttemptStartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("attempt_start_time", "body", "date-time", m.AttemptStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validatePlanAttemptLog(formats strfmt.Registry) error {

	if err := validate.Required("plan_attempt_log", "body", m.PlanAttemptLog); err != nil {
		return err
	}

	for i := 0; i < len(m.PlanAttemptLog); i++ {
		if swag.IsZero(m.PlanAttemptLog[i]) { // not required
			continue
		}

		if m.PlanAttemptLog[i] != nil {
			if err := m.PlanAttemptLog[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plan_attempt_log" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validatePlanEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("plan_end_time", "body", "date-time", m.PlanEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterPlanInfo) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterPlanInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterPlanInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterPlanInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}