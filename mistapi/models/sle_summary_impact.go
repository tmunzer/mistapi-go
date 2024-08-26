package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleSummaryImpact represents a SleSummaryImpact struct.
type SleSummaryImpact struct {
    NumAps               float64        `json:"num_aps"`
    NumUsers             float64        `json:"num_users"`
    TotalAps             float64        `json:"total_aps"`
    TotalUsers           float64        `json:"total_users"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleSummaryImpact.
// It customizes the JSON marshaling process for SleSummaryImpact objects.
func (s SleSummaryImpact) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleSummaryImpact object to a map representation for JSON marshaling.
func (s SleSummaryImpact) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["num_aps"] = s.NumAps
    structMap["num_users"] = s.NumUsers
    structMap["total_aps"] = s.TotalAps
    structMap["total_users"] = s.TotalUsers
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleSummaryImpact.
// It customizes the JSON unmarshaling process for SleSummaryImpact objects.
func (s *SleSummaryImpact) UnmarshalJSON(input []byte) error {
    var temp tempSleSummaryImpact
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_users", "total_aps", "total_users")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.NumAps = *temp.NumAps
    s.NumUsers = *temp.NumUsers
    s.TotalAps = *temp.TotalAps
    s.TotalUsers = *temp.TotalUsers
    return nil
}

// tempSleSummaryImpact is a temporary struct used for validating the fields of SleSummaryImpact.
type tempSleSummaryImpact  struct {
    NumAps     *float64 `json:"num_aps"`
    NumUsers   *float64 `json:"num_users"`
    TotalAps   *float64 `json:"total_aps"`
    TotalUsers *float64 `json:"total_users"`
}

func (s *tempSleSummaryImpact) validate() error {
    var errs []string
    if s.NumAps == nil {
        errs = append(errs, "required field `num_aps` is missing for type `sle_summary_impact`")
    }
    if s.NumUsers == nil {
        errs = append(errs, "required field `num_users` is missing for type `sle_summary_impact`")
    }
    if s.TotalAps == nil {
        errs = append(errs, "required field `total_aps` is missing for type `sle_summary_impact`")
    }
    if s.TotalUsers == nil {
        errs = append(errs, "required field `total_users` is missing for type `sle_summary_impact`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
