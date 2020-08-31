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

// RemoteClusterInfo Information about the remote cluster for cross-cluster search.
//
// swagger:model RemoteClusterInfo
type RemoteClusterInfo struct {

	// Flag indicating whether the remote cluster version is compatible
	// Required: true
	Compatible *bool `json:"compatible"`

	// The remote Elasticsearch cluster deployment id
	// Required: true
	ID *string `json:"id"`

	// Remote cluster Elasticsearch version
	Version string `json:"version,omitempty"`
}

// Validate validates this remote cluster info
func (m *RemoteClusterInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteClusterInfo) validateCompatible(formats strfmt.Registry) error {

	if err := validate.Required("compatible", "body", m.Compatible); err != nil {
		return err
	}

	return nil
}

func (m *RemoteClusterInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteClusterInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteClusterInfo) UnmarshalBinary(b []byte) error {
	var res RemoteClusterInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
