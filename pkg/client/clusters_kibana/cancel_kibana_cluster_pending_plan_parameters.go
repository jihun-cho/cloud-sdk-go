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

package clusters_kibana

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

// NewCancelKibanaClusterPendingPlanParams creates a new CancelKibanaClusterPendingPlanParams object
// with the default values initialized.
func NewCancelKibanaClusterPendingPlanParams() *CancelKibanaClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelKibanaClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelKibanaClusterPendingPlanParamsWithTimeout creates a new CancelKibanaClusterPendingPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelKibanaClusterPendingPlanParamsWithTimeout(timeout time.Duration) *CancelKibanaClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelKibanaClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		timeout: timeout,
	}
}

// NewCancelKibanaClusterPendingPlanParamsWithContext creates a new CancelKibanaClusterPendingPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelKibanaClusterPendingPlanParamsWithContext(ctx context.Context) *CancelKibanaClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelKibanaClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,

		Context: ctx,
	}
}

// NewCancelKibanaClusterPendingPlanParamsWithHTTPClient creates a new CancelKibanaClusterPendingPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelKibanaClusterPendingPlanParamsWithHTTPClient(client *http.Client) *CancelKibanaClusterPendingPlanParams {
	var (
		forceDeleteDefault   = bool(false)
		ignoreMissingDefault = bool(false)
	)
	return &CancelKibanaClusterPendingPlanParams{
		ForceDelete:   &forceDeleteDefault,
		IgnoreMissing: &ignoreMissingDefault,
		HTTPClient:    client,
	}
}

/*CancelKibanaClusterPendingPlanParams contains all the parameters to send to the API endpoint
for the cancel kibana cluster pending plan operation typically these are written to a http.Request
*/
type CancelKibanaClusterPendingPlanParams struct {

	/*ClusterID
	  The Kibana deployment identifier.

	*/
	ClusterID string
	/*ForceDelete
	  When `true`, deletes the pending plan instead of attempting a graceful cancellation. The default is `false`.

	*/
	ForceDelete *bool
	/*IgnoreMissing
	  When `true`, returns successfully, even when plans are pending. The default is `false`.

	*/
	IgnoreMissing *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithTimeout(timeout time.Duration) *CancelKibanaClusterPendingPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithContext(ctx context.Context) *CancelKibanaClusterPendingPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithHTTPClient(client *http.Client) *CancelKibanaClusterPendingPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithClusterID(clusterID string) *CancelKibanaClusterPendingPlanParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceDelete adds the forceDelete to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithForceDelete(forceDelete *bool) *CancelKibanaClusterPendingPlanParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithIgnoreMissing adds the ignoreMissing to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) WithIgnoreMissing(ignoreMissing *bool) *CancelKibanaClusterPendingPlanParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the cancel kibana cluster pending plan params
func (o *CancelKibanaClusterPendingPlanParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WriteToRequest writes these params to a swagger request
func (o *CancelKibanaClusterPendingPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ForceDelete != nil {

		// query param force_delete
		var qrForceDelete bool
		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {
			if err := r.SetQueryParam("force_delete", qForceDelete); err != nil {
				return err
			}
		}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
