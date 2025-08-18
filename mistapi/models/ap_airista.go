// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ApAirista represents a ApAirista struct.
type ApAirista struct {
    // Whether to enable Airista config
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Required if enabled, Airista server host
    Host                 Optional[string]       `json:"host"`
    Port                 Optional[int]          `json:"port"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApAirista,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApAirista) String() string {
    return fmt.Sprintf(
    	"ApAirista[Enabled=%v, Host=%v, Port=%v, AdditionalProperties=%v]",
    	a.Enabled, a.Host, a.Port, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApAirista.
// It customizes the JSON marshaling process for ApAirista objects.
func (a ApAirista) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "enabled", "host", "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApAirista object to a map representation for JSON marshaling.
func (a ApAirista) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Host.IsValueSet() {
        if a.Host.Value() != nil {
            structMap["host"] = a.Host.Value()
        } else {
            structMap["host"] = nil
        }
    }
    if a.Port.IsValueSet() {
        if a.Port.Value() != nil {
            structMap["port"] = a.Port.Value()
        } else {
            structMap["port"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApAirista.
// It customizes the JSON unmarshaling process for ApAirista objects.
func (a *ApAirista) UnmarshalJSON(input []byte) error {
    var temp tempApAirista
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "host", "port")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Enabled = temp.Enabled
    a.Host = temp.Host
    a.Port = temp.Port
    return nil
}

// tempApAirista is a temporary struct used for validating the fields of ApAirista.
type tempApAirista  struct {
    Enabled *bool            `json:"enabled,omitempty"`
    Host    Optional[string] `json:"host"`
    Port    Optional[int]    `json:"port"`
}
