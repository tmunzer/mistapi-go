// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAsyncLicense represents a ResponseAsyncLicense struct.
// Asynchronous license claim progress response
type ResponseAsyncLicense struct {
	// Device MAC addresses that completed asynchronous license claim processing
	Completed []string `json:"completed,omitempty"`
	// Per-device asynchronous license claim status details
	Details []ResponseAsyncLicenseDetail `json:"details,omitempty"`
	// Number of devices that failed license claim processing
	Failed *int `json:"failed,omitempty"`
	// Device MAC addresses not yet completed in asynchronous license claim processing
	Incompleted []string `json:"incompleted,omitempty"`
	// Number of devices processed so far by asynchronous license claim
	Processed *int `json:"processed,omitempty"`
	// Epoch timestamp when the asynchronous license claim was scheduled
	ScheduledAt *int `json:"scheduled_at,omitempty"`
	// Processing state for an asynchronous license claim. enum: `prepared`, `ongoing`, `done`
	Status *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	// Number of devices that successfully completed license claim processing
	Succeed *int `json:"succeed,omitempty"`
	// Epoch timestamp, in seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Number of devices included in the license claim
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncLicense,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncLicense) String() string {
	return fmt.Sprintf(
		"ResponseAsyncLicense[Completed=%v, Details=%v, Failed=%v, Incompleted=%v, Processed=%v, ScheduledAt=%v, Status=%v, Succeed=%v, Timestamp=%v, Total=%v, AdditionalProperties=%v]",
		r.Completed, r.Details, r.Failed, r.Incompleted, r.Processed, r.ScheduledAt, r.Status, r.Succeed, r.Timestamp, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncLicense.
// It customizes the JSON marshaling process for ResponseAsyncLicense objects.
func (r ResponseAsyncLicense) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"completed", "details", "failed", "incompleted", "processed", "scheduled_at", "status", "succeed", "timestamp", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncLicense object to a map representation for JSON marshaling.
func (r ResponseAsyncLicense) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Completed != nil {
		structMap["completed"] = r.Completed
	}
	if r.Details != nil {
		structMap["details"] = r.Details
	}
	if r.Failed != nil {
		structMap["failed"] = r.Failed
	}
	if r.Incompleted != nil {
		structMap["incompleted"] = r.Incompleted
	}
	if r.Processed != nil {
		structMap["processed"] = r.Processed
	}
	if r.ScheduledAt != nil {
		structMap["scheduled_at"] = r.ScheduledAt
	}
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.Succeed != nil {
		structMap["succeed"] = r.Succeed
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncLicense.
// It customizes the JSON unmarshaling process for ResponseAsyncLicense objects.
func (r *ResponseAsyncLicense) UnmarshalJSON(input []byte) error {
	var temp tempResponseAsyncLicense
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "completed", "details", "failed", "incompleted", "processed", "scheduled_at", "status", "succeed", "timestamp", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Completed = temp.Completed
	r.Details = temp.Details
	r.Failed = temp.Failed
	r.Incompleted = temp.Incompleted
	r.Processed = temp.Processed
	r.ScheduledAt = temp.ScheduledAt
	r.Status = temp.Status
	r.Succeed = temp.Succeed
	r.Timestamp = temp.Timestamp
	r.Total = temp.Total
	return nil
}

// tempResponseAsyncLicense is a temporary struct used for validating the fields of ResponseAsyncLicense.
type tempResponseAsyncLicense struct {
	Completed   []string                        `json:"completed,omitempty"`
	Details     []ResponseAsyncLicenseDetail    `json:"details,omitempty"`
	Failed      *int                            `json:"failed,omitempty"`
	Incompleted []string                        `json:"incompleted,omitempty"`
	Processed   *int                            `json:"processed,omitempty"`
	ScheduledAt *int                            `json:"scheduled_at,omitempty"`
	Status      *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	Succeed     *int                            `json:"succeed,omitempty"`
	Timestamp   *float64                        `json:"timestamp,omitempty"`
	Total       *int                            `json:"total,omitempty"`
}
