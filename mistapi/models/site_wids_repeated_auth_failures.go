package models

import (
    "encoding/json"
    "fmt"
)

// SiteWidsRepeatedAuthFailures represents a SiteWidsRepeatedAuthFailures struct.
type SiteWidsRepeatedAuthFailures struct {
    // Window where a trigger will be detected and action to be taken (in seconds)
    Duration             *int                   `json:"duration,omitempty"`
    // Count of events to trigger
    Threshold            *int                   `json:"threshold,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteWidsRepeatedAuthFailures,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteWidsRepeatedAuthFailures) String() string {
    return fmt.Sprintf(
    	"SiteWidsRepeatedAuthFailures[Duration=%v, Threshold=%v, AdditionalProperties=%v]",
    	s.Duration, s.Threshold, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteWidsRepeatedAuthFailures.
// It customizes the JSON marshaling process for SiteWidsRepeatedAuthFailures objects.
func (s SiteWidsRepeatedAuthFailures) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "duration", "threshold"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteWidsRepeatedAuthFailures object to a map representation for JSON marshaling.
func (s SiteWidsRepeatedAuthFailures) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Threshold != nil {
        structMap["threshold"] = s.Threshold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteWidsRepeatedAuthFailures.
// It customizes the JSON unmarshaling process for SiteWidsRepeatedAuthFailures objects.
func (s *SiteWidsRepeatedAuthFailures) UnmarshalJSON(input []byte) error {
    var temp tempSiteWidsRepeatedAuthFailures
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "threshold")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Duration = temp.Duration
    s.Threshold = temp.Threshold
    return nil
}

// tempSiteWidsRepeatedAuthFailures is a temporary struct used for validating the fields of SiteWidsRepeatedAuthFailures.
type tempSiteWidsRepeatedAuthFailures  struct {
    Duration  *int `json:"duration,omitempty"`
    Threshold *int `json:"threshold,omitempty"`
}
