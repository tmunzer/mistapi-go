package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseOrgSiteSle2 represents a ResponseOrgSiteSle2 struct.
type ResponseOrgSiteSle2 struct {
    value              any
    isOrgSiteSleWifi   bool
    isOrgSiteWiredWifi bool
    isOrgSiteWanWifi   bool
}

// String converts the ResponseOrgSiteSle2 object to a string representation.
func (r ResponseOrgSiteSle2) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSiteSle2.
// It customizes the JSON marshaling process for ResponseOrgSiteSle2 objects.
func (r ResponseOrgSiteSle2) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOrgSiteSleContainer.From*` functions to initialize the ResponseOrgSiteSle2 object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSiteSle2 object to a map representation for JSON marshaling.
func (r *ResponseOrgSiteSle2) toMap() any {
    switch obj := r.value.(type) {
    case *OrgSiteSleWifi:
        return obj.toMap()
    case *OrgSiteWiredWifi:
        return obj.toMap()
    case *OrgSiteWanWifi:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSiteSle2.
// It customizes the JSON unmarshaling process for ResponseOrgSiteSle2 objects.
func (r *ResponseOrgSiteSle2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&OrgSiteSleWifi{}, false, &r.isOrgSiteSleWifi),
        NewTypeHolder(&OrgSiteWiredWifi{}, false, &r.isOrgSiteWiredWifi),
        NewTypeHolder(&OrgSiteWanWifi{}, false, &r.isOrgSiteWanWifi),
    )
    
    r.value = result
    return err
}

func (r *ResponseOrgSiteSle2) AsOrgSiteSleWifi() (
    *OrgSiteSleWifi,
    bool) {
    if !r.isOrgSiteSleWifi {
        return nil, false
    }
    return r.value.(*OrgSiteSleWifi), true
}

func (r *ResponseOrgSiteSle2) AsOrgSiteWiredWifi() (
    *OrgSiteWiredWifi,
    bool) {
    if !r.isOrgSiteWiredWifi {
        return nil, false
    }
    return r.value.(*OrgSiteWiredWifi), true
}

func (r *ResponseOrgSiteSle2) AsOrgSiteWanWifi() (
    *OrgSiteWanWifi,
    bool) {
    if !r.isOrgSiteWanWifi {
        return nil, false
    }
    return r.value.(*OrgSiteWanWifi), true
}

// internalResponseOrgSiteSle2 represents a responseOrgSiteSle2 struct.
type internalResponseOrgSiteSle2 struct {}

var ResponseOrgSiteSleContainer internalResponseOrgSiteSle2

// The internalResponseOrgSiteSle2 instance, wrapping the provided OrgSiteSleWifi value.
func (r *internalResponseOrgSiteSle2) FromOrgSiteSleWifi(val OrgSiteSleWifi) ResponseOrgSiteSle2 {
    return ResponseOrgSiteSle2{value: &val}
}

// The internalResponseOrgSiteSle2 instance, wrapping the provided OrgSiteWiredWifi value.
func (r *internalResponseOrgSiteSle2) FromOrgSiteWiredWifi(val OrgSiteWiredWifi) ResponseOrgSiteSle2 {
    return ResponseOrgSiteSle2{value: &val}
}

// The internalResponseOrgSiteSle2 instance, wrapping the provided OrgSiteWanWifi value.
func (r *internalResponseOrgSiteSle2) FromOrgSiteWanWifi(val OrgSiteWanWifi) ResponseOrgSiteSle2 {
    return ResponseOrgSiteSle2{value: &val}
}
