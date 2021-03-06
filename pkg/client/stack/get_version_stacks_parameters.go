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
	"github.com/go-openapi/swag"
)

// NewGetVersionStacksParams creates a new GetVersionStacksParams object
// with the default values initialized.
func NewGetVersionStacksParams() *GetVersionStacksParams {
	var (
		showDeletedDefault  = bool(false)
		showUnusableDefault = bool(false)
	)
	return &GetVersionStacksParams{
		ShowDeleted:  &showDeletedDefault,
		ShowUnusable: &showUnusableDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionStacksParamsWithTimeout creates a new GetVersionStacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionStacksParamsWithTimeout(timeout time.Duration) *GetVersionStacksParams {
	var (
		showDeletedDefault  = bool(false)
		showUnusableDefault = bool(false)
	)
	return &GetVersionStacksParams{
		ShowDeleted:  &showDeletedDefault,
		ShowUnusable: &showUnusableDefault,

		timeout: timeout,
	}
}

// NewGetVersionStacksParamsWithContext creates a new GetVersionStacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionStacksParamsWithContext(ctx context.Context) *GetVersionStacksParams {
	var (
		showDeletedDefault  = bool(false)
		showUnusableDefault = bool(false)
	)
	return &GetVersionStacksParams{
		ShowDeleted:  &showDeletedDefault,
		ShowUnusable: &showUnusableDefault,

		Context: ctx,
	}
}

// NewGetVersionStacksParamsWithHTTPClient creates a new GetVersionStacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersionStacksParamsWithHTTPClient(client *http.Client) *GetVersionStacksParams {
	var (
		showDeletedDefault  = bool(false)
		showUnusableDefault = bool(false)
	)
	return &GetVersionStacksParams{
		ShowDeleted:  &showDeletedDefault,
		ShowUnusable: &showUnusableDefault,
		HTTPClient:   client,
	}
}

/*GetVersionStacksParams contains all the parameters to send to the API endpoint
for the get version stacks operation typically these are written to a http.Request
*/
type GetVersionStacksParams struct {

	/*ShowDeleted
	  Whether to show deleted stack versions or not

	*/
	ShowDeleted *bool
	/*ShowUnusable
	  Whether to show versions that are unusable by the authenticated user

	*/
	ShowUnusable *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get version stacks params
func (o *GetVersionStacksParams) WithTimeout(timeout time.Duration) *GetVersionStacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version stacks params
func (o *GetVersionStacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version stacks params
func (o *GetVersionStacksParams) WithContext(ctx context.Context) *GetVersionStacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version stacks params
func (o *GetVersionStacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version stacks params
func (o *GetVersionStacksParams) WithHTTPClient(client *http.Client) *GetVersionStacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version stacks params
func (o *GetVersionStacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShowDeleted adds the showDeleted to the get version stacks params
func (o *GetVersionStacksParams) WithShowDeleted(showDeleted *bool) *GetVersionStacksParams {
	o.SetShowDeleted(showDeleted)
	return o
}

// SetShowDeleted adds the showDeleted to the get version stacks params
func (o *GetVersionStacksParams) SetShowDeleted(showDeleted *bool) {
	o.ShowDeleted = showDeleted
}

// WithShowUnusable adds the showUnusable to the get version stacks params
func (o *GetVersionStacksParams) WithShowUnusable(showUnusable *bool) *GetVersionStacksParams {
	o.SetShowUnusable(showUnusable)
	return o
}

// SetShowUnusable adds the showUnusable to the get version stacks params
func (o *GetVersionStacksParams) SetShowUnusable(showUnusable *bool) {
	o.ShowUnusable = showUnusable
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionStacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ShowDeleted != nil {

		// query param show_deleted
		var qrShowDeleted bool
		if o.ShowDeleted != nil {
			qrShowDeleted = *o.ShowDeleted
		}
		qShowDeleted := swag.FormatBool(qrShowDeleted)
		if qShowDeleted != "" {
			if err := r.SetQueryParam("show_deleted", qShowDeleted); err != nil {
				return err
			}
		}

	}

	if o.ShowUnusable != nil {

		// query param show_unusable
		var qrShowUnusable bool
		if o.ShowUnusable != nil {
			qrShowUnusable = *o.ShowUnusable
		}
		qShowUnusable := swag.FormatBool(qrShowUnusable)
		if qShowUnusable != "" {
			if err := r.SetQueryParam("show_unusable", qShowUnusable); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
