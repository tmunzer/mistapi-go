package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GetSiteDeviceResponse represents a GetSiteDeviceResponse struct.
// This is a container for one-of cases.
type GetSiteDeviceResponse struct {
    value           any
    isDeviceAp      bool
    isDeviceSwitch  bool
    isDeviceGateway bool
}

// String converts the GetSiteDeviceResponse object to a string representation.
func (g GetSiteDeviceResponse) String() string {
    if bytes, err := json.Marshal(g.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for GetSiteDeviceResponse.
// It customizes the JSON marshaling process for GetSiteDeviceResponse objects.
func (g GetSiteDeviceResponse) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GetSiteDeviceResponseContainer.From*` functions to initialize the GetSiteDeviceResponse object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetSiteDeviceResponse object to a map representation for JSON marshaling.
func (g *GetSiteDeviceResponse) toMap() any {
    switch obj := g.value.(type) {
    case *DeviceAp:
        return obj.toMap()
    case *DeviceSwitch:
        return obj.toMap()
    case *DeviceGateway:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSiteDeviceResponse.
// It customizes the JSON unmarshaling process for GetSiteDeviceResponse objects.
func (g *GetSiteDeviceResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&DeviceAp{}, false, &g.isDeviceAp),
        NewTypeHolder(&DeviceSwitch{}, false, &g.isDeviceSwitch),
        NewTypeHolder(&DeviceGateway{}, false, &g.isDeviceGateway),
    )
    
    g.value = result
    return err
}

func (g *GetSiteDeviceResponse) AsDeviceAp() (
    *DeviceAp,
    bool) {
    if !g.isDeviceAp {
        return nil, false
    }
    return g.value.(*DeviceAp), true
}

func (g *GetSiteDeviceResponse) AsDeviceSwitch() (
    *DeviceSwitch,
    bool) {
    if !g.isDeviceSwitch {
        return nil, false
    }
    return g.value.(*DeviceSwitch), true
}

func (g *GetSiteDeviceResponse) AsDeviceGateway() (
    *DeviceGateway,
    bool) {
    if !g.isDeviceGateway {
        return nil, false
    }
    return g.value.(*DeviceGateway), true
}

// internalGetSiteDeviceResponse represents a getSiteDeviceResponse struct.
// This is a container for one-of cases.
type internalGetSiteDeviceResponse struct {}

var GetSiteDeviceResponseContainer internalGetSiteDeviceResponse

// The internalGetSiteDeviceResponse instance, wrapping the provided DeviceAp value.
func (g *internalGetSiteDeviceResponse) FromDeviceAp(val DeviceAp) GetSiteDeviceResponse {
    return GetSiteDeviceResponse{value: &val}
}

// The internalGetSiteDeviceResponse instance, wrapping the provided DeviceSwitch value.
func (g *internalGetSiteDeviceResponse) FromDeviceSwitch(val DeviceSwitch) GetSiteDeviceResponse {
    return GetSiteDeviceResponse{value: &val}
}

// The internalGetSiteDeviceResponse instance, wrapping the provided DeviceGateway value.
func (g *internalGetSiteDeviceResponse) FromDeviceGateway(val DeviceGateway) GetSiteDeviceResponse {
    return GetSiteDeviceResponse{value: &val}
}
