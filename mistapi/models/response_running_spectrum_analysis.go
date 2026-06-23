// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponseRunningSpectrumAnalysis represents a ResponseRunningSpectrumAnalysis struct.
// Running spectrum analysis session for a site
type ResponseRunningSpectrumAnalysis struct {
	// Radio band currently being scanned by spectrum analysis, such as 24, 5, or 6
	Band *string `json:"band,omitempty"`
	// List of channels being scanned in the spectrum analysis
	Channels []int `json:"channels,omitempty"`
	// Device ID of the AP that is running spectrum analysis
	DeviceId *uuid.UUID `json:"device_id,omitempty"`
	// Length of the running spectrum analysis session, in seconds
	Duration *int `json:"duration,omitempty"`
	// Output format for the running spectrum analysis data, such as json or stream
	Format *string `json:"format,omitempty"`
	// Timestamp when the spectrum analysis was started
	StartedTime *int `json:"started_time,omitempty"`
	// Channel width used during the spectrum analysis scan, in MHz
	Width                *int                   `json:"width,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRunningSpectrumAnalysis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRunningSpectrumAnalysis) String() string {
	return fmt.Sprintf(
		"ResponseRunningSpectrumAnalysis[Band=%v, Channels=%v, DeviceId=%v, Duration=%v, Format=%v, StartedTime=%v, Width=%v, AdditionalProperties=%v]",
		r.Band, r.Channels, r.DeviceId, r.Duration, r.Format, r.StartedTime, r.Width, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRunningSpectrumAnalysis.
// It customizes the JSON marshaling process for ResponseRunningSpectrumAnalysis objects.
func (r ResponseRunningSpectrumAnalysis) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"band", "channels", "device_id", "duration", "format", "started_time", "width"); err != nil {
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
	if r.Channels != nil {
		structMap["channels"] = r.Channels
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
	if r.Width != nil {
		structMap["width"] = r.Width
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channels", "device_id", "duration", "format", "started_time", "width")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Band = temp.Band
	r.Channels = temp.Channels
	r.DeviceId = temp.DeviceId
	r.Duration = temp.Duration
	r.Format = temp.Format
	r.StartedTime = temp.StartedTime
	r.Width = temp.Width
	return nil
}

// tempResponseRunningSpectrumAnalysis is a temporary struct used for validating the fields of ResponseRunningSpectrumAnalysis.
type tempResponseRunningSpectrumAnalysis struct {
	Band        *string    `json:"band,omitempty"`
	Channels    []int      `json:"channels,omitempty"`
	DeviceId    *uuid.UUID `json:"device_id,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
	Format      *string    `json:"format,omitempty"`
	StartedTime *int       `json:"started_time,omitempty"`
	Width       *int       `json:"width,omitempty"`
}
