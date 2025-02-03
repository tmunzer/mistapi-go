package models

import (
    "encoding/json"
    "fmt"
)

// WlanCiscoCwa represents a WlanCiscoCwa struct.
// Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html
type WlanCiscoCwa struct {
    // List of hostnames without http(s):// (matched by substring)
    AllowedHostnames     []string               `json:"allowed_hostnames,omitempty"`
    // List of CIDRs
    AllowedSubnets       []string               `json:"allowed_subnets,omitempty"`
    // List of blocked CIDRs
    BlockedSubnets       []string               `json:"blocked_subnets,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanCiscoCwa,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanCiscoCwa) String() string {
    return fmt.Sprintf(
    	"WlanCiscoCwa[AllowedHostnames=%v, AllowedSubnets=%v, BlockedSubnets=%v, Enabled=%v, AdditionalProperties=%v]",
    	w.AllowedHostnames, w.AllowedSubnets, w.BlockedSubnets, w.Enabled, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanCiscoCwa.
// It customizes the JSON marshaling process for WlanCiscoCwa objects.
func (w WlanCiscoCwa) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "allowed_hostnames", "allowed_subnets", "blocked_subnets", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanCiscoCwa object to a map representation for JSON marshaling.
func (w WlanCiscoCwa) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AllowedHostnames != nil {
        structMap["allowed_hostnames"] = w.AllowedHostnames
    }
    if w.AllowedSubnets != nil {
        structMap["allowed_subnets"] = w.AllowedSubnets
    }
    if w.BlockedSubnets != nil {
        structMap["blocked_subnets"] = w.BlockedSubnets
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanCiscoCwa.
// It customizes the JSON unmarshaling process for WlanCiscoCwa objects.
func (w *WlanCiscoCwa) UnmarshalJSON(input []byte) error {
    var temp tempWlanCiscoCwa
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allowed_hostnames", "allowed_subnets", "blocked_subnets", "enabled")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AllowedHostnames = temp.AllowedHostnames
    w.AllowedSubnets = temp.AllowedSubnets
    w.BlockedSubnets = temp.BlockedSubnets
    w.Enabled = temp.Enabled
    return nil
}

// tempWlanCiscoCwa is a temporary struct used for validating the fields of WlanCiscoCwa.
type tempWlanCiscoCwa  struct {
    AllowedHostnames []string `json:"allowed_hostnames,omitempty"`
    AllowedSubnets   []string `json:"allowed_subnets,omitempty"`
    BlockedSubnets   []string `json:"blocked_subnets,omitempty"`
    Enabled          *bool    `json:"enabled,omitempty"`
}
