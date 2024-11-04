package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ServiceMaxLatency represents a ServiceMaxLatency struct.
// for SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
type ServiceMaxLatency struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the ServiceMaxLatency object to a string representation.
func (s ServiceMaxLatency) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ServiceMaxLatency.
// It customizes the JSON marshaling process for ServiceMaxLatency objects.
func (s ServiceMaxLatency) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ServiceMaxLatencyContainer.From*` functions to initialize the ServiceMaxLatency object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceMaxLatency object to a map representation for JSON marshaling.
func (s *ServiceMaxLatency) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceMaxLatency.
// It customizes the JSON unmarshaling process for ServiceMaxLatency objects.
func (s *ServiceMaxLatency) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *ServiceMaxLatency) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *ServiceMaxLatency) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalServiceMaxLatency represents a serviceMaxLatency struct.
// for SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
type internalServiceMaxLatency struct {}

var ServiceMaxLatencyContainer internalServiceMaxLatency

// The internalServiceMaxLatency instance, wrapping the provided string value.
func (s *internalServiceMaxLatency) FromString(val string) ServiceMaxLatency {
    return ServiceMaxLatency{value: &val}
}

// The internalServiceMaxLatency instance, wrapping the provided int value.
func (s *internalServiceMaxLatency) FromNumber(val int) ServiceMaxLatency {
    return ServiceMaxLatency{value: &val}
}
