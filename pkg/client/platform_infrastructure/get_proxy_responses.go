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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetProxyReader is a Reader for the GetProxy structure.
type GetProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProxyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxyOK creates a GetProxyOK with default headers values
func NewGetProxyOK() *GetProxyOK {
	return &GetProxyOK{}
}

/*GetProxyOK handles this case with default header values.

The information for the proxy specified by {proxy_id}.
*/
type GetProxyOK struct {
	Payload *models.ProxyInfo
}

func (o *GetProxyOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/{proxy_id}][%d] getProxyOK  %+v", 200, o.Payload)
}

func (o *GetProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProxyNotFound creates a GetProxyNotFound with default headers values
func NewGetProxyNotFound() *GetProxyNotFound {
	return &GetProxyNotFound{}
}

/*GetProxyNotFound handles this case with default header values.

Unable to find the {proxy_id} specified proxy. Edit your request, then try again. (code: `proxies.proxy_not_found`)
*/
type GetProxyNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetProxyNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/{proxy_id}][%d] getProxyNotFound  %+v", 404, o.Payload)
}

func (o *GetProxyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}