// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RadiusCoaPort represents a RadiusCoaPort struct.
// Radius CoA Port, value from 1 to 65535, default is 3799
type RadiusCoaPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RadiusCoaPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadiusCoaPort) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RadiusCoaPort.
// It customizes the JSON marshaling process for RadiusCoaPort objects.
func (r RadiusCoaPort) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RadiusCoaPortContainer.From*` functions to initialize the RadiusCoaPort object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusCoaPort object to a map representation for JSON marshaling.
func (r *RadiusCoaPort) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusCoaPort.
// It customizes the JSON unmarshaling process for RadiusCoaPort objects.
func (r *RadiusCoaPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RadiusCoaPort) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RadiusCoaPort) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRadiusCoaPort represents a radiusCoaPort struct.
// Radius CoA Port, value from 1 to 65535, default is 3799
type internalRadiusCoaPort struct {}

var RadiusCoaPortContainer internalRadiusCoaPort

// The internalRadiusCoaPort instance, wrapping the provided int value.
func (r *internalRadiusCoaPort) FromNumber(val int) RadiusCoaPort {
    return RadiusCoaPort{value: &val}
}

// The internalRadiusCoaPort instance, wrapping the provided string value.
func (r *internalRadiusCoaPort) FromString(val string) RadiusCoaPort {
    return RadiusCoaPort{value: &val}
}
