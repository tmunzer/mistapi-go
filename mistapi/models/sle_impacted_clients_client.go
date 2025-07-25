// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedClientsClient represents a SleImpactedClientsClient struct.
type SleImpactedClientsClient struct {
    Degraded             *int                             `json:"degraded,omitempty"`
    Duration             *int                             `json:"duration,omitempty"`
    Mac                  *string                          `json:"mac,omitempty"`
    Name                 *string                          `json:"name,omitempty"`
    Switches             []SleImpactedClientsClientSwitch `json:"switches,omitempty"`
    Total                *int                             `json:"total,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedClientsClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedClientsClient) String() string {
    return fmt.Sprintf(
    	"SleImpactedClientsClient[Degraded=%v, Duration=%v, Mac=%v, Name=%v, Switches=%v, Total=%v, AdditionalProperties=%v]",
    	s.Degraded, s.Duration, s.Mac, s.Name, s.Switches, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedClientsClient.
// It customizes the JSON marshaling process for SleImpactedClientsClient objects.
func (s SleImpactedClientsClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "duration", "mac", "name", "switches", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedClientsClient object to a map representation for JSON marshaling.
func (s SleImpactedClientsClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Degraded != nil {
        structMap["degraded"] = s.Degraded
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Switches != nil {
        structMap["switches"] = s.Switches
    }
    if s.Total != nil {
        structMap["total"] = s.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedClientsClient.
// It customizes the JSON unmarshaling process for SleImpactedClientsClient objects.
func (s *SleImpactedClientsClient) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedClientsClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "mac", "name", "switches", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Degraded = temp.Degraded
    s.Duration = temp.Duration
    s.Mac = temp.Mac
    s.Name = temp.Name
    s.Switches = temp.Switches
    s.Total = temp.Total
    return nil
}

// tempSleImpactedClientsClient is a temporary struct used for validating the fields of SleImpactedClientsClient.
type tempSleImpactedClientsClient  struct {
    Degraded *int                             `json:"degraded,omitempty"`
    Duration *int                             `json:"duration,omitempty"`
    Mac      *string                          `json:"mac,omitempty"`
    Name     *string                          `json:"name,omitempty"`
    Switches []SleImpactedClientsClientSwitch `json:"switches,omitempty"`
    Total    *int                             `json:"total,omitempty"`
}
