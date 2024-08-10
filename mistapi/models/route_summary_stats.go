package models

import (
    "encoding/json"
)

// RouteSummaryStats represents a RouteSummaryStats struct.
type RouteSummaryStats struct {
    FibRoutes                 *int           `json:"fib_routes,omitempty"`
    MaxUnicastRoutesSupported *int           `json:"max_unicast_routes_supported,omitempty"`
    RibRoutes                 *int           `json:"rib_routes,omitempty"`
    TotalRoutes               *int           `json:"total_routes,omitempty"`
    AdditionalProperties      map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RouteSummaryStats.
// It customizes the JSON marshaling process for RouteSummaryStats objects.
func (r RouteSummaryStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RouteSummaryStats object to a map representation for JSON marshaling.
func (r RouteSummaryStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.FibRoutes != nil {
        structMap["fib_routes"] = r.FibRoutes
    }
    if r.MaxUnicastRoutesSupported != nil {
        structMap["max_unicast_routes_supported"] = r.MaxUnicastRoutesSupported
    }
    if r.RibRoutes != nil {
        structMap["rib_routes"] = r.RibRoutes
    }
    if r.TotalRoutes != nil {
        structMap["total_routes"] = r.TotalRoutes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RouteSummaryStats.
// It customizes the JSON unmarshaling process for RouteSummaryStats objects.
func (r *RouteSummaryStats) UnmarshalJSON(input []byte) error {
    var temp tempRouteSummaryStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "fib_routes", "max_unicast_routes_supported", "rib_routes", "total_routes")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.FibRoutes = temp.FibRoutes
    r.MaxUnicastRoutesSupported = temp.MaxUnicastRoutesSupported
    r.RibRoutes = temp.RibRoutes
    r.TotalRoutes = temp.TotalRoutes
    return nil
}

// tempRouteSummaryStats is a temporary struct used for validating the fields of RouteSummaryStats.
type tempRouteSummaryStats  struct {
    FibRoutes                 *int `json:"fib_routes,omitempty"`
    MaxUnicastRoutesSupported *int `json:"max_unicast_routes_supported,omitempty"`
    RibRoutes                 *int `json:"rib_routes,omitempty"`
    TotalRoutes               *int `json:"total_routes,omitempty"`
}
