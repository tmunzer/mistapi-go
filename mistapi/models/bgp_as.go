// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// BgpAs represents a BgpAs struct.
// BGP AS, value in range 1-4294967295
type BgpAs struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for BgpAs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BgpAs) String() string {
    return fmt.Sprintf("%v", b.value)
}

// MarshalJSON implements the json.Marshaler interface for BgpAs.
// It customizes the JSON marshaling process for BgpAs objects.
func (b BgpAs) MarshalJSON() (
    []byte,
    error) {
    if b.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.BgpAsContainer.From*` functions to initialize the BgpAs object.")
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BgpAs object to a map representation for JSON marshaling.
func (b *BgpAs) toMap() any {
    switch obj := b.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpAs.
// It customizes the JSON unmarshaling process for BgpAs objects.
func (b *BgpAs) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(string), false, &b.isString),
        NewTypeHolder(new(int), false, &b.isNumber),
    )
    
    b.value = result
    return err
}

func (b *BgpAs) AsString() (
    *string,
    bool) {
    if !b.isString {
        return nil, false
    }
    return b.value.(*string), true
}

func (b *BgpAs) AsNumber() (
    *int,
    bool) {
    if !b.isNumber {
        return nil, false
    }
    return b.value.(*int), true
}

// internalBgpAs represents a bgpAs struct.
// BGP AS, value in range 1-4294967295
type internalBgpAs struct {}

var BgpAsContainer internalBgpAs

// The internalBgpAs instance, wrapping the provided string value.
func (b *internalBgpAs) FromString(val string) BgpAs {
    return BgpAs{value: &val}
}

// The internalBgpAs instance, wrapping the provided int value.
func (b *internalBgpAs) FromNumber(val int) BgpAs {
    return BgpAs{value: &val}
}
