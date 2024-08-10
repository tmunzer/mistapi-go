package models

import (
    "encoding/json"
)

// SimpleAlertDnsFailure represents a SimpleAlertDnsFailure struct.
type SimpleAlertDnsFailure struct {
    ClientCount          *int           `json:"client_count,omitempty"`
    // failing within minutes
    Duration             *int           `json:"duration,omitempty"`
    IncidentCount        *int           `json:"incident_count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SimpleAlertDnsFailure.
// It customizes the JSON marshaling process for SimpleAlertDnsFailure objects.
func (s SimpleAlertDnsFailure) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SimpleAlertDnsFailure object to a map representation for JSON marshaling.
func (s SimpleAlertDnsFailure) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SimpleAlertDnsFailure.
// It customizes the JSON unmarshaling process for SimpleAlertDnsFailure objects.
func (s *SimpleAlertDnsFailure) UnmarshalJSON(input []byte) error {
    var temp tempSimpleAlertDnsFailure
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

// tempSimpleAlertDnsFailure is a temporary struct used for validating the fields of SimpleAlertDnsFailure.
type tempSimpleAlertDnsFailure  struct {
    ClientCount   *int `json:"client_count,omitempty"`
    Duration      *int `json:"duration,omitempty"`
    IncidentCount *int `json:"incident_count,omitempty"`
}
