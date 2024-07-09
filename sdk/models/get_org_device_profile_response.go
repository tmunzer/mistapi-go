package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GetOrgDeviceProfileResponse represents a GetOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type GetOrgDeviceProfileResponse struct {
    value             any
    isDeviceprofileAp bool
    isGatewayTemplate bool
}

// String converts the GetOrgDeviceProfileResponse object to a string representation.
func (g GetOrgDeviceProfileResponse) String() string {
    if bytes, err := json.Marshal(g.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for GetOrgDeviceProfileResponse.
// It customizes the JSON marshaling process for GetOrgDeviceProfileResponse objects.
func (g GetOrgDeviceProfileResponse) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GetOrgDeviceProfileResponseContainer.From*` functions to initialize the GetOrgDeviceProfileResponse object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetOrgDeviceProfileResponse object to a map representation for JSON marshaling.
func (g *GetOrgDeviceProfileResponse) toMap() any {
    switch obj := g.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *GatewayTemplate:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrgDeviceProfileResponse.
// It customizes the JSON unmarshaling process for GetOrgDeviceProfileResponse objects.
func (g *GetOrgDeviceProfileResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &g.isDeviceprofileAp),
        NewTypeHolder(&GatewayTemplate{}, false, &g.isGatewayTemplate),
    )
    
    g.value = result
    return err
}

func (g *GetOrgDeviceProfileResponse) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !g.isDeviceprofileAp {
        return nil, false
    }
    return g.value.(*DeviceprofileAp), true
}

func (g *GetOrgDeviceProfileResponse) AsGatewayTemplate() (
    *GatewayTemplate,
    bool) {
    if !g.isGatewayTemplate {
        return nil, false
    }
    return g.value.(*GatewayTemplate), true
}

// internalGetOrgDeviceProfileResponse represents a getOrgDeviceProfileResponse struct.
// This is a container for any-of cases.
type internalGetOrgDeviceProfileResponse struct {}

var GetOrgDeviceProfileResponseContainer internalGetOrgDeviceProfileResponse

// The internalGetOrgDeviceProfileResponse instance, wrapping the provided DeviceprofileAp value.
func (g *internalGetOrgDeviceProfileResponse) FromDeviceprofileAp(val DeviceprofileAp) GetOrgDeviceProfileResponse {
    return GetOrgDeviceProfileResponse{value: &val}
}

// The internalGetOrgDeviceProfileResponse instance, wrapping the provided GatewayTemplate value.
func (g *internalGetOrgDeviceProfileResponse) FromGatewayTemplate(val GatewayTemplate) GetOrgDeviceProfileResponse {
    return GetOrgDeviceProfileResponse{value: &val}
}
