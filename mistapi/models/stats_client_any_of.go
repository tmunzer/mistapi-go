package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// StatsClientAnyOf represents a StatsClientAnyOf struct.
type StatsClientAnyOf struct {
    value                        any
    isArrayOfClientWirelessStats bool
    isStatsWiredClient           bool
}

// String converts the StatsClientAnyOf object to a string representation.
func (s StatsClientAnyOf) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for StatsClientAnyOf.
// It customizes the JSON marshaling process for StatsClientAnyOf objects.
func (s StatsClientAnyOf) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.StatsClientAnyOfContainer.From*` functions to initialize the StatsClientAnyOf object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsClientAnyOf object to a map representation for JSON marshaling.
func (s *StatsClientAnyOf) toMap() any {
    switch obj := s.value.(type) {
    case *[]ClientWirelessStats:
        return *obj
    case *StatsWiredClient:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClientAnyOf.
// It customizes the JSON unmarshaling process for StatsClientAnyOf objects.
func (s *StatsClientAnyOf) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&[]ClientWirelessStats{}, false, &s.isArrayOfClientWirelessStats),
        NewTypeHolder(&StatsWiredClient{}, false, &s.isStatsWiredClient),
    )
    
    s.value = result
    return err
}

func (s *StatsClientAnyOf) AsArrayOfClientWirelessStats() (
    *[]ClientWirelessStats,
    bool) {
    if !s.isArrayOfClientWirelessStats {
        return nil, false
    }
    return s.value.(*[]ClientWirelessStats), true
}

func (s *StatsClientAnyOf) AsStatsWiredClient() (
    *StatsWiredClient,
    bool) {
    if !s.isStatsWiredClient {
        return nil, false
    }
    return s.value.(*StatsWiredClient), true
}

// internalStatsClientAnyOf represents a statsClientAnyOf struct.
type internalStatsClientAnyOf struct {}

var StatsClientAnyOfContainer internalStatsClientAnyOf

// The internalStatsClientAnyOf instance, wrapping the provided []ClientWirelessStats value.
func (s *internalStatsClientAnyOf) FromArrayOfClientWirelessStats(val []ClientWirelessStats) StatsClientAnyOf {
    return StatsClientAnyOf{value: &val}
}

// The internalStatsClientAnyOf instance, wrapping the provided StatsWiredClient value.
func (s *internalStatsClientAnyOf) FromStatsWiredClient(val StatsWiredClient) StatsClientAnyOf {
    return StatsClientAnyOf{value: &val}
}
