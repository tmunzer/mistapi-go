package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CoaPort represents a CoaPort struct.
// CoA Port, value from 1 to 65535
type CoaPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CoaPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CoaPort) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CoaPort.
// It customizes the JSON marshaling process for CoaPort objects.
func (c CoaPort) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CoaPortContainer.From*` functions to initialize the CoaPort object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CoaPort object to a map representation for JSON marshaling.
func (c *CoaPort) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoaPort.
// It customizes the JSON unmarshaling process for CoaPort objects.
func (c *CoaPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CoaPort) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CoaPort) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCoaPort represents a coaPort struct.
// CoA Port, value from 1 to 65535
type internalCoaPort struct {}

var CoaPortContainer internalCoaPort

// The internalCoaPort instance, wrapping the provided int value.
func (c *internalCoaPort) FromNumber(val int) CoaPort {
    return CoaPort{value: &val}
}

// The internalCoaPort instance, wrapping the provided string value.
func (c *internalCoaPort) FromString(val string) CoaPort {
    return CoaPort{value: &val}
}
