// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BgpLocalAs represents a BgpLocalAs struct.
// Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. BGP AS, value in range 1-4294967295
type BgpLocalAs struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for BgpLocalAs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BgpLocalAs) String() string {
	return fmt.Sprintf("%v", b.value)
}

// MarshalJSON implements the json.Marshaler interface for BgpLocalAs.
// It customizes the JSON marshaling process for BgpLocalAs objects.
func (b BgpLocalAs) MarshalJSON() (
	[]byte,
	error) {
	if b.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.BgpLocalAsContainer.From*` functions to initialize the BgpLocalAs object.")
	}
	return json.Marshal(b.toMap())
}

// toMap converts the BgpLocalAs object to a map representation for JSON marshaling.
func (b *BgpLocalAs) toMap() any {
	switch obj := b.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpLocalAs.
// It customizes the JSON unmarshaling process for BgpLocalAs objects.
func (b *BgpLocalAs) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &b.isString),
		NewTypeHolder(new(int), false, &b.isNumber),
	)

	b.value = result
	return err
}

func (b *BgpLocalAs) AsString() (
	*string,
	bool) {
	if !b.isString {
		return nil, false
	}
	return b.value.(*string), true
}

func (b *BgpLocalAs) AsNumber() (
	*int,
	bool) {
	if !b.isNumber {
		return nil, false
	}
	return b.value.(*int), true
}

// internalBgpLocalAs represents a bgpLocalAs struct.
// Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. BGP AS, value in range 1-4294967295
type internalBgpLocalAs struct{}

var BgpLocalAsContainer internalBgpLocalAs

// The internalBgpLocalAs instance, wrapping the provided string value.
func (b *internalBgpLocalAs) FromString(val string) BgpLocalAs {
	return BgpLocalAs{value: &val}
}

// The internalBgpLocalAs instance, wrapping the provided int value.
func (b *internalBgpLocalAs) FromNumber(val int) BgpLocalAs {
	return BgpLocalAs{value: &val}
}
