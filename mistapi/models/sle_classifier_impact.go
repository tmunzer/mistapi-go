package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleClassifierImpact represents a SleClassifierImpact struct.
type SleClassifierImpact struct {
    NumAps               float64                `json:"num_aps"`
    NumUsers             float64                `json:"num_users"`
    TotalAps             float64                `json:"total_aps"`
    TotalUsers           float64                `json:"total_users"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleClassifierImpact,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleClassifierImpact) String() string {
    return fmt.Sprintf(
    	"SleClassifierImpact[NumAps=%v, NumUsers=%v, TotalAps=%v, TotalUsers=%v, AdditionalProperties=%v]",
    	s.NumAps, s.NumUsers, s.TotalAps, s.TotalUsers, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierImpact.
// It customizes the JSON marshaling process for SleClassifierImpact objects.
func (s SleClassifierImpact) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "num_aps", "num_users", "total_aps", "total_users"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierImpact object to a map representation for JSON marshaling.
func (s SleClassifierImpact) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["num_aps"] = s.NumAps
    structMap["num_users"] = s.NumUsers
    structMap["total_aps"] = s.TotalAps
    structMap["total_users"] = s.TotalUsers
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierImpact.
// It customizes the JSON unmarshaling process for SleClassifierImpact objects.
func (s *SleClassifierImpact) UnmarshalJSON(input []byte) error {
    var temp tempSleClassifierImpact
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_aps", "num_users", "total_aps", "total_users")
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

// tempSleClassifierImpact is a temporary struct used for validating the fields of SleClassifierImpact.
type tempSleClassifierImpact  struct {
    NumAps     *float64 `json:"num_aps"`
    NumUsers   *float64 `json:"num_users"`
    TotalAps   *float64 `json:"total_aps"`
    TotalUsers *float64 `json:"total_users"`
}

func (s *tempSleClassifierImpact) validate() error {
    var errs []string
    if s.NumAps == nil {
        errs = append(errs, "required field `num_aps` is missing for type `sle_classifier_impact`")
    }
    if s.NumUsers == nil {
        errs = append(errs, "required field `num_users` is missing for type `sle_classifier_impact`")
    }
    if s.TotalAps == nil {
        errs = append(errs, "required field `total_aps` is missing for type `sle_classifier_impact`")
    }
    if s.TotalUsers == nil {
        errs = append(errs, "required field `total_users` is missing for type `sle_classifier_impact`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
