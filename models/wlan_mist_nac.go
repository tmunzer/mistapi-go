package models

import (
    "encoding/json"
)

// WlanMistNac represents a WlanMistNac struct.
type WlanMistNac struct {
    // when enabled:
    // * `auth_servers` is ignored
    // * `acct_servers` is ignored
    // * `auth_servers_*` are ignored
    // * `coa_servers` is ignored
    // * `radsec` is ignored
    // * `coa_enabled` is assumed
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanMistNac.
// It customizes the JSON marshaling process for WlanMistNac objects.
func (w WlanMistNac) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanMistNac object to a map representation for JSON marshaling.
func (w WlanMistNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanMistNac.
// It customizes the JSON unmarshaling process for WlanMistNac objects.
func (w *WlanMistNac) UnmarshalJSON(input []byte) error {
    var temp wlanMistNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Enabled = temp.Enabled
    return nil
}

// wlanMistNac is a temporary struct used for validating the fields of WlanMistNac.
type wlanMistNac  struct {
    Enabled *bool `json:"enabled,omitempty"`
}