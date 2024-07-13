package models

import (
    "encoding/json"
)

// MxtunnelIpsecExtraRoute represents a MxtunnelIpsecExtraRoute struct.
type MxtunnelIpsecExtraRoute struct {
    Dest                 *string        `json:"dest,omitempty"`
    NextHop              *string        `json:"next_hop,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxtunnelIpsecExtraRoute.
// It customizes the JSON marshaling process for MxtunnelIpsecExtraRoute objects.
func (m MxtunnelIpsecExtraRoute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxtunnelIpsecExtraRoute object to a map representation for JSON marshaling.
func (m MxtunnelIpsecExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp mxtunnelIpsecExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dest", "next_hop")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Dest = temp.Dest
    m.NextHop = temp.NextHop
    return nil
}

// mxtunnelIpsecExtraRoute is a temporary struct used for validating the fields of MxtunnelIpsecExtraRoute.
type mxtunnelIpsecExtraRoute  struct {
    Dest    *string `json:"dest,omitempty"`
    NextHop *string `json:"next_hop,omitempty"`
}
