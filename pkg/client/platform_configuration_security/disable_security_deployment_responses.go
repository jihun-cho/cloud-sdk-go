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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DisableSecurityDeploymentReader is a Reader for the DisableSecurityDeployment structure.
type DisableSecurityDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableSecurityDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableSecurityDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDisableSecurityDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDisableSecurityDeploymentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDisableSecurityDeploymentRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDisableSecurityDeploymentOK creates a DisableSecurityDeploymentOK with default headers values
func NewDisableSecurityDeploymentOK() *DisableSecurityDeploymentOK {
	return &DisableSecurityDeploymentOK{}
}

/*DisableSecurityDeploymentOK handles this case with default header values.

The security deployment was successfully disabled
*/
type DisableSecurityDeploymentOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload models.EmptyResponse
}

func (o *DisableSecurityDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment/_disable][%d] disableSecurityDeploymentOK  %+v", 200, o.Payload)
}

func (o *DisableSecurityDeploymentOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DisableSecurityDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableSecurityDeploymentNotFound creates a DisableSecurityDeploymentNotFound with default headers values
func NewDisableSecurityDeploymentNotFound() *DisableSecurityDeploymentNotFound {
	return &DisableSecurityDeploymentNotFound{}
}

/*DisableSecurityDeploymentNotFound handles this case with default header values.

The realm specified by {realm_id} cannot be found. (code: `security_deployment.not_found`)
*/
type DisableSecurityDeploymentNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DisableSecurityDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment/_disable][%d] disableSecurityDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *DisableSecurityDeploymentNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DisableSecurityDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableSecurityDeploymentConflict creates a DisableSecurityDeploymentConflict with default headers values
func NewDisableSecurityDeploymentConflict() *DisableSecurityDeploymentConflict {
	return &DisableSecurityDeploymentConflict{}
}

/*DisableSecurityDeploymentConflict handles this case with default header values.

There is a version conflict. (code: `security_deployment.version_conflict`)
*/
type DisableSecurityDeploymentConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DisableSecurityDeploymentConflict) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment/_disable][%d] disableSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *DisableSecurityDeploymentConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DisableSecurityDeploymentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableSecurityDeploymentRetryWith creates a DisableSecurityDeploymentRetryWith with default headers values
func NewDisableSecurityDeploymentRetryWith() *DisableSecurityDeploymentRetryWith {
	return &DisableSecurityDeploymentRetryWith{}
}

/*DisableSecurityDeploymentRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DisableSecurityDeploymentRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DisableSecurityDeploymentRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment/_disable][%d] disableSecurityDeploymentRetryWith  %+v", 449, o.Payload)
}

func (o *DisableSecurityDeploymentRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DisableSecurityDeploymentRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
