package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListSiteDevicesResponse represents a ListSiteDevicesResponse struct.
// This is Array of a container for any-of cases.
type ListSiteDevicesResponse struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String converts the ListSiteDevicesResponse object to a string representation.
func (l ListSiteDevicesResponse) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListSiteDevicesResponse.
// It customizes the JSON marshaling process for ListSiteDevicesResponse objects.
func (l ListSiteDevicesResponse) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListSiteDevicesResponseContainer.From*` functions to initialize the ListSiteDevicesResponse object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSiteDevicesResponse object to a map representation for JSON marshaling.
func (l *ListSiteDevicesResponse) toMap() any {
    switch obj := l.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSiteDevicesResponse.
// It customizes the JSON unmarshaling process for ListSiteDevicesResponse objects.
func (l *ListSiteDevicesResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceAp{}, false, &l.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &l.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &l.isDeviceGateway),
    )
    
    l.value = result
    return err
}

func (l *ListSiteDevicesResponse) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !l.isDeviceAp {
        return nil, false
    }
    return l.value.(*DeviceAp), true
}

func (l *ListSiteDevicesResponse) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !l.isDeviceSwitch {
        return nil, false
    }
    return l.value.(*DeviceSwitch), true
}

func (l *ListSiteDevicesResponse) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !l.isDeviceGateway {
        return nil, false
    }
    return l.value.(*DeviceGateway), true
}

// internalListSiteDevicesResponse represents a listSiteDevicesResponse struct.
// This is Array of a container for any-of cases.
type internalListSiteDevicesResponse struct {}

var ListSiteDevicesResponseContainer internalListSiteDevicesResponse

// The internalListSiteDevicesResponse instance, wrapping the provided DeviceAp value.
func (l *internalListSiteDevicesResponse) FromDeviceAp(val DeviceAp) ListSiteDevicesResponse {
    return ListSiteDevicesResponse{value: &val}
}

// The internalListSiteDevicesResponse instance, wrapping the provided DeviceSwitch value.
func (l *internalListSiteDevicesResponse) FromDeviceSwitch(val DeviceSwitch) ListSiteDevicesResponse {
    return ListSiteDevicesResponse{value: &val}
}

// The internalListSiteDevicesResponse instance, wrapping the provided DeviceGateway value.
func (l *internalListSiteDevicesResponse) FromDeviceGateway(val DeviceGateway) ListSiteDevicesResponse {
    return ListSiteDevicesResponse{value: &val}
}
