// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedApplicationsApp represents a SleImpactedApplicationsApp struct.
type SleImpactedApplicationsApp struct {
    App                  *string                `json:"app,omitempty"`
    Degraded             *int                   `json:"degraded,omitempty"`
    Duration             *int                   `json:"duration,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Threshold            *int                   `json:"threshold,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedApplicationsApp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedApplicationsApp) String() string {
    return fmt.Sprintf(
    	"SleImpactedApplicationsApp[App=%v, Degraded=%v, Duration=%v, Name=%v, Threshold=%v, Total=%v, AdditionalProperties=%v]",
    	s.App, s.Degraded, s.Duration, s.Name, s.Threshold, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedApplicationsApp.
// It customizes the JSON marshaling process for SleImpactedApplicationsApp objects.
func (s SleImpactedApplicationsApp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "app", "degraded", "duration", "name", "threshold", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedApplicationsApp object to a map representation for JSON marshaling.
func (s SleImpactedApplicationsApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSleImpactedApplicationsApp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "app", "degraded", "duration", "name", "threshold", "total")
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

// tempSleImpactedApplicationsApp is a temporary struct used for validating the fields of SleImpactedApplicationsApp.
type tempSleImpactedApplicationsApp  struct {
    App       *string `json:"app,omitempty"`
    Degraded  *int    `json:"degraded,omitempty"`
    Duration  *int    `json:"duration,omitempty"`
    Name      *string `json:"name,omitempty"`
    Threshold *int    `json:"threshold,omitempty"`
    Total     *int    `json:"total,omitempty"`
}
