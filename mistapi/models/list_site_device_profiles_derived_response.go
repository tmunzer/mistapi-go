package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListSiteDeviceProfilesDerivedResponse represents a ListSiteDeviceProfilesDerivedResponse struct.
// This is Array of a container for any-of cases.
type ListSiteDeviceProfilesDerivedResponse struct {
    value                  any
    isDeviceprofileAp      bool
    isDeviceprofileGateway bool
}

// String converts the ListSiteDeviceProfilesDerivedResponse object to a string representation.
func (l ListSiteDeviceProfilesDerivedResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListSiteDeviceProfilesDerivedResponse.
// It customizes the JSON marshaling process for ListSiteDeviceProfilesDerivedResponse objects.
func (l ListSiteDeviceProfilesDerivedResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListSiteDeviceProfilesDerivedResponseContainer.From*` functions to initialize the ListSiteDeviceProfilesDerivedResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSiteDeviceProfilesDerivedResponse object to a map representation for JSON marshaling.
func (l *ListSiteDeviceProfilesDerivedResponse) toMap() any {
    switch obj := l.value.(type) {
    case *DeviceprofileAp:
        return obj.toMap()
    case *DeviceprofileGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSiteDeviceProfilesDerivedResponse.
// It customizes the JSON unmarshaling process for ListSiteDeviceProfilesDerivedResponse objects.
func (l *ListSiteDeviceProfilesDerivedResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceprofileAp{}, false, &l.isDeviceprofileAp),
        NewTypeHolder(&DeviceprofileGateway{}, false, &l.isDeviceprofileGateway),
    )
    
    l.value = result
    return err
}

func (l *ListSiteDeviceProfilesDerivedResponse) AsDeviceprofileAp() (
    *DeviceprofileAp,
    bool) {
    if !l.isDeviceprofileAp {
        return nil, false
    }
    return l.value.(*DeviceprofileAp), true
}

func (l *ListSiteDeviceProfilesDerivedResponse) AsDeviceprofileGateway() (
    *DeviceprofileGateway,
    bool) {
    if !l.isDeviceprofileGateway {
        return nil, false
    }
    return l.value.(*DeviceprofileGateway), true
}

// internalListSiteDeviceProfilesDerivedResponse represents a listSiteDeviceProfilesDerivedResponse struct.
// This is Array of a container for any-of cases.
type internalListSiteDeviceProfilesDerivedResponse struct {}

var ListSiteDeviceProfilesDerivedResponseContainer internalListSiteDeviceProfilesDerivedResponse

// The internalListSiteDeviceProfilesDerivedResponse instance, wrapping the provided DeviceprofileAp value.
func (l *internalListSiteDeviceProfilesDerivedResponse) FromDeviceprofileAp(val DeviceprofileAp) ListSiteDeviceProfilesDerivedResponse {
    return ListSiteDeviceProfilesDerivedResponse{value: &val}
}

// The internalListSiteDeviceProfilesDerivedResponse instance, wrapping the provided DeviceprofileGateway value.
func (l *internalListSiteDeviceProfilesDerivedResponse) FromDeviceprofileGateway(val DeviceprofileGateway) ListSiteDeviceProfilesDerivedResponse {
    return ListSiteDeviceProfilesDerivedResponse{value: &val}
}
