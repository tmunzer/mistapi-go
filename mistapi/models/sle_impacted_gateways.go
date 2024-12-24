package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedGateways represents a SleImpactedGateways struct.
type SleImpactedGateways struct {
    Classifier           *string                      `json:"classifier,omitempty"`
    End                  *int                         `json:"end,omitempty"`
    Failure              *string                      `json:"failure,omitempty"`
    Gateways             []SleImpactedGatewaysGateway `json:"gateways,omitempty"`
    Limit                *int                         `json:"limit,omitempty"`
    Metric               *string                      `json:"metric,omitempty"`
    Page                 *int                         `json:"page,omitempty"`
    Start                *int                         `json:"start,omitempty"`
    TotalCount           *int                         `json:"total_count,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedGateways,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedGateways) String() string {
    return fmt.Sprintf(
    	"SleImpactedGateways[Classifier=%v, End=%v, Failure=%v, Gateways=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, AdditionalProperties=%v]",
    	s.Classifier, s.End, s.Failure, s.Gateways, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedGateways.
// It customizes the JSON marshaling process for SleImpactedGateways objects.
func (s SleImpactedGateways) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "classifier", "end", "failure", "gateways", "limit", "metric", "page", "start", "total_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedGateways object to a map representation for JSON marshaling.
func (s SleImpactedGateways) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Classifier != nil {
        structMap["classifier"] = s.Classifier
    }
    if s.End != nil {
        structMap["end"] = s.End
    }
    if s.Failure != nil {
        structMap["failure"] = s.Failure
    }
    if s.Gateways != nil {
        structMap["gateways"] = s.Gateways
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

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedGateways.
// It customizes the JSON unmarshaling process for SleImpactedGateways objects.
func (s *SleImpactedGateways) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedGateways
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifier", "end", "failure", "gateways", "limit", "metric", "page", "start", "total_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Classifier = temp.Classifier
    s.End = temp.End
    s.Failure = temp.Failure
    s.Gateways = temp.Gateways
    s.Limit = temp.Limit
    s.Metric = temp.Metric
    s.Page = temp.Page
    s.Start = temp.Start
    s.TotalCount = temp.TotalCount
    return nil
}

// tempSleImpactedGateways is a temporary struct used for validating the fields of SleImpactedGateways.
type tempSleImpactedGateways  struct {
    Classifier *string                      `json:"classifier,omitempty"`
    End        *int                         `json:"end,omitempty"`
    Failure    *string                      `json:"failure,omitempty"`
    Gateways   []SleImpactedGatewaysGateway `json:"gateways,omitempty"`
    Limit      *int                         `json:"limit,omitempty"`
    Metric     *string                      `json:"metric,omitempty"`
    Page       *int                         `json:"page,omitempty"`
    Start      *int                         `json:"start,omitempty"`
    TotalCount *int                         `json:"total_count,omitempty"`
}
