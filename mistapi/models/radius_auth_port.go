package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RadiusAuthPort represents a RadiusAuthPort struct.
// Radius Auth Port, value from 1 to 65535, default is 1812
type RadiusAuthPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RadiusAuthPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadiusAuthPort) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RadiusAuthPort.
// It customizes the JSON marshaling process for RadiusAuthPort objects.
func (r RadiusAuthPort) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RadiusAuthPortContainer.From*` functions to initialize the RadiusAuthPort object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAuthPort object to a map representation for JSON marshaling.
func (r *RadiusAuthPort) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusAuthPort.
// It customizes the JSON unmarshaling process for RadiusAuthPort objects.
func (r *RadiusAuthPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RadiusAuthPort) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RadiusAuthPort) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRadiusAuthPort represents a radiusAuthPort struct.
// Radius Auth Port, value from 1 to 65535, default is 1812
type internalRadiusAuthPort struct {}

var RadiusAuthPortContainer internalRadiusAuthPort

// The internalRadiusAuthPort instance, wrapping the provided int value.
func (r *internalRadiusAuthPort) FromNumber(val int) RadiusAuthPort {
    return RadiusAuthPort{value: &val}
}

// The internalRadiusAuthPort instance, wrapping the provided string value.
func (r *internalRadiusAuthPort) FromString(val string) RadiusAuthPort {
    return RadiusAuthPort{value: &val}
}
