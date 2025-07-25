// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseOrgDevicesSummary represents a ResponseOrgDevicesSummary struct.
type ResponseOrgDevicesSummary struct {
    NumAps                *int                   `json:"num_aps,omitempty"`
    NumGateways           *int                   `json:"num_gateways,omitempty"`
    NumMxedges            *int                   `json:"num_mxedges,omitempty"`
    NumSwitches           *int                   `json:"num_switches,omitempty"`
    NumUnassignedAps      *int                   `json:"num_unassigned_aps,omitempty"`
    NumUnassignedGateways *int                   `json:"num_unassigned_gateways,omitempty"`
    NumUnassignedSwitches *int                   `json:"num_unassigned_switches,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgDevicesSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgDevicesSummary) String() string {
    return fmt.Sprintf(
    	"ResponseOrgDevicesSummary[NumAps=%v, NumGateways=%v, NumMxedges=%v, NumSwitches=%v, NumUnassignedAps=%v, NumUnassignedGateways=%v, NumUnassignedSwitches=%v, AdditionalProperties=%v]",
    	r.NumAps, r.NumGateways, r.NumMxedges, r.NumSwitches, r.NumUnassignedAps, r.NumUnassignedGateways, r.NumUnassignedSwitches, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgDevicesSummary.
// It customizes the JSON marshaling process for ResponseOrgDevicesSummary objects.
func (r ResponseOrgDevicesSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "num_aps", "num_gateways", "num_mxedges", "num_switches", "num_unassigned_aps", "num_unassigned_gateways", "num_unassigned_switches"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgDevicesSummary object to a map representation for JSON marshaling.
func (r ResponseOrgDevicesSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.NumAps != nil {
        structMap["num_aps"] = r.NumAps
    }
    if r.NumGateways != nil {
        structMap["num_gateways"] = r.NumGateways
    }
    if r.NumMxedges != nil {
        structMap["num_mxedges"] = r.NumMxedges
    }
    if r.NumSwitches != nil {
        structMap["num_switches"] = r.NumSwitches
    }
    if r.NumUnassignedAps != nil {
        structMap["num_unassigned_aps"] = r.NumUnassignedAps
    }
    if r.NumUnassignedGateways != nil {
        structMap["num_unassigned_gateways"] = r.NumUnassignedGateways
    }
    if r.NumUnassignedSwitches != nil {
        structMap["num_unassigned_switches"] = r.NumUnassignedSwitches
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgDevicesSummary.
// It customizes the JSON unmarshaling process for ResponseOrgDevicesSummary objects.
func (r *ResponseOrgDevicesSummary) UnmarshalJSON(input []byte) error {
    var temp tempResponseOrgDevicesSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_aps", "num_gateways", "num_mxedges", "num_switches", "num_unassigned_aps", "num_unassigned_gateways", "num_unassigned_switches")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.NumAps = temp.NumAps
    r.NumGateways = temp.NumGateways
    r.NumMxedges = temp.NumMxedges
    r.NumSwitches = temp.NumSwitches
    r.NumUnassignedAps = temp.NumUnassignedAps
    r.NumUnassignedGateways = temp.NumUnassignedGateways
    r.NumUnassignedSwitches = temp.NumUnassignedSwitches
    return nil
}

// tempResponseOrgDevicesSummary is a temporary struct used for validating the fields of ResponseOrgDevicesSummary.
type tempResponseOrgDevicesSummary  struct {
    NumAps                *int `json:"num_aps,omitempty"`
    NumGateways           *int `json:"num_gateways,omitempty"`
    NumMxedges            *int `json:"num_mxedges,omitempty"`
    NumSwitches           *int `json:"num_switches,omitempty"`
    NumUnassignedAps      *int `json:"num_unassigned_aps,omitempty"`
    NumUnassignedGateways *int `json:"num_unassigned_gateways,omitempty"`
    NumUnassignedSwitches *int `json:"num_unassigned_switches,omitempty"`
}
