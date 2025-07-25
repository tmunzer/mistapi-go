// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseAutoOrientationInfo represents a ResponseAutoOrientationInfo struct.
type ResponseAutoOrientationInfo struct {
    // Only when `status`==`inprogress`, estimate of the time to completion
    EstTimeLeft          *float64                               `json:"est_time_left,omitempty"`
    // time when auto orient process was last queued for this map
    StartTime            *float64                               `json:"start_time,omitempty"`
    // The status of auto orient for a given map. enum:
    // * `pending`: Auto orient has not been requested for this map
    // * `inprogress`: Auto orient is currently processing
    // * `done`: The auto orient process has completed
    // * `error`: There was an error in the auto orient process
    Status               *ResponseAutoOrientationInfoStatusEnum `json:"status,omitempty"`
    // time when auto orient completed or was manually stopped
    StopTime             *float64                               `json:"stop_time,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoOrientationInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoOrientationInfo) String() string {
    return fmt.Sprintf(
    	"ResponseAutoOrientationInfo[EstTimeLeft=%v, StartTime=%v, Status=%v, StopTime=%v, AdditionalProperties=%v]",
    	r.EstTimeLeft, r.StartTime, r.Status, r.StopTime, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoOrientationInfo.
// It customizes the JSON marshaling process for ResponseAutoOrientationInfo objects.
func (r ResponseAutoOrientationInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "est_time_left", "start_time", "status", "stop_time"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoOrientationInfo object to a map representation for JSON marshaling.
func (r ResponseAutoOrientationInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.EstTimeLeft != nil {
        structMap["est_time_left"] = r.EstTimeLeft
    }
    if r.StartTime != nil {
        structMap["start_time"] = r.StartTime
    }
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.StopTime != nil {
        structMap["stop_time"] = r.StopTime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoOrientationInfo.
// It customizes the JSON unmarshaling process for ResponseAutoOrientationInfo objects.
func (r *ResponseAutoOrientationInfo) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoOrientationInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "est_time_left", "start_time", "status", "stop_time")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.EstTimeLeft = temp.EstTimeLeft
    r.StartTime = temp.StartTime
    r.Status = temp.Status
    r.StopTime = temp.StopTime
    return nil
}

// tempResponseAutoOrientationInfo is a temporary struct used for validating the fields of ResponseAutoOrientationInfo.
type tempResponseAutoOrientationInfo  struct {
    EstTimeLeft *float64                               `json:"est_time_left,omitempty"`
    StartTime   *float64                               `json:"start_time,omitempty"`
    Status      *ResponseAutoOrientationInfoStatusEnum `json:"status,omitempty"`
    StopTime    *float64                               `json:"stop_time,omitempty"`
}
