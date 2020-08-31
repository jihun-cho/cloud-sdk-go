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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetUserAPIKeysReader is a Reader for the GetUserAPIKeys structure.
type GetUserAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAPIKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserAPIKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserAPIKeysOK creates a GetUserAPIKeysOK with default headers values
func NewGetUserAPIKeysOK() *GetUserAPIKeysOK {
	return &GetUserAPIKeysOK{}
}

/*GetUserAPIKeysOK handles this case with default header values.

The API key metadata is retrieved.
*/
type GetUserAPIKeysOK struct {
	Payload *models.APIKeysResponse
}

func (o *GetUserAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/auth/keys][%d] getUserApiKeysOK  %+v", 200, o.Payload)
}

func (o *GetUserAPIKeysOK) GetPayload() *models.APIKeysResponse {
	return o.Payload
}

func (o *GetUserAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserAPIKeysNotFound creates a GetUserAPIKeysNotFound with default headers values
func NewGetUserAPIKeysNotFound() *GetUserAPIKeysNotFound {
	return &GetUserAPIKeysNotFound{}
}

/*GetUserAPIKeysNotFound handles this case with default header values.

The {user_id} can't be found. (code: `api_keys.user_not_found`)
*/
type GetUserAPIKeysNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetUserAPIKeysNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/auth/keys][%d] getUserApiKeysNotFound  %+v", 404, o.Payload)
}

func (o *GetUserAPIKeysNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetUserAPIKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
