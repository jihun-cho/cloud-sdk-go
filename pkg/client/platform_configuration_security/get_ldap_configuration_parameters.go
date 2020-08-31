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

package platform_configuration_security

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

// NewGetLdapConfigurationParams creates a new GetLdapConfigurationParams object
// with the default values initialized.
func NewGetLdapConfigurationParams() *GetLdapConfigurationParams {
	var ()
	return &GetLdapConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapConfigurationParamsWithTimeout creates a new GetLdapConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapConfigurationParamsWithTimeout(timeout time.Duration) *GetLdapConfigurationParams {
	var ()
	return &GetLdapConfigurationParams{

		timeout: timeout,
	}
}

// NewGetLdapConfigurationParamsWithContext creates a new GetLdapConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapConfigurationParamsWithContext(ctx context.Context) *GetLdapConfigurationParams {
	var ()
	return &GetLdapConfigurationParams{

		Context: ctx,
	}
}

// NewGetLdapConfigurationParamsWithHTTPClient creates a new GetLdapConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapConfigurationParamsWithHTTPClient(client *http.Client) *GetLdapConfigurationParams {
	var ()
	return &GetLdapConfigurationParams{
		HTTPClient: client,
	}
}

/*GetLdapConfigurationParams contains all the parameters to send to the API endpoint
for the get ldap configuration operation typically these are written to a http.Request
*/
type GetLdapConfigurationParams struct {

	/*RealmID
	  The Elasticsearch Security realm identifier.

	*/
	RealmID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap configuration params
func (o *GetLdapConfigurationParams) WithTimeout(timeout time.Duration) *GetLdapConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap configuration params
func (o *GetLdapConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap configuration params
func (o *GetLdapConfigurationParams) WithContext(ctx context.Context) *GetLdapConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap configuration params
func (o *GetLdapConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap configuration params
func (o *GetLdapConfigurationParams) WithHTTPClient(client *http.Client) *GetLdapConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap configuration params
func (o *GetLdapConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRealmID adds the realmID to the get ldap configuration params
func (o *GetLdapConfigurationParams) WithRealmID(realmID string) *GetLdapConfigurationParams {
	o.SetRealmID(realmID)
	return o
}

// SetRealmID adds the realmId to the get ldap configuration params
func (o *GetLdapConfigurationParams) SetRealmID(realmID string) {
	o.RealmID = realmID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param realm_id
	if err := r.SetPathParam("realm_id", o.RealmID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
