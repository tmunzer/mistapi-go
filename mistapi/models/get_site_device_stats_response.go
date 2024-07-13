package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GetSiteDeviceStatsResponse represents a GetSiteDeviceStatsResponse struct.
// This is a container for any-of cases.
type GetSiteDeviceStatsResponse struct {
    value          any
    isApStats      bool
    isSwitchStats  bool
    isGatewayStats bool
}

// String converts the GetSiteDeviceStatsResponse object to a string representation.
func (g GetSiteDeviceStatsResponse) String() string {
    if bytes, err := json.Marshal(g.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for GetSiteDeviceStatsResponse.
// It customizes the JSON marshaling process for GetSiteDeviceStatsResponse objects.
func (g GetSiteDeviceStatsResponse) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GetSiteDeviceStatsResponseContainer.From*` functions to initialize the GetSiteDeviceStatsResponse object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetSiteDeviceStatsResponse object to a map representation for JSON marshaling.
func (g *GetSiteDeviceStatsResponse) toMap() any {
    switch obj := g.value.(type) {
    case *ApStats:
        return obj.toMap()
    case *SwitchStats:
        return obj.toMap()
    case *GatewayStats:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSiteDeviceStatsResponse.
// It customizes the JSON unmarshaling process for GetSiteDeviceStatsResponse objects.
func (g *GetSiteDeviceStatsResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApStats{}, false, &g.isApStats),
        NewTypeHolder(&SwitchStats{}, false, &g.isSwitchStats),
        NewTypeHolder(&GatewayStats{}, false, &g.isGatewayStats),
    )
    
    g.value = result
    return err
}

func (g *GetSiteDeviceStatsResponse) AsApStats() (
    *ApStats,
    bool) {
    if !g.isApStats {
        return nil, false
    }
    return g.value.(*ApStats), true
}

func (g *GetSiteDeviceStatsResponse) AsSwitchStats() (
    *SwitchStats,
    bool) {
    if !g.isSwitchStats {
        return nil, false
    }
    return g.value.(*SwitchStats), true
}

func (g *GetSiteDeviceStatsResponse) AsGatewayStats() (
    *GatewayStats,
    bool) {
    if !g.isGatewayStats {
        return nil, false
    }
    return g.value.(*GatewayStats), true
}

// internalGetSiteDeviceStatsResponse represents a getSiteDeviceStatsResponse struct.
// This is a container for any-of cases.
type internalGetSiteDeviceStatsResponse struct {}

var GetSiteDeviceStatsResponseContainer internalGetSiteDeviceStatsResponse

// The internalGetSiteDeviceStatsResponse instance, wrapping the provided ApStats value.
func (g *internalGetSiteDeviceStatsResponse) FromApStats(val ApStats) GetSiteDeviceStatsResponse {
    return GetSiteDeviceStatsResponse{value: &val}
}

// The internalGetSiteDeviceStatsResponse instance, wrapping the provided SwitchStats value.
func (g *internalGetSiteDeviceStatsResponse) FromSwitchStats(val SwitchStats) GetSiteDeviceStatsResponse {
    return GetSiteDeviceStatsResponse{value: &val}
}

// The internalGetSiteDeviceStatsResponse instance, wrapping the provided GatewayStats value.
func (g *internalGetSiteDeviceStatsResponse) FromGatewayStats(val GatewayStats) GetSiteDeviceStatsResponse {
    return GetSiteDeviceStatsResponse{value: &val}
}
