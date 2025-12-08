// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// TuntermPortConfigUpstreamPortVlanId represents a TuntermPortConfigUpstreamPortVlanId struct.
// Native VLAN id for upstream ports
type TuntermPortConfigUpstreamPortVlanId struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for TuntermPortConfigUpstreamPortVlanId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TuntermPortConfigUpstreamPortVlanId) String() string {
	return fmt.Sprintf("%v", t.value)
}

// MarshalJSON implements the json.Marshaler interface for TuntermPortConfigUpstreamPortVlanId.
// It customizes the JSON marshaling process for TuntermPortConfigUpstreamPortVlanId objects.
func (t TuntermPortConfigUpstreamPortVlanId) MarshalJSON() (
	[]byte,
	error) {
	if t.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.TuntermPortConfigUpstreamPortVlanIdContainer.From*` functions to initialize the TuntermPortConfigUpstreamPortVlanId object.")
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TuntermPortConfigUpstreamPortVlanId object to a map representation for JSON marshaling.
func (t *TuntermPortConfigUpstreamPortVlanId) toMap() any {
	switch obj := t.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermPortConfigUpstreamPortVlanId.
// It customizes the JSON unmarshaling process for TuntermPortConfigUpstreamPortVlanId objects.
func (t *TuntermPortConfigUpstreamPortVlanId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &t.isString),
		NewTypeHolder(new(int), false, &t.isNumber),
	)

	t.value = result
	return err
}

func (t *TuntermPortConfigUpstreamPortVlanId) AsString() (
	*string,
	bool) {
	if !t.isString {
		return nil, false
	}
	return t.value.(*string), true
}

func (t *TuntermPortConfigUpstreamPortVlanId) AsNumber() (
	*int,
	bool) {
	if !t.isNumber {
		return nil, false
	}
	return t.value.(*int), true
}

// internalTuntermPortConfigUpstreamPortVlanId represents a tuntermPortConfigUpstreamPortVlanId struct.
// Native VLAN id for upstream ports
type internalTuntermPortConfigUpstreamPortVlanId struct{}

var TuntermPortConfigUpstreamPortVlanIdContainer internalTuntermPortConfigUpstreamPortVlanId

// The internalTuntermPortConfigUpstreamPortVlanId instance, wrapping the provided string value.
func (t *internalTuntermPortConfigUpstreamPortVlanId) FromString(val string) TuntermPortConfigUpstreamPortVlanId {
	return TuntermPortConfigUpstreamPortVlanId{value: &val}
}

// The internalTuntermPortConfigUpstreamPortVlanId instance, wrapping the provided int value.
func (t *internalTuntermPortConfigUpstreamPortVlanId) FromNumber(val int) TuntermPortConfigUpstreamPortVlanId {
	return TuntermPortConfigUpstreamPortVlanId{value: &val}
}
