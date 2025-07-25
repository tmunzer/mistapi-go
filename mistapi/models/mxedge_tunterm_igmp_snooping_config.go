// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermIgmpSnoopingConfig represents a MxedgeTuntermIgmpSnoopingConfig struct.
type MxedgeTuntermIgmpSnoopingConfig struct {
    Enabled              *bool                             `json:"enabled,omitempty"`
    Querier              *MxedgeTuntermIgmpSnoopingQuerier `json:"querier,omitempty"`
    // List of vlans on which tunterm performs IGMP snooping
    VlanIds              []int                             `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermIgmpSnoopingConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermIgmpSnoopingConfig) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermIgmpSnoopingConfig[Enabled=%v, Querier=%v, VlanIds=%v, AdditionalProperties=%v]",
    	m.Enabled, m.Querier, m.VlanIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIgmpSnoopingConfig.
// It customizes the JSON marshaling process for MxedgeTuntermIgmpSnoopingConfig objects.
func (m MxedgeTuntermIgmpSnoopingConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "enabled", "querier", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIgmpSnoopingConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermIgmpSnoopingConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.Querier != nil {
        structMap["querier"] = m.Querier.toMap()
    }
    if m.VlanIds != nil {
        structMap["vlan_ids"] = m.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermIgmpSnoopingConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermIgmpSnoopingConfig objects.
func (m *MxedgeTuntermIgmpSnoopingConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermIgmpSnoopingConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "querier", "vlan_ids")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Enabled = temp.Enabled
    m.Querier = temp.Querier
    m.VlanIds = temp.VlanIds
    return nil
}

// tempMxedgeTuntermIgmpSnoopingConfig is a temporary struct used for validating the fields of MxedgeTuntermIgmpSnoopingConfig.
type tempMxedgeTuntermIgmpSnoopingConfig  struct {
    Enabled *bool                             `json:"enabled,omitempty"`
    Querier *MxedgeTuntermIgmpSnoopingQuerier `json:"querier,omitempty"`
    VlanIds []int                             `json:"vlan_ids,omitempty"`
}
