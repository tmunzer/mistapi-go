package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ConfigDevice represents a ConfigDevice struct.
type ConfigDevice struct {
	value           any
	isDeviceAp      bool
	isDeviceSwitch  bool
	isDeviceGateway bool
}

// String converts the ConfigDevice object to a string representation.
func (c ConfigDevice) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ConfigDevice.
// It customizes the JSON marshaling process for ConfigDevice objects.
func (c ConfigDevice) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ConfigDeviceContainer.From*` functions to initialize the ConfigDevice object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConfigDevice object to a map representation for JSON marshaling.
func (c *ConfigDevice) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigDevice.
// It customizes the JSON unmarshaling process for ConfigDevice objects.
func (c *ConfigDevice) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&DeviceAp{}, false, &c.isDeviceAp),
		NewTypeHolder(&DeviceSwitch{}, false, &c.isDeviceSwitch),
		NewTypeHolder(&DeviceGateway{}, false, &c.isDeviceGateway),
	)

	c.value = result
	return err
}

func (c *ConfigDevice) AsDeviceAp() (
	*DeviceAp,
	bool) {
	if !c.isDeviceAp {
		return nil, false
	}
	return c.value.(*DeviceAp), true
}

func (c *ConfigDevice) AsDeviceSwitch() (
	*DeviceSwitch,
	bool) {
	if !c.isDeviceSwitch {
		return nil, false
	}
	return c.value.(*DeviceSwitch), true
}

func (c *ConfigDevice) AsDeviceGateway() (
	*DeviceGateway,
	bool) {
	if !c.isDeviceGateway {
		return nil, false
	}
	return c.value.(*DeviceGateway), true
}

// internalConfigDevice represents a configDevice struct.
type internalConfigDevice struct{}

var ConfigDeviceContainer internalConfigDevice

// The internalConfigDevice instance, wrapping the provided DeviceAp value.
func (c *internalConfigDevice) FromDeviceAp(val DeviceAp) ConfigDevice {
	return ConfigDevice{value: &val}
}

// The internalConfigDevice instance, wrapping the provided DeviceSwitch value.
func (c *internalConfigDevice) FromDeviceSwitch(val DeviceSwitch) ConfigDevice {
	return ConfigDevice{value: &val}
}

// The internalConfigDevice instance, wrapping the provided DeviceGateway value.
func (c *internalConfigDevice) FromDeviceGateway(val DeviceGateway) ConfigDevice {
	return ConfigDevice{value: &val}
}
