package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ServiceMaxLoss represents a ServiceMaxLoss struct.
// for SSR only, when `traffic_type`==`custom`, for uplink selection. 0-100 or variable
type ServiceMaxLoss struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the ServiceMaxLoss object to a string representation.
func (s ServiceMaxLoss) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ServiceMaxLoss.
// It customizes the JSON marshaling process for ServiceMaxLoss objects.
func (s ServiceMaxLoss) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ServiceMaxLossContainer.From*` functions to initialize the ServiceMaxLoss object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceMaxLoss object to a map representation for JSON marshaling.
func (s *ServiceMaxLoss) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceMaxLoss.
// It customizes the JSON unmarshaling process for ServiceMaxLoss objects.
func (s *ServiceMaxLoss) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *ServiceMaxLoss) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *ServiceMaxLoss) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalServiceMaxLoss represents a serviceMaxLoss struct.
// for SSR only, when `traffic_type`==`custom`, for uplink selection. 0-100 or variable
type internalServiceMaxLoss struct {}

var ServiceMaxLossContainer internalServiceMaxLoss

// The internalServiceMaxLoss instance, wrapping the provided string value.
func (s *internalServiceMaxLoss) FromString(val string) ServiceMaxLoss {
    return ServiceMaxLoss{value: &val}
}

// The internalServiceMaxLoss instance, wrapping the provided int value.
func (s *internalServiceMaxLoss) FromNumber(val int) ServiceMaxLoss {
    return ServiceMaxLoss{value: &val}
}
