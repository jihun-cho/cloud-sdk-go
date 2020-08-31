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

package telemetry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetTelemetryConfigReader is a Reader for the SetTelemetryConfig structure.
type SetTelemetryConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTelemetryConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetTelemetryConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSetTelemetryConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetTelemetryConfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetTelemetryConfigRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetTelemetryConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetTelemetryConfigOK creates a SetTelemetryConfigOK with default headers values
func NewSetTelemetryConfigOK() *SetTelemetryConfigOK {
	return &SetTelemetryConfigOK{}
}

/*SetTelemetryConfigOK handles this case with default header values.

Telemetry configuration updated successfully
*/
type SetTelemetryConfigOK struct {
	Payload *models.TelemetryConfig
}

func (o *SetTelemetryConfigOK) Error() string {
	return fmt.Sprintf("[PUT /phone-home/config][%d] setTelemetryConfigOK  %+v", 200, o.Payload)
}

func (o *SetTelemetryConfigOK) GetPayload() *models.TelemetryConfig {
	return o.Payload
}

func (o *SetTelemetryConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TelemetryConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTelemetryConfigForbidden creates a SetTelemetryConfigForbidden with default headers values
func NewSetTelemetryConfigForbidden() *SetTelemetryConfigForbidden {
	return &SetTelemetryConfigForbidden{}
}

/*SetTelemetryConfigForbidden handles this case with default header values.

User must have Platform level permissions. (code: `root.unauthorized.rbac`)
*/
type SetTelemetryConfigForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetTelemetryConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /phone-home/config][%d] setTelemetryConfigForbidden  %+v", 403, o.Payload)
}

func (o *SetTelemetryConfigForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetTelemetryConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTelemetryConfigConflict creates a SetTelemetryConfigConflict with default headers values
func NewSetTelemetryConfigConflict() *SetTelemetryConfigConflict {
	return &SetTelemetryConfigConflict{}
}

/*SetTelemetryConfigConflict handles this case with default header values.

The telemetry configuration did not exist so there was an attempt to create one. Another request resulted in the creation of a telemetry configuration before this request completed, resulting in the failure of this request to create a configuration. Please retry. (code: `telemetry.already_exists`)
*/
type SetTelemetryConfigConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetTelemetryConfigConflict) Error() string {
	return fmt.Sprintf("[PUT /phone-home/config][%d] setTelemetryConfigConflict  %+v", 409, o.Payload)
}

func (o *SetTelemetryConfigConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetTelemetryConfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTelemetryConfigRetryWith creates a SetTelemetryConfigRetryWith with default headers values
func NewSetTelemetryConfigRetryWith() *SetTelemetryConfigRetryWith {
	return &SetTelemetryConfigRetryWith{}
}

/*SetTelemetryConfigRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetTelemetryConfigRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetTelemetryConfigRetryWith) Error() string {
	return fmt.Sprintf("[PUT /phone-home/config][%d] setTelemetryConfigRetryWith  %+v", 449, o.Payload)
}

func (o *SetTelemetryConfigRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetTelemetryConfigRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTelemetryConfigInternalServerError creates a SetTelemetryConfigInternalServerError with default headers values
func NewSetTelemetryConfigInternalServerError() *SetTelemetryConfigInternalServerError {
	return &SetTelemetryConfigInternalServerError{}
}

/*SetTelemetryConfigInternalServerError handles this case with default header values.

Failed to set the configuration due to an internal server error. (code: `telemetry.request_execution_failed`)
*/
type SetTelemetryConfigInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetTelemetryConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /phone-home/config][%d] setTelemetryConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *SetTelemetryConfigInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetTelemetryConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
