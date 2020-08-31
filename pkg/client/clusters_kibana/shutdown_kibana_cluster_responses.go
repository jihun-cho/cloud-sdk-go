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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ShutdownKibanaClusterReader is a Reader for the ShutdownKibanaCluster structure.
type ShutdownKibanaClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownKibanaClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewShutdownKibanaClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShutdownKibanaClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewShutdownKibanaClusterRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownKibanaClusterAccepted creates a ShutdownKibanaClusterAccepted with default headers values
func NewShutdownKibanaClusterAccepted() *ShutdownKibanaClusterAccepted {
	return &ShutdownKibanaClusterAccepted{}
}

/*ShutdownKibanaClusterAccepted handles this case with default header values.

The shutdown command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type ShutdownKibanaClusterAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *ShutdownKibanaClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_shutdown][%d] shutdownKibanaClusterAccepted  %+v", 202, o.Payload)
}

func (o *ShutdownKibanaClusterAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *ShutdownKibanaClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownKibanaClusterNotFound creates a ShutdownKibanaClusterNotFound with default headers values
func NewShutdownKibanaClusterNotFound() *ShutdownKibanaClusterNotFound {
	return &ShutdownKibanaClusterNotFound{}
}

/*ShutdownKibanaClusterNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type ShutdownKibanaClusterNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownKibanaClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_shutdown][%d] shutdownKibanaClusterNotFound  %+v", 404, o.Payload)
}

func (o *ShutdownKibanaClusterNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownKibanaClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownKibanaClusterRetryWith creates a ShutdownKibanaClusterRetryWith with default headers values
func NewShutdownKibanaClusterRetryWith() *ShutdownKibanaClusterRetryWith {
	return &ShutdownKibanaClusterRetryWith{}
}

/*ShutdownKibanaClusterRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ShutdownKibanaClusterRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownKibanaClusterRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_shutdown][%d] shutdownKibanaClusterRetryWith  %+v", 449, o.Payload)
}

func (o *ShutdownKibanaClusterRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownKibanaClusterRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
