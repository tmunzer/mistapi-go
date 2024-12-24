package models

import (
    "encoding/json"
    "fmt"
)

// DhcpdStatLan represents a DhcpdStatLan struct.
type DhcpdStatLan struct {
    NumIps               *int                   `json:"num_ips,omitempty"`
    NumLeased            *int                   `json:"num_leased,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DhcpdStatLan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DhcpdStatLan) String() string {
    return fmt.Sprintf(
    	"DhcpdStatLan[NumIps=%v, NumLeased=%v, AdditionalProperties=%v]",
    	d.NumIps, d.NumLeased, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DhcpdStatLan.
// It customizes the JSON marshaling process for DhcpdStatLan objects.
func (d DhcpdStatLan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "num_ips", "num_leased"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdStatLan object to a map representation for JSON marshaling.
func (d DhcpdStatLan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
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
    var temp tempDhcpdStatLan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_ips", "num_leased")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.NumIps = temp.NumIps
    d.NumLeased = temp.NumLeased
    return nil
}

// tempDhcpdStatLan is a temporary struct used for validating the fields of DhcpdStatLan.
type tempDhcpdStatLan  struct {
    NumIps    *int `json:"num_ips,omitempty"`
    NumLeased *int `json:"num_leased,omitempty"`
}
