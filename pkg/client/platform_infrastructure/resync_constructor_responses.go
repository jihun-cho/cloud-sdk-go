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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ResyncConstructorReader is a Reader for the ResyncConstructor structure.
type ResyncConstructorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncConstructorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResyncConstructorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewResyncConstructorRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResyncConstructorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResyncConstructorOK creates a ResyncConstructorOK with default headers values
func NewResyncConstructorOK() *ResyncConstructorOK {
	return &ResyncConstructorOK{}
}

/*ResyncConstructorOK handles this case with default header values.

The constructor resync operation executed successfully
*/
type ResyncConstructorOK struct {
	Payload models.EmptyResponse
}

func (o *ResyncConstructorOK) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/_resync][%d] resyncConstructorOK  %+v", 200, o.Payload)
}

func (o *ResyncConstructorOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *ResyncConstructorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncConstructorRetryWith creates a ResyncConstructorRetryWith with default headers values
func NewResyncConstructorRetryWith() *ResyncConstructorRetryWith {
	return &ResyncConstructorRetryWith{}
}

/*ResyncConstructorRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncConstructorRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncConstructorRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/_resync][%d] resyncConstructorRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncConstructorRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncConstructorRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncConstructorInternalServerError creates a ResyncConstructorInternalServerError with default headers values
func NewResyncConstructorInternalServerError() *ResyncConstructorInternalServerError {
	return &ResyncConstructorInternalServerError{}
}

/*ResyncConstructorInternalServerError handles this case with default header values.

The constructor resync operation failed for allocator {constructor_id}. (code: `constructors.resync_failed`)
*/
type ResyncConstructorInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncConstructorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/{constructor_id}/_resync][%d] resyncConstructorInternalServerError  %+v", 500, o.Payload)
}

func (o *ResyncConstructorInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncConstructorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
