package models

import (
    "encoding/json"
)

// SiteEngagementDwellTags represents a SiteEngagementDwellTags struct.
// add tags to visits within the duration (in seconds), available tags (passerby, bounce, engaged, stationed)
type SiteEngagementDwellTags struct {
    Bounce               Optional[string] `json:"bounce"`
    Engaged              Optional[string] `json:"engaged"`
    Passerby             Optional[string] `json:"passerby"`
    Stationed            Optional[string] `json:"stationed"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagementDwellTags.
// It customizes the JSON marshaling process for SiteEngagementDwellTags objects.
func (s SiteEngagementDwellTags) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagementDwellTags object to a map representation for JSON marshaling.
func (s SiteEngagementDwellTags) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Bounce.IsValueSet() {
        if s.Bounce.Value() != nil {
            structMap["bounce"] = s.Bounce.Value()
        } else {
            structMap["bounce"] = nil
        }
    }
    if s.Engaged.IsValueSet() {
        if s.Engaged.Value() != nil {
            structMap["engaged"] = s.Engaged.Value()
        } else {
            structMap["engaged"] = nil
        }
    }
    if s.Passerby.IsValueSet() {
        if s.Passerby.Value() != nil {
            structMap["passerby"] = s.Passerby.Value()
        } else {
            structMap["passerby"] = nil
        }
    }
    if s.Stationed.IsValueSet() {
        if s.Stationed.Value() != nil {
            structMap["stationed"] = s.Stationed.Value()
        } else {
            structMap["stationed"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteEngagementDwellTags.
// It customizes the JSON unmarshaling process for SiteEngagementDwellTags objects.
func (s *SiteEngagementDwellTags) UnmarshalJSON(input []byte) error {
    var temp siteEngagementDwellTags
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

// siteEngagementDwellTags is a temporary struct used for validating the fields of SiteEngagementDwellTags.
type siteEngagementDwellTags  struct {
    Bounce    Optional[string] `json:"bounce"`
    Engaged   Optional[string] `json:"engaged"`
    Passerby  Optional[string] `json:"passerby"`
    Stationed Optional[string] `json:"stationed"`
}
