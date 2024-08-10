package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstDeviceModel2 represents a ConstDeviceModel2 struct.
type ConstDeviceModel2 struct {
    value                any
    isConstDeviceAp      bool
    isConstDeviceSwitch  bool
    isConstDeviceGateway bool
    isConstDeviceUnknown bool
}

// String converts the ConstDeviceModel2 object to a string representation.
func (c ConstDeviceModel2) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceModel2.
// It customizes the JSON marshaling process for ConstDeviceModel2 objects.
func (c ConstDeviceModel2) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ConstDeviceModelContainer.From*` functions to initialize the ConstDeviceModel2 object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceModel2 object to a map representation for JSON marshaling.
func (c *ConstDeviceModel2) toMap() any {
    switch obj := c.value.(type) {
    case *ConstDeviceAp:
        return obj.toMap()
    case *ConstDeviceSwitch:
        return obj.toMap()
    case *ConstDeviceGateway:
        return obj.toMap()
    case *ConstDeviceUnknown:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceModel2.
// It customizes the JSON unmarshaling process for ConstDeviceModel2 objects.
func (c *ConstDeviceModel2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&ConstDeviceAp{}, false, &c.isConstDeviceAp),
        NewTypeHolder(&ConstDeviceSwitch{}, false, &c.isConstDeviceSwitch),
        NewTypeHolder(&ConstDeviceGateway{}, false, &c.isConstDeviceGateway),
        NewTypeHolder(&ConstDeviceUnknown{}, false, &c.isConstDeviceUnknown),
    )
    
    c.value = result
    return err
}

func (c *ConstDeviceModel2) AsConstDeviceAp() (
    *ConstDeviceAp,
    bool) {
    if !c.isConstDeviceAp {
        return nil, false
    }
    return c.value.(*ConstDeviceAp), true
}

func (c *ConstDeviceModel2) AsConstDeviceSwitch() (
    *ConstDeviceSwitch,
    bool) {
    if !c.isConstDeviceSwitch {
        return nil, false
    }
    return c.value.(*ConstDeviceSwitch), true
}

func (c *ConstDeviceModel2) AsConstDeviceGateway() (
    *ConstDeviceGateway,
    bool) {
    if !c.isConstDeviceGateway {
        return nil, false
    }
    return c.value.(*ConstDeviceGateway), true
}

func (c *ConstDeviceModel2) AsConstDeviceUnknown() (
    *ConstDeviceUnknown,
    bool) {
    if !c.isConstDeviceUnknown {
        return nil, false
    }
    return c.value.(*ConstDeviceUnknown), true
}

// internalConstDeviceModel2 represents a constDeviceModel2 struct.
type internalConstDeviceModel2 struct {}

var ConstDeviceModelContainer internalConstDeviceModel2

// The internalConstDeviceModel2 instance, wrapping the provided ConstDeviceAp value.
func (c *internalConstDeviceModel2) FromConstDeviceAp(val ConstDeviceAp) ConstDeviceModel2 {
    return ConstDeviceModel2{value: &val}
}

// The internalConstDeviceModel2 instance, wrapping the provided ConstDeviceSwitch value.
func (c *internalConstDeviceModel2) FromConstDeviceSwitch(val ConstDeviceSwitch) ConstDeviceModel2 {
    return ConstDeviceModel2{value: &val}
}

// The internalConstDeviceModel2 instance, wrapping the provided ConstDeviceGateway value.
func (c *internalConstDeviceModel2) FromConstDeviceGateway(val ConstDeviceGateway) ConstDeviceModel2 {
    return ConstDeviceModel2{value: &val}
}

// The internalConstDeviceModel2 instance, wrapping the provided ConstDeviceUnknown value.
func (c *internalConstDeviceModel2) FromConstDeviceUnknown(val ConstDeviceUnknown) ConstDeviceModel2 {
    return ConstDeviceModel2{value: &val}
}
