package models

import (
    "encoding/json"
)

// SiteZoneOccupancyAlert represents a SiteZoneOccupancyAlert struct.
// Zone Occupancy alert site settings
type SiteZoneOccupancyAlert struct {
    // list of email addresses to send email notifications when the alert threshold is reached
    EmailNotifiers       []string               `json:"email_notifiers,omitempty"`
    // indicate whether zone occupancy alert is enabled for the site
    Enabled              *bool                  `json:"enabled,omitempty"`
    // sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes
    Threshold            *int                   `json:"threshold,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteZoneOccupancyAlert.
// It customizes the JSON marshaling process for SiteZoneOccupancyAlert objects.
func (s SiteZoneOccupancyAlert) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "email_notifiers", "enabled", "threshold"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteZoneOccupancyAlert object to a map representation for JSON marshaling.
func (s SiteZoneOccupancyAlert) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.EmailNotifiers != nil {
        structMap["email_notifiers"] = s.EmailNotifiers
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Threshold != nil {
        structMap["threshold"] = s.Threshold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteZoneOccupancyAlert.
// It customizes the JSON unmarshaling process for SiteZoneOccupancyAlert objects.
func (s *SiteZoneOccupancyAlert) UnmarshalJSON(input []byte) error {
    var temp tempSiteZoneOccupancyAlert
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "email_notifiers", "enabled", "threshold")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.EmailNotifiers = temp.EmailNotifiers
    s.Enabled = temp.Enabled
    s.Threshold = temp.Threshold
    return nil
}

// tempSiteZoneOccupancyAlert is a temporary struct used for validating the fields of SiteZoneOccupancyAlert.
type tempSiteZoneOccupancyAlert  struct {
    EmailNotifiers []string `json:"email_notifiers,omitempty"`
    Enabled        *bool    `json:"enabled,omitempty"`
    Threshold      *int     `json:"threshold,omitempty"`
}
