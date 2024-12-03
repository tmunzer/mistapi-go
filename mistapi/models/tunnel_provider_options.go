package models

import (
    "encoding/json"
)

// TunnelProviderOptions represents a TunnelProviderOptions struct.
type TunnelProviderOptions struct {
    // for jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added
    Jse                  *TunnelProviderOptionsJse     `json:"jse,omitempty"`
    // for zscaler-ipsec and zscaler-gre
    Zscaler              *TunnelProviderOptionsZscaler `json:"zscaler,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptions.
// It customizes the JSON marshaling process for TunnelProviderOptions objects.
func (t TunnelProviderOptions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "jse", "zscaler"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptions object to a map representation for JSON marshaling.
func (t TunnelProviderOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Jse != nil {
        structMap["jse"] = t.Jse.toMap()
    }
    if t.Zscaler != nil {
        structMap["zscaler"] = t.Zscaler.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptions.
// It customizes the JSON unmarshaling process for TunnelProviderOptions objects.
func (t *TunnelProviderOptions) UnmarshalJSON(input []byte) error {
    var temp tempTunnelProviderOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "jse", "zscaler")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Jse = temp.Jse
    t.Zscaler = temp.Zscaler
    return nil
}

// tempTunnelProviderOptions is a temporary struct used for validating the fields of TunnelProviderOptions.
type tempTunnelProviderOptions  struct {
    Jse     *TunnelProviderOptionsJse     `json:"jse,omitempty"`
    Zscaler *TunnelProviderOptionsZscaler `json:"zscaler,omitempty"`
}
