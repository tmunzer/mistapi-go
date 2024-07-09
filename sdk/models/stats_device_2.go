package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// StatsDevice2 represents a StatsDevice2 struct.
type StatsDevice2 struct {
    value          any
    isApStats      bool
    isSwitchStats  bool
    isGatewayStats bool
}

// String converts the StatsDevice2 object to a string representation.
func (s StatsDevice2) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for StatsDevice2.
// It customizes the JSON marshaling process for StatsDevice2 objects.
func (s StatsDevice2) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.StatsDeviceContainer.From*` functions to initialize the StatsDevice2 object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDevice2 object to a map representation for JSON marshaling.
func (s *StatsDevice2) toMap() any {
    switch obj := s.value.(type) {
    case *ApStats:
        return obj.toMap()
    case *SwitchStats:
        return obj.toMap()
    case *GatewayStats:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDevice2.
// It customizes the JSON unmarshaling process for StatsDevice2 objects.
func (s *StatsDevice2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ApStats{}, false, &s.isApStats),
        NewTypeHolder(&SwitchStats{}, false, &s.isSwitchStats),
        NewTypeHolder(&GatewayStats{}, false, &s.isGatewayStats),
    )
    
    s.value = result
    return err
}

func (s *StatsDevice2) AsApStats() (
    *ApStats,
    bool) {
    if !s.isApStats {
        return nil, false
    }
    return s.value.(*ApStats), true
}

func (s *StatsDevice2) AsSwitchStats() (
    *SwitchStats,
    bool) {
    if !s.isSwitchStats {
        return nil, false
    }
    return s.value.(*SwitchStats), true
}

func (s *StatsDevice2) AsGatewayStats() (
    *GatewayStats,
    bool) {
    if !s.isGatewayStats {
        return nil, false
    }
    return s.value.(*GatewayStats), true
}

// internalStatsDevice2 represents a statsDevice2 struct.
type internalStatsDevice2 struct {}

var StatsDeviceContainer internalStatsDevice2

// The internalStatsDevice2 instance, wrapping the provided ApStats value.
func (s *internalStatsDevice2) FromApStats(val ApStats) StatsDevice2 {
    return StatsDevice2{value: &val}
}

// The internalStatsDevice2 instance, wrapping the provided SwitchStats value.
func (s *internalStatsDevice2) FromSwitchStats(val SwitchStats) StatsDevice2 {
    return StatsDevice2{value: &val}
}

// The internalStatsDevice2 instance, wrapping the provided GatewayStats value.
func (s *internalStatsDevice2) FromGatewayStats(val GatewayStats) StatsDevice2 {
    return StatsDevice2{value: &val}
}
