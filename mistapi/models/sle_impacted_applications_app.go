package models

import (
    "encoding/json"
)

// SleImpactedApplicationsApp represents a SleImpactedApplicationsApp struct.
type SleImpactedApplicationsApp struct {
    App                  *string        `json:"app,omitempty"`
    Degraded             *int           `json:"degraded,omitempty"`
    Duration             *int           `json:"duration,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Threshold            *int           `json:"threshold,omitempty"`
    Total                *int           `json:"total,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedApplicationsApp.
// It customizes the JSON marshaling process for SleImpactedApplicationsApp objects.
func (s SleImpactedApplicationsApp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedApplicationsApp object to a map representation for JSON marshaling.
func (s SleImpactedApplicationsApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.App != nil {
        structMap["app"] = s.App
    }
    if s.Degraded != nil {
        structMap["degraded"] = s.Degraded
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Threshold != nil {
        structMap["threshold"] = s.Threshold
    }
    if s.Total != nil {
        structMap["total"] = s.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedApplicationsApp.
// It customizes the JSON unmarshaling process for SleImpactedApplicationsApp objects.
func (s *SleImpactedApplicationsApp) UnmarshalJSON(input []byte) error {
    var temp sleImpactedApplicationsApp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "app", "degraded", "duration", "name", "threshold", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.App = temp.App
    s.Degraded = temp.Degraded
    s.Duration = temp.Duration
    s.Name = temp.Name
    s.Threshold = temp.Threshold
    s.Total = temp.Total
    return nil
}

// sleImpactedApplicationsApp is a temporary struct used for validating the fields of SleImpactedApplicationsApp.
type sleImpactedApplicationsApp  struct {
    App       *string `json:"app,omitempty"`
    Degraded  *int    `json:"degraded,omitempty"`
    Duration  *int    `json:"duration,omitempty"`
    Name      *string `json:"name,omitempty"`
    Threshold *int    `json:"threshold,omitempty"`
    Total     *int    `json:"total,omitempty"`
}
