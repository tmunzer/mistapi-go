package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SiteSettingCriticalUrlMonitoringMonitorVlanId represents a SiteSettingCriticalUrlMonitoringMonitorVlanId struct.
// This is a container for one-of cases.
type SiteSettingCriticalUrlMonitoringMonitorVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SiteSettingCriticalUrlMonitoringMonitorVlanId object to a string representation.
func (s SiteSettingCriticalUrlMonitoringMonitorVlanId) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingCriticalUrlMonitoringMonitorVlanId.
// It customizes the JSON marshaling process for SiteSettingCriticalUrlMonitoringMonitorVlanId objects.
func (s SiteSettingCriticalUrlMonitoringMonitorVlanId) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SiteSettingCriticalUrlMonitoringMonitorVlanIdContainer.From*` functions to initialize the SiteSettingCriticalUrlMonitoringMonitorVlanId object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingCriticalUrlMonitoringMonitorVlanId object to a map representation for JSON marshaling.
func (s *SiteSettingCriticalUrlMonitoringMonitorVlanId) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingCriticalUrlMonitoringMonitorVlanId.
// It customizes the JSON unmarshaling process for SiteSettingCriticalUrlMonitoringMonitorVlanId objects.
func (s *SiteSettingCriticalUrlMonitoringMonitorVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SiteSettingCriticalUrlMonitoringMonitorVlanId) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SiteSettingCriticalUrlMonitoringMonitorVlanId) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSiteSettingCriticalUrlMonitoringMonitorVlanId represents a siteSettingCriticalUrlMonitoringMonitorVlanId struct.
// This is a container for one-of cases.
type internalSiteSettingCriticalUrlMonitoringMonitorVlanId struct {}

var SiteSettingCriticalUrlMonitoringMonitorVlanIdContainer internalSiteSettingCriticalUrlMonitoringMonitorVlanId

// The internalSiteSettingCriticalUrlMonitoringMonitorVlanId instance, wrapping the provided string value.
func (s *internalSiteSettingCriticalUrlMonitoringMonitorVlanId) FromString(val string) SiteSettingCriticalUrlMonitoringMonitorVlanId {
    return SiteSettingCriticalUrlMonitoringMonitorVlanId{value: &val}
}

// The internalSiteSettingCriticalUrlMonitoringMonitorVlanId instance, wrapping the provided int value.
func (s *internalSiteSettingCriticalUrlMonitoringMonitorVlanId) FromNumber(val int) SiteSettingCriticalUrlMonitoringMonitorVlanId {
    return SiteSettingCriticalUrlMonitoringMonitorVlanId{value: &val}
}
