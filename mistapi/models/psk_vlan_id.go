package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PskVlanId represents a PskVlanId struct.
// VLAN for this PSK key
type PskVlanId struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the PskVlanId object to a string representation.
func (p PskVlanId) String() string {
	if bytes, err := json.Marshal(p.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for PskVlanId.
// It customizes the JSON marshaling process for PskVlanId objects.
func (p PskVlanId) MarshalJSON() (
	[]byte,
	error) {
	if p.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.PskVlanIdContainer.From*` functions to initialize the PskVlanId object.")
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PskVlanId object to a map representation for JSON marshaling.
func (p *PskVlanId) toMap() any {
	switch obj := p.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskVlanId.
// It customizes the JSON unmarshaling process for PskVlanId objects.
func (p *PskVlanId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &p.isString),
		NewTypeHolder(new(int), false, &p.isNumber),
	)

	p.value = result
	return err
}

func (p *PskVlanId) AsString() (
	*string,
	bool) {
	if !p.isString {
		return nil, false
	}
	return p.value.(*string), true
}

func (p *PskVlanId) AsNumber() (
	*int,
	bool) {
	if !p.isNumber {
		return nil, false
	}
	return p.value.(*int), true
}

// internalPskVlanId represents a pskVlanId struct.
// VLAN for this PSK key
type internalPskVlanId struct{}

var PskVlanIdContainer internalPskVlanId

// The internalPskVlanId instance, wrapping the provided string value.
func (p *internalPskVlanId) FromString(val string) PskVlanId {
	return PskVlanId{value: &val}
}

// The internalPskVlanId instance, wrapping the provided int value.
func (p *internalPskVlanId) FromNumber(val int) PskVlanId {
	return PskVlanId{value: &val}
}
