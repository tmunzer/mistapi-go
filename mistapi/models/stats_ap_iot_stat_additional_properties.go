// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsApIotStatAdditionalProperties represents a StatsApIotStatAdditionalProperties struct.
type StatsApIotStatAdditionalProperties struct {
    Value                Optional[int]          `json:"value"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApIotStatAdditionalProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApIotStatAdditionalProperties) String() string {
    return fmt.Sprintf(
    	"StatsApIotStatAdditionalProperties[Value=%v, AdditionalProperties=%v]",
    	s.Value, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApIotStatAdditionalProperties.
// It customizes the JSON marshaling process for StatsApIotStatAdditionalProperties objects.
func (s StatsApIotStatAdditionalProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApIotStatAdditionalProperties object to a map representation for JSON marshaling.
func (s StatsApIotStatAdditionalProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Value.IsValueSet() {
        if s.Value.Value() != nil {
            structMap["value"] = s.Value.Value()
        } else {
            structMap["value"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApIotStatAdditionalProperties.
// It customizes the JSON unmarshaling process for StatsApIotStatAdditionalProperties objects.
func (s *StatsApIotStatAdditionalProperties) UnmarshalJSON(input []byte) error {
    var temp tempStatsApIotStatAdditionalProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "value")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Value = temp.Value
    return nil
}

// tempStatsApIotStatAdditionalProperties is a temporary struct used for validating the fields of StatsApIotStatAdditionalProperties.
type tempStatsApIotStatAdditionalProperties  struct {
    Value Optional[int] `json:"value"`
}
