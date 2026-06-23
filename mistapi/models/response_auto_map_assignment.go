// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAutoMapAssignment represents a ResponseAutoMapAssignment struct.
// Auto map assignment start response
type ResponseAutoMapAssignment struct {
	// Contains the validation status of each device. The property key is the device MAC address.
	Devices map[string]ResponseAutoMapAssignmentDevice `json:"devices,omitempty"`
	// Estimated runtime for the process in seconds
	EstimatedRuntime *int `json:"estimated_runtime,omitempty"`
	// Provides the reason for the status
	Reason *string `json:"reason,omitempty"`
	// Indicates whether the auto map assignment process has started
	Started *bool `json:"started,omitempty"`
	// Indicates whether the auto map assignment request is valid
	Valid                *bool                  `json:"valid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoMapAssignment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoMapAssignment) String() string {
	return fmt.Sprintf(
		"ResponseAutoMapAssignment[Devices=%v, EstimatedRuntime=%v, Reason=%v, Started=%v, Valid=%v, AdditionalProperties=%v]",
		r.Devices, r.EstimatedRuntime, r.Reason, r.Started, r.Valid, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoMapAssignment.
// It customizes the JSON marshaling process for ResponseAutoMapAssignment objects.
func (r ResponseAutoMapAssignment) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"devices", "estimated_runtime", "reason", "started", "valid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoMapAssignment object to a map representation for JSON marshaling.
func (r ResponseAutoMapAssignment) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Devices != nil {
		structMap["devices"] = r.Devices
	}
	if r.EstimatedRuntime != nil {
		structMap["estimated_runtime"] = r.EstimatedRuntime
	}
	if r.Reason != nil {
		structMap["reason"] = r.Reason
	}
	if r.Started != nil {
		structMap["started"] = r.Started
	}
	if r.Valid != nil {
		structMap["valid"] = r.Valid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoMapAssignment.
// It customizes the JSON unmarshaling process for ResponseAutoMapAssignment objects.
func (r *ResponseAutoMapAssignment) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoMapAssignment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "devices", "estimated_runtime", "reason", "started", "valid")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Devices = temp.Devices
	r.EstimatedRuntime = temp.EstimatedRuntime
	r.Reason = temp.Reason
	r.Started = temp.Started
	r.Valid = temp.Valid
	return nil
}

// tempResponseAutoMapAssignment is a temporary struct used for validating the fields of ResponseAutoMapAssignment.
type tempResponseAutoMapAssignment struct {
	Devices          map[string]ResponseAutoMapAssignmentDevice `json:"devices,omitempty"`
	EstimatedRuntime *int                                       `json:"estimated_runtime,omitempty"`
	Reason           *string                                    `json:"reason,omitempty"`
	Started          *bool                                      `json:"started,omitempty"`
	Valid            *bool                                      `json:"valid,omitempty"`
}
