package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// MistDevice represents a MistDevice struct.
type MistDevice struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String implements the fmt.Stringer interface for MistDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MistDevice) String() string {
    return fmt.Sprintf("%v", m.value)
}

// MarshalJSON implements the json.Marshaler interface for MistDevice.
// It customizes the JSON marshaling process for MistDevice objects.
func (m MistDevice) MarshalJSON() (
    []byte,
    error) {
    if m.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.MistDeviceContainer.From*` functions to initialize the MistDevice object.")
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MistDevice object to a map representation for JSON marshaling.
func (m *MistDevice) toMap() any {
    switch obj := m.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MistDevice.
// It customizes the JSON unmarshaling process for MistDevice objects.
func (m *MistDevice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&DeviceAp{}, false, &m.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &m.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &m.isDeviceGateway),
    )
    
    m.value = result
    return err
}

func (m *MistDevice) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !m.isDeviceAp {
        return nil, false
    }
    return m.value.(*DeviceAp), true
}

func (m *MistDevice) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !m.isDeviceSwitch {
        return nil, false
    }
    return m.value.(*DeviceSwitch), true
}

func (m *MistDevice) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !m.isDeviceGateway {
        return nil, false
    }
    return m.value.(*DeviceGateway), true
}

// internalMistDevice represents a mistDevice struct.
type internalMistDevice struct {}

var MistDeviceContainer internalMistDevice

// The internalMistDevice instance, wrapping the provided DeviceAp value.
func (m *internalMistDevice) FromDeviceAp(val DeviceAp) MistDevice {
    return MistDevice{value: &val}
}

// The internalMistDevice instance, wrapping the provided DeviceSwitch value.
func (m *internalMistDevice) FromDeviceSwitch(val DeviceSwitch) MistDevice {
    return MistDevice{value: &val}
}

// The internalMistDevice instance, wrapping the provided DeviceGateway value.
func (m *internalMistDevice) FromDeviceGateway(val DeviceGateway) MistDevice {
    return MistDevice{value: &val}
}
