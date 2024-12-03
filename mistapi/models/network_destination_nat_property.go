package models

import (
    "encoding/json"
)

// NetworkDestinationNatProperty represents a NetworkDestinationNatProperty struct.
type NetworkDestinationNatProperty struct {
    InternalIp           *string                `json:"internal_ip,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Port                 *int                   `json:"port,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkDestinationNatProperty.
// It customizes the JSON marshaling process for NetworkDestinationNatProperty objects.
func (n NetworkDestinationNatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "internal_ip", "name", "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkDestinationNatProperty object to a map representation for JSON marshaling.
func (n NetworkDestinationNatProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.InternalIp != nil {
        structMap["internal_ip"] = n.InternalIp
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.Port != nil {
        structMap["port"] = n.Port
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkDestinationNatProperty.
// It customizes the JSON unmarshaling process for NetworkDestinationNatProperty objects.
func (n *NetworkDestinationNatProperty) UnmarshalJSON(input []byte) error {
    var temp tempNetworkDestinationNatProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "internal_ip", "name", "port")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.InternalIp = temp.InternalIp
    n.Name = temp.Name
    n.Port = temp.Port
    return nil
}

// tempNetworkDestinationNatProperty is a temporary struct used for validating the fields of NetworkDestinationNatProperty.
type tempNetworkDestinationNatProperty  struct {
    InternalIp *string `json:"internal_ip,omitempty"`
    Name       *string `json:"name,omitempty"`
    Port       *int    `json:"port,omitempty"`
}
