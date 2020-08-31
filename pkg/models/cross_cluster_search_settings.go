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

// CrossClusterSearchSettings The configuration settings for the cross-cluster search.
//
// swagger:model CrossClusterSearchSettings
type CrossClusterSearchSettings struct {

	// Mapping of remote cluster references keyed by their respective aliases. Aliases must only contain letters, digits, dashes and underscores
	// Required: true
	RemoteClusters map[string]RemoteClusterRef `json:"remote_clusters"`
}

// Validate validates this cross cluster search settings
func (m *CrossClusterSearchSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossClusterSearchSettings) validateRemoteClusters(formats strfmt.Registry) error {

	for k := range m.RemoteClusters {

		if err := validate.Required("remote_clusters"+"."+k, "body", m.RemoteClusters[k]); err != nil {
			return err
		}
		if val, ok := m.RemoteClusters[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrossClusterSearchSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrossClusterSearchSettings) UnmarshalBinary(b []byte) error {
	var res CrossClusterSearchSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
