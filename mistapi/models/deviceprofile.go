package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// Deviceprofile represents a Deviceprofile struct.
type Deviceprofile struct {
    value                  any
    isDeviceprofileAp      bool
    isDeviceprofileGateway bool
}

// String implements the fmt.Stringer interface for Deviceprofile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d Deviceprofile) String() string {
    return fmt.Sprintf("%v", d.value)
}

// MarshalJSON implements the json.Marshaler interface for Deviceprofile.
// It customizes the JSON marshaling process for Deviceprofile objects.
func (d Deviceprofile) MarshalJSON() (
    []byte,
    error) {
    if d.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.DeviceprofileContainer.From*` functions to initialize the Deviceprofile object.")
    }
    return json.Marshal(d.toMap())
}

// toMap converts the Deviceprofile object to a map representation for JSON marshaling.
func (d *Deviceprofile) toMap() any {
    switch obj := d.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *DeviceprofileGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Deviceprofile.
// It customizes the JSON unmarshaling process for Deviceprofile objects.
func (d *Deviceprofile) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &d.isDeviceprofileAp),
        NewTypeHolder(&DeviceprofileGateway{}, false, &d.isDeviceprofileGateway),
    )
    
    d.value = result
    return err
}

func (d *Deviceprofile) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !d.isDeviceprofileAp {
        return nil, false
    }
    return d.value.(*DeviceprofileAp), true
}

func (d *Deviceprofile) AsDeviceprofileGateway() (
    *DeviceprofileGateway,
    bool) {
    if !d.isDeviceprofileGateway {
        return nil, false
    }
    return d.value.(*DeviceprofileGateway), true
}

// internalDeviceprofile represents a deviceprofile struct.
type internalDeviceprofile struct {}

var DeviceprofileContainer internalDeviceprofile

// The internalDeviceprofile instance, wrapping the provided DeviceprofileAp value.
func (d *internalDeviceprofile) FromDeviceprofileAp(val DeviceprofileAp) Deviceprofile {
    return Deviceprofile{value: &val}
}

// The internalDeviceprofile instance, wrapping the provided DeviceprofileGateway value.
func (d *internalDeviceprofile) FromDeviceprofileGateway(val DeviceprofileGateway) Deviceprofile {
    return Deviceprofile{value: &val}
}
