package models

import (
    "encoding/json"
)

// SiteWids represents a SiteWids struct.
// WIDS site settings
type SiteWids struct {
    RepeatedAuthFailures *SiteWidsRepeatedAuthFailures `json:"repeated_auth_failures,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteWids.
// It customizes the JSON marshaling process for SiteWids objects.
func (s SiteWids) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "repeated_auth_failures"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteWids object to a map representation for JSON marshaling.
func (s SiteWids) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.RepeatedAuthFailures != nil {
        structMap["repeated_auth_failures"] = s.RepeatedAuthFailures.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteWids.
// It customizes the JSON unmarshaling process for SiteWids objects.
func (s *SiteWids) UnmarshalJSON(input []byte) error {
    var temp tempSiteWids
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "repeated_auth_failures")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.RepeatedAuthFailures = temp.RepeatedAuthFailures
    return nil
}

// tempSiteWids is a temporary struct used for validating the fields of SiteWids.
type tempSiteWids  struct {
    RepeatedAuthFailures *SiteWidsRepeatedAuthFailures `json:"repeated_auth_failures,omitempty"`
}
