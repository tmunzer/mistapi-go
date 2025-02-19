package models

import (
    "encoding/json"
    "fmt"
)

// SiteEngagement represents a SiteEngagement struct.
// **Note**: if hours does not exist, it's treated as everyday of the week, 00:00-23:59. Currently, we don't allow multiple ranges for the same day
type SiteEngagement struct {
    // Name associated to each tag
    DwellTagNames        *SiteEngagementDwellTagNames `json:"dwell_tag_names,omitempty"`
    // add tags to visits within the duration (in seconds)
    DwellTags            *SiteEngagementDwellTags     `json:"dwell_tags,omitempty"`
    // Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun)
    Hours                *Hours                       `json:"hours,omitempty"`
    // Max time, default is 43200(12h), max is 68400 (18h)
    MaxDwell             *int                         `json:"max_dwell,omitempty"`
    // min time
    MinDwell             *int                         `json:"min_dwell,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SiteEngagement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteEngagement) String() string {
    return fmt.Sprintf(
    	"SiteEngagement[DwellTagNames=%v, DwellTags=%v, Hours=%v, MaxDwell=%v, MinDwell=%v, AdditionalProperties=%v]",
    	s.DwellTagNames, s.DwellTags, s.Hours, s.MaxDwell, s.MinDwell, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagement.
// It customizes the JSON marshaling process for SiteEngagement objects.
func (s SiteEngagement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "dwell_tag_names", "dwell_tags", "hours", "max_dwell", "min_dwell"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagement object to a map representation for JSON marshaling.
func (s SiteEngagement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSiteEngagement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dwell_tag_names", "dwell_tags", "hours", "max_dwell", "min_dwell")
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

// tempSiteEngagement is a temporary struct used for validating the fields of SiteEngagement.
type tempSiteEngagement  struct {
    DwellTagNames *SiteEngagementDwellTagNames `json:"dwell_tag_names,omitempty"`
    DwellTags     *SiteEngagementDwellTags     `json:"dwell_tags,omitempty"`
    Hours         *Hours                       `json:"hours,omitempty"`
    MaxDwell      *int                         `json:"max_dwell,omitempty"`
    MinDwell      *int                         `json:"min_dwell,omitempty"`
}
