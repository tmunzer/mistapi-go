package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateOrgDeviceProfileResponse represents a UpdateOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type UpdateOrgDeviceProfileResponse struct {
    value             any
    isDeviceprofileAp bool
    isGatewayTemplate bool
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
    case *GatewayTemplate:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateOrgDeviceProfileResponse.
// It customizes the JSON unmarshaling process for UpdateOrgDeviceProfileResponse objects.
func (u *UpdateOrgDeviceProfileResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &u.isDeviceprofileAp),
        NewTypeHolder(&GatewayTemplate{}, false, &u.isGatewayTemplate),
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

func (u *UpdateOrgDeviceProfileResponse) AsGatewayTemplate() (
    *GatewayTemplate,
    bool) {
    if !u.isGatewayTemplate {
        return nil, false
    }
    return u.value.(*GatewayTemplate), true
}

// internalUpdateOrgDeviceProfileResponse represents a updateOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type internalUpdateOrgDeviceProfileResponse struct {}

var UpdateOrgDeviceProfileResponseContainer internalUpdateOrgDeviceProfileResponse

// The internalUpdateOrgDeviceProfileResponse instance, wrapping the provided DeviceprofileAp value.
func (u *internalUpdateOrgDeviceProfileResponse) FromDeviceprofileAp(val DeviceprofileAp) UpdateOrgDeviceProfileResponse {
    return UpdateOrgDeviceProfileResponse{value: &val}
}

// The internalUpdateOrgDeviceProfileResponse instance, wrapping the provided GatewayTemplate value.
func (u *internalUpdateOrgDeviceProfileResponse) FromGatewayTemplate(val GatewayTemplate) UpdateOrgDeviceProfileResponse {
    return UpdateOrgDeviceProfileResponse{value: &val}
}
