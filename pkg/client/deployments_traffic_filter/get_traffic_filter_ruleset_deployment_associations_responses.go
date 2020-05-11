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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetTrafficFilterRulesetDeploymentAssociationsReader is a Reader for the GetTrafficFilterRulesetDeploymentAssociations structure.
type GetTrafficFilterRulesetDeploymentAssociationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTrafficFilterRulesetDeploymentAssociationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTrafficFilterRulesetDeploymentAssociationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTrafficFilterRulesetDeploymentAssociationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTrafficFilterRulesetDeploymentAssociationsOK creates a GetTrafficFilterRulesetDeploymentAssociationsOK with default headers values
func NewGetTrafficFilterRulesetDeploymentAssociationsOK() *GetTrafficFilterRulesetDeploymentAssociationsOK {
	return &GetTrafficFilterRulesetDeploymentAssociationsOK{}
}

/*GetTrafficFilterRulesetDeploymentAssociationsOK handles this case with default header values.

Associations referred by traffic filter rulesets were successfully returned
*/
type GetTrafficFilterRulesetDeploymentAssociationsOK struct {
	Payload *models.RulesetAssociations
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsOK) Error() string {
	return fmt.Sprintf("[GET /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] getTrafficFilterRulesetDeploymentAssociationsOK  %+v", 200, o.Payload)
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsOK) GetPayload() *models.RulesetAssociations {
	return o.Payload
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RulesetAssociations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTrafficFilterRulesetDeploymentAssociationsInternalServerError creates a GetTrafficFilterRulesetDeploymentAssociationsInternalServerError with default headers values
func NewGetTrafficFilterRulesetDeploymentAssociationsInternalServerError() *GetTrafficFilterRulesetDeploymentAssociationsInternalServerError {
	return &GetTrafficFilterRulesetDeploymentAssociationsInternalServerError{}
}

/*GetTrafficFilterRulesetDeploymentAssociationsInternalServerError handles this case with default header values.

Error fetching deployments. (code: `traffic_filter.request_execution_failed`)
*/
type GetTrafficFilterRulesetDeploymentAssociationsInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] getTrafficFilterRulesetDeploymentAssociationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetTrafficFilterRulesetDeploymentAssociationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}