package models

import (
    "encoding/json"
)

// DhcpdStatLan represents a DhcpdStatLan struct.
type DhcpdStatLan struct {
    NumIps               *int           `json:"num_ips,omitempty"`
    NumLeased            *int           `json:"num_leased,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdStatLan.
// It customizes the JSON marshaling process for DhcpdStatLan objects.
func (d DhcpdStatLan) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdStatLan object to a map representation for JSON marshaling.
func (d DhcpdStatLan) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.NumIps != nil {
        structMap["num_ips"] = d.NumIps
    }
    if d.NumLeased != nil {
        structMap["num_leased"] = d.NumLeased
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdStatLan.
// It customizes the JSON unmarshaling process for DhcpdStatLan objects.
func (d *DhcpdStatLan) UnmarshalJSON(input []byte) error {
    var temp dhcpdStatLan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_ips", "num_leased")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.NumIps = temp.NumIps
    d.NumLeased = temp.NumLeased
    return nil
}

// dhcpdStatLan is a temporary struct used for validating the fields of DhcpdStatLan.
type dhcpdStatLan  struct {
    NumIps    *int `json:"num_ips,omitempty"`
    NumLeased *int `json:"num_leased,omitempty"`
}
