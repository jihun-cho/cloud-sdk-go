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
)

// NewStartConstructorMaintenanceModeParams creates a new StartConstructorMaintenanceModeParams object
// with the default values initialized.
func NewStartConstructorMaintenanceModeParams() *StartConstructorMaintenanceModeParams {
	var ()
	return &StartConstructorMaintenanceModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartConstructorMaintenanceModeParamsWithTimeout creates a new StartConstructorMaintenanceModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartConstructorMaintenanceModeParamsWithTimeout(timeout time.Duration) *StartConstructorMaintenanceModeParams {
	var ()
	return &StartConstructorMaintenanceModeParams{

		timeout: timeout,
	}
}

// NewStartConstructorMaintenanceModeParamsWithContext creates a new StartConstructorMaintenanceModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartConstructorMaintenanceModeParamsWithContext(ctx context.Context) *StartConstructorMaintenanceModeParams {
	var ()
	return &StartConstructorMaintenanceModeParams{

		Context: ctx,
	}
}

// NewStartConstructorMaintenanceModeParamsWithHTTPClient creates a new StartConstructorMaintenanceModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartConstructorMaintenanceModeParamsWithHTTPClient(client *http.Client) *StartConstructorMaintenanceModeParams {
	var ()
	return &StartConstructorMaintenanceModeParams{
		HTTPClient: client,
	}
}

/*StartConstructorMaintenanceModeParams contains all the parameters to send to the API endpoint
for the start constructor maintenance mode operation typically these are written to a http.Request
*/
type StartConstructorMaintenanceModeParams struct {

	/*ConstructorID
	  Identifier for the constructor

	*/
	ConstructorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) WithTimeout(timeout time.Duration) *StartConstructorMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) WithContext(ctx context.Context) *StartConstructorMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) WithHTTPClient(client *http.Client) *StartConstructorMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstructorID adds the constructorID to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) WithConstructorID(constructorID string) *StartConstructorMaintenanceModeParams {
	o.SetConstructorID(constructorID)
	return o
}

// SetConstructorID adds the constructorId to the start constructor maintenance mode params
func (o *StartConstructorMaintenanceModeParams) SetConstructorID(constructorID string) {
	o.ConstructorID = constructorID
}

// WriteToRequest writes these params to a swagger request
func (o *StartConstructorMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param constructor_id
	if err := r.SetPathParam("constructor_id", o.ConstructorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
