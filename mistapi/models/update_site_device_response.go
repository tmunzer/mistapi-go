package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateSiteDeviceResponse represents a UpdateSiteDeviceResponse struct.
// This is a container for one-of cases.
type UpdateSiteDeviceResponse struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String converts the UpdateSiteDeviceResponse object to a string representation.
func (u UpdateSiteDeviceResponse) String() string {
    if bytes, err := json.Marshal(u.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateSiteDeviceResponse.
// It customizes the JSON marshaling process for UpdateSiteDeviceResponse objects.
func (u UpdateSiteDeviceResponse) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateSiteDeviceResponseContainer.From*` functions to initialize the UpdateSiteDeviceResponse object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSiteDeviceResponse object to a map representation for JSON marshaling.
func (u *UpdateSiteDeviceResponse) toMap() any {
    switch obj := u.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSiteDeviceResponse.
// It customizes the JSON unmarshaling process for UpdateSiteDeviceResponse objects.
func (u *UpdateSiteDeviceResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&DeviceAp{}, false, &u.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &u.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &u.isDeviceGateway),
    )
    
    u.value = result
    return err
}

func (u *UpdateSiteDeviceResponse) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !u.isDeviceAp {
        return nil, false
    }
    return u.value.(*DeviceAp), true
}

func (u *UpdateSiteDeviceResponse) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !u.isDeviceSwitch {
        return nil, false
    }
    return u.value.(*DeviceSwitch), true
}

func (u *UpdateSiteDeviceResponse) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !u.isDeviceGateway {
        return nil, false
    }
    return u.value.(*DeviceGateway), true
}

// internalUpdateSiteDeviceResponse represents a updateSiteDeviceResponse struct.
// This is a container for one-of cases.
type internalUpdateSiteDeviceResponse struct {}

var UpdateSiteDeviceResponseContainer internalUpdateSiteDeviceResponse

// The internalUpdateSiteDeviceResponse instance, wrapping the provided DeviceAp value.
func (u *internalUpdateSiteDeviceResponse) FromDeviceAp(val DeviceAp) UpdateSiteDeviceResponse {
    return UpdateSiteDeviceResponse{value: &val}
}

// The internalUpdateSiteDeviceResponse instance, wrapping the provided DeviceSwitch value.
func (u *internalUpdateSiteDeviceResponse) FromDeviceSwitch(val DeviceSwitch) UpdateSiteDeviceResponse {
    return UpdateSiteDeviceResponse{value: &val}
}

// The internalUpdateSiteDeviceResponse instance, wrapping the provided DeviceGateway value.
func (u *internalUpdateSiteDeviceResponse) FromDeviceGateway(val DeviceGateway) UpdateSiteDeviceResponse {
    return UpdateSiteDeviceResponse{value: &val}
}
