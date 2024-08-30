package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// StatsWirelessClientZone represents a StatsWirelessClientZone struct.
type StatsWirelessClientZone struct {
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Since                *int           `json:"since,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientZone.
// It customizes the JSON marshaling process for StatsWirelessClientZone objects.
func (s StatsWirelessClientZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientZone object to a map representation for JSON marshaling.
func (s StatsWirelessClientZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Since != nil {
        structMap["since"] = s.Since
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientZone.
// It customizes the JSON unmarshaling process for StatsWirelessClientZone objects.
func (s *StatsWirelessClientZone) UnmarshalJSON(input []byte) error {
    var temp tempStatsWirelessClientZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "since")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.Since = temp.Since
    return nil
}

// tempStatsWirelessClientZone is a temporary struct used for validating the fields of StatsWirelessClientZone.
type tempStatsWirelessClientZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}