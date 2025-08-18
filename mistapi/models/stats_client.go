// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// StatsClient represents a StatsClient struct.
type StatsClient struct {
	value                 any
	isStatsWirelessClient bool
	isStatsWiredClient    bool
}

// String implements the fmt.Stringer interface for StatsClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsClient) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for StatsClient.
// It customizes the JSON marshaling process for StatsClient objects.
func (s StatsClient) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.StatsClientContainer.From*` functions to initialize the StatsClient object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsClient object to a map representation for JSON marshaling.
func (s *StatsClient) toMap() any {
	switch obj := s.value.(type) {
	case *StatsWirelessClient:
		return obj.toMap()
	case *StatsWiredClient:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClient.
// It customizes the JSON unmarshaling process for StatsClient objects.
func (s *StatsClient) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&StatsWirelessClient{}, false, &s.isStatsWirelessClient),
		NewTypeHolder(&StatsWiredClient{}, false, &s.isStatsWiredClient),
	)

	s.value = result
	return err
}

func (s *StatsClient) AsStatsWirelessClient() (
	*StatsWirelessClient,
	bool) {
	if !s.isStatsWirelessClient {
		return nil, false
	}
	return s.value.(*StatsWirelessClient), true
}

func (s *StatsClient) AsStatsWiredClient() (
	*StatsWiredClient,
	bool) {
	if !s.isStatsWiredClient {
		return nil, false
	}
	return s.value.(*StatsWiredClient), true
}

// internalStatsClient represents a statsClient struct.
type internalStatsClient struct{}

var StatsClientContainer internalStatsClient

// The internalStatsClient instance, wrapping the provided StatsWirelessClient value.
func (s *internalStatsClient) FromStatsWirelessClient(val StatsWirelessClient) StatsClient {
	return StatsClient{value: &val}
}

// The internalStatsClient instance, wrapping the provided StatsWiredClient value.
func (s *internalStatsClient) FromStatsWiredClient(val StatsWiredClient) StatsClient {
	return StatsClient{value: &val}
}
