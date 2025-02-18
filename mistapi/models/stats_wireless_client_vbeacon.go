package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsWirelessClientVbeacon represents a StatsWirelessClientVbeacon struct.
type StatsWirelessClientVbeacon struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Since                *int                   `json:"since,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClientVbeacon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClientVbeacon) String() string {
    return fmt.Sprintf(
    	"StatsWirelessClientVbeacon[Id=%v, Since=%v, AdditionalProperties=%v]",
    	s.Id, s.Since, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientVbeacon.
// It customizes the JSON marshaling process for StatsWirelessClientVbeacon objects.
func (s StatsWirelessClientVbeacon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientVbeacon object to a map representation for JSON marshaling.
func (s StatsWirelessClientVbeacon) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientVbeacon.
// It customizes the JSON unmarshaling process for StatsWirelessClientVbeacon objects.
func (s *StatsWirelessClientVbeacon) UnmarshalJSON(input []byte) error {
    var temp tempStatsWirelessClientVbeacon
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

// tempStatsWirelessClientVbeacon is a temporary struct used for validating the fields of StatsWirelessClientVbeacon.
type tempStatsWirelessClientVbeacon  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}
