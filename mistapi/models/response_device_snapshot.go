// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseDeviceSnapshot represents a ResponseDeviceSnapshot struct.
type ResponseDeviceSnapshot struct {
    // enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
    Status               *ResponseDeviceSnapshotStatusEnum `json:"status,omitempty"`
    // Internal status id
    StatusId             *string                           `json:"status_id,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64                          `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceSnapshot,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceSnapshot) String() string {
    return fmt.Sprintf(
    	"ResponseDeviceSnapshot[Status=%v, StatusId=%v, Timestamp=%v, AdditionalProperties=%v]",
    	r.Status, r.StatusId, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSnapshot.
// It customizes the JSON marshaling process for ResponseDeviceSnapshot objects.
func (r ResponseDeviceSnapshot) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "status", "status_id", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSnapshot object to a map representation for JSON marshaling.
func (r ResponseDeviceSnapshot) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.StatusId != nil {
        structMap["status_id"] = r.StatusId
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSnapshot.
// It customizes the JSON unmarshaling process for ResponseDeviceSnapshot objects.
func (r *ResponseDeviceSnapshot) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceSnapshot
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "status", "status_id", "timestamp")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Status = temp.Status
    r.StatusId = temp.StatusId
    r.Timestamp = temp.Timestamp
    return nil
}

// tempResponseDeviceSnapshot is a temporary struct used for validating the fields of ResponseDeviceSnapshot.
type tempResponseDeviceSnapshot  struct {
    Status    *ResponseDeviceSnapshotStatusEnum `json:"status,omitempty"`
    StatusId  *string                           `json:"status_id,omitempty"`
    Timestamp *float64                          `json:"timestamp,omitempty"`
}
