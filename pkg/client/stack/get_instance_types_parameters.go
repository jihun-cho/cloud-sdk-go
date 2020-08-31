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

package stack

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
)

// NewGetInstanceTypesParams creates a new GetInstanceTypesParams object
// with the default values initialized.
func NewGetInstanceTypesParams() *GetInstanceTypesParams {

	return &GetInstanceTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceTypesParamsWithTimeout creates a new GetInstanceTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceTypesParamsWithTimeout(timeout time.Duration) *GetInstanceTypesParams {

	return &GetInstanceTypesParams{

		timeout: timeout,
	}
}

// NewGetInstanceTypesParamsWithContext creates a new GetInstanceTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceTypesParamsWithContext(ctx context.Context) *GetInstanceTypesParams {

	return &GetInstanceTypesParams{

		Context: ctx,
	}
}

// NewGetInstanceTypesParamsWithHTTPClient creates a new GetInstanceTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceTypesParamsWithHTTPClient(client *http.Client) *GetInstanceTypesParams {

	return &GetInstanceTypesParams{
		HTTPClient: client,
	}
}

/*GetInstanceTypesParams contains all the parameters to send to the API endpoint
for the get instance types operation typically these are written to a http.Request
*/
type GetInstanceTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance types params
func (o *GetInstanceTypesParams) WithTimeout(timeout time.Duration) *GetInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance types params
func (o *GetInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance types params
func (o *GetInstanceTypesParams) WithContext(ctx context.Context) *GetInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance types params
func (o *GetInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance types params
func (o *GetInstanceTypesParams) WithHTTPClient(client *http.Client) *GetInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance types params
func (o *GetInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
