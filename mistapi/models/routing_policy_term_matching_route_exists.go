package models

import (
    "encoding/json"
    "fmt"
)

// RoutingPolicyTermMatchingRouteExists represents a RoutingPolicyTermMatchingRouteExists struct.
type RoutingPolicyTermMatchingRouteExists struct {
    Route                *string                `json:"route,omitempty"`
    // Name of the vrf instance, it can also be the name of the VPN or wan if they
    VrfName              *string                `json:"vrf_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RoutingPolicyTermMatchingRouteExists,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicyTermMatchingRouteExists) String() string {
    return fmt.Sprintf(
    	"RoutingPolicyTermMatchingRouteExists[Route=%v, VrfName=%v, AdditionalProperties=%v]",
    	r.Route, r.VrfName, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermMatchingRouteExists.
// It customizes the JSON marshaling process for RoutingPolicyTermMatchingRouteExists objects.
func (r RoutingPolicyTermMatchingRouteExists) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "route", "vrf_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermMatchingRouteExists object to a map representation for JSON marshaling.
func (r RoutingPolicyTermMatchingRouteExists) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "route", "vrf_name")
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
