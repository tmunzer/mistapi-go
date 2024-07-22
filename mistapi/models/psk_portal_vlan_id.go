package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PskPortalVlanId represents a PskPortalVlanId struct.
// This is a container for one-of cases.
type PskPortalVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the PskPortalVlanId object to a string representation.
func (p PskPortalVlanId) String() string {
    if bytes, err := json.Marshal(p.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for PskPortalVlanId.
// It customizes the JSON marshaling process for PskPortalVlanId objects.
func (p PskPortalVlanId) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PskPortalVlanIdContainer.From*` functions to initialize the PskPortalVlanId object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalVlanId object to a map representation for JSON marshaling.
func (p *PskPortalVlanId) toMap() any {
    switch obj := p.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalVlanId.
// It customizes the JSON unmarshaling process for PskPortalVlanId objects.
func (p *PskPortalVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &p.isString),
        NewTypeHolder(new(int), false, &p.isNumber),
    )
    
    p.value = result
    return err
}

func (p *PskPortalVlanId) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

func (p *PskPortalVlanId) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

// internalPskPortalVlanId represents a pskPortalVlanId struct.
// This is a container for one-of cases.
type internalPskPortalVlanId struct {}

var PskPortalVlanIdContainer internalPskPortalVlanId

// The internalPskPortalVlanId instance, wrapping the provided string value.
func (p *internalPskPortalVlanId) FromString(val string) PskPortalVlanId {
    return PskPortalVlanId{value: &val}
}

// The internalPskPortalVlanId instance, wrapping the provided int value.
func (p *internalPskPortalVlanId) FromNumber(val int) PskPortalVlanId {
    return PskPortalVlanId{value: &val}
}
