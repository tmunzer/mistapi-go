package models

import (
    "encoding/json"
    "fmt"
)

// ProtectReCustom represents a ProtectReCustom struct.
// Custom acls
type ProtectReCustom struct {
    // Matched dst port, "0" means any
    PortRange            *string                      `json:"port_range,omitempty"`
    // enum: `any`, `icmp`, `tcp`, `udp`
    Protocol             *ProtectReCustomProtocolEnum `json:"protocol,omitempty"`
    Subnets              []string                     `json:"subnets,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for ProtectReCustom,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p ProtectReCustom) String() string {
    return fmt.Sprintf(
    	"ProtectReCustom[PortRange=%v, Protocol=%v, Subnets=%v, AdditionalProperties=%v]",
    	p.PortRange, p.Protocol, p.Subnets, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ProtectReCustom.
// It customizes the JSON marshaling process for ProtectReCustom objects.
func (p ProtectReCustom) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "port_range", "protocol", "subnets"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProtectReCustom object to a map representation for JSON marshaling.
func (p ProtectReCustom) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PortRange != nil {
        structMap["port_range"] = p.PortRange
    }
    if p.Protocol != nil {
        structMap["protocol"] = p.Protocol
    }
    if p.Subnets != nil {
        structMap["subnets"] = p.Subnets
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProtectReCustom.
// It customizes the JSON unmarshaling process for ProtectReCustom objects.
func (p *ProtectReCustom) UnmarshalJSON(input []byte) error {
    var temp tempProtectReCustom
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port_range", "protocol", "subnets")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.PortRange = temp.PortRange
    p.Protocol = temp.Protocol
    p.Subnets = temp.Subnets
    return nil
}

// tempProtectReCustom is a temporary struct used for validating the fields of ProtectReCustom.
type tempProtectReCustom  struct {
    PortRange *string                      `json:"port_range,omitempty"`
    Protocol  *ProtectReCustomProtocolEnum `json:"protocol,omitempty"`
    Subnets   []string                     `json:"subnets,omitempty"`
}
