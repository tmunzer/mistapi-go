// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// UpgradeBios represents a UpgradeBios struct.
type UpgradeBios struct {
    // Reboot device immediately after upgrade is completed
    Reboot               *bool                  `json:"reboot,omitempty"`
    // Specific bios version
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeBios,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeBios) String() string {
    return fmt.Sprintf(
    	"UpgradeBios[Reboot=%v, Version=%v, AdditionalProperties=%v]",
    	u.Reboot, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeBios.
// It customizes the JSON marshaling process for UpgradeBios objects.
func (u UpgradeBios) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "reboot", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeBios object to a map representation for JSON marshaling.
func (u UpgradeBios) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Reboot != nil {
        structMap["reboot"] = u.Reboot
    }
    if u.Version != nil {
        structMap["version"] = u.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeBios.
// It customizes the JSON unmarshaling process for UpgradeBios objects.
func (u *UpgradeBios) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeBios
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reboot", "version")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Reboot = temp.Reboot
    u.Version = temp.Version
    return nil
}

// tempUpgradeBios is a temporary struct used for validating the fields of UpgradeBios.
type tempUpgradeBios  struct {
    Reboot  *bool   `json:"reboot,omitempty"`
    Version *string `json:"version,omitempty"`
}
