// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseRunningSpectrumAnalysis represents a ResponseRunningSpectrumAnalysis struct.
type ResponseRunningSpectrumAnalysis struct {
    // Band on which the spectrum analysis is running (e.g., 24, 5, 6)
    Band                 *string                `json:"band,omitempty"`
    // Device ID of the AP that is running spectrum analysis
    DeviceId             *uuid.UUID             `json:"device_id,omitempty"`
    // Duration of the spectrum analysis in seconds
    Duration             *int                   `json:"duration,omitempty"`
    // Format of the spectrum analysis data (e.g., json, stream)
    Format               *string                `json:"format,omitempty"`
    // Time when the spectrum analysis was started
    StartedTime          *int                   `json:"started_time,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRunningSpectrumAnalysis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRunningSpectrumAnalysis) String() string {
    return fmt.Sprintf(
    	"ResponseRunningSpectrumAnalysis[Band=%v, DeviceId=%v, Duration=%v, Format=%v, StartedTime=%v, AdditionalProperties=%v]",
    	r.Band, r.DeviceId, r.Duration, r.Format, r.StartedTime, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRunningSpectrumAnalysis.
// It customizes the JSON marshaling process for ResponseRunningSpectrumAnalysis objects.
func (r ResponseRunningSpectrumAnalysis) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "band", "device_id", "duration", "format", "started_time"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseRunningSpectrumAnalysis object to a map representation for JSON marshaling.
func (r ResponseRunningSpectrumAnalysis) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Band != nil {
        structMap["band"] = r.Band
    }
    if r.DeviceId != nil {
        structMap["device_id"] = r.DeviceId
    }
    if r.Duration != nil {
        structMap["duration"] = r.Duration
    }
    if r.Format != nil {
        structMap["format"] = r.Format
    }
    if r.StartedTime != nil {
        structMap["started_time"] = r.StartedTime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRunningSpectrumAnalysis.
// It customizes the JSON unmarshaling process for ResponseRunningSpectrumAnalysis objects.
func (r *ResponseRunningSpectrumAnalysis) UnmarshalJSON(input []byte) error {
    var temp tempResponseRunningSpectrumAnalysis
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "device_id", "duration", "format", "started_time")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Band = temp.Band
    r.DeviceId = temp.DeviceId
    r.Duration = temp.Duration
    r.Format = temp.Format
    r.StartedTime = temp.StartedTime
    return nil
}

// tempResponseRunningSpectrumAnalysis is a temporary struct used for validating the fields of ResponseRunningSpectrumAnalysis.
type tempResponseRunningSpectrumAnalysis  struct {
    Band        *string    `json:"band,omitempty"`
    DeviceId    *uuid.UUID `json:"device_id,omitempty"`
    Duration    *int       `json:"duration,omitempty"`
    Format      *string    `json:"format,omitempty"`
    StartedTime *int       `json:"started_time,omitempty"`
}
