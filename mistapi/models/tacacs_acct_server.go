// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// TacacsAcctServer represents a TacacsAcctServer struct.
type TacacsAcctServer struct {
    Host                 *string                `json:"host,omitempty"`
    Port                 *string                `json:"port,omitempty"`
    Secret               *string                `json:"secret,omitempty"`
    Timeout              *int                   `json:"timeout,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TacacsAcctServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TacacsAcctServer) String() string {
    return fmt.Sprintf(
    	"TacacsAcctServer[Host=%v, Port=%v, Secret=%v, Timeout=%v, AdditionalProperties=%v]",
    	t.Host, t.Port, t.Secret, t.Timeout, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TacacsAcctServer.
// It customizes the JSON marshaling process for TacacsAcctServer objects.
func (t TacacsAcctServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "host", "port", "secret", "timeout"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TacacsAcctServer object to a map representation for JSON marshaling.
func (t TacacsAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for TacacsAcctServer.
// It customizes the JSON unmarshaling process for TacacsAcctServer objects.
func (t *TacacsAcctServer) UnmarshalJSON(input []byte) error {
    var temp tempTacacsAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "port", "secret", "timeout")
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

// tempTacacsAcctServer is a temporary struct used for validating the fields of TacacsAcctServer.
type tempTacacsAcctServer  struct {
    Host    *string `json:"host,omitempty"`
    Port    *string `json:"port,omitempty"`
    Secret  *string `json:"secret,omitempty"`
    Timeout *int    `json:"timeout,omitempty"`
}
