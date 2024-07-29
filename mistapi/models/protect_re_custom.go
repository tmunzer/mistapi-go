package models

import (
    "encoding/json"
)

// ProtectReCustom represents a ProtectReCustom struct.
// custom acls
type ProtectReCustom struct {
    // matched dst port, "0" means any
    PortRange            *string                      `json:"port_range,omitempty"`
    // enum: `any`, `icmp`, `tcp`, `udp`
    Protocol             *ProtectReCustomProtocolEnum `json:"protocol,omitempty"`
    Subnet               []string                     `json:"subnet,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProtectReCustom.
// It customizes the JSON marshaling process for ProtectReCustom objects.
func (p ProtectReCustom) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProtectReCustom object to a map representation for JSON marshaling.
func (p ProtectReCustom) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PortRange != nil {
        structMap["port_range"] = p.PortRange
    }
    if p.Protocol != nil {
        structMap["protocol"] = p.Protocol
    }
    if p.Subnet != nil {
        structMap["subnet"] = p.Subnet
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProtectReCustom.
// It customizes the JSON unmarshaling process for ProtectReCustom objects.
func (p *ProtectReCustom) UnmarshalJSON(input []byte) error {
    var temp protectReCustom
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port_range", "protocol", "subnet")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.PortRange = temp.PortRange
    p.Protocol = temp.Protocol
    p.Subnet = temp.Subnet
    return nil
}

// protectReCustom is a temporary struct used for validating the fields of ProtectReCustom.
type protectReCustom  struct {
    PortRange *string                      `json:"port_range,omitempty"`
    Protocol  *ProtectReCustomProtocolEnum `json:"protocol,omitempty"`
    Subnet    []string                     `json:"subnet,omitempty"`
}
