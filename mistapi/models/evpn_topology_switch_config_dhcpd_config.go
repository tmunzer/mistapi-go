package models

import (
    "encoding/json"
    "fmt"
)

// EvpnTopologySwitchConfigDhcpdConfig represents a EvpnTopologySwitchConfigDhcpdConfig struct.
type EvpnTopologySwitchConfigDhcpdConfig struct {
    // If DHCPD is enabled on the switch
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologySwitchConfigDhcpdConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologySwitchConfigDhcpdConfig) String() string {
    return fmt.Sprintf(
    	"EvpnTopologySwitchConfigDhcpdConfig[Enabled=%v, AdditionalProperties=%v]",
    	e.Enabled, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitchConfigDhcpdConfig.
// It customizes the JSON marshaling process for EvpnTopologySwitchConfigDhcpdConfig objects.
func (e EvpnTopologySwitchConfigDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitchConfigDhcpdConfig object to a map representation for JSON marshaling.
func (e EvpnTopologySwitchConfigDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Enabled != nil {
        structMap["enabled"] = e.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitchConfigDhcpdConfig.
// It customizes the JSON unmarshaling process for EvpnTopologySwitchConfigDhcpdConfig objects.
func (e *EvpnTopologySwitchConfigDhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitchConfigDhcpdConfig
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

// tempEvpnTopologySwitchConfigDhcpdConfig is a temporary struct used for validating the fields of EvpnTopologySwitchConfigDhcpdConfig.
type tempEvpnTopologySwitchConfigDhcpdConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
