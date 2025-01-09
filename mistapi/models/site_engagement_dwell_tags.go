package models

import (
    "encoding/json"
    "fmt"
)

// SiteEngagementDwellTags represents a SiteEngagementDwellTags struct.
// add tags to visits within the duration (in seconds)
type SiteEngagementDwellTags struct {
    Bounce               Optional[string]       `json:"bounce"`
    Engaged              Optional[string]       `json:"engaged"`
    Passerby             Optional[string]       `json:"passerby"`
    Stationed            Optional[string]       `json:"stationed"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteEngagementDwellTags,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteEngagementDwellTags) String() string {
    return fmt.Sprintf(
    	"SiteEngagementDwellTags[Bounce=%v, Engaged=%v, Passerby=%v, Stationed=%v, AdditionalProperties=%v]",
    	s.Bounce, s.Engaged, s.Passerby, s.Stationed, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagementDwellTags.
// It customizes the JSON marshaling process for SiteEngagementDwellTags objects.
func (s SiteEngagementDwellTags) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "bounce", "engaged", "passerby", "stationed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagementDwellTags object to a map representation for JSON marshaling.
func (s SiteEngagementDwellTags) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSiteEngagementDwellTags
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bounce", "engaged", "passerby", "stationed")
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

// tempSiteEngagementDwellTags is a temporary struct used for validating the fields of SiteEngagementDwellTags.
type tempSiteEngagementDwellTags  struct {
    Bounce    Optional[string] `json:"bounce"`
    Engaged   Optional[string] `json:"engaged"`
    Passerby  Optional[string] `json:"passerby"`
    Stationed Optional[string] `json:"stationed"`
}
