package models

import (
    "encoding/json"
    "fmt"
)

// NetworkSourceNat represents a NetworkSourceNat struct.
// if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub
type NetworkSourceNat struct {
    ExternalIp           *string                `json:"external_ip,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkSourceNat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkSourceNat) String() string {
    return fmt.Sprintf(
    	"NetworkSourceNat[ExternalIp=%v, AdditionalProperties=%v]",
    	n.ExternalIp, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkSourceNat.
// It customizes the JSON marshaling process for NetworkSourceNat objects.
func (n NetworkSourceNat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "external_ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkSourceNat object to a map representation for JSON marshaling.
func (n NetworkSourceNat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.ExternalIp != nil {
        structMap["external_ip"] = n.ExternalIp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkSourceNat.
// It customizes the JSON unmarshaling process for NetworkSourceNat objects.
func (n *NetworkSourceNat) UnmarshalJSON(input []byte) error {
    var temp tempNetworkSourceNat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "external_ip")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.ExternalIp = temp.ExternalIp
    return nil
}

// tempNetworkSourceNat is a temporary struct used for validating the fields of NetworkSourceNat.
type tempNetworkSourceNat  struct {
    ExternalIp *string `json:"external_ip,omitempty"`
}
