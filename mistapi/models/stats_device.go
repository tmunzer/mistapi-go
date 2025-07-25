// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// StatsDevice represents a StatsDevice struct.
type StatsDevice struct {
    value          any
    isStatsAp      bool
    isStatsSwitch  bool
    isStatsGateway bool
}

// String implements the fmt.Stringer interface for StatsDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsDevice) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for StatsDevice.
// It customizes the JSON marshaling process for StatsDevice objects.
func (s StatsDevice) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.StatsDeviceContainer.From*` functions to initialize the StatsDevice object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDevice object to a map representation for JSON marshaling.
func (s *StatsDevice) toMap() any {
    switch obj := s.value.(type) {
    case *StatsAp:
        return obj.toMap()
    case *StatsSwitch:
        return obj.toMap()
    case *StatsGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDevice.
// It customizes the JSON unmarshaling process for StatsDevice objects.
func (s *StatsDevice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&StatsAp{}, false, &s.isStatsAp),
        NewTypeHolder(&StatsSwitch{}, false, &s.isStatsSwitch),
        NewTypeHolder(&StatsGateway{}, false, &s.isStatsGateway),
    )
    
    s.value = result
    return err
}

func (s *StatsDevice) AsStatsAp() (
    *StatsAp,
    bool) {
    if !s.isStatsAp {
        return nil, false
    }
    return s.value.(*StatsAp), true
}

func (s *StatsDevice) AsStatsSwitch() (
    *StatsSwitch,
    bool) {
    if !s.isStatsSwitch {
        return nil, false
    }
    return s.value.(*StatsSwitch), true
}

func (s *StatsDevice) AsStatsGateway() (
    *StatsGateway,
    bool) {
    if !s.isStatsGateway {
        return nil, false
    }
    return s.value.(*StatsGateway), true
}

// internalStatsDevice represents a statsDevice struct.
type internalStatsDevice struct {}

var StatsDeviceContainer internalStatsDevice

// The internalStatsDevice instance, wrapping the provided StatsAp value.
func (s *internalStatsDevice) FromStatsAp(val StatsAp) StatsDevice {
    return StatsDevice{value: &val}
}

// The internalStatsDevice instance, wrapping the provided StatsSwitch value.
func (s *internalStatsDevice) FromStatsSwitch(val StatsSwitch) StatsDevice {
    return StatsDevice{value: &val}
}

// The internalStatsDevice instance, wrapping the provided StatsGateway value.
func (s *internalStatsDevice) FromStatsGateway(val StatsGateway) StatsDevice {
    return StatsDevice{value: &val}
}
