package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedChassis represents a SleImpactedChassis struct.
type SleImpactedChassis struct {
    Chassis              []SleImpactedChassisChassisItem `json:"chassis,omitempty"`
    Classifier           *string                         `json:"classifier,omitempty"`
    End                  *int                            `json:"end,omitempty"`
    Failure              *string                         `json:"failure,omitempty"`
    Limit                *int                            `json:"limit,omitempty"`
    Metric               *string                         `json:"metric,omitempty"`
    Page                 *int                            `json:"page,omitempty"`
    Start                *int                            `json:"start,omitempty"`
    TotalCount           *int                            `json:"total_count,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedChassis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedChassis) String() string {
    return fmt.Sprintf(
    	"SleImpactedChassis[Chassis=%v, Classifier=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, AdditionalProperties=%v]",
    	s.Chassis, s.Classifier, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedChassis.
// It customizes the JSON marshaling process for SleImpactedChassis objects.
func (s SleImpactedChassis) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "chassis", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedChassis object to a map representation for JSON marshaling.
func (s SleImpactedChassis) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Chassis != nil {
        structMap["chassis"] = s.Chassis
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

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedChassis.
// It customizes the JSON unmarshaling process for SleImpactedChassis objects.
func (s *SleImpactedChassis) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedChassis
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chassis", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Chassis = temp.Chassis
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

// tempSleImpactedChassis is a temporary struct used for validating the fields of SleImpactedChassis.
type tempSleImpactedChassis  struct {
    Chassis    []SleImpactedChassisChassisItem `json:"chassis,omitempty"`
    Classifier *string                         `json:"classifier,omitempty"`
    End        *int                            `json:"end,omitempty"`
    Failure    *string                         `json:"failure,omitempty"`
    Limit      *int                            `json:"limit,omitempty"`
    Metric     *string                         `json:"metric,omitempty"`
    Page       *int                            `json:"page,omitempty"`
    Start      *int                            `json:"start,omitempty"`
    TotalCount *int                            `json:"total_count,omitempty"`
}
