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

package platform_infrastructure

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

// NewDeleteProxiesFilteredGroupParams creates a new DeleteProxiesFilteredGroupParams object
// with the default values initialized.
func NewDeleteProxiesFilteredGroupParams() *DeleteProxiesFilteredGroupParams {
	var ()
	return &DeleteProxiesFilteredGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProxiesFilteredGroupParamsWithTimeout creates a new DeleteProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProxiesFilteredGroupParamsWithTimeout(timeout time.Duration) *DeleteProxiesFilteredGroupParams {
	var ()
	return &DeleteProxiesFilteredGroupParams{

		timeout: timeout,
	}
}

// NewDeleteProxiesFilteredGroupParamsWithContext creates a new DeleteProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProxiesFilteredGroupParamsWithContext(ctx context.Context) *DeleteProxiesFilteredGroupParams {
	var ()
	return &DeleteProxiesFilteredGroupParams{

		Context: ctx,
	}
}

// NewDeleteProxiesFilteredGroupParamsWithHTTPClient creates a new DeleteProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProxiesFilteredGroupParamsWithHTTPClient(client *http.Client) *DeleteProxiesFilteredGroupParams {
	var ()
	return &DeleteProxiesFilteredGroupParams{
		HTTPClient: client,
	}
}

/*DeleteProxiesFilteredGroupParams contains all the parameters to send to the API endpoint
for the delete proxies filtered group operation typically these are written to a http.Request
*/
type DeleteProxiesFilteredGroupParams struct {

	/*ProxiesFilteredGroupID
	  "The identifier for the filtered group of proxies

	*/
	ProxiesFilteredGroupID string
	/*Version
	  Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) WithTimeout(timeout time.Duration) *DeleteProxiesFilteredGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) WithContext(ctx context.Context) *DeleteProxiesFilteredGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) WithHTTPClient(client *http.Client) *DeleteProxiesFilteredGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProxiesFilteredGroupID adds the proxiesFilteredGroupID to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) WithProxiesFilteredGroupID(proxiesFilteredGroupID string) *DeleteProxiesFilteredGroupParams {
	o.SetProxiesFilteredGroupID(proxiesFilteredGroupID)
	return o
}

// SetProxiesFilteredGroupID adds the proxiesFilteredGroupId to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) SetProxiesFilteredGroupID(proxiesFilteredGroupID string) {
	o.ProxiesFilteredGroupID = proxiesFilteredGroupID
}

// WithVersion adds the version to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) WithVersion(version *int64) *DeleteProxiesFilteredGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete proxies filtered group params
func (o *DeleteProxiesFilteredGroupParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProxiesFilteredGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proxies_filtered_group_id
	if err := r.SetPathParam("proxies_filtered_group_id", o.ProxiesFilteredGroupID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
