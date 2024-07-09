package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListOrgDevicesStatsResponse represents a ListOrgDevicesStatsResponse struct.
// This is Array of a container for any-of cases.
type ListOrgDevicesStatsResponse struct {
    value          any
    isApStats      bool
    isSwitchStats  bool
    isGatewayStats bool
}

// String converts the ListOrgDevicesStatsResponse object to a string representation.
func (l ListOrgDevicesStatsResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListOrgDevicesStatsResponse.
// It customizes the JSON marshaling process for ListOrgDevicesStatsResponse objects.
func (l ListOrgDevicesStatsResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListOrgDevicesStatsResponseContainer.From*` functions to initialize the ListOrgDevicesStatsResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListOrgDevicesStatsResponse object to a map representation for JSON marshaling.
func (l *ListOrgDevicesStatsResponse) toMap() any {
    switch obj := l.value.(type) {
    case *ApStats:
        return obj.toMap()
    case *SwitchStats:
        return obj.toMap()
    case *GatewayStats:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListOrgDevicesStatsResponse.
// It customizes the JSON unmarshaling process for ListOrgDevicesStatsResponse objects.
func (l *ListOrgDevicesStatsResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApStats{}, false, &l.isApStats),
        NewTypeHolder(&SwitchStats{}, false, &l.isSwitchStats),
        NewTypeHolder(&GatewayStats{}, false, &l.isGatewayStats),
    )
    
    l.value = result
    return err
}

func (l *ListOrgDevicesStatsResponse) AsApStats() (
    *ApStats,
    bool) {
    if !l.isApStats {
        return nil, false
    }
    return l.value.(*ApStats), true
}

func (l *ListOrgDevicesStatsResponse) AsSwitchStats() (
    *SwitchStats,
    bool) {
    if !l.isSwitchStats {
        return nil, false
    }
    return l.value.(*SwitchStats), true
}

func (l *ListOrgDevicesStatsResponse) AsGatewayStats() (
    *GatewayStats,
    bool) {
    if !l.isGatewayStats {
        return nil, false
    }
    return l.value.(*GatewayStats), true
}

// internalListOrgDevicesStatsResponse represents a listOrgDevicesStatsResponse struct.
// This is Array of a container for any-of cases.
type internalListOrgDevicesStatsResponse struct {}

var ListOrgDevicesStatsResponseContainer internalListOrgDevicesStatsResponse

// The internalListOrgDevicesStatsResponse instance, wrapping the provided ApStats value.
func (l *internalListOrgDevicesStatsResponse) FromApStats(val ApStats) ListOrgDevicesStatsResponse {
    return ListOrgDevicesStatsResponse{value: &val}
}

// The internalListOrgDevicesStatsResponse instance, wrapping the provided SwitchStats value.
func (l *internalListOrgDevicesStatsResponse) FromSwitchStats(val SwitchStats) ListOrgDevicesStatsResponse {
    return ListOrgDevicesStatsResponse{value: &val}
}

// The internalListOrgDevicesStatsResponse instance, wrapping the provided GatewayStats value.
func (l *internalListOrgDevicesStatsResponse) FromGatewayStats(val GatewayStats) ListOrgDevicesStatsResponse {
    return ListOrgDevicesStatsResponse{value: &val}
}
