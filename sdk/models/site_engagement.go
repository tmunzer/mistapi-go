package models

import (
    "encoding/json"
)

// SiteEngagement represents a SiteEngagement struct.
// **Note**: if hours does not exist, it’s treated as everyday of the week, 00:00-23:59. Currently we don’t allow multiple ranges for the same day
// **Note**: default values for `dwell_tags`: passerby (1,300) bounce (301, 14400) engaged (14401, 28800) stationed (28801, 42000)
// **Note**: default values for `dwell_tag_names`: passerby = “Passerby”, bounce = “Visitor”, engaged = “Associates”, stationed = “Assets”
type SiteEngagement struct {
    DwellTagNames        *SiteEngagementDwellTagNames `json:"dwell_tag_names,omitempty"`
    // add tags to visits within the duration (in seconds), available tags (passerby, bounce, engaged, stationed)
    DwellTags            *SiteEngagementDwellTags     `json:"dwell_tags,omitempty"`
    // hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun). 
    // **Note**: If the dow is not defined then it’s treated as 00:00-23:59.
    Hours                *Hours                       `json:"hours,omitempty"`
    // max time, default is 43200(12h), max is 68400 (18h)
    MaxDwell             *int                         `json:"max_dwell,omitempty"`
    // min time
    MinDwell             *int                         `json:"min_dwell,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagement.
// It customizes the JSON marshaling process for SiteEngagement objects.
func (s SiteEngagement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagement object to a map representation for JSON marshaling.
func (s SiteEngagement) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DwellTagNames != nil {
        structMap["dwell_tag_names"] = s.DwellTagNames.toMap()
    }
    if s.DwellTags != nil {
        structMap["dwell_tags"] = s.DwellTags.toMap()
    }
    if s.Hours != nil {
        structMap["hours"] = s.Hours.toMap()
    }
    if s.MaxDwell != nil {
        structMap["max_dwell"] = s.MaxDwell
    }
    if s.MinDwell != nil {
        structMap["min_dwell"] = s.MinDwell
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteEngagement.
// It customizes the JSON unmarshaling process for SiteEngagement objects.
func (s *SiteEngagement) UnmarshalJSON(input []byte) error {
    var temp siteEngagement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dwell_tag_names", "dwell_tags", "hours", "max_dwell", "min_dwell")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.DwellTagNames = temp.DwellTagNames
    s.DwellTags = temp.DwellTags
    s.Hours = temp.Hours
    s.MaxDwell = temp.MaxDwell
    s.MinDwell = temp.MinDwell
    return nil
}

// siteEngagement is a temporary struct used for validating the fields of SiteEngagement.
type siteEngagement  struct {
    DwellTagNames *SiteEngagementDwellTagNames `json:"dwell_tag_names,omitempty"`
    DwellTags     *SiteEngagementDwellTags     `json:"dwell_tags,omitempty"`
    Hours         *Hours                       `json:"hours,omitempty"`
    MaxDwell      *int                         `json:"max_dwell,omitempty"`
    MinDwell      *int                         `json:"min_dwell,omitempty"`
}
