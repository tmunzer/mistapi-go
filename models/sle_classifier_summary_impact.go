package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleClassifierSummaryImpact represents a SleClassifierSummaryImpact struct.
type SleClassifierSummaryImpact struct {
    NumAps               float64        `json:"num_aps"`
    NumUsers             float64        `json:"num_users"`
    TotalAps             float64        `json:"total_aps"`
    TotalUsers           float64        `json:"total_users"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierSummaryImpact.
// It customizes the JSON marshaling process for SleClassifierSummaryImpact objects.
func (s SleClassifierSummaryImpact) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierSummaryImpact object to a map representation for JSON marshaling.
func (s SleClassifierSummaryImpact) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["num_aps"] = s.NumAps
    structMap["num_users"] = s.NumUsers
    structMap["total_aps"] = s.TotalAps
    structMap["total_users"] = s.TotalUsers
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierSummaryImpact.
// It customizes the JSON unmarshaling process for SleClassifierSummaryImpact objects.
func (s *SleClassifierSummaryImpact) UnmarshalJSON(input []byte) error {
    var temp sleClassifierSummaryImpact
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

// sleClassifierSummaryImpact is a temporary struct used for validating the fields of SleClassifierSummaryImpact.
type sleClassifierSummaryImpact  struct {
    NumAps     *float64 `json:"num_aps"`
    NumUsers   *float64 `json:"num_users"`
    TotalAps   *float64 `json:"total_aps"`
    TotalUsers *float64 `json:"total_users"`
}

func (s *sleClassifierSummaryImpact) validate() error {
    var errs []string
    if s.NumAps == nil {
        errs = append(errs, "required field `num_aps` is missing for type `Sle_Classifier_Summary_Impact`")
    }
    if s.NumUsers == nil {
        errs = append(errs, "required field `num_users` is missing for type `Sle_Classifier_Summary_Impact`")
    }
    if s.TotalAps == nil {
        errs = append(errs, "required field `total_aps` is missing for type `Sle_Classifier_Summary_Impact`")
    }
    if s.TotalUsers == nil {
        errs = append(errs, "required field `total_users` is missing for type `Sle_Classifier_Summary_Impact`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
