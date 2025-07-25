// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// AutoPreemption represents a AutoPreemption struct.
// Schedule to preempt ap’s which are not connected to preferred peer
type AutoPreemption struct {
    // enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
    DayOfWeek            *DayOfWeekEnum         `json:"day_of_week,omitempty"`
    // Whether auto preemption should happen
    Enabled              *bool                  `json:"enabled,omitempty"`
    // `any` / HH:MM (24-hour format)
    TimeOfDay            *string                `json:"time_of_day,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoPreemption,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoPreemption) String() string {
    return fmt.Sprintf(
    	"AutoPreemption[DayOfWeek=%v, Enabled=%v, TimeOfDay=%v, AdditionalProperties=%v]",
    	a.DayOfWeek, a.Enabled, a.TimeOfDay, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoPreemption.
// It customizes the JSON marshaling process for AutoPreemption objects.
func (a AutoPreemption) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "day_of_week", "enabled", "time_of_day"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AutoPreemption object to a map representation for JSON marshaling.
func (a AutoPreemption) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.DayOfWeek != nil {
        structMap["day_of_week"] = a.DayOfWeek
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.TimeOfDay != nil {
        structMap["time_of_day"] = a.TimeOfDay
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoPreemption.
// It customizes the JSON unmarshaling process for AutoPreemption objects.
func (a *AutoPreemption) UnmarshalJSON(input []byte) error {
    var temp tempAutoPreemption
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "day_of_week", "enabled", "time_of_day")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.DayOfWeek = temp.DayOfWeek
    a.Enabled = temp.Enabled
    a.TimeOfDay = temp.TimeOfDay
    return nil
}

// tempAutoPreemption is a temporary struct used for validating the fields of AutoPreemption.
type tempAutoPreemption  struct {
    DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
    Enabled   *bool          `json:"enabled,omitempty"`
    TimeOfDay *string        `json:"time_of_day,omitempty"`
}
