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

package platform

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

// NewGetPlatformParams creates a new GetPlatformParams object
// with the default values initialized.
func NewGetPlatformParams() *GetPlatformParams {

	return &GetPlatformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformParamsWithTimeout creates a new GetPlatformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformParamsWithTimeout(timeout time.Duration) *GetPlatformParams {

	return &GetPlatformParams{

		timeout: timeout,
	}
}

// NewGetPlatformParamsWithContext creates a new GetPlatformParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformParamsWithContext(ctx context.Context) *GetPlatformParams {

	return &GetPlatformParams{

		Context: ctx,
	}
}

// NewGetPlatformParamsWithHTTPClient creates a new GetPlatformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformParamsWithHTTPClient(client *http.Client) *GetPlatformParams {

	return &GetPlatformParams{
		HTTPClient: client,
	}
}

/*GetPlatformParams contains all the parameters to send to the API endpoint
for the get platform operation typically these are written to a http.Request
*/
type GetPlatformParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform params
func (o *GetPlatformParams) WithTimeout(timeout time.Duration) *GetPlatformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform params
func (o *GetPlatformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform params
func (o *GetPlatformParams) WithContext(ctx context.Context) *GetPlatformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform params
func (o *GetPlatformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform params
func (o *GetPlatformParams) WithHTTPClient(client *http.Client) *GetPlatformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform params
func (o *GetPlatformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
