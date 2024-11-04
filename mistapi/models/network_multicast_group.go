package models

import (
    "encoding/json"
)

// NetworkMulticastGroup represents a NetworkMulticastGroup struct.
type NetworkMulticastGroup struct {
    // RP (rendezvous point) IP Address
    RpIp                 *string        `json:"rp_ip,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkMulticastGroup.
// It customizes the JSON marshaling process for NetworkMulticastGroup objects.
func (n NetworkMulticastGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkMulticastGroup object to a map representation for JSON marshaling.
func (n NetworkMulticastGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.RpIp != nil {
        structMap["rp_ip"] = n.RpIp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkMulticastGroup.
// It customizes the JSON unmarshaling process for NetworkMulticastGroup objects.
func (n *NetworkMulticastGroup) UnmarshalJSON(input []byte) error {
    var temp tempNetworkMulticastGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "rp_ip")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.RpIp = temp.RpIp
    return nil
}

// tempNetworkMulticastGroup is a temporary struct used for validating the fields of NetworkMulticastGroup.
type tempNetworkMulticastGroup  struct {
    RpIp *string `json:"rp_ip,omitempty"`
}
