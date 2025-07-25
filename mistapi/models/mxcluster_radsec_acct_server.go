// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxclusterRadsecAcctServer represents a MxclusterRadsecAcctServer struct.
type MxclusterRadsecAcctServer struct {
    // IP / hostname of RADIUS server
    Host                 *string                `json:"host,omitempty"`
    // Acct port of RADIUS server
    Port                 *int                   `json:"port,omitempty"`
    // Secret of RADIUS server
    Secret               *string                `json:"secret,omitempty"`
    // List of ssids that will use this server if match_ssid is true and match is found
    Ssids                []string               `json:"ssids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterRadsecAcctServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterRadsecAcctServer) String() string {
    return fmt.Sprintf(
    	"MxclusterRadsecAcctServer[Host=%v, Port=%v, Secret=%v, Ssids=%v, AdditionalProperties=%v]",
    	m.Host, m.Port, m.Secret, m.Ssids, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsecAcctServer.
// It customizes the JSON marshaling process for MxclusterRadsecAcctServer objects.
func (m MxclusterRadsecAcctServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "host", "port", "secret", "ssids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsecAcctServer object to a map representation for JSON marshaling.
func (m MxclusterRadsecAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMxclusterRadsecAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "port", "secret", "ssids")
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

// tempMxclusterRadsecAcctServer is a temporary struct used for validating the fields of MxclusterRadsecAcctServer.
type tempMxclusterRadsecAcctServer  struct {
    Host   *string  `json:"host,omitempty"`
    Port   *int     `json:"port,omitempty"`
    Secret *string  `json:"secret,omitempty"`
    Ssids  []string `json:"ssids,omitempty"`
}
