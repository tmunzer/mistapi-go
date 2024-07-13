package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListSiteDevicesStatsResponse represents a ListSiteDevicesStatsResponse struct.
// This is Array of a container for any-of cases.
type ListSiteDevicesStatsResponse struct {
    value          any
    isApStats      bool
    isSwitchStats  bool
    isGatewayStats bool
}

// String converts the ListSiteDevicesStatsResponse object to a string representation.
func (l ListSiteDevicesStatsResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListSiteDevicesStatsResponse.
// It customizes the JSON marshaling process for ListSiteDevicesStatsResponse objects.
func (l ListSiteDevicesStatsResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListSiteDevicesStatsResponseContainer.From*` functions to initialize the ListSiteDevicesStatsResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSiteDevicesStatsResponse object to a map representation for JSON marshaling.
func (l *ListSiteDevicesStatsResponse) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ListSiteDevicesStatsResponse.
// It customizes the JSON unmarshaling process for ListSiteDevicesStatsResponse objects.
func (l *ListSiteDevicesStatsResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApStats{}, false, &l.isApStats),
        NewTypeHolder(&SwitchStats{}, false, &l.isSwitchStats),
        NewTypeHolder(&GatewayStats{}, false, &l.isGatewayStats),
    )
    
    l.value = result
    return err
}

func (l *ListSiteDevicesStatsResponse) AsApStats() (
    *ApStats,
    bool) {
    if !l.isApStats {
        return nil, false
    }
    return l.value.(*ApStats), true
}

func (l *ListSiteDevicesStatsResponse) AsSwitchStats() (
    *SwitchStats,
    bool) {
    if !l.isSwitchStats {
        return nil, false
    }
    return l.value.(*SwitchStats), true
}

func (l *ListSiteDevicesStatsResponse) AsGatewayStats() (
    *GatewayStats,
    bool) {
    if !l.isGatewayStats {
        return nil, false
    }
    return l.value.(*GatewayStats), true
}

// internalListSiteDevicesStatsResponse represents a listSiteDevicesStatsResponse struct.
// This is Array of a container for any-of cases.
type internalListSiteDevicesStatsResponse struct {}

var ListSiteDevicesStatsResponseContainer internalListSiteDevicesStatsResponse

// The internalListSiteDevicesStatsResponse instance, wrapping the provided ApStats value.
func (l *internalListSiteDevicesStatsResponse) FromApStats(val ApStats) ListSiteDevicesStatsResponse {
    return ListSiteDevicesStatsResponse{value: &val}
}

// The internalListSiteDevicesStatsResponse instance, wrapping the provided SwitchStats value.
func (l *internalListSiteDevicesStatsResponse) FromSwitchStats(val SwitchStats) ListSiteDevicesStatsResponse {
    return ListSiteDevicesStatsResponse{value: &val}
}

// The internalListSiteDevicesStatsResponse instance, wrapping the provided GatewayStats value.
func (l *internalListSiteDevicesStatsResponse) FromGatewayStats(val GatewayStats) ListSiteDevicesStatsResponse {
    return ListSiteDevicesStatsResponse{value: &val}
}
