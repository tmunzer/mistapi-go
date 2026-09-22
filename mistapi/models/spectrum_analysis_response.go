// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// SpectrumAnalysisResponse represents a SpectrumAnalysisResponse struct.
// Response returned after starting AP spectrum analysis
type SpectrumAnalysisResponse struct {
	// Band for spectrum analysis. enum: `24`, `5`, `6`
	Band     *SpectrumAnalysisBandEnum `json:"band,omitempty"`
	Channels []int                     `json:"channels,omitempty"`
	// AP device UUID used for a single-AP spectrum analysis
	DeviceId *uuid.UUID `json:"device_id,omitempty"`
	// AP device UUIDs used for a multi-AP spectrum analysis; maximum 5 devices
	DeviceIds []uuid.UUID `json:"device_ids,omitempty"`
	Duration  *int        `json:"duration,omitempty"`
	// Format of the spectrum analysis data. enum: `json`, `stream`
	Format *SpectrumAnalysisFormatEnum `json:"format,omitempty"`
	// AP device IDs that failed validation, grouped by reason; available for multi-AP requests
	InvalidDeviceIds map[string][]uuid.UUID `json:"invalid_device_ids,omitempty"`
	// Spectrum analysis session identifier used to correlate WebSocket output
	SessionId uuid.UUID `json:"session_id"`
	// Epoch timestamp when spectrum analysis started
	StartedTime *int `json:"started_time,omitempty"`
	// Channel width used during spectrum analysis, in MHz
	Width                *int                   `json:"width,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SpectrumAnalysisResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SpectrumAnalysisResponse) String() string {
	return fmt.Sprintf(
		"SpectrumAnalysisResponse[Band=%v, Channels=%v, DeviceId=%v, DeviceIds=%v, Duration=%v, Format=%v, InvalidDeviceIds=%v, SessionId=%v, StartedTime=%v, Width=%v, AdditionalProperties=%v]",
		s.Band, s.Channels, s.DeviceId, s.DeviceIds, s.Duration, s.Format, s.InvalidDeviceIds, s.SessionId, s.StartedTime, s.Width, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SpectrumAnalysisResponse.
// It customizes the JSON marshaling process for SpectrumAnalysisResponse objects.
func (s SpectrumAnalysisResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"band", "channels", "device_id", "device_ids", "duration", "format", "invalid_device_ids", "session_id", "started_time", "width"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SpectrumAnalysisResponse object to a map representation for JSON marshaling.
func (s SpectrumAnalysisResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Band != nil {
		structMap["band"] = s.Band
	}
	if s.Channels != nil {
		structMap["channels"] = s.Channels
	}
	if s.DeviceId != nil {
		structMap["device_id"] = s.DeviceId
	}
	if s.DeviceIds != nil {
		structMap["device_ids"] = s.DeviceIds
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.Format != nil {
		structMap["format"] = s.Format
	}
	if s.InvalidDeviceIds != nil {
		structMap["invalid_device_ids"] = s.InvalidDeviceIds
	}
	structMap["session_id"] = s.SessionId
	if s.StartedTime != nil {
		structMap["started_time"] = s.StartedTime
	}
	if s.Width != nil {
		structMap["width"] = s.Width
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SpectrumAnalysisResponse.
// It customizes the JSON unmarshaling process for SpectrumAnalysisResponse objects.
func (s *SpectrumAnalysisResponse) UnmarshalJSON(input []byte) error {
	var temp tempSpectrumAnalysisResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channels", "device_id", "device_ids", "duration", "format", "invalid_device_ids", "session_id", "started_time", "width")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Band = temp.Band
	s.Channels = temp.Channels
	s.DeviceId = temp.DeviceId
	s.DeviceIds = temp.DeviceIds
	s.Duration = temp.Duration
	s.Format = temp.Format
	s.InvalidDeviceIds = temp.InvalidDeviceIds
	s.SessionId = *temp.SessionId
	s.StartedTime = temp.StartedTime
	s.Width = temp.Width
	return nil
}

// tempSpectrumAnalysisResponse is a temporary struct used for validating the fields of SpectrumAnalysisResponse.
type tempSpectrumAnalysisResponse struct {
	Band             *SpectrumAnalysisBandEnum   `json:"band,omitempty"`
	Channels         []int                       `json:"channels,omitempty"`
	DeviceId         *uuid.UUID                  `json:"device_id,omitempty"`
	DeviceIds        []uuid.UUID                 `json:"device_ids,omitempty"`
	Duration         *int                        `json:"duration,omitempty"`
	Format           *SpectrumAnalysisFormatEnum `json:"format,omitempty"`
	InvalidDeviceIds map[string][]uuid.UUID      `json:"invalid_device_ids,omitempty"`
	SessionId        *uuid.UUID                  `json:"session_id"`
	StartedTime      *int                        `json:"started_time,omitempty"`
	Width            *int                        `json:"width,omitempty"`
}

func (s *tempSpectrumAnalysisResponse) validate() error {
	var errs []string
	if s.SessionId == nil {
		errs = append(errs, "required field `session_id` is missing for type `spectrum_analysis_response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
