package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedApplications represents a SleImpactedApplications struct.
type SleImpactedApplications struct {
    Apps                 []SleImpactedApplicationsApp `json:"apps,omitempty"`
    Classifier           *string                      `json:"classifier,omitempty"`
    End                  *int                         `json:"end,omitempty"`
    Failure              *string                      `json:"failure,omitempty"`
    Limit                *string                      `json:"limit,omitempty"`
    Metric               *string                      `json:"metric,omitempty"`
    Page                 *int                         `json:"page,omitempty"`
    Start                *int                         `json:"start,omitempty"`
    TotalCount           *int                         `json:"total_count,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedApplications,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedApplications) String() string {
    return fmt.Sprintf(
    	"SleImpactedApplications[Apps=%v, Classifier=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, AdditionalProperties=%v]",
    	s.Apps, s.Classifier, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedApplications.
// It customizes the JSON marshaling process for SleImpactedApplications objects.
func (s SleImpactedApplications) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "apps", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedApplications object to a map representation for JSON marshaling.
func (s SleImpactedApplications) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Apps != nil {
        structMap["apps"] = s.Apps
    }
    if s.Classifier != nil {
        structMap["classifier"] = s.Classifier
    }
    if s.End != nil {
        structMap["end"] = s.End
    }
    if s.Failure != nil {
        structMap["failure"] = s.Failure
    }
    if s.Limit != nil {
        structMap["limit"] = s.Limit
    }
    if s.Metric != nil {
        structMap["metric"] = s.Metric
    }
    if s.Page != nil {
        structMap["page"] = s.Page
    }
    if s.Start != nil {
        structMap["start"] = s.Start
    }
    if s.TotalCount != nil {
        structMap["total_count"] = s.TotalCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedApplications.
// It customizes the JSON unmarshaling process for SleImpactedApplications objects.
func (s *SleImpactedApplications) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedApplications
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apps", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Apps = temp.Apps
    s.Classifier = temp.Classifier
    s.End = temp.End
    s.Failure = temp.Failure
    s.Limit = temp.Limit
    s.Metric = temp.Metric
    s.Page = temp.Page
    s.Start = temp.Start
    s.TotalCount = temp.TotalCount
    return nil
}

// tempSleImpactedApplications is a temporary struct used for validating the fields of SleImpactedApplications.
type tempSleImpactedApplications  struct {
    Apps       []SleImpactedApplicationsApp `json:"apps,omitempty"`
    Classifier *string                      `json:"classifier,omitempty"`
    End        *int                         `json:"end,omitempty"`
    Failure    *string                      `json:"failure,omitempty"`
    Limit      *string                      `json:"limit,omitempty"`
    Metric     *string                      `json:"metric,omitempty"`
    Page       *int                         `json:"page,omitempty"`
    Start      *int                         `json:"start,omitempty"`
    TotalCount *int                         `json:"total_count,omitempty"`
}
