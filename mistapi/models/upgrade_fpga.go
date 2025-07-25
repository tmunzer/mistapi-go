// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// UpgradeFpga represents a UpgradeFpga struct.
type UpgradeFpga struct {
    // Reboot device immediately after upgrade is completed
    Reboot               *bool                  `json:"reboot,omitempty"`
    // Specific fpga version
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeFpga,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeFpga) String() string {
    return fmt.Sprintf(
    	"UpgradeFpga[Reboot=%v, Version=%v, AdditionalProperties=%v]",
    	u.Reboot, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeFpga.
// It customizes the JSON marshaling process for UpgradeFpga objects.
func (u UpgradeFpga) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "reboot", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeFpga object to a map representation for JSON marshaling.
func (u UpgradeFpga) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeFpga.
// It customizes the JSON unmarshaling process for UpgradeFpga objects.
func (u *UpgradeFpga) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeFpga
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

// tempUpgradeFpga is a temporary struct used for validating the fields of UpgradeFpga.
type tempUpgradeFpga  struct {
    Reboot  *bool   `json:"reboot,omitempty"`
    Version *string `json:"version,omitempty"`
}
