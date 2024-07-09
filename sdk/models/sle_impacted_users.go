package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactedUsers represents a SleImpactedUsers struct.
type SleImpactedUsers struct {
    Classifier           string                 `json:"classifier"`
    End                  float64                `json:"end"`
    Failure              string                 `json:"failure"`
    Limit                float64                `json:"limit"`
    Metric               string                 `json:"metric"`
    Page                 float64                `json:"page"`
    Start                float64                `json:"start"`
    TotalCount           float64                `json:"total_count"`
    Users                []SleImpactedUsersUser `json:"users"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedUsers.
// It customizes the JSON marshaling process for SleImpactedUsers objects.
func (s SleImpactedUsers) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedUsers object to a map representation for JSON marshaling.
func (s SleImpactedUsers) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["classifier"] = s.Classifier
    structMap["end"] = s.End
    structMap["failure"] = s.Failure
    structMap["limit"] = s.Limit
    structMap["metric"] = s.Metric
    structMap["page"] = s.Page
    structMap["start"] = s.Start
    structMap["total_count"] = s.TotalCount
    structMap["users"] = s.Users
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedUsers.
// It customizes the JSON unmarshaling process for SleImpactedUsers objects.
func (s *SleImpactedUsers) UnmarshalJSON(input []byte) error {
    var temp sleImpactedUsers
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count", "users")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Classifier = *temp.Classifier
    s.End = *temp.End
    s.Failure = *temp.Failure
    s.Limit = *temp.Limit
    s.Metric = *temp.Metric
    s.Page = *temp.Page
    s.Start = *temp.Start
    s.TotalCount = *temp.TotalCount
    s.Users = *temp.Users
    return nil
}

// sleImpactedUsers is a temporary struct used for validating the fields of SleImpactedUsers.
type sleImpactedUsers  struct {
    Classifier *string                 `json:"classifier"`
    End        *float64                `json:"end"`
    Failure    *string                 `json:"failure"`
    Limit      *float64                `json:"limit"`
    Metric     *string                 `json:"metric"`
    Page       *float64                `json:"page"`
    Start      *float64                `json:"start"`
    TotalCount *float64                `json:"total_count"`
    Users      *[]SleImpactedUsersUser `json:"users"`
}

func (s *sleImpactedUsers) validate() error {
    var errs []string
    if s.Classifier == nil {
        errs = append(errs, "required field `classifier` is missing for type `Sle_Impacted_Users`")
    }
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `Sle_Impacted_Users`")
    }
    if s.Failure == nil {
        errs = append(errs, "required field `failure` is missing for type `Sle_Impacted_Users`")
    }
    if s.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Sle_Impacted_Users`")
    }
    if s.Metric == nil {
        errs = append(errs, "required field `metric` is missing for type `Sle_Impacted_Users`")
    }
    if s.Page == nil {
        errs = append(errs, "required field `page` is missing for type `Sle_Impacted_Users`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Sle_Impacted_Users`")
    }
    if s.TotalCount == nil {
        errs = append(errs, "required field `total_count` is missing for type `Sle_Impacted_Users`")
    }
    if s.Users == nil {
        errs = append(errs, "required field `users` is missing for type `Sle_Impacted_Users`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
