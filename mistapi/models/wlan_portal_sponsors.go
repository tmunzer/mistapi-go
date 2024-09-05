package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanPortalSponsors represents a WlanPortalSponsors struct.
// object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty.
// Property key is the sponsor email, Property value is the sponsor name
// List of email allowed for backward compatibility
type WlanPortalSponsors struct {
    value           any
    isArrayOfString bool
    isMapOfString   bool
}

// String converts the WlanPortalSponsors object to a string representation.
func (w WlanPortalSponsors) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalSponsors.
// It customizes the JSON marshaling process for WlanPortalSponsors objects.
func (w WlanPortalSponsors) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanPortalSponsorsContainer.From*` functions to initialize the WlanPortalSponsors object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalSponsors object to a map representation for JSON marshaling.
func (w *WlanPortalSponsors) toMap() any {
    switch obj := w.value.(type) {
    case *[]string:
        return *obj
    case *map[string]string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalSponsors.
// It customizes the JSON unmarshaling process for WlanPortalSponsors objects.
func (w *WlanPortalSponsors) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new([]string), false, &w.isArrayOfString),
        NewTypeHolder(new(map[string]string), false, &w.isMapOfString),
    )
    
    w.value = result
    return err
}

func (w *WlanPortalSponsors) AsArrayOfString() (
    *[]string,
    bool) {
    if !w.isArrayOfString {
        return nil, false
    }
    return w.value.(*[]string), true
}

func (w *WlanPortalSponsors) AsMapOfString() (
    *map[string]string,
    bool) {
    if !w.isMapOfString {
        return nil, false
    }
    return w.value.(*map[string]string), true
}

// internalWlanPortalSponsors represents a wlanPortalSponsors struct.
// object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty.
// Property key is the sponsor email, Property value is the sponsor name
// List of email allowed for backward compatibility
type internalWlanPortalSponsors struct {}

var WlanPortalSponsorsContainer internalWlanPortalSponsors

// The internalWlanPortalSponsors instance, wrapping the provided []string value.
func (w *internalWlanPortalSponsors) FromArrayOfString(val []string) WlanPortalSponsors {
    return WlanPortalSponsors{value: &val}
}

// The internalWlanPortalSponsors instance, wrapping the provided map[string]string value.
func (w *internalWlanPortalSponsors) FromMapOfString(val map[string]string) WlanPortalSponsors {
    return WlanPortalSponsors{value: &val}
}
