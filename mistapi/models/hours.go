// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// Hours represents a Hours struct.
// Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun)
type Hours struct {
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Fri                  *string                `json:"fri,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Mon                  *string                `json:"mon,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Sat                  *string                `json:"sat,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Sun                  *string                `json:"sun,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Thu                  *string                `json:"thu,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Tue                  *string                `json:"tue,omitempty"`
    // Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
    Wed                  *string                `json:"wed,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Hours,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h Hours) String() string {
    return fmt.Sprintf(
    	"Hours[Fri=%v, Mon=%v, Sat=%v, Sun=%v, Thu=%v, Tue=%v, Wed=%v, AdditionalProperties=%v]",
    	h.Fri, h.Mon, h.Sat, h.Sun, h.Thu, h.Tue, h.Wed, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Hours.
// It customizes the JSON marshaling process for Hours objects.
func (h Hours) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(h.AdditionalProperties,
        "fri", "mon", "sat", "sun", "thu", "tue", "wed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(h.toMap())
}

// toMap converts the Hours object to a map representation for JSON marshaling.
func (h Hours) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, h.AdditionalProperties)
    if h.Fri != nil {
        structMap["fri"] = h.Fri
    }
    if h.Mon != nil {
        structMap["mon"] = h.Mon
    }
    if h.Sat != nil {
        structMap["sat"] = h.Sat
    }
    if h.Sun != nil {
        structMap["sun"] = h.Sun
    }
    if h.Thu != nil {
        structMap["thu"] = h.Thu
    }
    if h.Tue != nil {
        structMap["tue"] = h.Tue
    }
    if h.Wed != nil {
        structMap["wed"] = h.Wed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Hours.
// It customizes the JSON unmarshaling process for Hours objects.
func (h *Hours) UnmarshalJSON(input []byte) error {
    var temp tempHours
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "fri", "mon", "sat", "sun", "thu", "tue", "wed")
    if err != nil {
    	return err
    }
    h.AdditionalProperties = additionalProperties
    
    h.Fri = temp.Fri
    h.Mon = temp.Mon
    h.Sat = temp.Sat
    h.Sun = temp.Sun
    h.Thu = temp.Thu
    h.Tue = temp.Tue
    h.Wed = temp.Wed
    return nil
}

// tempHours is a temporary struct used for validating the fields of Hours.
type tempHours  struct {
    Fri *string `json:"fri,omitempty"`
    Mon *string `json:"mon,omitempty"`
    Sat *string `json:"sat,omitempty"`
    Sun *string `json:"sun,omitempty"`
    Thu *string `json:"thu,omitempty"`
    Tue *string `json:"tue,omitempty"`
    Wed *string `json:"wed,omitempty"`
}
