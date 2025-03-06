package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// IntegerOrString represents a IntegerOrString struct.
type IntegerOrString struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for IntegerOrString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IntegerOrString) String() string {
    return fmt.Sprintf("%v", i.value)
}

// MarshalJSON implements the json.Marshaler interface for IntegerOrString.
// It customizes the JSON marshaling process for IntegerOrString objects.
func (i IntegerOrString) MarshalJSON() (
    []byte,
    error) {
    if i.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.IntegerOrStringContainer.From*` functions to initialize the IntegerOrString object.")
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IntegerOrString object to a map representation for JSON marshaling.
func (i *IntegerOrString) toMap() any {
    switch obj := i.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for IntegerOrString.
// It customizes the JSON unmarshaling process for IntegerOrString objects.
func (i *IntegerOrString) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &i.isNumber),
        NewTypeHolder(new(string), false, &i.isString),
    )
    
    i.value = result
    return err
}

func (i *IntegerOrString) AsNumber() (
    *int,
    bool) {
    if !i.isNumber {
        return nil, false
    }
    return i.value.(*int), true
}

func (i *IntegerOrString) AsString() (
    *string,
    bool) {
    if !i.isString {
        return nil, false
    }
    return i.value.(*string), true
}

// internalIntegerOrString represents a integerOrString struct.
type internalIntegerOrString struct {}

var IntegerOrStringContainer internalIntegerOrString

// The internalIntegerOrString instance, wrapping the provided int value.
func (i *internalIntegerOrString) FromNumber(val int) IntegerOrString {
    return IntegerOrString{value: &val}
}

// The internalIntegerOrString instance, wrapping the provided string value.
func (i *internalIntegerOrString) FromString(val string) IntegerOrString {
    return IntegerOrString{value: &val}
}
