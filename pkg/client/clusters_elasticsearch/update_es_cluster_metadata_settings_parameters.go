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

package clusters_elasticsearch

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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateEsClusterMetadataSettingsParams creates a new UpdateEsClusterMetadataSettingsParams object
// with the default values initialized.
func NewUpdateEsClusterMetadataSettingsParams() *UpdateEsClusterMetadataSettingsParams {
	var ()
	return &UpdateEsClusterMetadataSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEsClusterMetadataSettingsParamsWithTimeout creates a new UpdateEsClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEsClusterMetadataSettingsParamsWithTimeout(timeout time.Duration) *UpdateEsClusterMetadataSettingsParams {
	var ()
	return &UpdateEsClusterMetadataSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateEsClusterMetadataSettingsParamsWithContext creates a new UpdateEsClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEsClusterMetadataSettingsParamsWithContext(ctx context.Context) *UpdateEsClusterMetadataSettingsParams {
	var ()
	return &UpdateEsClusterMetadataSettingsParams{

		Context: ctx,
	}
}

// NewUpdateEsClusterMetadataSettingsParamsWithHTTPClient creates a new UpdateEsClusterMetadataSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEsClusterMetadataSettingsParamsWithHTTPClient(client *http.Client) *UpdateEsClusterMetadataSettingsParams {
	var ()
	return &UpdateEsClusterMetadataSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateEsClusterMetadataSettingsParams contains all the parameters to send to the API endpoint
for the update es cluster metadata settings operation typically these are written to a http.Request
*/
type UpdateEsClusterMetadataSettingsParams struct {

	/*Body
	  The cluster settings including updated values

	*/
	Body *models.ClusterMetadataSettings
	/*ClusterID
	  Elasticsearch cluster identifier

	*/
	ClusterID string
	/*Version
	  Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithTimeout(timeout time.Duration) *UpdateEsClusterMetadataSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithContext(ctx context.Context) *UpdateEsClusterMetadataSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithHTTPClient(client *http.Client) *UpdateEsClusterMetadataSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithBody(body *models.ClusterMetadataSettings) *UpdateEsClusterMetadataSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetBody(body *models.ClusterMetadataSettings) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithClusterID(clusterID string) *UpdateEsClusterMetadataSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithVersion adds the version to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) WithVersion(version *int64) *UpdateEsClusterMetadataSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update es cluster metadata settings params
func (o *UpdateEsClusterMetadataSettingsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEsClusterMetadataSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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
