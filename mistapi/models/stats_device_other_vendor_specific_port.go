package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// StatsDeviceOtherVendorSpecificPort represents a StatsDeviceOtherVendorSpecificPort struct.
type StatsDeviceOtherVendorSpecificPort struct {
    BytesIn              *int           `json:"bytes_in,omitempty"`
    BytesOut             *int           `json:"bytes_out,omitempty"`
    HealthCategory       *string        `json:"health_category,omitempty"`
    HealthScore          *int           `json:"health_score,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Mode                 *string        `json:"mode,omitempty"`
    Model                *string        `json:"model,omitempty"`
    State                *string        `json:"state,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Uptime               *float64       `json:"uptime,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOtherVendorSpecificPort.
// It customizes the JSON marshaling process for StatsDeviceOtherVendorSpecificPort objects.
func (s StatsDeviceOtherVendorSpecificPort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOtherVendorSpecificPort object to a map representation for JSON marshaling.
func (s StatsDeviceOtherVendorSpecificPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BytesIn != nil {
        structMap["bytes_in"] = s.BytesIn
    }
    if s.BytesOut != nil {
        structMap["bytes_out"] = s.BytesOut
    }
    if s.HealthCategory != nil {
        structMap["health_category"] = s.HealthCategory
    }
    if s.HealthScore != nil {
        structMap["health_score"] = s.HealthScore
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Mode != nil {
        structMap["mode"] = s.Mode
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.State != nil {
        structMap["state"] = s.State
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDeviceOtherVendorSpecificPort.
// It customizes the JSON unmarshaling process for StatsDeviceOtherVendorSpecificPort objects.
func (s *StatsDeviceOtherVendorSpecificPort) UnmarshalJSON(input []byte) error {
    var temp tempStatsDeviceOtherVendorSpecificPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bytes_in", "bytes_out", "health_category", "health_score", "id", "mode", "model", "state", "type", "uptime")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.BytesIn = temp.BytesIn
    s.BytesOut = temp.BytesOut
    s.HealthCategory = temp.HealthCategory
    s.HealthScore = temp.HealthScore
    s.Id = temp.Id
    s.Mode = temp.Mode
    s.Model = temp.Model
    s.State = temp.State
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    return nil
}

// tempStatsDeviceOtherVendorSpecificPort is a temporary struct used for validating the fields of StatsDeviceOtherVendorSpecificPort.
type tempStatsDeviceOtherVendorSpecificPort  struct {
    BytesIn        *int       `json:"bytes_in,omitempty"`
    BytesOut       *int       `json:"bytes_out,omitempty"`
    HealthCategory *string    `json:"health_category,omitempty"`
    HealthScore    *int       `json:"health_score,omitempty"`
    Id             *uuid.UUID `json:"id,omitempty"`
    Mode           *string    `json:"mode,omitempty"`
    Model          *string    `json:"model,omitempty"`
    State          *string    `json:"state,omitempty"`
    Type           *string    `json:"type,omitempty"`
    Uptime         *float64   `json:"uptime,omitempty"`
}
