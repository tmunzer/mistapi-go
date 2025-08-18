// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponsePastSpectrumAnalysisFftSample represents a ResponsePastSpectrumAnalysisFftSample struct.
// FFT sample data for a specific frequency
type ResponsePastSpectrumAnalysisFftSample struct {
    // Frequency in MHz
    Frequency            *float64               `json:"frequency,omitempty"`
    // RSSI in dBm
    Rssi                 *float64               `json:"rssi,omitempty"`
    // RSSI in dBm
    Signal7              *float64               `json:"signal7,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePastSpectrumAnalysisFftSample,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePastSpectrumAnalysisFftSample) String() string {
    return fmt.Sprintf(
    	"ResponsePastSpectrumAnalysisFftSample[Frequency=%v, Rssi=%v, Signal7=%v, AdditionalProperties=%v]",
    	r.Frequency, r.Rssi, r.Signal7, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePastSpectrumAnalysisFftSample.
// It customizes the JSON marshaling process for ResponsePastSpectrumAnalysisFftSample objects.
func (r ResponsePastSpectrumAnalysisFftSample) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "frequency", "rssi", "signal7"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePastSpectrumAnalysisFftSample object to a map representation for JSON marshaling.
func (r ResponsePastSpectrumAnalysisFftSample) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Frequency != nil {
        structMap["frequency"] = r.Frequency
    }
    if r.Rssi != nil {
        structMap["rssi"] = r.Rssi
    }
    if r.Signal7 != nil {
        structMap["signal7"] = r.Signal7
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePastSpectrumAnalysisFftSample.
// It customizes the JSON unmarshaling process for ResponsePastSpectrumAnalysisFftSample objects.
func (r *ResponsePastSpectrumAnalysisFftSample) UnmarshalJSON(input []byte) error {
    var temp tempResponsePastSpectrumAnalysisFftSample
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "frequency", "rssi", "signal7")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Frequency = temp.Frequency
    r.Rssi = temp.Rssi
    r.Signal7 = temp.Signal7
    return nil
}

// tempResponsePastSpectrumAnalysisFftSample is a temporary struct used for validating the fields of ResponsePastSpectrumAnalysisFftSample.
type tempResponsePastSpectrumAnalysisFftSample  struct {
    Frequency *float64 `json:"frequency,omitempty"`
    Rssi      *float64 `json:"rssi,omitempty"`
    Signal7   *float64 `json:"signal7,omitempty"`
}
