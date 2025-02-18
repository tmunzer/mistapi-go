package models

import (
    "encoding/json"
    "fmt"
)

// DhcpdConfig represents a DhcpdConfig struct.
type DhcpdConfig struct {
    // If set to `false`, disable the DHCP server
    Enabled              *bool                          `json:"enabled,omitempty"`
    AdditionalProperties map[string]DhcpdConfigProperty `json:"_"`
}

// String implements the fmt.Stringer interface for DhcpdConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DhcpdConfig) String() string {
    return fmt.Sprintf(
    	"DhcpdConfig[Enabled=%v, AdditionalProperties=%v]",
    	d.Enabled, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfig.
// It customizes the JSON marshaling process for DhcpdConfig objects.
func (d DhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfig object to a map representation for JSON marshaling.
func (d DhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Enabled != nil {
        structMap["enabled"] = d.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfig.
// It customizes the JSON unmarshaling process for DhcpdConfig objects.
func (d *DhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempDhcpdConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[DhcpdConfigProperty](input, "enabled")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Enabled = temp.Enabled
    return nil
}

// tempDhcpdConfig is a temporary struct used for validating the fields of DhcpdConfig.
type tempDhcpdConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
