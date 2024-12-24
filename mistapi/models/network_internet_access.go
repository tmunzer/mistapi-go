package models

import (
    "encoding/json"
    "fmt"
)

// NetworkInternetAccess represents a NetworkInternetAccess struct.
// whether this network has direct internet access
type NetworkInternetAccess struct {
    CreateSimpleServicePolicy *bool                                    `json:"create_simple_service_policy,omitempty"`
    // Property key may be an IP/Port (i.e. "63.16.0.3:443"), or a port (i.e. ":2222")
    DestinationNat            map[string]NetworkDestinationNatProperty `json:"destination_nat,omitempty"`
    Enabled                   *bool                                    `json:"enabled,omitempty"`
    // by default, all access is allowed, to only allow certain traffic, make `restricted`=`true` and define service_policies
    Restricted                *bool                                    `json:"restricted,omitempty"`
    // Property key may be an IP Address (i.e. "172.16.0.1"), and IP Address and Port (i.e. "172.16.0.1:8443") or a CIDR (i.e. "172.16.0.12/20")
    StaticNat                 map[string]NetworkStaticNatProperty      `json:"static_nat,omitempty"`
    AdditionalProperties      map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkInternetAccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkInternetAccess) String() string {
    return fmt.Sprintf(
    	"NetworkInternetAccess[CreateSimpleServicePolicy=%v, DestinationNat=%v, Enabled=%v, Restricted=%v, StaticNat=%v, AdditionalProperties=%v]",
    	n.CreateSimpleServicePolicy, n.DestinationNat, n.Enabled, n.Restricted, n.StaticNat, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkInternetAccess.
// It customizes the JSON marshaling process for NetworkInternetAccess objects.
func (n NetworkInternetAccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "create_simple_service_policy", "destination_nat", "enabled", "restricted", "static_nat"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkInternetAccess object to a map representation for JSON marshaling.
func (n NetworkInternetAccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.CreateSimpleServicePolicy != nil {
        structMap["create_simple_service_policy"] = n.CreateSimpleServicePolicy
    }
    if n.DestinationNat != nil {
        structMap["destination_nat"] = n.DestinationNat
    }
    if n.Enabled != nil {
        structMap["enabled"] = n.Enabled
    }
    if n.Restricted != nil {
        structMap["restricted"] = n.Restricted
    }
    if n.StaticNat != nil {
        structMap["static_nat"] = n.StaticNat
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkInternetAccess.
// It customizes the JSON unmarshaling process for NetworkInternetAccess objects.
func (n *NetworkInternetAccess) UnmarshalJSON(input []byte) error {
    var temp tempNetworkInternetAccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "create_simple_service_policy", "destination_nat", "enabled", "restricted", "static_nat")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.CreateSimpleServicePolicy = temp.CreateSimpleServicePolicy
    n.DestinationNat = temp.DestinationNat
    n.Enabled = temp.Enabled
    n.Restricted = temp.Restricted
    n.StaticNat = temp.StaticNat
    return nil
}

// tempNetworkInternetAccess is a temporary struct used for validating the fields of NetworkInternetAccess.
type tempNetworkInternetAccess  struct {
    CreateSimpleServicePolicy *bool                                    `json:"create_simple_service_policy,omitempty"`
    DestinationNat            map[string]NetworkDestinationNatProperty `json:"destination_nat,omitempty"`
    Enabled                   *bool                                    `json:"enabled,omitempty"`
    Restricted                *bool                                    `json:"restricted,omitempty"`
    StaticNat                 map[string]NetworkStaticNatProperty      `json:"static_nat,omitempty"`
}
