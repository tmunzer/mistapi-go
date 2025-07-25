// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SimpleAlertDnsFailure represents a SimpleAlertDnsFailure struct.
type SimpleAlertDnsFailure struct {
    ClientCount          *int                   `json:"client_count,omitempty"`
    // failing within minutes
    Duration             *int                   `json:"duration,omitempty"`
    IncidentCount        *int                   `json:"incident_count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SimpleAlertDnsFailure,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SimpleAlertDnsFailure) String() string {
    return fmt.Sprintf(
    	"SimpleAlertDnsFailure[ClientCount=%v, Duration=%v, IncidentCount=%v, AdditionalProperties=%v]",
    	s.ClientCount, s.Duration, s.IncidentCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SimpleAlertDnsFailure.
// It customizes the JSON marshaling process for SimpleAlertDnsFailure objects.
func (s SimpleAlertDnsFailure) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "client_count", "duration", "incident_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SimpleAlertDnsFailure object to a map representation for JSON marshaling.
func (s SimpleAlertDnsFailure) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ClientCount != nil {
        structMap["client_count"] = s.ClientCount
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.IncidentCount != nil {
        structMap["incident_count"] = s.IncidentCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SimpleAlertDnsFailure.
// It customizes the JSON unmarshaling process for SimpleAlertDnsFailure objects.
func (s *SimpleAlertDnsFailure) UnmarshalJSON(input []byte) error {
    var temp tempSimpleAlertDnsFailure
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_count", "duration", "incident_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ClientCount = temp.ClientCount
    s.Duration = temp.Duration
    s.IncidentCount = temp.IncidentCount
    return nil
}

// tempSimpleAlertDnsFailure is a temporary struct used for validating the fields of SimpleAlertDnsFailure.
type tempSimpleAlertDnsFailure  struct {
    ClientCount   *int `json:"client_count,omitempty"`
    Duration      *int `json:"duration,omitempty"`
    IncidentCount *int `json:"incident_count,omitempty"`
}
