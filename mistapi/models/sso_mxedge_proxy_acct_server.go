package models

import (
    "encoding/json"
)

// SsoMxedgeProxyAcctServer represents a SsoMxedgeProxyAcctServer struct.
type SsoMxedgeProxyAcctServer struct {
    Host                 *string        `json:"host,omitempty"`
    Port                 *int           `json:"port,omitempty"`
    Secret               *string        `json:"secret,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsoMxedgeProxyAcctServer.
// It customizes the JSON marshaling process for SsoMxedgeProxyAcctServer objects.
func (s SsoMxedgeProxyAcctServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SsoMxedgeProxyAcctServer object to a map representation for JSON marshaling.
func (s SsoMxedgeProxyAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Host != nil {
        structMap["host"] = s.Host
    }
    if s.Port != nil {
        structMap["port"] = s.Port
    }
    if s.Secret != nil {
        structMap["secret"] = s.Secret
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoMxedgeProxyAcctServer.
// It customizes the JSON unmarshaling process for SsoMxedgeProxyAcctServer objects.
func (s *SsoMxedgeProxyAcctServer) UnmarshalJSON(input []byte) error {
    var temp tempSsoMxedgeProxyAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "port", "secret")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Host = temp.Host
    s.Port = temp.Port
    s.Secret = temp.Secret
    return nil
}

// tempSsoMxedgeProxyAcctServer is a temporary struct used for validating the fields of SsoMxedgeProxyAcctServer.
type tempSsoMxedgeProxyAcctServer  struct {
    Host   *string `json:"host,omitempty"`
    Port   *int    `json:"port,omitempty"`
    Secret *string `json:"secret,omitempty"`
}
