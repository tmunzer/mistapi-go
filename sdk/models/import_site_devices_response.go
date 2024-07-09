package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ImportSiteDevicesResponse represents a ImportSiteDevicesResponse struct.
// This is Array of a container for any-of cases.
type ImportSiteDevicesResponse struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String converts the ImportSiteDevicesResponse object to a string representation.
func (i ImportSiteDevicesResponse) String() string {
    if bytes, err := json.Marshal(i.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ImportSiteDevicesResponse.
// It customizes the JSON marshaling process for ImportSiteDevicesResponse objects.
func (i ImportSiteDevicesResponse) MarshalJSON() (
    []byte,
    error) {
    if i.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ImportSiteDevicesResponseContainer.From*` functions to initialize the ImportSiteDevicesResponse object.")
    }
    return json.Marshal(i.toMap())
}

// toMap converts the ImportSiteDevicesResponse object to a map representation for JSON marshaling.
func (i *ImportSiteDevicesResponse) toMap() any {
    switch obj := i.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ImportSiteDevicesResponse.
// It customizes the JSON unmarshaling process for ImportSiteDevicesResponse objects.
func (i *ImportSiteDevicesResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&DeviceAp{}, false, &i.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &i.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &i.isDeviceGateway),
    )
    
    i.value = result
    return err
}

func (i *ImportSiteDevicesResponse) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !i.isDeviceAp {
        return nil, false
    }
    return i.value.(*DeviceAp), true
}

func (i *ImportSiteDevicesResponse) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !i.isDeviceSwitch {
        return nil, false
    }
    return i.value.(*DeviceSwitch), true
}

func (i *ImportSiteDevicesResponse) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !i.isDeviceGateway {
        return nil, false
    }
    return i.value.(*DeviceGateway), true
}

// internalImportSiteDevicesResponse represents a importSiteDevicesResponse struct.
// This is Array of a container for any-of cases.
type internalImportSiteDevicesResponse struct {}

var ImportSiteDevicesResponseContainer internalImportSiteDevicesResponse

// The internalImportSiteDevicesResponse instance, wrapping the provided DeviceAp value.
func (i *internalImportSiteDevicesResponse) FromDeviceAp(val DeviceAp) ImportSiteDevicesResponse {
    return ImportSiteDevicesResponse{value: &val}
}

// The internalImportSiteDevicesResponse instance, wrapping the provided DeviceSwitch value.
func (i *internalImportSiteDevicesResponse) FromDeviceSwitch(val DeviceSwitch) ImportSiteDevicesResponse {
    return ImportSiteDevicesResponse{value: &val}
}

// The internalImportSiteDevicesResponse instance, wrapping the provided DeviceGateway value.
func (i *internalImportSiteDevicesResponse) FromDeviceGateway(val DeviceGateway) ImportSiteDevicesResponse {
    return ImportSiteDevicesResponse{value: &val}
}
