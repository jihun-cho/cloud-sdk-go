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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetEsClusterMonitoringReader is a Reader for the SetEsClusterMonitoring structure.
type SetEsClusterMonitoringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEsClusterMonitoringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSetEsClusterMonitoringAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetEsClusterMonitoringNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetEsClusterMonitoringRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetEsClusterMonitoringAccepted creates a SetEsClusterMonitoringAccepted with default headers values
func NewSetEsClusterMonitoringAccepted() *SetEsClusterMonitoringAccepted {
	return &SetEsClusterMonitoringAccepted{}
}

/*SetEsClusterMonitoringAccepted handles this case with default header values.

The destination monitoring cluster creation/overwrite request was initiated
*/
type SetEsClusterMonitoringAccepted struct {
	Payload models.EmptyResponse
}

func (o *SetEsClusterMonitoringAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/monitoring/{dest_cluster_id}][%d] setEsClusterMonitoringAccepted  %+v", 202, o.Payload)
}

func (o *SetEsClusterMonitoringAccepted) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *SetEsClusterMonitoringAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEsClusterMonitoringNotFound creates a SetEsClusterMonitoringNotFound with default headers values
func NewSetEsClusterMonitoringNotFound() *SetEsClusterMonitoringNotFound {
	return &SetEsClusterMonitoringNotFound{}
}

/*SetEsClusterMonitoringNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type SetEsClusterMonitoringNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *SetEsClusterMonitoringNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/monitoring/{dest_cluster_id}][%d] setEsClusterMonitoringNotFound  %+v", 404, o.Payload)
}

func (o *SetEsClusterMonitoringNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetEsClusterMonitoringNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEsClusterMonitoringRetryWith creates a SetEsClusterMonitoringRetryWith with default headers values
func NewSetEsClusterMonitoringRetryWith() *SetEsClusterMonitoringRetryWith {
	return &SetEsClusterMonitoringRetryWith{}
}

/*SetEsClusterMonitoringRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type SetEsClusterMonitoringRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *SetEsClusterMonitoringRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/monitoring/{dest_cluster_id}][%d] setEsClusterMonitoringRetryWith  %+v", 449, o.Payload)
}

func (o *SetEsClusterMonitoringRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetEsClusterMonitoringRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
