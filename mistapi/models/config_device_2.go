package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConfigDevice2 represents a ConfigDevice2 struct.
type ConfigDevice2 struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String converts the ConfigDevice2 object to a string representation.
func (c ConfigDevice2) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ConfigDevice2.
// It customizes the JSON marshaling process for ConfigDevice2 objects.
func (c ConfigDevice2) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ConfigDeviceContainer.From*` functions to initialize the ConfigDevice2 object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigDevice2 object to a map representation for JSON marshaling.
func (c *ConfigDevice2) toMap() any {
    switch obj := c.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigDevice2.
// It customizes the JSON unmarshaling process for ConfigDevice2 objects.
func (c *ConfigDevice2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceAp{}, false, &c.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &c.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &c.isDeviceGateway),
    )
    
    c.value = result
    return err
}

func (c *ConfigDevice2) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !c.isDeviceAp {
        return nil, false
    }
    return c.value.(*DeviceAp), true
}

func (c *ConfigDevice2) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !c.isDeviceSwitch {
        return nil, false
    }
    return c.value.(*DeviceSwitch), true
}

func (c *ConfigDevice2) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !c.isDeviceGateway {
        return nil, false
    }
    return c.value.(*DeviceGateway), true
}

// internalConfigDevice2 represents a configDevice2 struct.
type internalConfigDevice2 struct {}

var ConfigDeviceContainer internalConfigDevice2

// The internalConfigDevice2 instance, wrapping the provided DeviceAp value.
func (c *internalConfigDevice2) FromDeviceAp(val DeviceAp) ConfigDevice2 {
    return ConfigDevice2{value: &val}
}

// The internalConfigDevice2 instance, wrapping the provided DeviceSwitch value.
func (c *internalConfigDevice2) FromDeviceSwitch(val DeviceSwitch) ConfigDevice2 {
    return ConfigDevice2{value: &val}
}

// The internalConfigDevice2 instance, wrapping the provided DeviceGateway value.
func (c *internalConfigDevice2) FromDeviceGateway(val DeviceGateway) ConfigDevice2 {
    return ConfigDevice2{value: &val}
}
