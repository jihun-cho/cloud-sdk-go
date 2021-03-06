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

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SearchClustersReader is a Reader for the SearchClusters structure.
type SearchClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchClustersOK creates a SearchClustersOK with default headers values
func NewSearchClustersOK() *SearchClustersOK {
	return &SearchClustersOK{}
}

/*SearchClustersOK handles this case with default header values.

A list of clusters that matched the given search query.
*/
type SearchClustersOK struct {
	Payload *models.ClustersInfo
}

func (o *SearchClustersOK) Error() string {
	return fmt.Sprintf("[POST /clusters/_search][%d] searchClustersOK  %+v", 200, o.Payload)
}

func (o *SearchClustersOK) GetPayload() *models.ClustersInfo {
	return o.Payload
}

func (o *SearchClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustersInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchClustersBadRequest creates a SearchClustersBadRequest with default headers values
func NewSearchClustersBadRequest() *SearchClustersBadRequest {
	return &SearchClustersBadRequest{}
}

/*SearchClustersBadRequest handles this case with default header values.

The search request failed. (code: `search.invalid_request`)
*/
type SearchClustersBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SearchClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/_search][%d] searchClustersBadRequest  %+v", 400, o.Payload)
}

func (o *SearchClustersBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SearchClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
