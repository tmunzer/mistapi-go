package models

import (
    "encoding/json"
)

// ApClientBridge represents a ApClientBridge struct.
type ApClientBridge struct {
    Auth                 *ApClientBridgeAuth `json:"auth,omitempty"`
    // when acted as client bridge:
    // * only 5G radio can be used
    // * will not serve as AP on any radios
    Enabled              *bool               `json:"enabled,omitempty"`
    Ssid                 *string             `json:"ssid,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApClientBridge.
// It customizes the JSON marshaling process for ApClientBridge objects.
func (a ApClientBridge) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApClientBridge object to a map representation for JSON marshaling.
func (a ApClientBridge) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Auth != nil {
        structMap["auth"] = a.Auth.toMap()
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Ssid != nil {
        structMap["ssid"] = a.Ssid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApClientBridge.
// It customizes the JSON unmarshaling process for ApClientBridge objects.
func (a *ApClientBridge) UnmarshalJSON(input []byte) error {
    var temp apClientBridge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth", "enabled", "ssid")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Auth = temp.Auth
    a.Enabled = temp.Enabled
    a.Ssid = temp.Ssid
    return nil
}

// apClientBridge is a temporary struct used for validating the fields of ApClientBridge.
type apClientBridge  struct {
    Auth    *ApClientBridgeAuth `json:"auth,omitempty"`
    Enabled *bool               `json:"enabled,omitempty"`
    Ssid    *string             `json:"ssid,omitempty"`
}
