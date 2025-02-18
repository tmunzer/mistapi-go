package models

import (
    "encoding/json"
    "fmt"
)

// EvpnTopologySwitchConfigVrfConfig represents a EvpnTopologySwitchConfigVrfConfig struct.
type EvpnTopologySwitchConfigVrfConfig struct {
    // Whether to enable VRF (when supported on the device)
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologySwitchConfigVrfConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologySwitchConfigVrfConfig) String() string {
    return fmt.Sprintf(
    	"EvpnTopologySwitchConfigVrfConfig[Enabled=%v, AdditionalProperties=%v]",
    	e.Enabled, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitchConfigVrfConfig.
// It customizes the JSON marshaling process for EvpnTopologySwitchConfigVrfConfig objects.
func (e EvpnTopologySwitchConfigVrfConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitchConfigVrfConfig object to a map representation for JSON marshaling.
func (e EvpnTopologySwitchConfigVrfConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Enabled != nil {
        structMap["enabled"] = e.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitchConfigVrfConfig.
// It customizes the JSON unmarshaling process for EvpnTopologySwitchConfigVrfConfig objects.
func (e *EvpnTopologySwitchConfigVrfConfig) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitchConfigVrfConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Enabled = temp.Enabled
    return nil
}

// tempEvpnTopologySwitchConfigVrfConfig is a temporary struct used for validating the fields of EvpnTopologySwitchConfigVrfConfig.
type tempEvpnTopologySwitchConfigVrfConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
