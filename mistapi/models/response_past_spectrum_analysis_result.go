// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponsePastSpectrumAnalysisResult represents a ResponsePastSpectrumAnalysisResult struct.
// Result of a past spectrum analysis
type ResponsePastSpectrumAnalysisResult struct {
	// Band on which the spectrum analysis was run (e.g., 24, 5, 6)
	Band         *string                                    `json:"band,omitempty"`
	ChannelUsage []ResponsePastSpectrumAnalysisChannelUsage `json:"channel_usage,omitempty"`
	// List of FFT samples for the spectrum analysis
	FftSamples []ResponsePastSpectrumAnalysisFftSample `json:"fft_samples,omitempty"`
	// MAC Address of the AP that ran the spectrum analysis
	Mac   *string    `json:"mac,omitempty"`
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Timestamp when the spectrum analysis was run in epoch seconds
	Timestamp            *int                   `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePastSpectrumAnalysisResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePastSpectrumAnalysisResult) String() string {
	return fmt.Sprintf(
		"ResponsePastSpectrumAnalysisResult[Band=%v, ChannelUsage=%v, FftSamples=%v, Mac=%v, OrgId=%v, Timestamp=%v, AdditionalProperties=%v]",
		r.Band, r.ChannelUsage, r.FftSamples, r.Mac, r.OrgId, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePastSpectrumAnalysisResult.
// It customizes the JSON marshaling process for ResponsePastSpectrumAnalysisResult objects.
func (r ResponsePastSpectrumAnalysisResult) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"band", "channel_usage", "fft_samples", "mac", "org_id", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePastSpectrumAnalysisResult object to a map representation for JSON marshaling.
func (r ResponsePastSpectrumAnalysisResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Band != nil {
		structMap["band"] = r.Band
	}
	if r.ChannelUsage != nil {
		structMap["channel_usage"] = r.ChannelUsage
	}
	if r.FftSamples != nil {
		structMap["fft_samples"] = r.FftSamples
	}
	if r.Mac != nil {
		structMap["mac"] = r.Mac
	}
	if r.OrgId != nil {
		structMap["org_id"] = r.OrgId
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePastSpectrumAnalysisResult.
// It customizes the JSON unmarshaling process for ResponsePastSpectrumAnalysisResult objects.
func (r *ResponsePastSpectrumAnalysisResult) UnmarshalJSON(input []byte) error {
	var temp tempResponsePastSpectrumAnalysisResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channel_usage", "fft_samples", "mac", "org_id", "timestamp")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Band = temp.Band
	r.ChannelUsage = temp.ChannelUsage
	r.FftSamples = temp.FftSamples
	r.Mac = temp.Mac
	r.OrgId = temp.OrgId
	r.Timestamp = temp.Timestamp
	return nil
}

// tempResponsePastSpectrumAnalysisResult is a temporary struct used for validating the fields of ResponsePastSpectrumAnalysisResult.
type tempResponsePastSpectrumAnalysisResult struct {
	Band         *string                                    `json:"band,omitempty"`
	ChannelUsage []ResponsePastSpectrumAnalysisChannelUsage `json:"channel_usage,omitempty"`
	FftSamples   []ResponsePastSpectrumAnalysisFftSample    `json:"fft_samples,omitempty"`
	Mac          *string                                    `json:"mac,omitempty"`
	OrgId        *uuid.UUID                                 `json:"org_id,omitempty"`
	Timestamp    *int                                       `json:"timestamp,omitempty"`
}
