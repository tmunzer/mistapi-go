package models

import (
    "encoding/json"
)

// SynthetictestConfigWanSpeedtest represents a SynthetictestConfigWanSpeedtest struct.
type SynthetictestConfigWanSpeedtest struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    // any / HH:MM (24-hour format)
    TimeOfDay            *string        `json:"time_of_day,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfigWanSpeedtest.
// It customizes the JSON marshaling process for SynthetictestConfigWanSpeedtest objects.
func (s SynthetictestConfigWanSpeedtest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfigWanSpeedtest object to a map representation for JSON marshaling.
func (s SynthetictestConfigWanSpeedtest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.TimeOfDay != nil {
        structMap["time_of_day"] = s.TimeOfDay
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfigWanSpeedtest.
// It customizes the JSON unmarshaling process for SynthetictestConfigWanSpeedtest objects.
func (s *SynthetictestConfigWanSpeedtest) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestConfigWanSpeedtest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "time_of_day")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    s.TimeOfDay = temp.TimeOfDay
    return nil
}

// tempSynthetictestConfigWanSpeedtest is a temporary struct used for validating the fields of SynthetictestConfigWanSpeedtest.
type tempSynthetictestConfigWanSpeedtest  struct {
    Enabled   *bool   `json:"enabled,omitempty"`
    TimeOfDay *string `json:"time_of_day,omitempty"`
}
