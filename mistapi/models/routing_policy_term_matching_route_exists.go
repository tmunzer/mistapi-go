package models

import (
    "encoding/json"
)

// RoutingPolicyTermMatchingRouteExists represents a RoutingPolicyTermMatchingRouteExists struct.
type RoutingPolicyTermMatchingRouteExists struct {
    Route                *string        `json:"route,omitempty"`
    // name of the vrf instance
    // it can also be the name of the VPN or wan if they
    VrfName              *string        `json:"vrf_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermMatchingRouteExists.
// It customizes the JSON marshaling process for RoutingPolicyTermMatchingRouteExists objects.
func (r RoutingPolicyTermMatchingRouteExists) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermMatchingRouteExists object to a map representation for JSON marshaling.
func (r RoutingPolicyTermMatchingRouteExists) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Route != nil {
        structMap["route"] = r.Route
    }
    if r.VrfName != nil {
        structMap["vrf_name"] = r.VrfName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyTermMatchingRouteExists.
// It customizes the JSON unmarshaling process for RoutingPolicyTermMatchingRouteExists objects.
func (r *RoutingPolicyTermMatchingRouteExists) UnmarshalJSON(input []byte) error {
    var temp tempRoutingPolicyTermMatchingRouteExists
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "route", "vrf_name")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Route = temp.Route
    r.VrfName = temp.VrfName
    return nil
}

// tempRoutingPolicyTermMatchingRouteExists is a temporary struct used for validating the fields of RoutingPolicyTermMatchingRouteExists.
type tempRoutingPolicyTermMatchingRouteExists  struct {
    Route   *string `json:"route,omitempty"`
    VrfName *string `json:"vrf_name,omitempty"`
}
