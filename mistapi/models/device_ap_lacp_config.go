package models

import (
    "encoding/json"
    "fmt"
)

// DeviceApLacpConfig represents a DeviceApLacpConfig struct.
type DeviceApLacpConfig struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceApLacpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceApLacpConfig) String() string {
    return fmt.Sprintf(
    	"DeviceApLacpConfig[Enabled=%v, AdditionalProperties=%v]",
    	d.Enabled, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceApLacpConfig.
// It customizes the JSON marshaling process for DeviceApLacpConfig objects.
func (d DeviceApLacpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceApLacpConfig object to a map representation for JSON marshaling.
func (d DeviceApLacpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Enabled != nil {
        structMap["enabled"] = d.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceApLacpConfig.
// It customizes the JSON unmarshaling process for DeviceApLacpConfig objects.
func (d *DeviceApLacpConfig) UnmarshalJSON(input []byte) error {
    var temp tempDeviceApLacpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Enabled = temp.Enabled
    return nil
}

// tempDeviceApLacpConfig is a temporary struct used for validating the fields of DeviceApLacpConfig.
type tempDeviceApLacpConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
