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

// ElasticsearchClusterSecurityInfo For 2.x Elasticsearch clusters, specifies the information about the users and roles. For 5.x Elasticsearch clusters, use the Kibana management UI.
//
// swagger:model ElasticsearchClusterSecurityInfo
type ElasticsearchClusterSecurityInfo struct {

	// The most recent time the security settings were changed (ISO format in UTC)
	// Required: true
	// Format: date-time
	LastModified *strfmt.DateTime `json:"last_modified"`

	// An arbitrarily nested JSON object mapping roles to sets of resources and permissions - see the Elasticsearch security documentation for more details on roles
	// Required: true
	Roles interface{} `json:"roles"`

	// users
	// Required: true
	Users []*ElasticsearchClusterUser `json:"users"`

	// users roles
	// Required: true
	UsersRoles []*ElasticsearchClusterRole `json:"users_roles"`

	// The resource version number of the security settings
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this elasticsearch cluster security info
func (m *ElasticsearchClusterSecurityInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterSecurityInfo) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("last_modified", "body", m.LastModified); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterSecurityInfo) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterSecurityInfo) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterSecurityInfo) validateUsersRoles(formats strfmt.Registry) error {

	if err := validate.Required("users_roles", "body", m.UsersRoles); err != nil {
		return err
	}

	for i := 0; i < len(m.UsersRoles); i++ {
		if swag.IsZero(m.UsersRoles[i]) { // not required
			continue
		}

		if m.UsersRoles[i] != nil {
			if err := m.UsersRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterSecurityInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterSecurityInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterSecurityInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterSecurityInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
