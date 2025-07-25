// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanMistNac represents a WlanMistNac struct.
type WlanMistNac struct {
    // When enabled:
    // * `auth_servers` is ignored
    // * `acct_servers` is ignored
    // * `auth_servers_*` are ignored
    // * `coa_servers` is ignored
    // * `radsec` is ignored
    // * `coa_enabled` is assumed
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanMistNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanMistNac) String() string {
    return fmt.Sprintf(
    	"WlanMistNac[Enabled=%v, AdditionalProperties=%v]",
    	w.Enabled, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanMistNac.
// It customizes the JSON marshaling process for WlanMistNac objects.
func (w WlanMistNac) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanMistNac object to a map representation for JSON marshaling.
func (w WlanMistNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanMistNac.
// It customizes the JSON unmarshaling process for WlanMistNac objects.
func (w *WlanMistNac) UnmarshalJSON(input []byte) error {
    var temp tempWlanMistNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Enabled = temp.Enabled
    return nil
}

// tempWlanMistNac is a temporary struct used for validating the fields of WlanMistNac.
type tempWlanMistNac  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
