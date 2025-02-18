package models

import (
    "encoding/json"
    "fmt"
)

// NetworkInternetAccessStaticNatProperty represents a NetworkInternetAccessStaticNatProperty struct.
type NetworkInternetAccessStaticNatProperty struct {
    // The Static NAT destination IP Address. Must be an IP Address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}")
    InternalIp           *string                `json:"internal_ip,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity. Can be a Variable (i.e. "{{myvar}}")
    WanName              *string                `json:"wan_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkInternetAccessStaticNatProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkInternetAccessStaticNatProperty) String() string {
    return fmt.Sprintf(
    	"NetworkInternetAccessStaticNatProperty[InternalIp=%v, Name=%v, WanName=%v, AdditionalProperties=%v]",
    	n.InternalIp, n.Name, n.WanName, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkInternetAccessStaticNatProperty.
// It customizes the JSON marshaling process for NetworkInternetAccessStaticNatProperty objects.
func (n NetworkInternetAccessStaticNatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "internal_ip", "name", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkInternetAccessStaticNatProperty object to a map representation for JSON marshaling.
func (n NetworkInternetAccessStaticNatProperty) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkInternetAccessStaticNatProperty.
// It customizes the JSON unmarshaling process for NetworkInternetAccessStaticNatProperty objects.
func (n *NetworkInternetAccessStaticNatProperty) UnmarshalJSON(input []byte) error {
    var temp tempNetworkInternetAccessStaticNatProperty
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

// tempNetworkInternetAccessStaticNatProperty is a temporary struct used for validating the fields of NetworkInternetAccessStaticNatProperty.
type tempNetworkInternetAccessStaticNatProperty  struct {
    InternalIp *string `json:"internal_ip,omitempty"`
    Name       *string `json:"name,omitempty"`
    WanName    *string `json:"wan_name,omitempty"`
}
