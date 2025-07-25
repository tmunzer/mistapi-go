// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RadiusAcctPort represents a RadiusAcctPort struct.
// Radius Auth Port, value from 1 to 65535, default is 1813
type RadiusAcctPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RadiusAcctPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadiusAcctPort) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RadiusAcctPort.
// It customizes the JSON marshaling process for RadiusAcctPort objects.
func (r RadiusAcctPort) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RadiusAcctPortContainer.From*` functions to initialize the RadiusAcctPort object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAcctPort object to a map representation for JSON marshaling.
func (r *RadiusAcctPort) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusAcctPort.
// It customizes the JSON unmarshaling process for RadiusAcctPort objects.
func (r *RadiusAcctPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RadiusAcctPort) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RadiusAcctPort) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRadiusAcctPort represents a radiusAcctPort struct.
// Radius Auth Port, value from 1 to 65535, default is 1813
type internalRadiusAcctPort struct {}

var RadiusAcctPortContainer internalRadiusAcctPort

// The internalRadiusAcctPort instance, wrapping the provided int value.
func (r *internalRadiusAcctPort) FromNumber(val int) RadiusAcctPort {
    return RadiusAcctPort{value: &val}
}

// The internalRadiusAcctPort instance, wrapping the provided string value.
func (r *internalRadiusAcctPort) FromString(val string) RadiusAcctPort {
    return RadiusAcctPort{value: &val}
}
