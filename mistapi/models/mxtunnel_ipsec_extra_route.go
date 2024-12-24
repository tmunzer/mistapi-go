package models

import (
    "encoding/json"
    "fmt"
)

// MxtunnelIpsecExtraRoute represents a MxtunnelIpsecExtraRoute struct.
type MxtunnelIpsecExtraRoute struct {
    Dest                 *string                `json:"dest,omitempty"`
    NextHop              *string                `json:"next_hop,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxtunnelIpsecExtraRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxtunnelIpsecExtraRoute) String() string {
    return fmt.Sprintf(
    	"MxtunnelIpsecExtraRoute[Dest=%v, NextHop=%v, AdditionalProperties=%v]",
    	m.Dest, m.NextHop, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxtunnelIpsecExtraRoute.
// It customizes the JSON marshaling process for MxtunnelIpsecExtraRoute objects.
func (m MxtunnelIpsecExtraRoute) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "dest", "next_hop"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxtunnelIpsecExtraRoute object to a map representation for JSON marshaling.
func (m MxtunnelIpsecExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Dest != nil {
        structMap["dest"] = m.Dest
    }
    if m.NextHop != nil {
        structMap["next_hop"] = m.NextHop
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxtunnelIpsecExtraRoute.
// It customizes the JSON unmarshaling process for MxtunnelIpsecExtraRoute objects.
func (m *MxtunnelIpsecExtraRoute) UnmarshalJSON(input []byte) error {
    var temp tempMxtunnelIpsecExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dest", "next_hop")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Dest = temp.Dest
    m.NextHop = temp.NextHop
    return nil
}

// tempMxtunnelIpsecExtraRoute is a temporary struct used for validating the fields of MxtunnelIpsecExtraRoute.
type tempMxtunnelIpsecExtraRoute  struct {
    Dest    *string `json:"dest,omitempty"`
    NextHop *string `json:"next_hop,omitempty"`
}
