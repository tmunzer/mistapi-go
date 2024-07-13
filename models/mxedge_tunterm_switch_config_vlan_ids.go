package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxedgeTuntermSwitchConfigVlanIds represents a MxedgeTuntermSwitchConfigVlanIds struct.
// This is Array of a container for any-of cases.
type MxedgeTuntermSwitchConfigVlanIds struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the MxedgeTuntermSwitchConfigVlanIds object to a string representation.
func (m MxedgeTuntermSwitchConfigVlanIds) String() string {
    if bytes, err := json.Marshal(m.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfigVlanIds.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfigVlanIds objects.
func (m MxedgeTuntermSwitchConfigVlanIds) MarshalJSON() (
    []byte,
    error) {
    if m.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.MxedgeTuntermSwitchConfigVlanIdsContainer.From*` functions to initialize the MxedgeTuntermSwitchConfigVlanIds object.")
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfigVlanIds object to a map representation for JSON marshaling.
func (m *MxedgeTuntermSwitchConfigVlanIds) toMap() any {
    switch obj := m.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermSwitchConfigVlanIds.
// It customizes the JSON unmarshaling process for MxedgeTuntermSwitchConfigVlanIds objects.
func (m *MxedgeTuntermSwitchConfigVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &m.isNumber),
        NewTypeHolder(new(string), false, &m.isString),
    )
    
    m.value = result
    return err
}

func (m *MxedgeTuntermSwitchConfigVlanIds) AsNumber() (
    *int,
    bool) {
    if !m.isNumber {
        return nil, false
    }
    return m.value.(*int), true
}

func (m *MxedgeTuntermSwitchConfigVlanIds) AsString() (
    *string,
    bool) {
    if !m.isString {
        return nil, false
    }
    return m.value.(*string), true
}

// internalMxedgeTuntermSwitchConfigVlanIds represents a mxedgeTuntermSwitchConfigVlanIds struct.
// This is Array of a container for any-of cases.
type internalMxedgeTuntermSwitchConfigVlanIds struct {}

var MxedgeTuntermSwitchConfigVlanIdsContainer internalMxedgeTuntermSwitchConfigVlanIds

// The internalMxedgeTuntermSwitchConfigVlanIds instance, wrapping the provided int value.
func (m *internalMxedgeTuntermSwitchConfigVlanIds) FromNumber(val int) MxedgeTuntermSwitchConfigVlanIds {
    return MxedgeTuntermSwitchConfigVlanIds{value: &val}
}

// The internalMxedgeTuntermSwitchConfigVlanIds instance, wrapping the provided string value.
func (m *internalMxedgeTuntermSwitchConfigVlanIds) FromString(val string) MxedgeTuntermSwitchConfigVlanIds {
    return MxedgeTuntermSwitchConfigVlanIds{value: &val}
}
