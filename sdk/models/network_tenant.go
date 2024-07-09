package models

import (
    "encoding/json"
)

// NetworkTenant represents a NetworkTenant struct.
type NetworkTenant struct {
    Addresses            []string       `json:"addresses,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkTenant.
// It customizes the JSON marshaling process for NetworkTenant objects.
func (n NetworkTenant) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkTenant object to a map representation for JSON marshaling.
func (n NetworkTenant) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.Addresses != nil {
        structMap["addresses"] = n.Addresses
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkTenant.
// It customizes the JSON unmarshaling process for NetworkTenant objects.
func (n *NetworkTenant) UnmarshalJSON(input []byte) error {
    var temp networkTenant
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "addresses")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.Addresses = temp.Addresses
    return nil
}

// networkTenant is a temporary struct used for validating the fields of NetworkTenant.
type networkTenant  struct {
    Addresses []string `json:"addresses,omitempty"`
}
