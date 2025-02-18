package models

import (
    "encoding/json"
    "fmt"
)

// NetworkInternetAccessDestinationNatProperty represents a NetworkInternetAccessDestinationNatProperty struct.
type NetworkInternetAccessDestinationNatProperty struct {
    // The Destination NAT destination IP Address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}")
    InternalIp           *string                `json:"internal_ip,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // The Destination NAT destination IP Address. Must be a Port (i.e. "443") or a Variable (i.e. "{{myvar}}")
    Port                 *string                `json:"port,omitempty"`
    // SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity
    WanName              *string                `json:"wan_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkInternetAccessDestinationNatProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkInternetAccessDestinationNatProperty) String() string {
    return fmt.Sprintf(
    	"NetworkInternetAccessDestinationNatProperty[InternalIp=%v, Name=%v, Port=%v, WanName=%v, AdditionalProperties=%v]",
    	n.InternalIp, n.Name, n.Port, n.WanName, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkInternetAccessDestinationNatProperty.
// It customizes the JSON marshaling process for NetworkInternetAccessDestinationNatProperty objects.
func (n NetworkInternetAccessDestinationNatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "internal_ip", "name", "port", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkInternetAccessDestinationNatProperty object to a map representation for JSON marshaling.
func (n NetworkInternetAccessDestinationNatProperty) toMap() map[string]any {
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
    if n.WanName != nil {
        structMap["wan_name"] = n.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkInternetAccessDestinationNatProperty.
// It customizes the JSON unmarshaling process for NetworkInternetAccessDestinationNatProperty objects.
func (n *NetworkInternetAccessDestinationNatProperty) UnmarshalJSON(input []byte) error {
    var temp tempNetworkInternetAccessDestinationNatProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "internal_ip", "name", "port", "wan_name")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.InternalIp = temp.InternalIp
    n.Name = temp.Name
    n.Port = temp.Port
    n.WanName = temp.WanName
    return nil
}

// tempNetworkInternetAccessDestinationNatProperty is a temporary struct used for validating the fields of NetworkInternetAccessDestinationNatProperty.
type tempNetworkInternetAccessDestinationNatProperty  struct {
    InternalIp *string `json:"internal_ip,omitempty"`
    Name       *string `json:"name,omitempty"`
    Port       *string `json:"port,omitempty"`
    WanName    *string `json:"wan_name,omitempty"`
}
