// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsWirelessClientWxruleUsage represents a StatsWirelessClientWxruleUsage struct.
type StatsWirelessClientWxruleUsage struct {
    TagId                *uuid.UUID             `json:"tag_id,omitempty"`
    Usage                *int                   `json:"usage,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClientWxruleUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClientWxruleUsage) String() string {
    return fmt.Sprintf(
    	"StatsWirelessClientWxruleUsage[TagId=%v, Usage=%v, AdditionalProperties=%v]",
    	s.TagId, s.Usage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientWxruleUsage.
// It customizes the JSON marshaling process for StatsWirelessClientWxruleUsage objects.
func (s StatsWirelessClientWxruleUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "tag_id", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientWxruleUsage object to a map representation for JSON marshaling.
func (s StatsWirelessClientWxruleUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.TagId != nil {
        structMap["tag_id"] = s.TagId
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientWxruleUsage.
// It customizes the JSON unmarshaling process for StatsWirelessClientWxruleUsage objects.
func (s *StatsWirelessClientWxruleUsage) UnmarshalJSON(input []byte) error {
    var temp tempStatsWirelessClientWxruleUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tag_id", "usage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.TagId = temp.TagId
    s.Usage = temp.Usage
    return nil
}

// tempStatsWirelessClientWxruleUsage is a temporary struct used for validating the fields of StatsWirelessClientWxruleUsage.
type tempStatsWirelessClientWxruleUsage  struct {
    TagId *uuid.UUID `json:"tag_id,omitempty"`
    Usage *int       `json:"usage,omitempty"`
}
