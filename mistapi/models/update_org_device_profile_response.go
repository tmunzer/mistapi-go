package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateOrgDeviceProfileResponse represents a UpdateOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type UpdateOrgDeviceProfileResponse struct {
    value                  any
    isDeviceprofileAp      bool
    isDeviceprofileGateway bool
}

// String converts the UpdateOrgDeviceProfileResponse object to a string representation.
func (u UpdateOrgDeviceProfileResponse) String() string {
    if bytes, err := json.Marshal(u.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateOrgDeviceProfileResponse.
// It customizes the JSON marshaling process for UpdateOrgDeviceProfileResponse objects.
func (u UpdateOrgDeviceProfileResponse) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateOrgDeviceProfileResponseContainer.From*` functions to initialize the UpdateOrgDeviceProfileResponse object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateOrgDeviceProfileResponse object to a map representation for JSON marshaling.
func (u *UpdateOrgDeviceProfileResponse) toMap() any {
    switch obj := u.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *DeviceprofileGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateOrgDeviceProfileResponse.
// It customizes the JSON unmarshaling process for UpdateOrgDeviceProfileResponse objects.
func (u *UpdateOrgDeviceProfileResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &u.isDeviceprofileAp),
        NewTypeHolder(&DeviceprofileGateway{}, false, &u.isDeviceprofileGateway),
    )
    
    u.value = result
    return err
}

func (u *UpdateOrgDeviceProfileResponse) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !u.isDeviceprofileAp {
        return nil, false
    }
    return u.value.(*DeviceprofileAp), true
}

func (u *UpdateOrgDeviceProfileResponse) AsDeviceprofileGateway() (
    *DeviceprofileGateway,
    bool) {
    if !u.isDeviceprofileGateway {
        return nil, false
    }
    return u.value.(*DeviceprofileGateway), true
}

// internalUpdateOrgDeviceProfileResponse represents a updateOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type internalUpdateOrgDeviceProfileResponse struct {}

var UpdateOrgDeviceProfileResponseContainer internalUpdateOrgDeviceProfileResponse

// The internalUpdateOrgDeviceProfileResponse instance, wrapping the provided DeviceprofileAp value.
func (u *internalUpdateOrgDeviceProfileResponse) FromDeviceprofileAp(val DeviceprofileAp) UpdateOrgDeviceProfileResponse {
    return UpdateOrgDeviceProfileResponse{value: &val}
}

// The internalUpdateOrgDeviceProfileResponse instance, wrapping the provided DeviceprofileGateway value.
func (u *internalUpdateOrgDeviceProfileResponse) FromDeviceprofileGateway(val DeviceprofileGateway) UpdateOrgDeviceProfileResponse {
    return UpdateOrgDeviceProfileResponse{value: &val}
}
