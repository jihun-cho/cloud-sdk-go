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

package deployments_ip_filtering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteIPFilterRulesetReader is a Reader for the DeleteIPFilterRuleset structure.
type DeleteIPFilterRulesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPFilterRulesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIPFilterRulesetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteIPFilterRulesetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteIPFilterRulesetRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIPFilterRulesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIPFilterRulesetOK creates a DeleteIPFilterRulesetOK with default headers values
func NewDeleteIPFilterRulesetOK() *DeleteIPFilterRulesetOK {
	return &DeleteIPFilterRulesetOK{}
}

/*DeleteIPFilterRulesetOK handles this case with default header values.

The IP filter ruleset was successfully deleted
*/
type DeleteIPFilterRulesetOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteIPFilterRulesetOK) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}][%d] deleteIpFilterRulesetOK  %+v", 200, o.Payload)
}

func (o *DeleteIPFilterRulesetOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteIPFilterRulesetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPFilterRulesetNotFound creates a DeleteIPFilterRulesetNotFound with default headers values
func NewDeleteIPFilterRulesetNotFound() *DeleteIPFilterRulesetNotFound {
	return &DeleteIPFilterRulesetNotFound{}
}

/*DeleteIPFilterRulesetNotFound handles this case with default header values.

The IP filter ruleset specified by {ruleset_id} cannot be found (code: 'ip_filtering.ruleset_not_found')
*/
type DeleteIPFilterRulesetNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteIPFilterRulesetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}][%d] deleteIpFilterRulesetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIPFilterRulesetNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteIPFilterRulesetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPFilterRulesetRetryWith creates a DeleteIPFilterRulesetRetryWith with default headers values
func NewDeleteIPFilterRulesetRetryWith() *DeleteIPFilterRulesetRetryWith {
	return &DeleteIPFilterRulesetRetryWith{}
}

/*DeleteIPFilterRulesetRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type DeleteIPFilterRulesetRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteIPFilterRulesetRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}][%d] deleteIpFilterRulesetRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteIPFilterRulesetRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteIPFilterRulesetRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPFilterRulesetInternalServerError creates a DeleteIPFilterRulesetInternalServerError with default headers values
func NewDeleteIPFilterRulesetInternalServerError() *DeleteIPFilterRulesetInternalServerError {
	return &DeleteIPFilterRulesetInternalServerError{}
}

/*DeleteIPFilterRulesetInternalServerError handles this case with default header values.

The request execution failed (code: 'ip_filtering.request_execution_failed')
*/
type DeleteIPFilterRulesetInternalServerError struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteIPFilterRulesetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}][%d] deleteIpFilterRulesetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPFilterRulesetInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteIPFilterRulesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
