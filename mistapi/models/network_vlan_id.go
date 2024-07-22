package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// NetworkVlanId represents a NetworkVlanId struct.
// This is a container for one-of cases.
type NetworkVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the NetworkVlanId object to a string representation.
func (n NetworkVlanId) String() string {
    if bytes, err := json.Marshal(n.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for NetworkVlanId.
// It customizes the JSON marshaling process for NetworkVlanId objects.
func (n NetworkVlanId) MarshalJSON() (
    []byte,
    error) {
    if n.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.NetworkVlanIdContainer.From*` functions to initialize the NetworkVlanId object.")
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkVlanId object to a map representation for JSON marshaling.
func (n *NetworkVlanId) toMap() any {
    switch obj := n.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkVlanId.
// It customizes the JSON unmarshaling process for NetworkVlanId objects.
func (n *NetworkVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &n.isString),
        NewTypeHolder(new(int), false, &n.isNumber),
    )
    
    n.value = result
    return err
}

func (n *NetworkVlanId) AsString() (
    *string,
    bool) {
    if !n.isString {
        return nil, false
    }
    return n.value.(*string), true
}

func (n *NetworkVlanId) AsNumber() (
    *int,
    bool) {
    if !n.isNumber {
        return nil, false
    }
    return n.value.(*int), true
}

// internalNetworkVlanId represents a networkVlanId struct.
// This is a container for one-of cases.
type internalNetworkVlanId struct {}

var NetworkVlanIdContainer internalNetworkVlanId

// The internalNetworkVlanId instance, wrapping the provided string value.
func (n *internalNetworkVlanId) FromString(val string) NetworkVlanId {
    return NetworkVlanId{value: &val}
}

// The internalNetworkVlanId instance, wrapping the provided int value.
func (n *internalNetworkVlanId) FromNumber(val int) NetworkVlanId {
    return NetworkVlanId{value: &val}
}
