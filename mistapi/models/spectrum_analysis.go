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

// SpectrumAnalysis represents a SpectrumAnalysis struct.
type SpectrumAnalysis struct {
    // Band for spectrum analysis. enum: `24`, `5`, `6`
    Band                 SpectrumAnalysisBandEnum    `json:"band"`
    // Device ID of the AP that is performing spectrum analysis
    DeviceId             *uuid.UUID                  `json:"device_id,omitempty"`
    // Duration of the spectrum analysis in seconds
    Duration             *int                        `json:"duration,omitempty"`
    // Format of the spectrum analysis data. enum: `json`, `stream`
    Format               *SpectrumAnalysisFormatEnum `json:"format,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SpectrumAnalysis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SpectrumAnalysis) String() string {
    return fmt.Sprintf(
    	"SpectrumAnalysis[Band=%v, DeviceId=%v, Duration=%v, Format=%v, AdditionalProperties=%v]",
    	s.Band, s.DeviceId, s.Duration, s.Format, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SpectrumAnalysis.
// It customizes the JSON marshaling process for SpectrumAnalysis objects.
func (s SpectrumAnalysis) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "band", "device_id", "duration", "format"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SpectrumAnalysis object to a map representation for JSON marshaling.
func (s SpectrumAnalysis) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["band"] = s.Band
    if s.DeviceId != nil {
        structMap["device_id"] = s.DeviceId
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Format != nil {
        structMap["format"] = s.Format
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SpectrumAnalysis.
// It customizes the JSON unmarshaling process for SpectrumAnalysis objects.
func (s *SpectrumAnalysis) UnmarshalJSON(input []byte) error {
    var temp tempSpectrumAnalysis
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "device_id", "duration", "format")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Band = *temp.Band
    s.DeviceId = temp.DeviceId
    s.Duration = temp.Duration
    s.Format = temp.Format
    return nil
}

// tempSpectrumAnalysis is a temporary struct used for validating the fields of SpectrumAnalysis.
type tempSpectrumAnalysis  struct {
    Band     *SpectrumAnalysisBandEnum   `json:"band"`
    DeviceId *uuid.UUID                  `json:"device_id,omitempty"`
    Duration *int                        `json:"duration,omitempty"`
    Format   *SpectrumAnalysisFormatEnum `json:"format,omitempty"`
}

func (s *tempSpectrumAnalysis) validate() error {
    var errs []string
    if s.Band == nil {
        errs = append(errs, "required field `band` is missing for type `spectrum_analysis`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
