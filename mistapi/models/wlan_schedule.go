// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WlanSchedule represents a WlanSchedule struct.
// WLAN operating schedule, default is disabled
type WlanSchedule struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun)
	Hours                *Hours                 `json:"hours,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanSchedule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanSchedule) String() string {
	return fmt.Sprintf(
		"WlanSchedule[Enabled=%v, Hours=%v, AdditionalProperties=%v]",
		w.Enabled, w.Hours, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanSchedule.
// It customizes the JSON marshaling process for WlanSchedule objects.
func (w WlanSchedule) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"enabled", "hours"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WlanSchedule object to a map representation for JSON marshaling.
func (w WlanSchedule) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Enabled != nil {
		structMap["enabled"] = w.Enabled
	}
	if w.Hours != nil {
		structMap["hours"] = w.Hours.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanSchedule.
// It customizes the JSON unmarshaling process for WlanSchedule objects.
func (w *WlanSchedule) UnmarshalJSON(input []byte) error {
	var temp tempWlanSchedule
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "hours")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Enabled = temp.Enabled
	w.Hours = temp.Hours
	return nil
}

// tempWlanSchedule is a temporary struct used for validating the fields of WlanSchedule.
type tempWlanSchedule struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Hours   *Hours `json:"hours,omitempty"`
}
