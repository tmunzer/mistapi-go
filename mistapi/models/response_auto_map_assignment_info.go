// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseAutoMapAssignmentInfo represents a ResponseAutoMapAssignmentInfo struct.
// Auto map assignment status response
type ResponseAutoMapAssignmentInfo struct {
	// Only when `status`==`in_progress`, estimated seconds remaining
	EstTimeLeft *float64 `json:"est_time_left,omitempty"`
	// Unix timestamp when auto map assignment was started
	StartTime *float64 `json:"start_time,omitempty"`
	// The status of auto map assignment for a given site. enum:
	// * `not_started`: Auto map assignment has not been requested
	// * `in_progress`: Auto map assignment is currently processing
	// * `completed`: The auto map assignment process has completed
	// * `error`: There was an error in the auto map assignment process
	Status ResponseAutoMapAssignmentInfoStatusEnum `json:"status"`
	// Only when `status`==`completed`, Unix timestamp when auto map assignment stopped
	StopTime *float64 `json:"stop_time,omitempty"`
	// Unix timestamp when status was last updated
	TimeUpdated          *float64               `json:"time_updated,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoMapAssignmentInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoMapAssignmentInfo) String() string {
	return fmt.Sprintf(
		"ResponseAutoMapAssignmentInfo[EstTimeLeft=%v, StartTime=%v, Status=%v, StopTime=%v, TimeUpdated=%v, AdditionalProperties=%v]",
		r.EstTimeLeft, r.StartTime, r.Status, r.StopTime, r.TimeUpdated, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoMapAssignmentInfo.
// It customizes the JSON marshaling process for ResponseAutoMapAssignmentInfo objects.
func (r ResponseAutoMapAssignmentInfo) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"est_time_left", "start_time", "status", "stop_time", "time_updated"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoMapAssignmentInfo object to a map representation for JSON marshaling.
func (r ResponseAutoMapAssignmentInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.EstTimeLeft != nil {
		structMap["est_time_left"] = r.EstTimeLeft
	}
	if r.StartTime != nil {
		structMap["start_time"] = r.StartTime
	}
	structMap["status"] = r.Status
	if r.StopTime != nil {
		structMap["stop_time"] = r.StopTime
	}
	if r.TimeUpdated != nil {
		structMap["time_updated"] = r.TimeUpdated
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoMapAssignmentInfo.
// It customizes the JSON unmarshaling process for ResponseAutoMapAssignmentInfo objects.
func (r *ResponseAutoMapAssignmentInfo) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoMapAssignmentInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "est_time_left", "start_time", "status", "stop_time", "time_updated")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.EstTimeLeft = temp.EstTimeLeft
	r.StartTime = temp.StartTime
	r.Status = *temp.Status
	r.StopTime = temp.StopTime
	r.TimeUpdated = temp.TimeUpdated
	return nil
}

// tempResponseAutoMapAssignmentInfo is a temporary struct used for validating the fields of ResponseAutoMapAssignmentInfo.
type tempResponseAutoMapAssignmentInfo struct {
	EstTimeLeft *float64                                 `json:"est_time_left,omitempty"`
	StartTime   *float64                                 `json:"start_time,omitempty"`
	Status      *ResponseAutoMapAssignmentInfoStatusEnum `json:"status"`
	StopTime    *float64                                 `json:"stop_time,omitempty"`
	TimeUpdated *float64                                 `json:"time_updated,omitempty"`
}

func (r *tempResponseAutoMapAssignmentInfo) validate() error {
	var errs []string
	if r.Status == nil {
		errs = append(errs, "required field `status` is missing for type `response_auto_map_assignment_info`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
