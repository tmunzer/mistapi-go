// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeFpgaMulti represents a UpgradeFpgaMulti struct.
type UpgradeFpgaMulti struct {
    // List of device id to upgrade bios
    DeviceIds            []uuid.UUID            `json:"device_ids,omitempty"`
    // List of device model to upgrade bios
    Models               []string               `json:"models,omitempty"`
    // Reboot device immediately after upgrade is completed
    Reboot               *bool                  `json:"reboot,omitempty"`
    // Specific FPGA version
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeFpgaMulti,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeFpgaMulti) String() string {
    return fmt.Sprintf(
    	"UpgradeFpgaMulti[DeviceIds=%v, Models=%v, Reboot=%v, Version=%v, AdditionalProperties=%v]",
    	u.DeviceIds, u.Models, u.Reboot, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeFpgaMulti.
// It customizes the JSON marshaling process for UpgradeFpgaMulti objects.
func (u UpgradeFpgaMulti) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "device_ids", "models", "reboot", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeFpgaMulti object to a map representation for JSON marshaling.
func (u UpgradeFpgaMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.DeviceIds != nil {
        structMap["device_ids"] = u.DeviceIds
    }
    if u.Models != nil {
        structMap["models"] = u.Models
    }
    if u.Reboot != nil {
        structMap["reboot"] = u.Reboot
    }
    if u.Version != nil {
        structMap["version"] = u.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeFpgaMulti.
// It customizes the JSON unmarshaling process for UpgradeFpgaMulti objects.
func (u *UpgradeFpgaMulti) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeFpgaMulti
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_ids", "models", "reboot", "version")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.DeviceIds = temp.DeviceIds
    u.Models = temp.Models
    u.Reboot = temp.Reboot
    u.Version = temp.Version
    return nil
}

// tempUpgradeFpgaMulti is a temporary struct used for validating the fields of UpgradeFpgaMulti.
type tempUpgradeFpgaMulti  struct {
    DeviceIds []uuid.UUID `json:"device_ids,omitempty"`
    Models    []string    `json:"models,omitempty"`
    Reboot    *bool       `json:"reboot,omitempty"`
    Version   *string     `json:"version,omitempty"`
}
