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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RollingStrategyConfig Performs inline, rolling configuration changes that mutate existing containers. TIP: This is the fastest way to update a plan, but can fail for complex plan changes, such as topology changes. Also, this is less safe for configuration changes that leave a cluster in a non running state. NOTE: When you perform a major version upgrade, and 'group_by' is set to 'pass:macros[__all__]';, rolling is required.
//
// swagger:model RollingStrategyConfig
type RollingStrategyConfig struct {

	// Whether we allow changing the capacity of instances (default false). This is currently implemented by stopping, re-creating then starting the affected instance on its associated allocator when performing the changes. NOTES: This requires a round-trip through the allocation infrastructure of the active constructor, as it has to reserve the target capacity without over-committing
	AllowInlineResize *bool `json:"allow_inline_resize,omitempty"`

	// Specifies the grouping attribute to use when rolling several instances. Instances that share the same value for the provided attribute key are rolled together as a unit. Examples that make sense to use are '\_\_all\_\_' (roll all instances as a single unit), 'logical_zone_name' (roll instances by zone), '\_\_name\_\_' (roll one instance at a time, the default if not specified). Note that '\_\_all\_\_' is required when performing a major version upgrade
	GroupBy string `json:"group_by,omitempty"`

	// The time, in seconds, to wait for shards that show no progress of initializing before rolling the next group (default: 10 minutes)
	ShardInitWaitTime int64 `json:"shard_init_wait_time,omitempty"`

	// Whether to skip attempting to do a synced flush on the filesystem of the container (default: false), which is less safe but may be required if the container is unhealthy
	SkipSyncedFlush *bool `json:"skip_synced_flush,omitempty"`
}

// Validate validates this rolling strategy config
func (m *RollingStrategyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RollingStrategyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RollingStrategyConfig) UnmarshalBinary(b []byte) error {
	var res RollingStrategyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
