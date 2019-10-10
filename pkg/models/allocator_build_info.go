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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AllocatorBuildInfo Information about the build that is running on the allocator.
// swagger:model AllocatorBuildInfo
type AllocatorBuildInfo struct {

	// Git commit hash
	CommitHash string `json:"commit_hash,omitempty"`

	// Version of the artifact
	Version string `json:"version,omitempty"`
}

// Validate validates this allocator build info
func (m *AllocatorBuildInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllocatorBuildInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatorBuildInfo) UnmarshalBinary(b []byte) error {
	var res AllocatorBuildInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}