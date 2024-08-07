package models

import (
    "encoding/json"
)

// SsoMxedgeProxyAuthServer represents a SsoMxedgeProxyAuthServer struct.
type SsoMxedgeProxyAuthServer struct {
    Host                 *string        `json:"host,omitempty"`
    Port                 *int           `json:"port,omitempty"`
    Secret               *string        `json:"secret,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsoMxedgeProxyAuthServer.
// It customizes the JSON marshaling process for SsoMxedgeProxyAuthServer objects.
func (s SsoMxedgeProxyAuthServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SsoMxedgeProxyAuthServer object to a map representation for JSON marshaling.
func (s SsoMxedgeProxyAuthServer) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SsoMxedgeProxyAuthServer.
// It customizes the JSON unmarshaling process for SsoMxedgeProxyAuthServer objects.
func (s *SsoMxedgeProxyAuthServer) UnmarshalJSON(input []byte) error {
    var temp ssoMxedgeProxyAuthServer
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

// ssoMxedgeProxyAuthServer is a temporary struct used for validating the fields of SsoMxedgeProxyAuthServer.
type ssoMxedgeProxyAuthServer  struct {
    Host   *string `json:"host,omitempty"`
    Port   *int    `json:"port,omitempty"`
    Secret *string `json:"secret,omitempty"`
}
