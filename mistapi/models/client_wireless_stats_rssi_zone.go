package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientWirelessStatsRssiZone represents a ClientWirelessStatsRssiZone struct.
type ClientWirelessStatsRssiZone struct {
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Since                *int           `json:"since,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsRssiZone.
// It customizes the JSON marshaling process for ClientWirelessStatsRssiZone objects.
func (c ClientWirelessStatsRssiZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsRssiZone object to a map representation for JSON marshaling.
func (c ClientWirelessStatsRssiZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Since != nil {
        structMap["since"] = c.Since
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsRssiZone.
// It customizes the JSON unmarshaling process for ClientWirelessStatsRssiZone objects.
func (c *ClientWirelessStatsRssiZone) UnmarshalJSON(input []byte) error {
    var temp tempClientWirelessStatsRssiZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "since")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Id = temp.Id
    c.Since = temp.Since
    return nil
}

// tempClientWirelessStatsRssiZone is a temporary struct used for validating the fields of ClientWirelessStatsRssiZone.
type tempClientWirelessStatsRssiZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}
