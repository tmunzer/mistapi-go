// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
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
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedUsers,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedUsers) String() string {
	return fmt.Sprintf(
		"SleImpactedUsers[Classifier=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, Users=%v, AdditionalProperties=%v]",
		s.Classifier, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.Users, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedUsers.
// It customizes the JSON marshaling process for SleImpactedUsers objects.
func (s SleImpactedUsers) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"classifier", "end", "failure", "limit", "metric", "page", "start", "total_count", "users"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedUsers object to a map representation for JSON marshaling.
func (s SleImpactedUsers) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
	var temp tempSleImpactedUsers
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count", "users")
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

// tempSleImpactedUsers is a temporary struct used for validating the fields of SleImpactedUsers.
type tempSleImpactedUsers struct {
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

func (s *tempSleImpactedUsers) validate() error {
	var errs []string
	if s.Classifier == nil {
		errs = append(errs, "required field `classifier` is missing for type `sle_impacted_users`")
	}
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `sle_impacted_users`")
	}
	if s.Failure == nil {
		errs = append(errs, "required field `failure` is missing for type `sle_impacted_users`")
	}
	if s.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `sle_impacted_users`")
	}
	if s.Metric == nil {
		errs = append(errs, "required field `metric` is missing for type `sle_impacted_users`")
	}
	if s.Page == nil {
		errs = append(errs, "required field `page` is missing for type `sle_impacted_users`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `sle_impacted_users`")
	}
	if s.TotalCount == nil {
		errs = append(errs, "required field `total_count` is missing for type `sle_impacted_users`")
	}
	if s.Users == nil {
		errs = append(errs, "required field `users` is missing for type `sle_impacted_users`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
