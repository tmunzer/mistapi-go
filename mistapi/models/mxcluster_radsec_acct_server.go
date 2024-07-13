package models

import (
    "encoding/json"
)

// MxclusterRadsecAcctServer represents a MxclusterRadsecAcctServer struct.
type MxclusterRadsecAcctServer struct {
    // ip / hostname of RADIUS server
    Host                 *string        `json:"host,omitempty"`
    // Acct port of RADIUS server
    Port                 *int           `json:"port,omitempty"`
    // secret of RADIUS server
    Secret               *string        `json:"secret,omitempty"`
    // list of ssids that will use this server if match_ssid is true and match is found
    Ssids                []string       `json:"ssids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsecAcctServer.
// It customizes the JSON marshaling process for MxclusterRadsecAcctServer objects.
func (m MxclusterRadsecAcctServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsecAcctServer object to a map representation for JSON marshaling.
func (m MxclusterRadsecAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Host != nil {
        structMap["host"] = m.Host
    }
    if m.Port != nil {
        structMap["port"] = m.Port
    }
    if m.Secret != nil {
        structMap["secret"] = m.Secret
    }
    if m.Ssids != nil {
        structMap["ssids"] = m.Ssids
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsecAcctServer.
// It customizes the JSON unmarshaling process for MxclusterRadsecAcctServer objects.
func (m *MxclusterRadsecAcctServer) UnmarshalJSON(input []byte) error {
    var temp mxclusterRadsecAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "port", "secret", "ssids")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Host = temp.Host
    m.Port = temp.Port
    m.Secret = temp.Secret
    m.Ssids = temp.Ssids
    return nil
}

// mxclusterRadsecAcctServer is a temporary struct used for validating the fields of MxclusterRadsecAcctServer.
type mxclusterRadsecAcctServer  struct {
    Host   *string  `json:"host,omitempty"`
    Port   *int     `json:"port,omitempty"`
    Secret *string  `json:"secret,omitempty"`
    Ssids  []string `json:"ssids,omitempty"`
}
