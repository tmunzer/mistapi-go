// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsWirelessClientZone represents a StatsWirelessClientZone struct.
type StatsWirelessClientZone struct {
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Since                *int                   `json:"since,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClientZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClientZone) String() string {
    return fmt.Sprintf(
    	"StatsWirelessClientZone[Id=%v, Since=%v, AdditionalProperties=%v]",
    	s.Id, s.Since, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientZone.
// It customizes the JSON marshaling process for StatsWirelessClientZone objects.
func (s StatsWirelessClientZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientZone object to a map representation for JSON marshaling.
func (s StatsWirelessClientZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "since")
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
