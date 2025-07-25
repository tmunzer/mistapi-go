// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// NetworkMulticast represents a NetworkMulticast struct.
// Whether to enable multicast support (only PIM-sparse mode is supported)
type NetworkMulticast struct {
    // If the network will only be the source of the multicast traffic, IGMP can be disabled
    DisableIgmp          *bool                            `json:"disable_igmp,omitempty"`
    Enabled              *bool                            `json:"enabled,omitempty"`
    // Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32")
    Groups               map[string]NetworkMulticastGroup `json:"groups,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkMulticast,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkMulticast) String() string {
    return fmt.Sprintf(
    	"NetworkMulticast[DisableIgmp=%v, Enabled=%v, Groups=%v, AdditionalProperties=%v]",
    	n.DisableIgmp, n.Enabled, n.Groups, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkMulticast.
// It customizes the JSON marshaling process for NetworkMulticast objects.
func (n NetworkMulticast) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "disable_igmp", "enabled", "groups"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkMulticast object to a map representation for JSON marshaling.
func (n NetworkMulticast) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_igmp", "enabled", "groups")
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
