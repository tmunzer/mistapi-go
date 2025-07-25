// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxclusterNac represents a MxclusterNac struct.
type MxclusterNac struct {
    AcctServerPort       *int                            `json:"acct_server_port,omitempty"`
    AuthServerPort       *int                            `json:"auth_server_port,omitempty"`
    // Property key is the RADIUS Client IP/Subnet.
    ClientIps            map[string]MxclusterNacClientIp `json:"client_ips,omitempty"`
    Enabled              *bool                           `json:"enabled,omitempty"`
    Secret               *string                         `json:"secret,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterNac) String() string {
    return fmt.Sprintf(
    	"MxclusterNac[AcctServerPort=%v, AuthServerPort=%v, ClientIps=%v, Enabled=%v, Secret=%v, AdditionalProperties=%v]",
    	m.AcctServerPort, m.AuthServerPort, m.ClientIps, m.Enabled, m.Secret, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterNac.
// It customizes the JSON marshaling process for MxclusterNac objects.
func (m MxclusterNac) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "acct_server_port", "auth_server_port", "client_ips", "enabled", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterNac object to a map representation for JSON marshaling.
func (m MxclusterNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AcctServerPort != nil {
        structMap["acct_server_port"] = m.AcctServerPort
    }
    if m.AuthServerPort != nil {
        structMap["auth_server_port"] = m.AuthServerPort
    }
    if m.ClientIps != nil {
        structMap["client_ips"] = m.ClientIps
    }
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.Secret != nil {
        structMap["secret"] = m.Secret
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterNac.
// It customizes the JSON unmarshaling process for MxclusterNac objects.
func (m *MxclusterNac) UnmarshalJSON(input []byte) error {
    var temp tempMxclusterNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_server_port", "auth_server_port", "client_ips", "enabled", "secret")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AcctServerPort = temp.AcctServerPort
    m.AuthServerPort = temp.AuthServerPort
    m.ClientIps = temp.ClientIps
    m.Enabled = temp.Enabled
    m.Secret = temp.Secret
    return nil
}

// tempMxclusterNac is a temporary struct used for validating the fields of MxclusterNac.
type tempMxclusterNac  struct {
    AcctServerPort *int                            `json:"acct_server_port,omitempty"`
    AuthServerPort *int                            `json:"auth_server_port,omitempty"`
    ClientIps      map[string]MxclusterNacClientIp `json:"client_ips,omitempty"`
    Enabled        *bool                           `json:"enabled,omitempty"`
    Secret         *string                         `json:"secret,omitempty"`
}
