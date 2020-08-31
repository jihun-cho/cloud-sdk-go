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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewStartApmMaintenanceModeParams creates a new StartApmMaintenanceModeParams object
// with the default values initialized.
func NewStartApmMaintenanceModeParams() *StartApmMaintenanceModeParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StartApmMaintenanceModeParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStartApmMaintenanceModeParamsWithTimeout creates a new StartApmMaintenanceModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartApmMaintenanceModeParamsWithTimeout(timeout time.Duration) *StartApmMaintenanceModeParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StartApmMaintenanceModeParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: timeout,
	}
}

// NewStartApmMaintenanceModeParamsWithContext creates a new StartApmMaintenanceModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartApmMaintenanceModeParamsWithContext(ctx context.Context) *StartApmMaintenanceModeParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StartApmMaintenanceModeParams{
		IgnoreMissing: &ignoreMissingDefault,

		Context: ctx,
	}
}

// NewStartApmMaintenanceModeParamsWithHTTPClient creates a new StartApmMaintenanceModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartApmMaintenanceModeParamsWithHTTPClient(client *http.Client) *StartApmMaintenanceModeParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StartApmMaintenanceModeParams{
		IgnoreMissing: &ignoreMissingDefault,
		HTTPClient:    client,
	}
}

/*StartApmMaintenanceModeParams contains all the parameters to send to the API endpoint
for the start apm maintenance mode operation typically these are written to a http.Request
*/
type StartApmMaintenanceModeParams struct {

	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*IgnoreMissing
	  When `true` and the instance does not exist, proceeds to the next instance, or treats the instance as an error.

	*/
	IgnoreMissing *bool
	/*InstanceIds
	  A comma-separated list of instance identifiers.

	*/
	InstanceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithTimeout(timeout time.Duration) *StartApmMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithContext(ctx context.Context) *StartApmMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithHTTPClient(client *http.Client) *StartApmMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithClusterID(clusterID string) *StartApmMaintenanceModeParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithIgnoreMissing adds the ignoreMissing to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithIgnoreMissing(ignoreMissing *bool) *StartApmMaintenanceModeParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) WithInstanceIds(instanceIds []string) *StartApmMaintenanceModeParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the start apm maintenance mode params
func (o *StartApmMaintenanceModeParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WriteToRequest writes these params to a swagger request
func (o *StartApmMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	valuesInstanceIds := o.InstanceIds

	joinedInstanceIds := swag.JoinByFormat(valuesInstanceIds, "csv")
	// path array param instance_ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedInstanceIds) > 0 {
		if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
