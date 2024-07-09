package models

import (
    "encoding/json"
)

// TacacsAuthServer represents a TacacsAuthServer struct.
type TacacsAuthServer struct {
    Host                 *string        `json:"host,omitempty"`
    Port                 *string        `json:"port,omitempty"`
    Secret               *string        `json:"secret,omitempty"`
    Timeout              *int           `json:"timeout,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TacacsAuthServer.
// It customizes the JSON marshaling process for TacacsAuthServer objects.
func (t TacacsAuthServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TacacsAuthServer object to a map representation for JSON marshaling.
func (t TacacsAuthServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Host != nil {
        structMap["host"] = t.Host
    }
    if t.Port != nil {
        structMap["port"] = t.Port
    }
    if t.Secret != nil {
        structMap["secret"] = t.Secret
    }
    if t.Timeout != nil {
        structMap["timeout"] = t.Timeout
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TacacsAuthServer.
// It customizes the JSON unmarshaling process for TacacsAuthServer objects.
func (t *TacacsAuthServer) UnmarshalJSON(input []byte) error {
    var temp tacacsAuthServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "port", "secret", "timeout")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Host = temp.Host
    t.Port = temp.Port
    t.Secret = temp.Secret
    t.Timeout = temp.Timeout
    return nil
}

// tacacsAuthServer is a temporary struct used for validating the fields of TacacsAuthServer.
type tacacsAuthServer  struct {
    Host    *string `json:"host,omitempty"`
    Port    *string `json:"port,omitempty"`
    Secret  *string `json:"secret,omitempty"`
    Timeout *int    `json:"timeout,omitempty"`
}
