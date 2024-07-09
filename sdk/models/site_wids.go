package models

import (
    "encoding/json"
)

// SiteWids represents a SiteWids struct.
// WIDS site settings
type SiteWids struct {
    RepeatedAuthFailures *SiteWidsRepeatedAuthFailures `json:"repeated_auth_failures,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteWids.
// It customizes the JSON marshaling process for SiteWids objects.
func (s SiteWids) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteWids object to a map representation for JSON marshaling.
func (s SiteWids) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.RepeatedAuthFailures != nil {
        structMap["repeated_auth_failures"] = s.RepeatedAuthFailures.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteWids.
// It customizes the JSON unmarshaling process for SiteWids objects.
func (s *SiteWids) UnmarshalJSON(input []byte) error {
    var temp siteWids
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "repeated_auth_failures")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.RepeatedAuthFailures = temp.RepeatedAuthFailures
    return nil
}

// siteWids is a temporary struct used for validating the fields of SiteWids.
type siteWids  struct {
    RepeatedAuthFailures *SiteWidsRepeatedAuthFailures `json:"repeated_auth_failures,omitempty"`
}
