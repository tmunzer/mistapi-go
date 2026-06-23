// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NextHopVia represents a NextHopVia struct.
// Next-hop IP address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops.
type NextHopVia struct {
	value           any
	isString        bool
	isArrayOfString bool
}

// String implements the fmt.Stringer interface for NextHopVia,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NextHopVia) String() string {
	return fmt.Sprintf("%v", n.value)
}

// MarshalJSON implements the json.Marshaler interface for NextHopVia.
// It customizes the JSON marshaling process for NextHopVia objects.
func (n NextHopVia) MarshalJSON() (
	[]byte,
	error) {
	if n.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.NextHopViaContainer.From*` functions to initialize the NextHopVia object.")
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NextHopVia object to a map representation for JSON marshaling.
func (n *NextHopVia) toMap() any {
	switch obj := n.value.(type) {
	case *string:
		return *obj
	case *[]string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for NextHopVia.
// It customizes the JSON unmarshaling process for NextHopVia objects.
func (n *NextHopVia) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &n.isString),
		NewTypeHolder(new([]string), false, &n.isArrayOfString),
	)

	n.value = result
	return err
}

func (n *NextHopVia) AsString() (
	*string,
	bool) {
	if !n.isString {
		return nil, false
	}
	return n.value.(*string), true
}

func (n *NextHopVia) AsArrayOfString() (
	*[]string,
	bool) {
	if !n.isArrayOfString {
		return nil, false
	}
	return n.value.(*[]string), true
}

// internalNextHopVia represents a nextHopVia struct.
// Next-hop IP address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops.
type internalNextHopVia struct{}

var NextHopViaContainer internalNextHopVia

// The internalNextHopVia instance, wrapping the provided string value.
func (n *internalNextHopVia) FromString(val string) NextHopVia {
	return NextHopVia{value: &val}
}

// The internalNextHopVia instance, wrapping the provided []string value.
func (n *internalNextHopVia) FromArrayOfString(val []string) NextHopVia {
	return NextHopVia{value: &val}
}
