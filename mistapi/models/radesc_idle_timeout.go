package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// RadescIdleTimeout represents a RadescIdleTimeout struct.
// Radsec Idle Timeout in seconds. Default is 60
type RadescIdleTimeout struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for RadescIdleTimeout,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadescIdleTimeout) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RadescIdleTimeout.
// It customizes the JSON marshaling process for RadescIdleTimeout objects.
func (r RadescIdleTimeout) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RadescIdleTimeoutContainer.From*` functions to initialize the RadescIdleTimeout object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadescIdleTimeout object to a map representation for JSON marshaling.
func (r *RadescIdleTimeout) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadescIdleTimeout.
// It customizes the JSON unmarshaling process for RadescIdleTimeout objects.
func (r *RadescIdleTimeout) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RadescIdleTimeout) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *RadescIdleTimeout) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRadescIdleTimeout represents a radescIdleTimeout struct.
// Radsec Idle Timeout in seconds. Default is 60
type internalRadescIdleTimeout struct {}

var RadescIdleTimeoutContainer internalRadescIdleTimeout

// The internalRadescIdleTimeout instance, wrapping the provided int value.
func (r *internalRadescIdleTimeout) FromNumber(val int) RadescIdleTimeout {
    return RadescIdleTimeout{value: &val}
}

// The internalRadescIdleTimeout instance, wrapping the provided string value.
func (r *internalRadescIdleTimeout) FromString(val string) RadescIdleTimeout {
    return RadescIdleTimeout{value: &val}
}
