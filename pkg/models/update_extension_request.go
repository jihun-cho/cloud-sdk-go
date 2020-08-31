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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateExtensionRequest The body of a request to update an extension
//
// swagger:model UpdateExtensionRequest
type UpdateExtensionRequest struct {

	// The extension description.
	Description string `json:"description,omitempty"`

	// The URL to download the extension archive.
	DownloadURL string `json:"download_url,omitempty"`

	// The extension type.
	// Required: true
	// Enum: [plugin bundle]
	ExtensionType *string `json:"extension_type"`

	// The extension name.
	// Required: true
	Name *string `json:"name"`

	// The Elasticsearch version.
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this update extension request
func (m *UpdateExtensionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

var updateExtensionRequestTypeExtensionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["plugin","bundle"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateExtensionRequestTypeExtensionTypePropEnum = append(updateExtensionRequestTypeExtensionTypePropEnum, v)
	}
}

const (

	// UpdateExtensionRequestExtensionTypePlugin captures enum value "plugin"
	UpdateExtensionRequestExtensionTypePlugin string = "plugin"

	// UpdateExtensionRequestExtensionTypeBundle captures enum value "bundle"
	UpdateExtensionRequestExtensionTypeBundle string = "bundle"
)

// prop value enum
func (m *UpdateExtensionRequest) validateExtensionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateExtensionRequestTypeExtensionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateExtensionRequest) validateExtensionType(formats strfmt.Registry) error {

	if err := validate.Required("extension_type", "body", m.ExtensionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateExtensionTypeEnum("extension_type", "body", *m.ExtensionType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExtensionRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExtensionRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateExtensionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateExtensionRequest) UnmarshalBinary(b []byte) error {
	var res UpdateExtensionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
