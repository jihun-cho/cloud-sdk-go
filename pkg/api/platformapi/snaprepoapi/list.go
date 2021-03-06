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

package snaprepoapi

import (
	"context"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_snapshots"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// ListParams is embedded in all of the specific action functions
type ListParams struct {
	*api.API
	Region string
}

// Validate ensures that parameters are correct
func (params ListParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid snapshot repository list params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// List obtains all the configured platform snapshot repositories
func List(params ListParams) (*models.RepositoryConfigs, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	repo, err := params.V1API.PlatformConfigurationSnapshots.GetSnapshotRepositories(
		platform_configuration_snapshots.NewGetSnapshotRepositoriesParams().
			WithContext(api.WithRegion(context.Background(), params.Region)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, api.UnwrapError(err)
	}

	return repo.Payload, nil
}
