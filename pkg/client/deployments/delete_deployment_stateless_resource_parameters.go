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

package deployments

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

// NewDeleteDeploymentStatelessResourceParams creates a new DeleteDeploymentStatelessResourceParams object
// with the default values initialized.
func NewDeleteDeploymentStatelessResourceParams() *DeleteDeploymentStatelessResourceParams {
	var ()
	return &DeleteDeploymentStatelessResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentStatelessResourceParamsWithTimeout creates a new DeleteDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentStatelessResourceParamsWithTimeout(timeout time.Duration) *DeleteDeploymentStatelessResourceParams {
	var ()
	return &DeleteDeploymentStatelessResourceParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentStatelessResourceParamsWithContext creates a new DeleteDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentStatelessResourceParamsWithContext(ctx context.Context) *DeleteDeploymentStatelessResourceParams {
	var ()
	return &DeleteDeploymentStatelessResourceParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentStatelessResourceParamsWithHTTPClient creates a new DeleteDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentStatelessResourceParamsWithHTTPClient(client *http.Client) *DeleteDeploymentStatelessResourceParams {
	var ()
	return &DeleteDeploymentStatelessResourceParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentStatelessResourceParams contains all the parameters to send to the API endpoint
for the delete deployment stateless resource operation typically these are written to a http.Request
*/
type DeleteDeploymentStatelessResourceParams struct {

	/*DeploymentID
	  Identifier for the Deployment

	*/
	DeploymentID string
	/*RefID
	  User-specified RefId for the Resource

	*/
	RefID string
	/*StatelessResourceKind
	  The kind of stateless resource

	*/
	StatelessResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithTimeout(timeout time.Duration) *DeleteDeploymentStatelessResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithContext(ctx context.Context) *DeleteDeploymentStatelessResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithHTTPClient(client *http.Client) *DeleteDeploymentStatelessResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithDeploymentID(deploymentID string) *DeleteDeploymentStatelessResourceParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithRefID(refID string) *DeleteDeploymentStatelessResourceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithStatelessResourceKind adds the statelessResourceKind to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) WithStatelessResourceKind(statelessResourceKind string) *DeleteDeploymentStatelessResourceParams {
	o.SetStatelessResourceKind(statelessResourceKind)
	return o
}

// SetStatelessResourceKind adds the statelessResourceKind to the delete deployment stateless resource params
func (o *DeleteDeploymentStatelessResourceParams) SetStatelessResourceKind(statelessResourceKind string) {
	o.StatelessResourceKind = statelessResourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentStatelessResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	// path param stateless_resource_kind
	if err := r.SetPathParam("stateless_resource_kind", o.StatelessResourceKind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
