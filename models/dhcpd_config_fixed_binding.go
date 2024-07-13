package models

import (
    "encoding/json"
)

// DhcpdConfigFixedBinding represents a DhcpdConfigFixedBinding struct.
type DhcpdConfigFixedBinding struct {
    Ip                   *string        `json:"ip,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigFixedBinding.
// It customizes the JSON marshaling process for DhcpdConfigFixedBinding objects.
func (d DhcpdConfigFixedBinding) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigFixedBinding object to a map representation for JSON marshaling.
func (d DhcpdConfigFixedBinding) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Ip != nil {
        structMap["ip"] = d.Ip
    }
    if d.Name != nil {
        structMap["name"] = d.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigFixedBinding.
// It customizes the JSON unmarshaling process for DhcpdConfigFixedBinding objects.
func (d *DhcpdConfigFixedBinding) UnmarshalJSON(input []byte) error {
    var temp dhcpdConfigFixedBinding
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "name")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Ip = temp.Ip
    d.Name = temp.Name
    return nil
}

// dhcpdConfigFixedBinding is a temporary struct used for validating the fields of DhcpdConfigFixedBinding.
type dhcpdConfigFixedBinding  struct {
    Ip   *string `json:"ip,omitempty"`
    Name *string `json:"name,omitempty"`
}