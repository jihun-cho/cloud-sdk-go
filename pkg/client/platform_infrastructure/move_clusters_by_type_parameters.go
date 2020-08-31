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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewMoveClustersByTypeParams creates a new MoveClustersByTypeParams object
// with the default values initialized.
func NewMoveClustersByTypeParams() *MoveClustersByTypeParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersByTypeParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveClustersByTypeParamsWithTimeout creates a new MoveClustersByTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveClustersByTypeParamsWithTimeout(timeout time.Duration) *MoveClustersByTypeParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersByTypeParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewMoveClustersByTypeParamsWithContext creates a new MoveClustersByTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveClustersByTypeParamsWithContext(ctx context.Context) *MoveClustersByTypeParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersByTypeParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		Context: ctx,
	}
}

// NewMoveClustersByTypeParamsWithHTTPClient creates a new MoveClustersByTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveClustersByTypeParamsWithHTTPClient(client *http.Client) *MoveClustersByTypeParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersByTypeParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,
		HTTPClient:   client,
	}
}

/*MoveClustersByTypeParams contains all the parameters to send to the API endpoint
for the move clusters by type operation typically these are written to a http.Request
*/
type MoveClustersByTypeParams struct {

	/*AllocatorDown
	  When `true`, considers all instances on the allocator as permanently shut down when deciding how to migrate data to new nodes.When left blank, the system automatically decides. NOTE: The default treats the allocator as up.

	*/
	AllocatorDown *bool
	/*AllocatorID
	  The allocator identifier.

	*/
	AllocatorID string
	/*Body
	  Overrides defaults for the move of each cluster

	*/
	Body *models.MoveClustersRequest
	/*ClusterType
	  The cluster types to move off of the allocator. NOTE: When unspecified, all clusters are moved.

	*/
	ClusterType string
	/*ForceUpdate
	  When true, cancels and overwrites pending plans, or treats instance as an error

	*/
	ForceUpdate *bool
	/*MoveOnly
	  When true, bypasses the cluster state changes, but continues to move the specified instances

	*/
	MoveOnly *bool
	/*ValidateOnly
	  When `true`, validates the plan overrides, then returns the plan without performing the move.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move clusters by type params
func (o *MoveClustersByTypeParams) WithTimeout(timeout time.Duration) *MoveClustersByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move clusters by type params
func (o *MoveClustersByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move clusters by type params
func (o *MoveClustersByTypeParams) WithContext(ctx context.Context) *MoveClustersByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move clusters by type params
func (o *MoveClustersByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move clusters by type params
func (o *MoveClustersByTypeParams) WithHTTPClient(client *http.Client) *MoveClustersByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move clusters by type params
func (o *MoveClustersByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorDown adds the allocatorDown to the move clusters by type params
func (o *MoveClustersByTypeParams) WithAllocatorDown(allocatorDown *bool) *MoveClustersByTypeParams {
	o.SetAllocatorDown(allocatorDown)
	return o
}

// SetAllocatorDown adds the allocatorDown to the move clusters by type params
func (o *MoveClustersByTypeParams) SetAllocatorDown(allocatorDown *bool) {
	o.AllocatorDown = allocatorDown
}

// WithAllocatorID adds the allocatorID to the move clusters by type params
func (o *MoveClustersByTypeParams) WithAllocatorID(allocatorID string) *MoveClustersByTypeParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the move clusters by type params
func (o *MoveClustersByTypeParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithBody adds the body to the move clusters by type params
func (o *MoveClustersByTypeParams) WithBody(body *models.MoveClustersRequest) *MoveClustersByTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move clusters by type params
func (o *MoveClustersByTypeParams) SetBody(body *models.MoveClustersRequest) {
	o.Body = body
}

// WithClusterType adds the clusterType to the move clusters by type params
func (o *MoveClustersByTypeParams) WithClusterType(clusterType string) *MoveClustersByTypeParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the move clusters by type params
func (o *MoveClustersByTypeParams) SetClusterType(clusterType string) {
	o.ClusterType = clusterType
}

// WithForceUpdate adds the forceUpdate to the move clusters by type params
func (o *MoveClustersByTypeParams) WithForceUpdate(forceUpdate *bool) *MoveClustersByTypeParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move clusters by type params
func (o *MoveClustersByTypeParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithMoveOnly adds the moveOnly to the move clusters by type params
func (o *MoveClustersByTypeParams) WithMoveOnly(moveOnly *bool) *MoveClustersByTypeParams {
	o.SetMoveOnly(moveOnly)
	return o
}

// SetMoveOnly adds the moveOnly to the move clusters by type params
func (o *MoveClustersByTypeParams) SetMoveOnly(moveOnly *bool) {
	o.MoveOnly = moveOnly
}

// WithValidateOnly adds the validateOnly to the move clusters by type params
func (o *MoveClustersByTypeParams) WithValidateOnly(validateOnly *bool) *MoveClustersByTypeParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move clusters by type params
func (o *MoveClustersByTypeParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveClustersByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllocatorDown != nil {

		// query param allocator_down
		var qrAllocatorDown bool
		if o.AllocatorDown != nil {
			qrAllocatorDown = *o.AllocatorDown
		}
		qAllocatorDown := swag.FormatBool(qrAllocatorDown)
		if qAllocatorDown != "" {
			if err := r.SetQueryParam("allocator_down", qAllocatorDown); err != nil {
				return err
			}
		}

	}

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_type
	if err := r.SetPathParam("cluster_type", o.ClusterType); err != nil {
		return err
	}

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool
		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {
			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
		}

	}

	if o.MoveOnly != nil {

		// query param move_only
		var qrMoveOnly bool
		if o.MoveOnly != nil {
			qrMoveOnly = *o.MoveOnly
		}
		qMoveOnly := swag.FormatBool(qrMoveOnly)
		if qMoveOnly != "" {
			if err := r.SetQueryParam("move_only", qMoveOnly); err != nil {
				return err
			}
		}

	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
