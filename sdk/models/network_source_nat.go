package models

import (
    "encoding/json"
)

// NetworkSourceNat represents a NetworkSourceNat struct.
// if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub
type NetworkSourceNat struct {
    ExteralIp            *string        `json:"exteral_ip,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkSourceNat.
// It customizes the JSON marshaling process for NetworkSourceNat objects.
func (n NetworkSourceNat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkSourceNat object to a map representation for JSON marshaling.
func (n NetworkSourceNat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.ExteralIp != nil {
        structMap["exteral_ip"] = n.ExteralIp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkSourceNat.
// It customizes the JSON unmarshaling process for NetworkSourceNat objects.
func (n *NetworkSourceNat) UnmarshalJSON(input []byte) error {
    var temp networkSourceNat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "exteral_ip")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.ExteralIp = temp.ExteralIp
    return nil
}

// networkSourceNat is a temporary struct used for validating the fields of NetworkSourceNat.
type networkSourceNat  struct {
    ExteralIp *string `json:"exteral_ip,omitempty"`
}
