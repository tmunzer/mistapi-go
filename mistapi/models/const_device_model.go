package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstDeviceModel represents a ConstDeviceModel struct.
type ConstDeviceModel struct {
    value                any
    isConstDeviceAp      bool
    isConstDeviceSwitch  bool
    isConstDeviceGateway bool
}

// String converts the ConstDeviceModel object to a string representation.
func (c ConstDeviceModel) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceModel.
// It customizes the JSON marshaling process for ConstDeviceModel objects.
func (c ConstDeviceModel) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ConstDeviceModelContainer.From*` functions to initialize the ConstDeviceModel object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceModel object to a map representation for JSON marshaling.
func (c *ConstDeviceModel) toMap() any {
    switch obj := c.value.(type) {
    case *ConstDeviceAp:
        return obj.toMap()
    case *ConstDeviceSwitch:
        return obj.toMap()
    case *ConstDeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceModel.
// It customizes the JSON unmarshaling process for ConstDeviceModel objects.
func (c *ConstDeviceModel) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&ConstDeviceAp{}, false, &c.isConstDeviceAp),
        NewTypeHolder(&ConstDeviceSwitch{}, false, &c.isConstDeviceSwitch),
        NewTypeHolder(&ConstDeviceGateway{}, false, &c.isConstDeviceGateway),
    )
    
    c.value = result
    return err
}

func (c *ConstDeviceModel) AsConstDeviceAp() (
    *ConstDeviceAp,
    bool) {
    if !c.isConstDeviceAp {
        return nil, false
    }
    return c.value.(*ConstDeviceAp), true
}

func (c *ConstDeviceModel) AsConstDeviceSwitch() (
    *ConstDeviceSwitch,
    bool) {
    if !c.isConstDeviceSwitch {
        return nil, false
    }
    return c.value.(*ConstDeviceSwitch), true
}

func (c *ConstDeviceModel) AsConstDeviceGateway() (
    *ConstDeviceGateway,
    bool) {
    if !c.isConstDeviceGateway {
        return nil, false
    }
    return c.value.(*ConstDeviceGateway), true
}

// internalConstDeviceModel represents a constDeviceModel struct.
type internalConstDeviceModel struct {}

var ConstDeviceModelContainer internalConstDeviceModel

// The internalConstDeviceModel instance, wrapping the provided ConstDeviceAp value.
func (c *internalConstDeviceModel) FromConstDeviceAp(val ConstDeviceAp) ConstDeviceModel {
    return ConstDeviceModel{value: &val}
}

// The internalConstDeviceModel instance, wrapping the provided ConstDeviceSwitch value.
func (c *internalConstDeviceModel) FromConstDeviceSwitch(val ConstDeviceSwitch) ConstDeviceModel {
    return ConstDeviceModel{value: &val}
}

// The internalConstDeviceModel instance, wrapping the provided ConstDeviceGateway value.
func (c *internalConstDeviceModel) FromConstDeviceGateway(val ConstDeviceGateway) ConstDeviceModel {
    return ConstDeviceModel{value: &val}
}
