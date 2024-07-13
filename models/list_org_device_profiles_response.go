package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListOrgDeviceProfilesResponse represents a ListOrgDeviceProfilesResponse struct.
// This is Array of a container for any-of cases.
type ListOrgDeviceProfilesResponse struct {
    value             any
    isDeviceprofileAp bool
    isGatewayTemplate bool
}

// String converts the ListOrgDeviceProfilesResponse object to a string representation.
func (l ListOrgDeviceProfilesResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListOrgDeviceProfilesResponse.
// It customizes the JSON marshaling process for ListOrgDeviceProfilesResponse objects.
func (l ListOrgDeviceProfilesResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListOrgDeviceProfilesResponseContainer.From*` functions to initialize the ListOrgDeviceProfilesResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListOrgDeviceProfilesResponse object to a map representation for JSON marshaling.
func (l *ListOrgDeviceProfilesResponse) toMap() any {
    switch obj := l.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *GatewayTemplate:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListOrgDeviceProfilesResponse.
// It customizes the JSON unmarshaling process for ListOrgDeviceProfilesResponse objects.
func (l *ListOrgDeviceProfilesResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &l.isDeviceprofileAp),
        NewTypeHolder(&GatewayTemplate{}, false, &l.isGatewayTemplate),
    )
    
    l.value = result
    return err
}

func (l *ListOrgDeviceProfilesResponse) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !l.isDeviceprofileAp {
        return nil, false
    }
    return l.value.(*DeviceprofileAp), true
}

func (l *ListOrgDeviceProfilesResponse) AsGatewayTemplate() (
    *GatewayTemplate,
    bool) {
    if !l.isGatewayTemplate {
        return nil, false
    }
    return l.value.(*GatewayTemplate), true
}

// internalListOrgDeviceProfilesResponse represents a listOrgDeviceProfilesResponse struct.
// This is Array of a container for any-of cases.
type internalListOrgDeviceProfilesResponse struct {}

var ListOrgDeviceProfilesResponseContainer internalListOrgDeviceProfilesResponse

// The internalListOrgDeviceProfilesResponse instance, wrapping the provided DeviceprofileAp value.
func (l *internalListOrgDeviceProfilesResponse) FromDeviceprofileAp(val DeviceprofileAp) ListOrgDeviceProfilesResponse {
    return ListOrgDeviceProfilesResponse{value: &val}
}

// The internalListOrgDeviceProfilesResponse instance, wrapping the provided GatewayTemplate value.
func (l *internalListOrgDeviceProfilesResponse) FromGatewayTemplate(val GatewayTemplate) ListOrgDeviceProfilesResponse {
    return ListOrgDeviceProfilesResponse{value: &val}
}
