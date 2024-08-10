package models

import (
    "encoding/json"
)

// SynthetictestConfigWanSpeedtest represents a SynthetictestConfigWanSpeedtest struct.
type SynthetictestConfigWanSpeedtest struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    // any / HH:MM (24-hour format)
    TimeOdFay            *string        `json:"time_od_fay,omitempty"`
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
    if s.TimeOdFay != nil {
        structMap["time_od_fay"] = s.TimeOdFay
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "time_od_fay")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    s.TimeOdFay = temp.TimeOdFay
    return nil
}

// tempSynthetictestConfigWanSpeedtest is a temporary struct used for validating the fields of SynthetictestConfigWanSpeedtest.
type tempSynthetictestConfigWanSpeedtest  struct {
    Enabled   *bool   `json:"enabled,omitempty"`
    TimeOdFay *string `json:"time_od_fay,omitempty"`
}
