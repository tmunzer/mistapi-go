package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientWirelessStatsWxruleUsage represents a ClientWirelessStatsWxruleUsage struct.
type ClientWirelessStatsWxruleUsage struct {
    TagId                *uuid.UUID     `json:"tag_id,omitempty"`
    Usage                *int           `json:"usage,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsWxruleUsage.
// It customizes the JSON marshaling process for ClientWirelessStatsWxruleUsage objects.
func (c ClientWirelessStatsWxruleUsage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsWxruleUsage object to a map representation for JSON marshaling.
func (c ClientWirelessStatsWxruleUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.TagId != nil {
        structMap["tag_id"] = c.TagId
    }
    if c.Usage != nil {
        structMap["usage"] = c.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsWxruleUsage.
// It customizes the JSON unmarshaling process for ClientWirelessStatsWxruleUsage objects.
func (c *ClientWirelessStatsWxruleUsage) UnmarshalJSON(input []byte) error {
    var temp tempClientWirelessStatsWxruleUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "tag_id", "usage")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.TagId = temp.TagId
    c.Usage = temp.Usage
    return nil
}

// tempClientWirelessStatsWxruleUsage is a temporary struct used for validating the fields of ClientWirelessStatsWxruleUsage.
type tempClientWirelessStatsWxruleUsage  struct {
    TagId *uuid.UUID `json:"tag_id,omitempty"`
    Usage *int       `json:"usage,omitempty"`
}
