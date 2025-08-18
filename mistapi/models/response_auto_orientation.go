// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAutoOrientation represents a ResponseAutoOrientation struct.
type ResponseAutoOrientation struct {
	// Contains the validation status of each device. The Property Key is the device MAC Address.
	Devices map[string]ResponseAutoOrientationDevice `json:"devices,omitempty"`
	// Estimated runtime for the process in seconds
	EstimatedRuntime *int `json:"estimated_runtime,omitempty"`
	// Provides the reason for the status.
	Reason *string `json:"reason,omitempty"`
	// Indicates whether the auto orient process has started.
	Started *bool `json:"started,omitempty"`
	// Indicates whether the auto orient request is valid.
	Valid *bool `json:"valid,omitempty"`
	// Indicates whether the auto orient process will interrupt WiFi traffic.
	WifiInterrupting     *bool                  `json:"wifi_interrupting,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoOrientation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoOrientation) String() string {
	return fmt.Sprintf(
		"ResponseAutoOrientation[Devices=%v, EstimatedRuntime=%v, Reason=%v, Started=%v, Valid=%v, WifiInterrupting=%v, AdditionalProperties=%v]",
		r.Devices, r.EstimatedRuntime, r.Reason, r.Started, r.Valid, r.WifiInterrupting, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoOrientation.
// It customizes the JSON marshaling process for ResponseAutoOrientation objects.
func (r ResponseAutoOrientation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"devices", "estimated_runtime", "reason", "started", "valid", "wifi_interrupting"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoOrientation object to a map representation for JSON marshaling.
func (r ResponseAutoOrientation) toMap() map[string]any {
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
	if r.WifiInterrupting != nil {
		structMap["wifi_interrupting"] = r.WifiInterrupting
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoOrientation.
// It customizes the JSON unmarshaling process for ResponseAutoOrientation objects.
func (r *ResponseAutoOrientation) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoOrientation
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "devices", "estimated_runtime", "reason", "started", "valid", "wifi_interrupting")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Devices = temp.Devices
	r.EstimatedRuntime = temp.EstimatedRuntime
	r.Reason = temp.Reason
	r.Started = temp.Started
	r.Valid = temp.Valid
	r.WifiInterrupting = temp.WifiInterrupting
	return nil
}

// tempResponseAutoOrientation is a temporary struct used for validating the fields of ResponseAutoOrientation.
type tempResponseAutoOrientation struct {
	Devices          map[string]ResponseAutoOrientationDevice `json:"devices,omitempty"`
	EstimatedRuntime *int                                     `json:"estimated_runtime,omitempty"`
	Reason           *string                                  `json:"reason,omitempty"`
	Started          *bool                                    `json:"started,omitempty"`
	Valid            *bool                                    `json:"valid,omitempty"`
	WifiInterrupting *bool                                    `json:"wifi_interrupting,omitempty"`
}
