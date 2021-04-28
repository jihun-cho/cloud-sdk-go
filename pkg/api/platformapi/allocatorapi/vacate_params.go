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

package allocatorapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/output"
	"github.com/elastic/cloud-sdk-go/pkg/sync/pool"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

var (
	errCannotFilterByIDAndKind      = errors.New(`only one of "clusters" or "kind" can be specified`)
	errMustSpecifyAtLeast1Allocator = errors.New("must specify at least one allocator")
	errAPIMustNotBeNil              = errors.New("api must not be nil")
	errConcurrencyCannotBeZero      = errors.New("concurrency cannot be 0")
	errOutputDeviceCannotBeNil      = errors.New("output device cannot be nil")
	errCannotOverrideAllocatorDown  = errors.New("cannot set the AllocatorDown when multiple allocators are specified")

	allowedClusterKinds = []string{
		util.Apm,
		util.Appsearch,
		util.Elasticsearch,
		util.EnterpriseSearch,
		util.Kibana,
	}
)

// VacateParams used to vacate N allocators or clusters.
//nolint
type VacateParams struct {
	*api.API

	Region string

	// List of allocators that will be marked to be vacated.
	Allocators []string

	// List of allocators to be used as potential targets.
	PreferredAllocators []string

	// Can specify a list of cluster IDs that will be moved
	// of N numbers of allocators.
	ClusterFilter []string

	// If specified it will only move clusters that match the kind
	// (elasticsearch, kibana).
	KindFilter string

	// Maximum number of concurrent cluster moves at any time.
	Concurrency uint16

	// Output device where the progress will be sent.
	Output *output.Device

	// OutputFormat to use
	OutputFormat string

	// Maximum number of errors to allow the plan status poller to tolerate.
	MaxPollRetries uint8

	// Poll frequency
	TrackFrequency time.Duration

	// Optional value to be set to the pool on construction.
	PoolTimeout pool.Timeout

	// Optional value to be set to override the default autodiscovery of an
	// allocator's health. This can only be used when a single allocator is
	// specified.
	AllocatorDown *bool

	// Optional value to be set to keep the cluster in its current -possibly broken- state and just does the
	// bare minimum to move the requested instances across to another allocator.
	MoveOnly *bool

	// SkipTracking skips displaying and waiting for the individual vacates to complete.
	// Setting it to true will render the concurrency flag pretty much ineffective since
	// the vacate action is asynchronous and the only thing keeping the working items in
	// the pool is the tracking function call which synchronously waits until the vacate
	// has effectively finished.
	SkipTracking bool

	// Plan body overrides to place in all of the vacate clusters.
	PlanOverrides
}

// Validate validates the parameters
func (params VacateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator vacate params")
	if params.API == nil {
		merr = merr.Append(errAPIMustNotBeNil)
	}

	if len(params.Allocators) == 0 {
		merr = merr.Append(errMustSpecifyAtLeast1Allocator)
	}

	if len(params.ClusterFilter) > 0 && len(params.KindFilter) > 0 {
		merr = merr.Append(errCannotFilterByIDAndKind)
	}

	if params.Concurrency == 0 {
		merr = merr.Append(errConcurrencyCannotBeZero)
	}

	for i := range params.ClusterFilter {
		if len(params.ClusterFilter[i]) != 32 {
			merr = merr.Append(fmt.Errorf(
				"cluster filter: id \"%s\" is invalid, must be 32 characters long", params.ClusterFilter[i],
			))
		}
	}

	if params.Output == nil {
		merr = merr.Append(errOutputDeviceCannotBeNil)
	}

	if params.AllocatorDown != nil && len(params.Allocators) > 1 {
		merr = merr.Append(errCannotOverrideAllocatorDown)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// VacateClusterParams is used by VacateCluster to move a cluster node
// from an allocator.
type VacateClusterParams struct {
	PreferredAllocators []string
	ClusterFilter       []string
	// Plan body overrides to place in all of the vacate clusters.
	PlanOverrides
	Region    string
	ID        string // allocatorID
	ClusterID string
	Kind      string
	*api.API
	TrackFrequency time.Duration
	AllocatorDown  *bool
	MoveOnly       *bool
	Output         *output.Device
	OutputFormat   string
	MaxPollRetries uint8
	SkipTracking   bool
	Moves          *models.MoveClustersDetails
}

// Validate validates the parameters
func (params VacateClusterParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator vacate params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(
			fmt.Errorf("invalid allocator ID %s", params.ID),
		)
	}

	if len(params.ClusterID) != 32 {
		merr = merr.Append(
			fmt.Errorf("invalid cluster ID %s", params.ClusterID),
		)
	}

	if !slice.HasString(allowedClusterKinds, params.Kind) {
		merr = merr.Append(
			fmt.Errorf("invalid kind %s", params.Kind),
		)
	}

	if params.Output == nil {
		merr = merr.Append(errOutputDeviceCannotBeNil)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	if params.Moves == nil {
		merr = merr.Append(
			fmt.Errorf("cluster move plan should not be empty"),
		)
	}

	return merr.ErrorOrNil()
}

type addAllocatorMovesToPoolParams struct {
	ID           string
	Moves        *models.MoveClustersDetails
	Pool         *pool.Pool
	VacateParams *VacateParams
}

// PlanOverrides is used to override any API value that is returned by default
// with the specified value.
type PlanOverrides struct {
	// SkipSnapshot overwrites the Transient part of an Elastisearch vacate.
	SkipSnapshot      *bool
	SkipDataMigration *bool
	OverrideFailsafe  *bool
}
