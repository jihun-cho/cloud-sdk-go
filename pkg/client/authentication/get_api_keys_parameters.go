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

package authentication

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

// NewGetAPIKeysParams creates a new GetAPIKeysParams object
// with the default values initialized.
func NewGetAPIKeysParams() *GetAPIKeysParams {

	return &GetAPIKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKeysParamsWithTimeout creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIKeysParamsWithTimeout(timeout time.Duration) *GetAPIKeysParams {

	return &GetAPIKeysParams{

		timeout: timeout,
	}
}

// NewGetAPIKeysParamsWithContext creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIKeysParamsWithContext(ctx context.Context) *GetAPIKeysParams {

	return &GetAPIKeysParams{

		Context: ctx,
	}
}

// NewGetAPIKeysParamsWithHTTPClient creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIKeysParamsWithHTTPClient(client *http.Client) *GetAPIKeysParams {

	return &GetAPIKeysParams{
		HTTPClient: client,
	}
}

/*GetAPIKeysParams contains all the parameters to send to the API endpoint
for the get api keys operation typically these are written to a http.Request
*/
type GetAPIKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get api keys params
func (o *GetAPIKeysParams) WithTimeout(timeout time.Duration) *GetAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get api keys params
func (o *GetAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get api keys params
func (o *GetAPIKeysParams) WithContext(ctx context.Context) *GetAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get api keys params
func (o *GetAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get api keys params
func (o *GetAPIKeysParams) WithHTTPClient(client *http.Client) *GetAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get api keys params
func (o *GetAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
