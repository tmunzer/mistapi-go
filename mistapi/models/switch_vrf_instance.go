// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchVrfInstance represents a SwitchVrfInstance struct.
type SwitchVrfInstance struct {
    // Property key is the destination subnet (e.g. "172.16.3.0/24")
    AggregateRoutes         map[string]AggregateRoute `json:"aggregate_routes,omitempty"`
    // Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64")
    AggregateRoutes6        map[string]AggregateRoute `json:"aggregate_routes6,omitempty"`
    EvpnAutoLoopbackSubnet  *string                   `json:"evpn_auto_loopback_subnet,omitempty"`
    EvpnAutoLoopbackSubnet6 *string                   `json:"evpn_auto_loopback_subnet6,omitempty"`
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes             map[string]VrfExtraRoute  `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6            map[string]VrfExtraRoute  `json:"extra_routes6,omitempty"`
    Networks                []string                  `json:"networks,omitempty"`
    AdditionalProperties    map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchVrfInstance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchVrfInstance) String() string {
    return fmt.Sprintf(
    	"SwitchVrfInstance[AggregateRoutes=%v, AggregateRoutes6=%v, EvpnAutoLoopbackSubnet=%v, EvpnAutoLoopbackSubnet6=%v, ExtraRoutes=%v, ExtraRoutes6=%v, Networks=%v, AdditionalProperties=%v]",
    	s.AggregateRoutes, s.AggregateRoutes6, s.EvpnAutoLoopbackSubnet, s.EvpnAutoLoopbackSubnet6, s.ExtraRoutes, s.ExtraRoutes6, s.Networks, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchVrfInstance.
// It customizes the JSON marshaling process for SwitchVrfInstance objects.
func (s SwitchVrfInstance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "aggregate_routes", "aggregate_routes6", "evpn_auto_loopback_subnet", "evpn_auto_loopback_subnet6", "extra_routes", "extra_routes6", "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchVrfInstance object to a map representation for JSON marshaling.
func (s SwitchVrfInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AggregateRoutes != nil {
        structMap["aggregate_routes"] = s.AggregateRoutes
    }
    if s.AggregateRoutes6 != nil {
        structMap["aggregate_routes6"] = s.AggregateRoutes6
    }
    if s.EvpnAutoLoopbackSubnet != nil {
        structMap["evpn_auto_loopback_subnet"] = s.EvpnAutoLoopbackSubnet
    }
    if s.EvpnAutoLoopbackSubnet6 != nil {
        structMap["evpn_auto_loopback_subnet6"] = s.EvpnAutoLoopbackSubnet6
    }
    if s.ExtraRoutes != nil {
        structMap["extra_routes"] = s.ExtraRoutes
    }
    if s.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = s.ExtraRoutes6
    }
    if s.Networks != nil {
        structMap["networks"] = s.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVrfInstance.
// It customizes the JSON unmarshaling process for SwitchVrfInstance objects.
func (s *SwitchVrfInstance) UnmarshalJSON(input []byte) error {
    var temp tempSwitchVrfInstance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aggregate_routes", "aggregate_routes6", "evpn_auto_loopback_subnet", "evpn_auto_loopback_subnet6", "extra_routes", "extra_routes6", "networks")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AggregateRoutes = temp.AggregateRoutes
    s.AggregateRoutes6 = temp.AggregateRoutes6
    s.EvpnAutoLoopbackSubnet = temp.EvpnAutoLoopbackSubnet
    s.EvpnAutoLoopbackSubnet6 = temp.EvpnAutoLoopbackSubnet6
    s.ExtraRoutes = temp.ExtraRoutes
    s.ExtraRoutes6 = temp.ExtraRoutes6
    s.Networks = temp.Networks
    return nil
}

// tempSwitchVrfInstance is a temporary struct used for validating the fields of SwitchVrfInstance.
type tempSwitchVrfInstance  struct {
    AggregateRoutes         map[string]AggregateRoute `json:"aggregate_routes,omitempty"`
    AggregateRoutes6        map[string]AggregateRoute `json:"aggregate_routes6,omitempty"`
    EvpnAutoLoopbackSubnet  *string                   `json:"evpn_auto_loopback_subnet,omitempty"`
    EvpnAutoLoopbackSubnet6 *string                   `json:"evpn_auto_loopback_subnet6,omitempty"`
    ExtraRoutes             map[string]VrfExtraRoute  `json:"extra_routes,omitempty"`
    ExtraRoutes6            map[string]VrfExtraRoute  `json:"extra_routes6,omitempty"`
    Networks                []string                  `json:"networks,omitempty"`
}
