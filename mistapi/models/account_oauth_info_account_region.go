// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AccountOauthInfoAccountRegion represents a AccountOauthInfoAccountRegion struct.
type AccountOauthInfoAccountRegion struct {
	// Bandwidth Aggregate region for this region
	AggregateRegion *string `json:"aggregate_region,omitempty"`
	// Allocated bandwidth for the region, in Mbps
	AllocatedBandwidth *int `json:"allocated_bandwidth,omitempty"`
	// Display name for this region
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountOauthInfoAccountRegion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountOauthInfoAccountRegion) String() string {
	return fmt.Sprintf(
		"AccountOauthInfoAccountRegion[AggregateRegion=%v, AllocatedBandwidth=%v, Name=%v, AdditionalProperties=%v]",
		a.AggregateRegion, a.AllocatedBandwidth, a.Name, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthInfoAccountRegion.
// It customizes the JSON marshaling process for AccountOauthInfoAccountRegion objects.
func (a AccountOauthInfoAccountRegion) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"aggregate_region", "allocated_bandwidth", "name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthInfoAccountRegion object to a map representation for JSON marshaling.
func (a AccountOauthInfoAccountRegion) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AggregateRegion != nil {
		structMap["aggregate_region"] = a.AggregateRegion
	}
	if a.AllocatedBandwidth != nil {
		structMap["allocated_bandwidth"] = a.AllocatedBandwidth
	}
	if a.Name != nil {
		structMap["name"] = a.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthInfoAccountRegion.
// It customizes the JSON unmarshaling process for AccountOauthInfoAccountRegion objects.
func (a *AccountOauthInfoAccountRegion) UnmarshalJSON(input []byte) error {
	var temp tempAccountOauthInfoAccountRegion
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aggregate_region", "allocated_bandwidth", "name")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.AggregateRegion = temp.AggregateRegion
	a.AllocatedBandwidth = temp.AllocatedBandwidth
	a.Name = temp.Name
	return nil
}

// tempAccountOauthInfoAccountRegion is a temporary struct used for validating the fields of AccountOauthInfoAccountRegion.
type tempAccountOauthInfoAccountRegion struct {
	AggregateRegion    *string `json:"aggregate_region,omitempty"`
	AllocatedBandwidth *int    `json:"allocated_bandwidth,omitempty"`
	Name               *string `json:"name,omitempty"`
}
