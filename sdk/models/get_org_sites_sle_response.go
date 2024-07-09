package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GetOrgSitesSleResponse represents a GetOrgSitesSleResponse struct.
// This is a container for any-of cases.
type GetOrgSitesSleResponse struct {
    value              any
    isOrgSiteSleWifi   bool
    isOrgSiteWiredWifi bool
    isOrgSiteWanWifi   bool
}

// String converts the GetOrgSitesSleResponse object to a string representation.
func (g GetOrgSitesSleResponse) String() string {
    if bytes, err := json.Marshal(g.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for GetOrgSitesSleResponse.
// It customizes the JSON marshaling process for GetOrgSitesSleResponse objects.
func (g GetOrgSitesSleResponse) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GetOrgSitesSleResponseContainer.From*` functions to initialize the GetOrgSitesSleResponse object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetOrgSitesSleResponse object to a map representation for JSON marshaling.
func (g *GetOrgSitesSleResponse) toMap() any {
    switch obj := g.value.(type) {
    case *OrgSiteSleWifi:
        return obj.toMap()
    case *OrgSiteWiredWifi:
        return obj.toMap()
    case *OrgSiteWanWifi:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrgSitesSleResponse.
// It customizes the JSON unmarshaling process for GetOrgSitesSleResponse objects.
func (g *GetOrgSitesSleResponse) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&OrgSiteSleWifi{}, false, &g.isOrgSiteSleWifi),
        NewTypeHolder(&OrgSiteWiredWifi{}, false, &g.isOrgSiteWiredWifi),
        NewTypeHolder(&OrgSiteWanWifi{}, false, &g.isOrgSiteWanWifi),
    )
    
    g.value = result
    return err
}

func (g *GetOrgSitesSleResponse) AsOrgSiteSleWifi() (
    *OrgSiteSleWifi,
    bool) {
    if !g.isOrgSiteSleWifi {
        return nil, false
    }
    return g.value.(*OrgSiteSleWifi), true
}

func (g *GetOrgSitesSleResponse) AsOrgSiteWiredWifi() (
    *OrgSiteWiredWifi,
    bool) {
    if !g.isOrgSiteWiredWifi {
        return nil, false
    }
    return g.value.(*OrgSiteWiredWifi), true
}

func (g *GetOrgSitesSleResponse) AsOrgSiteWanWifi() (
    *OrgSiteWanWifi,
    bool) {
    if !g.isOrgSiteWanWifi {
        return nil, false
    }
    return g.value.(*OrgSiteWanWifi), true
}

// internalGetOrgSitesSleResponse represents a getOrgSitesSleResponse struct.
// This is a container for any-of cases.
type internalGetOrgSitesSleResponse struct {}

var GetOrgSitesSleResponseContainer internalGetOrgSitesSleResponse

// The internalGetOrgSitesSleResponse instance, wrapping the provided OrgSiteSleWifi value.
func (g *internalGetOrgSitesSleResponse) FromOrgSiteSleWifi(val OrgSiteSleWifi) GetOrgSitesSleResponse {
    return GetOrgSitesSleResponse{value: &val}
}

// The internalGetOrgSitesSleResponse instance, wrapping the provided OrgSiteWiredWifi value.
func (g *internalGetOrgSitesSleResponse) FromOrgSiteWiredWifi(val OrgSiteWiredWifi) GetOrgSitesSleResponse {
    return GetOrgSitesSleResponse{value: &val}
}

// The internalGetOrgSitesSleResponse instance, wrapping the provided OrgSiteWanWifi value.
func (g *internalGetOrgSitesSleResponse) FromOrgSiteWanWifi(val OrgSiteWanWifi) GetOrgSitesSleResponse {
    return GetOrgSitesSleResponse{value: &val}
}
