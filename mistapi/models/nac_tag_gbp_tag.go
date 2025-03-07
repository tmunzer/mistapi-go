package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// NacTagGbpTag represents a NacTagGbpTag struct.
// If `type`==`gbp_tag`
type NacTagGbpTag struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for NacTagGbpTag,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacTagGbpTag) String() string {
    return fmt.Sprintf("%v", n.value)
}

// MarshalJSON implements the json.Marshaler interface for NacTagGbpTag.
// It customizes the JSON marshaling process for NacTagGbpTag objects.
func (n NacTagGbpTag) MarshalJSON() (
    []byte,
    error) {
    if n.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.NacTagGbpTagContainer.From*` functions to initialize the NacTagGbpTag object.")
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NacTagGbpTag object to a map representation for JSON marshaling.
func (n *NacTagGbpTag) toMap() any {
    switch obj := n.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacTagGbpTag.
// It customizes the JSON unmarshaling process for NacTagGbpTag objects.
func (n *NacTagGbpTag) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &n.isNumber),
        NewTypeHolder(new(string), false, &n.isString),
    )
    
    n.value = result
    return err
}

func (n *NacTagGbpTag) AsNumber() (
    *int,
    bool) {
    if !n.isNumber {
        return nil, false
    }
    return n.value.(*int), true
}

func (n *NacTagGbpTag) AsString() (
    *string,
    bool) {
    if !n.isString {
        return nil, false
    }
    return n.value.(*string), true
}

// internalNacTagGbpTag represents a nacTagGbpTag struct.
// If `type`==`gbp_tag`
type internalNacTagGbpTag struct {}

var NacTagGbpTagContainer internalNacTagGbpTag

// The internalNacTagGbpTag instance, wrapping the provided int value.
func (n *internalNacTagGbpTag) FromNumber(val int) NacTagGbpTag {
    return NacTagGbpTag{value: &val}
}

// The internalNacTagGbpTag instance, wrapping the provided string value.
func (n *internalNacTagGbpTag) FromString(val string) NacTagGbpTag {
    return NacTagGbpTag{value: &val}
}
