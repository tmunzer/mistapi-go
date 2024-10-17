package models

import (
	"encoding/json"
)

// WlanSchedule represents a WlanSchedule struct.
// WLAN operating schedule, default is disabled
type WlanSchedule struct {
	Enabled *bool `json:"enabled,omitempty"`
	// hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).
	// **Note**: If the dow is not defined then it\u2019\ s treated as 00:00-23:59.
	Hours                *Hours         `json:"hours,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanSchedule.
// It customizes the JSON marshaling process for WlanSchedule objects.
func (w WlanSchedule) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WlanSchedule object to a map representation for JSON marshaling.
func (w WlanSchedule) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "hours")
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
