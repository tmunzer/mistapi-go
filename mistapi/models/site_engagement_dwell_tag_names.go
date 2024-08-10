package models

import (
    "encoding/json"
)

// SiteEngagementDwellTagNames represents a SiteEngagementDwellTagNames struct.
type SiteEngagementDwellTagNames struct {
    Bounce               *string        `json:"bounce,omitempty"`
    Engaged              *string        `json:"engaged,omitempty"`
    Passerby             *string        `json:"passerby,omitempty"`
    Stationed            *string        `json:"stationed,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagementDwellTagNames.
// It customizes the JSON marshaling process for SiteEngagementDwellTagNames objects.
func (s SiteEngagementDwellTagNames) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagementDwellTagNames object to a map representation for JSON marshaling.
func (s SiteEngagementDwellTagNames) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Bounce != nil {
        structMap["bounce"] = s.Bounce
    }
    if s.Engaged != nil {
        structMap["engaged"] = s.Engaged
    }
    if s.Passerby != nil {
        structMap["passerby"] = s.Passerby
    }
    if s.Stationed != nil {
        structMap["stationed"] = s.Stationed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteEngagementDwellTagNames.
// It customizes the JSON unmarshaling process for SiteEngagementDwellTagNames objects.
func (s *SiteEngagementDwellTagNames) UnmarshalJSON(input []byte) error {
    var temp tempSiteEngagementDwellTagNames
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bounce", "engaged", "passerby", "stationed")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Bounce = temp.Bounce
    s.Engaged = temp.Engaged
    s.Passerby = temp.Passerby
    s.Stationed = temp.Stationed
    return nil
}

// tempSiteEngagementDwellTagNames is a temporary struct used for validating the fields of SiteEngagementDwellTagNames.
type tempSiteEngagementDwellTagNames  struct {
    Bounce    *string `json:"bounce,omitempty"`
    Engaged   *string `json:"engaged,omitempty"`
    Passerby  *string `json:"passerby,omitempty"`
    Stationed *string `json:"stationed,omitempty"`
}
