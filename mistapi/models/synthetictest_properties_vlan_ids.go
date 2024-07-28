package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SynthetictestPropertiesVlanIds represents a SynthetictestPropertiesVlanIds struct.
// This is Array of a container for one-of cases.
type SynthetictestPropertiesVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SynthetictestPropertiesVlanIds object to a string representation.
func (s SynthetictestPropertiesVlanIds) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestPropertiesVlanIds.
// It customizes the JSON marshaling process for SynthetictestPropertiesVlanIds objects.
func (s SynthetictestPropertiesVlanIds) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SynthetictestPropertiesVlanIdsContainer.From*` functions to initialize the SynthetictestPropertiesVlanIds object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestPropertiesVlanIds object to a map representation for JSON marshaling.
func (s *SynthetictestPropertiesVlanIds) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestPropertiesVlanIds.
// It customizes the JSON unmarshaling process for SynthetictestPropertiesVlanIds objects.
func (s *SynthetictestPropertiesVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SynthetictestPropertiesVlanIds) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SynthetictestPropertiesVlanIds) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSynthetictestPropertiesVlanIds represents a synthetictestPropertiesVlanIds struct.
// This is Array of a container for one-of cases.
type internalSynthetictestPropertiesVlanIds struct {}

var SynthetictestPropertiesVlanIdsContainer internalSynthetictestPropertiesVlanIds

// The internalSynthetictestPropertiesVlanIds instance, wrapping the provided string value.
func (s *internalSynthetictestPropertiesVlanIds) FromString(val string) SynthetictestPropertiesVlanIds {
    return SynthetictestPropertiesVlanIds{value: &val}
}

// The internalSynthetictestPropertiesVlanIds instance, wrapping the provided int value.
func (s *internalSynthetictestPropertiesVlanIds) FromNumber(val int) SynthetictestPropertiesVlanIds {
    return SynthetictestPropertiesVlanIds{value: &val}
}
