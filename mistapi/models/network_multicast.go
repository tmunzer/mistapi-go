package models

import (
    "encoding/json"
)

// NetworkMulticast represents a NetworkMulticast struct.
// whether to enable multicast support (only PIM-sparse mode is supported)
type NetworkMulticast struct {
    // if the network will only be the soruce of the multicast traffic, IGMP can be disabled
    DisableIgmp          *bool                            `json:"disable_igmp,omitempty"`
    Enabled              *bool                            `json:"enabled,omitempty"`
    // Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32")
    Groups               map[string]NetworkMulticastGroup `json:"groups,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkMulticast.
// It customizes the JSON marshaling process for NetworkMulticast objects.
func (n NetworkMulticast) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkMulticast object to a map representation for JSON marshaling.
func (n NetworkMulticast) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.DisableIgmp != nil {
        structMap["disable_igmp"] = n.DisableIgmp
    }
    if n.Enabled != nil {
        structMap["enabled"] = n.Enabled
    }
    if n.Groups != nil {
        structMap["groups"] = n.Groups
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkMulticast.
// It customizes the JSON unmarshaling process for NetworkMulticast objects.
func (n *NetworkMulticast) UnmarshalJSON(input []byte) error {
    var temp tempNetworkMulticast
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disable_igmp", "enabled", "groups")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.DisableIgmp = temp.DisableIgmp
    n.Enabled = temp.Enabled
    n.Groups = temp.Groups
    return nil
}

// tempNetworkMulticast is a temporary struct used for validating the fields of NetworkMulticast.
type tempNetworkMulticast  struct {
    DisableIgmp *bool                            `json:"disable_igmp,omitempty"`
    Enabled     *bool                            `json:"enabled,omitempty"`
    Groups      map[string]NetworkMulticastGroup `json:"groups,omitempty"`
}
