package models

import (
    "encoding/json"
)

// NetworkStaticNatProperty represents a NetworkStaticNatProperty struct.
type NetworkStaticNatProperty struct {
    InternalIp           *string                `json:"internal_ip,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // If not set, we configure the nat policies against all WAN ports for simplicity
    WanName              *string                `json:"wan_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkStaticNatProperty.
// It customizes the JSON marshaling process for NetworkStaticNatProperty objects.
func (n NetworkStaticNatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "internal_ip", "name", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkStaticNatProperty object to a map representation for JSON marshaling.
func (n NetworkStaticNatProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.InternalIp != nil {
        structMap["internal_ip"] = n.InternalIp
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.WanName != nil {
        structMap["wan_name"] = n.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkStaticNatProperty.
// It customizes the JSON unmarshaling process for NetworkStaticNatProperty objects.
func (n *NetworkStaticNatProperty) UnmarshalJSON(input []byte) error {
    var temp tempNetworkStaticNatProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "internal_ip", "name", "wan_name")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.InternalIp = temp.InternalIp
    n.Name = temp.Name
    n.WanName = temp.WanName
    return nil
}

// tempNetworkStaticNatProperty is a temporary struct used for validating the fields of NetworkStaticNatProperty.
type tempNetworkStaticNatProperty  struct {
    InternalIp *string `json:"internal_ip,omitempty"`
    Name       *string `json:"name,omitempty"`
    WanName    *string `json:"wan_name,omitempty"`
}
