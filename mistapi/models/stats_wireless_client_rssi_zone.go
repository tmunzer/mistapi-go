package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsWirelessClientRssiZone represents a StatsWirelessClientRssiZone struct.
type StatsWirelessClientRssiZone struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Since                *int                   `json:"since,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClientRssiZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClientRssiZone) String() string {
    return fmt.Sprintf(
    	"StatsWirelessClientRssiZone[Id=%v, Since=%v, AdditionalProperties=%v]",
    	s.Id, s.Since, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientRssiZone.
// It customizes the JSON marshaling process for StatsWirelessClientRssiZone objects.
func (s StatsWirelessClientRssiZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientRssiZone object to a map representation for JSON marshaling.
func (s StatsWirelessClientRssiZone) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientRssiZone.
// It customizes the JSON unmarshaling process for StatsWirelessClientRssiZone objects.
func (s *StatsWirelessClientRssiZone) UnmarshalJSON(input []byte) error {
    var temp tempStatsWirelessClientRssiZone
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

// tempStatsWirelessClientRssiZone is a temporary struct used for validating the fields of StatsWirelessClientRssiZone.
type tempStatsWirelessClientRssiZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}
