package models

import (
    "encoding/json"
)

// ApClientBridgeAuth represents a ApClientBridgeAuth struct.
type ApClientBridgeAuth struct {
    Psk                  *string                     `json:"psk,omitempty"`
    // wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`
    Type                 *ApClientBridgeAuthTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApClientBridgeAuth.
// It customizes the JSON marshaling process for ApClientBridgeAuth objects.
func (a ApClientBridgeAuth) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApClientBridgeAuth object to a map representation for JSON marshaling.
func (a ApClientBridgeAuth) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Psk != nil {
        structMap["psk"] = a.Psk
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApClientBridgeAuth.
// It customizes the JSON unmarshaling process for ApClientBridgeAuth objects.
func (a *ApClientBridgeAuth) UnmarshalJSON(input []byte) error {
    var temp apClientBridgeAuth
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "psk", "type")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Psk = temp.Psk
    a.Type = temp.Type
    return nil
}

// apClientBridgeAuth is a temporary struct used for validating the fields of ApClientBridgeAuth.
type apClientBridgeAuth  struct {
    Psk  *string                     `json:"psk,omitempty"`
    Type *ApClientBridgeAuthTypeEnum `json:"type,omitempty"`
}
