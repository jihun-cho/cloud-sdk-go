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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteTrafficFilterRulesetReader is a Reader for the DeleteTrafficFilterRuleset structure.
type DeleteTrafficFilterRulesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTrafficFilterRulesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTrafficFilterRulesetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTrafficFilterRulesetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTrafficFilterRulesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTrafficFilterRulesetOK creates a DeleteTrafficFilterRulesetOK with default headers values
func NewDeleteTrafficFilterRulesetOK() *DeleteTrafficFilterRulesetOK {
	return &DeleteTrafficFilterRulesetOK{}
}

/*DeleteTrafficFilterRulesetOK handles this case with default header values.

The traffic filter ruleset was successfully deleted.
*/
type DeleteTrafficFilterRulesetOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteTrafficFilterRulesetOK) Error() string {
	return fmt.Sprintf("[DELETE /deployments/traffic-filter/rulesets/{ruleset_id}][%d] deleteTrafficFilterRulesetOK  %+v", 200, o.Payload)
}

func (o *DeleteTrafficFilterRulesetOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteTrafficFilterRulesetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTrafficFilterRulesetNotFound creates a DeleteTrafficFilterRulesetNotFound with default headers values
func NewDeleteTrafficFilterRulesetNotFound() *DeleteTrafficFilterRulesetNotFound {
	return &DeleteTrafficFilterRulesetNotFound{}
}

/*DeleteTrafficFilterRulesetNotFound handles this case with default header values.

The traffic filter ruleset specified by {ruleset_id} cannot be found. (code: `traffic_filter.not_found`)
*/
type DeleteTrafficFilterRulesetNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteTrafficFilterRulesetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployments/traffic-filter/rulesets/{ruleset_id}][%d] deleteTrafficFilterRulesetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTrafficFilterRulesetNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteTrafficFilterRulesetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTrafficFilterRulesetInternalServerError creates a DeleteTrafficFilterRulesetInternalServerError with default headers values
func NewDeleteTrafficFilterRulesetInternalServerError() *DeleteTrafficFilterRulesetInternalServerError {
	return &DeleteTrafficFilterRulesetInternalServerError{}
}

/*DeleteTrafficFilterRulesetInternalServerError handles this case with default header values.

Error deleting the traffic filter ruleset. (code: `traffic_filter.request_execution_failed`)
*/
type DeleteTrafficFilterRulesetInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteTrafficFilterRulesetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /deployments/traffic-filter/rulesets/{ruleset_id}][%d] deleteTrafficFilterRulesetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTrafficFilterRulesetInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteTrafficFilterRulesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
