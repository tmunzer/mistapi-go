package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// Deviceprofile2 represents a Deviceprofile2 struct.
type Deviceprofile2 struct {
    value             any
    isDeviceprofileAp bool
    isGatewayTemplate bool
}

// String converts the Deviceprofile2 object to a string representation.
func (d Deviceprofile2) String() string {
    if bytes, err := json.Marshal(d.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for Deviceprofile2.
// It customizes the JSON marshaling process for Deviceprofile2 objects.
func (d Deviceprofile2) MarshalJSON() (
    []byte,
    error) {
    if d.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.DeviceprofileContainer.From*` functions to initialize the Deviceprofile2 object.")
    }
    return json.Marshal(d.toMap())
}

// toMap converts the Deviceprofile2 object to a map representation for JSON marshaling.
func (d *Deviceprofile2) toMap() any {
    switch obj := d.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *GatewayTemplate:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Deviceprofile2.
// It customizes the JSON unmarshaling process for Deviceprofile2 objects.
func (d *Deviceprofile2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &d.isDeviceprofileAp),
        NewTypeHolder(&GatewayTemplate{}, false, &d.isGatewayTemplate),
    )
    
    d.value = result
    return err
}

func (d *Deviceprofile2) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !d.isDeviceprofileAp {
        return nil, false
    }
    return d.value.(*DeviceprofileAp), true
}

func (d *Deviceprofile2) AsGatewayTemplate() (
    *GatewayTemplate,
    bool) {
    if !d.isGatewayTemplate {
        return nil, false
    }
    return d.value.(*GatewayTemplate), true
}

// internalDeviceprofile2 represents a deviceprofile2 struct.
type internalDeviceprofile2 struct {}

var DeviceprofileContainer internalDeviceprofile2

// The internalDeviceprofile2 instance, wrapping the provided DeviceprofileAp value.
func (d *internalDeviceprofile2) FromDeviceprofileAp(val DeviceprofileAp) Deviceprofile2 {
    return Deviceprofile2{value: &val}
}

// The internalDeviceprofile2 instance, wrapping the provided GatewayTemplate value.
func (d *internalDeviceprofile2) FromGatewayTemplate(val GatewayTemplate) Deviceprofile2 {
    return Deviceprofile2{value: &val}
}
