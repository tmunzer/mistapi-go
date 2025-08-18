// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponsePastSpectrumAnalysisChannelUsage represents a ResponsePastSpectrumAnalysisChannelUsage struct.
// Channel usage data for a specific channel
type ResponsePastSpectrumAnalysisChannelUsage struct {
    // Channel number
    Channel              *int                   `json:"channel,omitempty"`
    // Noise level in dBm
    Noise                *float64               `json:"noise,omitempty"`
    // Percentage of channel usage by non-WiFi signals in the range [0, 1]
    NonWifi              *float64               `json:"non_wifi,omitempty"`
    // Percentage of channel usage by WiFi in the range [0, 1]
    Wifi                 *float64               `json:"wifi,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePastSpectrumAnalysisChannelUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePastSpectrumAnalysisChannelUsage) String() string {
    return fmt.Sprintf(
    	"ResponsePastSpectrumAnalysisChannelUsage[Channel=%v, Noise=%v, NonWifi=%v, Wifi=%v, AdditionalProperties=%v]",
    	r.Channel, r.Noise, r.NonWifi, r.Wifi, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePastSpectrumAnalysisChannelUsage.
// It customizes the JSON marshaling process for ResponsePastSpectrumAnalysisChannelUsage objects.
func (r ResponsePastSpectrumAnalysisChannelUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "channel", "noise", "non_wifi", "wifi"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePastSpectrumAnalysisChannelUsage object to a map representation for JSON marshaling.
func (r ResponsePastSpectrumAnalysisChannelUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Channel != nil {
        structMap["channel"] = r.Channel
    }
    if r.Noise != nil {
        structMap["noise"] = r.Noise
    }
    if r.NonWifi != nil {
        structMap["non_wifi"] = r.NonWifi
    }
    if r.Wifi != nil {
        structMap["wifi"] = r.Wifi
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePastSpectrumAnalysisChannelUsage.
// It customizes the JSON unmarshaling process for ResponsePastSpectrumAnalysisChannelUsage objects.
func (r *ResponsePastSpectrumAnalysisChannelUsage) UnmarshalJSON(input []byte) error {
    var temp tempResponsePastSpectrumAnalysisChannelUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "noise", "non_wifi", "wifi")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Channel = temp.Channel
    r.Noise = temp.Noise
    r.NonWifi = temp.NonWifi
    r.Wifi = temp.Wifi
    return nil
}

// tempResponsePastSpectrumAnalysisChannelUsage is a temporary struct used for validating the fields of ResponsePastSpectrumAnalysisChannelUsage.
type tempResponsePastSpectrumAnalysisChannelUsage  struct {
    Channel *int     `json:"channel,omitempty"`
    Noise   *float64 `json:"noise,omitempty"`
    NonWifi *float64 `json:"non_wifi,omitempty"`
    Wifi    *float64 `json:"wifi,omitempty"`
}
