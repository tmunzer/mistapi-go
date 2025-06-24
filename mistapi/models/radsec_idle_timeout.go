package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RadsecIdleTimeout represents a RadsecIdleTimeout struct.
// Radsec Idle Timeout in seconds. Default is 60
type RadsecIdleTimeout struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RadsecIdleTimeout,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadsecIdleTimeout) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RadsecIdleTimeout.
// It customizes the JSON marshaling process for RadsecIdleTimeout objects.
func (r RadsecIdleTimeout) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RadsecIdleTimeoutContainer.From*` functions to initialize the RadsecIdleTimeout object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadsecIdleTimeout object to a map representation for JSON marshaling.
func (r *RadsecIdleTimeout) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadsecIdleTimeout.
// It customizes the JSON unmarshaling process for RadsecIdleTimeout objects.
func (r *RadsecIdleTimeout) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RadsecIdleTimeout) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RadsecIdleTimeout) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRadsecIdleTimeout represents a radsecIdleTimeout struct.
// Radsec Idle Timeout in seconds. Default is 60
type internalRadsecIdleTimeout struct {}

var RadsecIdleTimeoutContainer internalRadsecIdleTimeout

// The internalRadsecIdleTimeout instance, wrapping the provided int value.
func (r *internalRadsecIdleTimeout) FromNumber(val int) RadsecIdleTimeout {
    return RadsecIdleTimeout{value: &val}
}

// The internalRadsecIdleTimeout instance, wrapping the provided string value.
func (r *internalRadsecIdleTimeout) FromString(val string) RadsecIdleTimeout {
    return RadsecIdleTimeout{value: &val}
}
