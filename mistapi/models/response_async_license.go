// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAsyncLicense represents a ResponseAsyncLicense struct.
type ResponseAsyncLicense struct {
	Completed []string `json:"completed,omitempty"`
	// detail claim status per device
	Detail *ResponseAsyncLicenseDetail `json:"detail,omitempty"`
	// Current failed number of device
	Failed *int `json:"failed,omitempty"`
	// Current incompleted lists (macs)
	Incompleted []string `json:"incompleted,omitempty"`
	// Current processed number of device
	Processed *int `json:"processed,omitempty"`
	// epoch time of aysnc claim scheduled
	ScheduledAt *int `json:"scheduled_at,omitempty"`
	// processing status of async. enum: `prepared`, `ongoing`, `done`
	Status *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	// Current succeed number of device
	Succeed *int `json:"succeed,omitempty"`
	// Epoch (seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// total number of device included in claim
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncLicense,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncLicense) String() string {
	return fmt.Sprintf(
		"ResponseAsyncLicense[Completed=%v, Detail=%v, Failed=%v, Incompleted=%v, Processed=%v, ScheduledAt=%v, Status=%v, Succeed=%v, Timestamp=%v, Total=%v, AdditionalProperties=%v]",
		r.Completed, r.Detail, r.Failed, r.Incompleted, r.Processed, r.ScheduledAt, r.Status, r.Succeed, r.Timestamp, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncLicense.
// It customizes the JSON marshaling process for ResponseAsyncLicense objects.
func (r ResponseAsyncLicense) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"completed", "detail", "failed", "incompleted", "processed", "scheduled_at", "status", "succeed", "timestamp", "total"); err != nil {
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
	if r.Detail != nil {
		structMap["detail"] = r.Detail.toMap()
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "completed", "detail", "failed", "incompleted", "processed", "scheduled_at", "status", "succeed", "timestamp", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Completed = temp.Completed
	r.Detail = temp.Detail
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
	Detail      *ResponseAsyncLicenseDetail     `json:"detail,omitempty"`
	Failed      *int                            `json:"failed,omitempty"`
	Incompleted []string                        `json:"incompleted,omitempty"`
	Processed   *int                            `json:"processed,omitempty"`
	ScheduledAt *int                            `json:"scheduled_at,omitempty"`
	Status      *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	Succeed     *int                            `json:"succeed,omitempty"`
	Timestamp   *float64                        `json:"timestamp,omitempty"`
	Total       *int                            `json:"total,omitempty"`
}
