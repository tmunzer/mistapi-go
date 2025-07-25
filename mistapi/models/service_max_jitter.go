// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ServiceMaxJitter represents a ServiceMaxJitter struct.
// For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
type ServiceMaxJitter struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for ServiceMaxJitter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServiceMaxJitter) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ServiceMaxJitter.
// It customizes the JSON marshaling process for ServiceMaxJitter objects.
func (s ServiceMaxJitter) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ServiceMaxJitterContainer.From*` functions to initialize the ServiceMaxJitter object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceMaxJitter object to a map representation for JSON marshaling.
func (s *ServiceMaxJitter) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceMaxJitter.
// It customizes the JSON unmarshaling process for ServiceMaxJitter objects.
func (s *ServiceMaxJitter) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *ServiceMaxJitter) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *ServiceMaxJitter) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalServiceMaxJitter represents a serviceMaxJitter struct.
// For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
type internalServiceMaxJitter struct {}

var ServiceMaxJitterContainer internalServiceMaxJitter

// The internalServiceMaxJitter instance, wrapping the provided string value.
func (s *internalServiceMaxJitter) FromString(val string) ServiceMaxJitter {
    return ServiceMaxJitter{value: &val}
}

// The internalServiceMaxJitter instance, wrapping the provided int value.
func (s *internalServiceMaxJitter) FromNumber(val int) ServiceMaxJitter {
    return ServiceMaxJitter{value: &val}
}
