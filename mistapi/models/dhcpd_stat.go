package models

import (
    "encoding/json"
)

// DhcpdStat represents a DhcpdStat struct.
// Property key is the network name
type DhcpdStat struct {
    NumIps               *int           `json:"num_ips,omitempty"`
    NumLeased            *int           `json:"num_leased,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdStat.
// It customizes the JSON marshaling process for DhcpdStat objects.
func (d DhcpdStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdStat object to a map representation for JSON marshaling.
func (d DhcpdStat) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdStat.
// It customizes the JSON unmarshaling process for DhcpdStat objects.
func (d *DhcpdStat) UnmarshalJSON(input []byte) error {
    var temp tempDhcpdStat
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

// tempDhcpdStat is a temporary struct used for validating the fields of DhcpdStat.
type tempDhcpdStat  struct {
    NumIps    *int `json:"num_ips,omitempty"`
    NumLeased *int `json:"num_leased,omitempty"`
}
