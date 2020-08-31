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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// MigrateEsClusterPlanReader is a Reader for the MigrateEsClusterPlan structure.
type MigrateEsClusterPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateEsClusterPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateEsClusterPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewMigrateEsClusterPlanAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMigrateEsClusterPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMigrateEsClusterPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewMigrateEsClusterPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewMigrateEsClusterPlanRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigrateEsClusterPlanOK creates a MigrateEsClusterPlanOK with default headers values
func NewMigrateEsClusterPlanOK() *MigrateEsClusterPlanOK {
	return &MigrateEsClusterPlanOK{}
}

/*MigrateEsClusterPlanOK handles this case with default header values.

The current cluster plan migrated to the specified deployment template.
*/
type MigrateEsClusterPlanOK struct {
	Payload *models.ClusterPlanMigrationResponse
}

func (o *MigrateEsClusterPlanOK) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanOK  %+v", 200, o.Payload)
}

func (o *MigrateEsClusterPlanOK) GetPayload() *models.ClusterPlanMigrationResponse {
	return o.Payload
}

func (o *MigrateEsClusterPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPlanMigrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateEsClusterPlanAccepted creates a MigrateEsClusterPlanAccepted with default headers values
func NewMigrateEsClusterPlanAccepted() *MigrateEsClusterPlanAccepted {
	return &MigrateEsClusterPlanAccepted{}
}

/*MigrateEsClusterPlanAccepted handles this case with default header values.

The plan definition was valid and the updated plan is in progress
*/
type MigrateEsClusterPlanAccepted struct {
	Payload *models.ClusterPlanMigrationResponse
}

func (o *MigrateEsClusterPlanAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanAccepted  %+v", 202, o.Payload)
}

func (o *MigrateEsClusterPlanAccepted) GetPayload() *models.ClusterPlanMigrationResponse {
	return o.Payload
}

func (o *MigrateEsClusterPlanAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPlanMigrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateEsClusterPlanBadRequest creates a MigrateEsClusterPlanBadRequest with default headers values
func NewMigrateEsClusterPlanBadRequest() *MigrateEsClusterPlanBadRequest {
	return &MigrateEsClusterPlanBadRequest{}
}

/*MigrateEsClusterPlanBadRequest handles this case with default header values.

Migrating to the specified template would lead to an invalid or unsupported cluster definition. (code: `clusters.cluster_invalid_plan`)
*/
type MigrateEsClusterPlanBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MigrateEsClusterPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanBadRequest  %+v", 400, o.Payload)
}

func (o *MigrateEsClusterPlanBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateEsClusterPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateEsClusterPlanNotFound creates a MigrateEsClusterPlanNotFound with default headers values
func NewMigrateEsClusterPlanNotFound() *MigrateEsClusterPlanNotFound {
	return &MigrateEsClusterPlanNotFound{}
}

/*MigrateEsClusterPlanNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type MigrateEsClusterPlanNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MigrateEsClusterPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanNotFound  %+v", 404, o.Payload)
}

func (o *MigrateEsClusterPlanNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateEsClusterPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateEsClusterPlanPreconditionFailed creates a MigrateEsClusterPlanPreconditionFailed with default headers values
func NewMigrateEsClusterPlanPreconditionFailed() *MigrateEsClusterPlanPreconditionFailed {
	return &MigrateEsClusterPlanPreconditionFailed{}
}

/*MigrateEsClusterPlanPreconditionFailed handles this case with default header values.

There is not currently applied plan - eg the cluster has not finished provisioning, or the provisioning failed. (code: `clusters.cluster_plan_state_error`)
*/
type MigrateEsClusterPlanPreconditionFailed struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MigrateEsClusterPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanPreconditionFailed  %+v", 412, o.Payload)
}

func (o *MigrateEsClusterPlanPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateEsClusterPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateEsClusterPlanRetryWith creates a MigrateEsClusterPlanRetryWith with default headers values
func NewMigrateEsClusterPlanRetryWith() *MigrateEsClusterPlanRetryWith {
	return &MigrateEsClusterPlanRetryWith{}
}

/*MigrateEsClusterPlanRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type MigrateEsClusterPlanRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MigrateEsClusterPlanRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/plan/_migrate][%d] migrateEsClusterPlanRetryWith  %+v", 449, o.Payload)
}

func (o *MigrateEsClusterPlanRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateEsClusterPlanRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
