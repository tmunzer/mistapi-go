package models

import (
    "encoding/json"
)

// SimpleAlertDhcpFailure represents a SimpleAlertDhcpFailure struct.
type SimpleAlertDhcpFailure struct {
    ClientCount          *int           `json:"client_count,omitempty"`
    // failing within minutes
    Duration             *int           `json:"duration,omitempty"`
    IncidentCount        *int           `json:"incident_count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SimpleAlertDhcpFailure.
// It customizes the JSON marshaling process for SimpleAlertDhcpFailure objects.
func (s SimpleAlertDhcpFailure) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SimpleAlertDhcpFailure object to a map representation for JSON marshaling.
func (s SimpleAlertDhcpFailure) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for SimpleAlertDhcpFailure.
// It customizes the JSON unmarshaling process for SimpleAlertDhcpFailure objects.
func (s *SimpleAlertDhcpFailure) UnmarshalJSON(input []byte) error {
    var temp simpleAlertDhcpFailure
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "client_count", "duration", "incident_count")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ClientCount = temp.ClientCount
    s.Duration = temp.Duration
    s.IncidentCount = temp.IncidentCount
    return nil
}

// simpleAlertDhcpFailure is a temporary struct used for validating the fields of SimpleAlertDhcpFailure.
type simpleAlertDhcpFailure  struct {
    ClientCount   *int `json:"client_count,omitempty"`
    Duration      *int `json:"duration,omitempty"`
    IncidentCount *int `json:"incident_count,omitempty"`
}