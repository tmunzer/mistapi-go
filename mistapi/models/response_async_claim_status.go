// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponseAsyncClaimStatus represents a ResponseAsyncClaimStatus struct.
// Async inventory claim job status
type ResponseAsyncClaimStatus struct {
	// Unique identifier of the async claim job
	ClaimId *uuid.UUID `json:"claim_id,omitempty"`
	// Device MAC addresses that completed asynchronous license claim processing
	Completed []string `json:"completed,omitempty"`
	// Per-device asynchronous license claim status details
	Details []ResponseAsyncLicenseDetail `json:"details,omitempty"`
	// Number of devices that failed claim processing
	Failed *int `json:"failed,omitempty"`
	// Device MAC addresses not yet completed in asynchronous license claim processing
	Incompleted []string `json:"incompleted,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Number of devices processed so far
	Processed *int `json:"processed,omitempty"`
	// Epoch timestamp when the async claim was scheduled
	ScheduledAt *int `json:"scheduled_at,omitempty"`
	// Processing state for an asynchronous license claim. enum: `prepared`, `ongoing`, `done`
	Status *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	// Number of devices that successfully completed claim processing
	Succeed *int `json:"succeed,omitempty"`
	// Epoch timestamp, in seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Total number of devices included in the claim
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncClaimStatus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncClaimStatus) String() string {
	return fmt.Sprintf(
		"ResponseAsyncClaimStatus[ClaimId=%v, Completed=%v, Details=%v, Failed=%v, Incompleted=%v, OrgId=%v, Processed=%v, ScheduledAt=%v, Status=%v, Succeed=%v, Timestamp=%v, Total=%v, AdditionalProperties=%v]",
		r.ClaimId, r.Completed, r.Details, r.Failed, r.Incompleted, r.OrgId, r.Processed, r.ScheduledAt, r.Status, r.Succeed, r.Timestamp, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncClaimStatus.
// It customizes the JSON marshaling process for ResponseAsyncClaimStatus objects.
func (r ResponseAsyncClaimStatus) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"claim_id", "completed", "details", "failed", "incompleted", "org_id", "processed", "scheduled_at", "status", "succeed", "timestamp", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncClaimStatus object to a map representation for JSON marshaling.
func (r ResponseAsyncClaimStatus) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ClaimId != nil {
		structMap["claim_id"] = r.ClaimId
	}
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
	if r.OrgId != nil {
		structMap["org_id"] = r.OrgId
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncClaimStatus.
// It customizes the JSON unmarshaling process for ResponseAsyncClaimStatus objects.
func (r *ResponseAsyncClaimStatus) UnmarshalJSON(input []byte) error {
	var temp tempResponseAsyncClaimStatus
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "claim_id", "completed", "details", "failed", "incompleted", "org_id", "processed", "scheduled_at", "status", "succeed", "timestamp", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.ClaimId = temp.ClaimId
	r.Completed = temp.Completed
	r.Details = temp.Details
	r.Failed = temp.Failed
	r.Incompleted = temp.Incompleted
	r.OrgId = temp.OrgId
	r.Processed = temp.Processed
	r.ScheduledAt = temp.ScheduledAt
	r.Status = temp.Status
	r.Succeed = temp.Succeed
	r.Timestamp = temp.Timestamp
	r.Total = temp.Total
	return nil
}

// tempResponseAsyncClaimStatus is a temporary struct used for validating the fields of ResponseAsyncClaimStatus.
type tempResponseAsyncClaimStatus struct {
	ClaimId     *uuid.UUID                      `json:"claim_id,omitempty"`
	Completed   []string                        `json:"completed,omitempty"`
	Details     []ResponseAsyncLicenseDetail    `json:"details,omitempty"`
	Failed      *int                            `json:"failed,omitempty"`
	Incompleted []string                        `json:"incompleted,omitempty"`
	OrgId       *uuid.UUID                      `json:"org_id,omitempty"`
	Processed   *int                            `json:"processed,omitempty"`
	ScheduledAt *int                            `json:"scheduled_at,omitempty"`
	Status      *ResponseAsyncLicenseStatusEnum `json:"status,omitempty"`
	Succeed     *int                            `json:"succeed,omitempty"`
	Timestamp   *float64                        `json:"timestamp,omitempty"`
	Total       *int                            `json:"total,omitempty"`
}
