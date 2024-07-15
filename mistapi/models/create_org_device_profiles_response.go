package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateOrgDeviceProfilesResponse represents a CreateOrgDeviceProfilesResponse struct.
// This is a container for any-of cases.
type CreateOrgDeviceProfilesResponse struct {
    value                  any
    isDeviceprofileAp      bool
    isDeviceprofileGateway bool
}

// String converts the CreateOrgDeviceProfilesResponse object to a string representation.
func (c CreateOrgDeviceProfilesResponse) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateOrgDeviceProfilesResponse.
// It customizes the JSON marshaling process for CreateOrgDeviceProfilesResponse objects.
func (c CreateOrgDeviceProfilesResponse) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateOrgDeviceProfilesResponseContainer.From*` functions to initialize the CreateOrgDeviceProfilesResponse object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrgDeviceProfilesResponse object to a map representation for JSON marshaling.
func (c *CreateOrgDeviceProfilesResponse) toMap() any {
    switch obj := c.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *DeviceprofileGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrgDeviceProfilesResponse.
// It customizes the JSON unmarshaling process for CreateOrgDeviceProfilesResponse objects.
func (c *CreateOrgDeviceProfilesResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &c.isDeviceprofileAp),
        NewTypeHolder(&DeviceprofileGateway{}, false, &c.isDeviceprofileGateway),
    )
    
    c.value = result
    return err
}

func (c *CreateOrgDeviceProfilesResponse) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !c.isDeviceprofileAp {
        return nil, false
    }
    return c.value.(*DeviceprofileAp), true
}

func (c *CreateOrgDeviceProfilesResponse) AsDeviceprofileGateway() (
    *DeviceprofileGateway,
    bool) {
    if !c.isDeviceprofileGateway {
        return nil, false
    }
    return c.value.(*DeviceprofileGateway), true
}

// internalCreateOrgDeviceProfilesResponse represents a createOrgDeviceProfilesResponse struct.
// This is a container for any-of cases.
type internalCreateOrgDeviceProfilesResponse struct {}

var CreateOrgDeviceProfilesResponseContainer internalCreateOrgDeviceProfilesResponse

// The internalCreateOrgDeviceProfilesResponse instance, wrapping the provided DeviceprofileAp value.
func (c *internalCreateOrgDeviceProfilesResponse) FromDeviceprofileAp(val DeviceprofileAp) CreateOrgDeviceProfilesResponse {
    return CreateOrgDeviceProfilesResponse{value: &val}
}

// The internalCreateOrgDeviceProfilesResponse instance, wrapping the provided DeviceprofileGateway value.
func (c *internalCreateOrgDeviceProfilesResponse) FromDeviceprofileGateway(val DeviceprofileGateway) CreateOrgDeviceProfilesResponse {
    return CreateOrgDeviceProfilesResponse{value: &val}
}
