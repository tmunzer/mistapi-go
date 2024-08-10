package models

import (
    "encoding/json"
)

// SleImpactedClientsClient represents a SleImpactedClientsClient struct.
type SleImpactedClientsClient struct {
    Degraded             *int                             `json:"degraded,omitempty"`
    Duration             *int                             `json:"duration,omitempty"`
    Mac                  *string                          `json:"mac,omitempty"`
    Name                 *string                          `json:"name,omitempty"`
    Switches             []SleImpactedClientsClientSwitch `json:"switches,omitempty"`
    Total                *int                             `json:"total,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedClientsClient.
// It customizes the JSON marshaling process for SleImpactedClientsClient objects.
func (s SleImpactedClientsClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedClientsClient object to a map representation for JSON marshaling.
func (s SleImpactedClientsClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "duration", "mac", "name", "switches", "total")
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
