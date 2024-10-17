package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// MxedgeTuntermSwitchConfigVlanId represents a MxedgeTuntermSwitchConfigVlanId struct.
type MxedgeTuntermSwitchConfigVlanId struct {
	value    any
	isNumber bool
	isString bool
}

// String converts the MxedgeTuntermSwitchConfigVlanId object to a string representation.
func (m MxedgeTuntermSwitchConfigVlanId) String() string {
	if bytes, err := json.Marshal(m.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfigVlanId.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfigVlanId objects.
func (m MxedgeTuntermSwitchConfigVlanId) MarshalJSON() (
	[]byte,
	error) {
	if m.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.MxedgeTuntermSwitchConfigVlanIdContainer.From*` functions to initialize the MxedgeTuntermSwitchConfigVlanId object.")
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfigVlanId object to a map representation for JSON marshaling.
func (m *MxedgeTuntermSwitchConfigVlanId) toMap() any {
	switch obj := m.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermSwitchConfigVlanId.
// It customizes the JSON unmarshaling process for MxedgeTuntermSwitchConfigVlanId objects.
func (m *MxedgeTuntermSwitchConfigVlanId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(int), false, &m.isNumber),
		NewTypeHolder(new(string), false, &m.isString),
	)

	m.value = result
	return err
}

func (m *MxedgeTuntermSwitchConfigVlanId) AsNumber() (
	*int,
	bool) {
	if !m.isNumber {
		return nil, false
	}
	return m.value.(*int), true
}

func (m *MxedgeTuntermSwitchConfigVlanId) AsString() (
	*string,
	bool) {
	if !m.isString {
		return nil, false
	}
	return m.value.(*string), true
}

// internalMxedgeTuntermSwitchConfigVlanId represents a mxedgeTuntermSwitchConfigVlanId struct.
type internalMxedgeTuntermSwitchConfigVlanId struct{}

var MxedgeTuntermSwitchConfigVlanIdContainer internalMxedgeTuntermSwitchConfigVlanId

// The internalMxedgeTuntermSwitchConfigVlanId instance, wrapping the provided int value.
func (m *internalMxedgeTuntermSwitchConfigVlanId) FromNumber(val int) MxedgeTuntermSwitchConfigVlanId {
	return MxedgeTuntermSwitchConfigVlanId{value: &val}
}

// The internalMxedgeTuntermSwitchConfigVlanId instance, wrapping the provided string value.
func (m *internalMxedgeTuntermSwitchConfigVlanId) FromString(val string) MxedgeTuntermSwitchConfigVlanId {
	return MxedgeTuntermSwitchConfigVlanId{value: &val}
}
