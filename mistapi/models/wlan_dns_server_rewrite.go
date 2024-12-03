package models

import (
    "encoding/json"
)

// WlanDnsServerRewrite represents a WlanDnsServerRewrite struct.
// for radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns)
type WlanDnsServerRewrite struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    // map between radius_group and the desired DNS server (IPv4 only)
    // Property key is the RADIUS group, property value is the desired DNS Server
    RadiusGroups         map[string]string      `json:"radius_groups,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanDnsServerRewrite.
// It customizes the JSON marshaling process for WlanDnsServerRewrite objects.
func (w WlanDnsServerRewrite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "enabled", "radius_groups"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDnsServerRewrite object to a map representation for JSON marshaling.
func (w WlanDnsServerRewrite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.RadiusGroups != nil {
        structMap["radius_groups"] = w.RadiusGroups
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDnsServerRewrite.
// It customizes the JSON unmarshaling process for WlanDnsServerRewrite objects.
func (w *WlanDnsServerRewrite) UnmarshalJSON(input []byte) error {
    var temp tempWlanDnsServerRewrite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "radius_groups")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Enabled = temp.Enabled
    w.RadiusGroups = temp.RadiusGroups
    return nil
}

// tempWlanDnsServerRewrite is a temporary struct used for validating the fields of WlanDnsServerRewrite.
type tempWlanDnsServerRewrite  struct {
    Enabled      *bool             `json:"enabled,omitempty"`
    RadiusGroups map[string]string `json:"radius_groups,omitempty"`
}
