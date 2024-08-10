package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// UpgradeBiosMulti represents a UpgradeBiosMulti struct.
type UpgradeBiosMulti struct {
    // list of device id to upgrade bios
    DeviceIds            []uuid.UUID    `json:"device_ids,omitempty"`
    // list of device model to upgrade bios
    Models               []string       `json:"models,omitempty"`
    // Reboot device immediately after upgrade is completed
    Reboot               *bool          `json:"reboot,omitempty"`
    // specific bios version
    Version              *string        `json:"version,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpgradeBiosMulti.
// It customizes the JSON marshaling process for UpgradeBiosMulti objects.
func (u UpgradeBiosMulti) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeBiosMulti object to a map representation for JSON marshaling.
func (u UpgradeBiosMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeBiosMulti.
// It customizes the JSON unmarshaling process for UpgradeBiosMulti objects.
func (u *UpgradeBiosMulti) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeBiosMulti
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_ids", "models", "reboot", "version")
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

// tempUpgradeBiosMulti is a temporary struct used for validating the fields of UpgradeBiosMulti.
type tempUpgradeBiosMulti  struct {
    DeviceIds []uuid.UUID `json:"device_ids,omitempty"`
    Models    []string    `json:"models,omitempty"`
    Reboot    *bool       `json:"reboot,omitempty"`
    Version   *string     `json:"version,omitempty"`
}
