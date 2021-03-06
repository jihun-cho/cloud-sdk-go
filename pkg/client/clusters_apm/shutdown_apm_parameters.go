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

// NewShutdownApmParams creates a new ShutdownApmParams object
// with the default values initialized.
func NewShutdownApmParams() *ShutdownApmParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownApmParams{
		Hide: &hideDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShutdownApmParamsWithTimeout creates a new ShutdownApmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShutdownApmParamsWithTimeout(timeout time.Duration) *ShutdownApmParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownApmParams{
		Hide: &hideDefault,

		timeout: timeout,
	}
}

// NewShutdownApmParamsWithContext creates a new ShutdownApmParams object
// with the default values initialized, and the ability to set a context for a request
func NewShutdownApmParamsWithContext(ctx context.Context) *ShutdownApmParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownApmParams{
		Hide: &hideDefault,

		Context: ctx,
	}
}

// NewShutdownApmParamsWithHTTPClient creates a new ShutdownApmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShutdownApmParamsWithHTTPClient(client *http.Client) *ShutdownApmParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownApmParams{
		Hide:       &hideDefault,
		HTTPClient: client,
	}
}

/*ShutdownApmParams contains all the parameters to send to the API endpoint
for the shutdown apm operation typically these are written to a http.Request
*/
type ShutdownApmParams struct {

	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*Hide
	  Hides the clusters during shutdown. NOTE: By default, hidden clusters are not listed.

	*/
	Hide *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shutdown apm params
func (o *ShutdownApmParams) WithTimeout(timeout time.Duration) *ShutdownApmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shutdown apm params
func (o *ShutdownApmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shutdown apm params
func (o *ShutdownApmParams) WithContext(ctx context.Context) *ShutdownApmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shutdown apm params
func (o *ShutdownApmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shutdown apm params
func (o *ShutdownApmParams) WithHTTPClient(client *http.Client) *ShutdownApmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shutdown apm params
func (o *ShutdownApmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the shutdown apm params
func (o *ShutdownApmParams) WithClusterID(clusterID string) *ShutdownApmParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the shutdown apm params
func (o *ShutdownApmParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithHide adds the hide to the shutdown apm params
func (o *ShutdownApmParams) WithHide(hide *bool) *ShutdownApmParams {
	o.SetHide(hide)
	return o
}

// SetHide adds the hide to the shutdown apm params
func (o *ShutdownApmParams) SetHide(hide *bool) {
	o.Hide = hide
}

// WriteToRequest writes these params to a swagger request
func (o *ShutdownApmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Hide != nil {

		// query param hide
		var qrHide bool
		if o.Hide != nil {
			qrHide = *o.Hide
		}
		qHide := swag.FormatBool(qrHide)
		if qHide != "" {
			if err := r.SetQueryParam("hide", qHide); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
