// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchAutoUpgrade represents a SwitchAutoUpgrade struct.
type SwitchAutoUpgrade struct {
    // Custom version to be used. The Property Key is the switch hardware and the property value is the firmware version
    CustomVersions       map[string]string      `json:"custom_versions,omitempty"`
    // Enable auto upgrade for the switch
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Enable snapshot during the upgrade process
    Snapshot             *bool                  `json:"snapshot,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchAutoUpgrade) String() string {
    return fmt.Sprintf(
    	"SwitchAutoUpgrade[CustomVersions=%v, Enabled=%v, Snapshot=%v, AdditionalProperties=%v]",
    	s.CustomVersions, s.Enabled, s.Snapshot, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchAutoUpgrade.
// It customizes the JSON marshaling process for SwitchAutoUpgrade objects.
func (s SwitchAutoUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "custom_versions", "enabled", "snapshot"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchAutoUpgrade object to a map representation for JSON marshaling.
func (s SwitchAutoUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomVersions != nil {
        structMap["custom_versions"] = s.CustomVersions
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Snapshot != nil {
        structMap["snapshot"] = s.Snapshot
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchAutoUpgrade.
// It customizes the JSON unmarshaling process for SwitchAutoUpgrade objects.
func (s *SwitchAutoUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempSwitchAutoUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_versions", "enabled", "snapshot")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CustomVersions = temp.CustomVersions
    s.Enabled = temp.Enabled
    s.Snapshot = temp.Snapshot
    return nil
}

// tempSwitchAutoUpgrade is a temporary struct used for validating the fields of SwitchAutoUpgrade.
type tempSwitchAutoUpgrade  struct {
    CustomVersions map[string]string `json:"custom_versions,omitempty"`
    Enabled        *bool             `json:"enabled,omitempty"`
    Snapshot       *bool             `json:"snapshot,omitempty"`
}
