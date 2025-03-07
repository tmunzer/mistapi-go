package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// Dscp represents a Dscp struct.
// DSCP value range between 0 and 63
type Dscp struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for Dscp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d Dscp) String() string {
    return fmt.Sprintf("%v", d.value)
}

// MarshalJSON implements the json.Marshaler interface for Dscp.
// It customizes the JSON marshaling process for Dscp objects.
func (d Dscp) MarshalJSON() (
    []byte,
    error) {
    if d.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.DscpContainer.From*` functions to initialize the Dscp object.")
    }
    return json.Marshal(d.toMap())
}

// toMap converts the Dscp object to a map representation for JSON marshaling.
func (d *Dscp) toMap() any {
    switch obj := d.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Dscp.
// It customizes the JSON unmarshaling process for Dscp objects.
func (d *Dscp) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(string), false, &d.isString),
        NewTypeHolder(new(int), false, &d.isNumber),
    )
    
    d.value = result
    return err
}

func (d *Dscp) AsString() (
    *string,
    bool) {
    if !d.isString {
        return nil, false
    }
    return d.value.(*string), true
}

func (d *Dscp) AsNumber() (
    *int,
    bool) {
    if !d.isNumber {
        return nil, false
    }
    return d.value.(*int), true
}

// internalDscp represents a dscp struct.
// DSCP value range between 0 and 63
type internalDscp struct {}

var DscpContainer internalDscp

// The internalDscp instance, wrapping the provided string value.
func (d *internalDscp) FromString(val string) Dscp {
    return Dscp{value: &val}
}

// The internalDscp instance, wrapping the provided int value.
func (d *internalDscp) FromNumber(val int) Dscp {
    return Dscp{value: &val}
}
