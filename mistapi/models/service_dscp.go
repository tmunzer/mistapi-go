// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ServiceDscp represents a ServiceDscp struct.
// For SSR only, when `traffic_type`==`custom`. 0-63 or variable
type ServiceDscp struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for ServiceDscp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServiceDscp) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ServiceDscp.
// It customizes the JSON marshaling process for ServiceDscp objects.
func (s ServiceDscp) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ServiceDscpContainer.From*` functions to initialize the ServiceDscp object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServiceDscp object to a map representation for JSON marshaling.
func (s *ServiceDscp) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceDscp.
// It customizes the JSON unmarshaling process for ServiceDscp objects.
func (s *ServiceDscp) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *ServiceDscp) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *ServiceDscp) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalServiceDscp represents a serviceDscp struct.
// For SSR only, when `traffic_type`==`custom`. 0-63 or variable
type internalServiceDscp struct{}

var ServiceDscpContainer internalServiceDscp

// The internalServiceDscp instance, wrapping the provided string value.
func (s *internalServiceDscp) FromString(val string) ServiceDscp {
	return ServiceDscp{value: &val}
}

// The internalServiceDscp instance, wrapping the provided int value.
func (s *internalServiceDscp) FromNumber(val int) ServiceDscp {
	return ServiceDscp{value: &val}
}
