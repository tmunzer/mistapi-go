package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientWirelessStatsVbeacon represents a ClientWirelessStatsVbeacon struct.
type ClientWirelessStatsVbeacon struct {
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Since                *int           `json:"since,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsVbeacon.
// It customizes the JSON marshaling process for ClientWirelessStatsVbeacon objects.
func (c ClientWirelessStatsVbeacon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsVbeacon object to a map representation for JSON marshaling.
func (c ClientWirelessStatsVbeacon) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsVbeacon.
// It customizes the JSON unmarshaling process for ClientWirelessStatsVbeacon objects.
func (c *ClientWirelessStatsVbeacon) UnmarshalJSON(input []byte) error {
    var temp tempClientWirelessStatsVbeacon
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

// tempClientWirelessStatsVbeacon is a temporary struct used for validating the fields of ClientWirelessStatsVbeacon.
type tempClientWirelessStatsVbeacon  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}
