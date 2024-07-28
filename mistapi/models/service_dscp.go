package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ServiceDscp represents a ServiceDscp struct.
// This is a container for one-of cases.
type ServiceDscp struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the ServiceDscp object to a string representation.
func (s ServiceDscp) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
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
// This is a container for one-of cases.
type internalServiceDscp struct {}

var ServiceDscpContainer internalServiceDscp

// The internalServiceDscp instance, wrapping the provided string value.
func (s *internalServiceDscp) FromString(val string) ServiceDscp {
    return ServiceDscp{value: &val}
}

// The internalServiceDscp instance, wrapping the provided int value.
func (s *internalServiceDscp) FromNumber(val int) ServiceDscp {
    return ServiceDscp{value: &val}
}
