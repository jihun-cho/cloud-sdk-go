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

// ElasticsearchClustersInfo Information about a set of Elasticsearch clusters.
//
// swagger:model ElasticsearchClustersInfo
type ElasticsearchClustersInfo struct {

	// elasticsearch clusters
	// Required: true
	ElasticsearchClusters []*ElasticsearchClusterInfo `json:"elasticsearch_clusters"`

	// If a query is supplied, then the total number of clusters that matched
	MatchCount int32 `json:"match_count,omitempty"`

	// The number of clusters actually returned
	// Required: true
	ReturnCount *int32 `json:"return_count"`
}

// Validate validates this elasticsearch clusters info
func (m *ElasticsearchClustersInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElasticsearchClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClustersInfo) validateElasticsearchClusters(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch_clusters", "body", m.ElasticsearchClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.ElasticsearchClusters); i++ {
		if swag.IsZero(m.ElasticsearchClusters[i]) { // not required
			continue
		}

		if m.ElasticsearchClusters[i] != nil {
			if err := m.ElasticsearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClustersInfo) validateReturnCount(formats strfmt.Registry) error {

	if err := validate.Required("return_count", "body", m.ReturnCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClustersInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClustersInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClustersInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
