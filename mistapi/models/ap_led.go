package models

import (
	"encoding/json"
)

// ApLed represents a ApLed struct.
// LED AP settings
type ApLed struct {
	Brightness           *int           `json:"brightness,omitempty"`
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApLed.
// It customizes the JSON marshaling process for ApLed objects.
func (a ApLed) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApLed object to a map representation for JSON marshaling.
func (a ApLed) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Brightness != nil {
		structMap["brightness"] = a.Brightness
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApLed.
// It customizes the JSON unmarshaling process for ApLed objects.
func (a *ApLed) UnmarshalJSON(input []byte) error {
	var temp tempApLed
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "brightness", "enabled")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Brightness = temp.Brightness
	a.Enabled = temp.Enabled
	return nil
}

// tempApLed is a temporary struct used for validating the fields of ApLed.
type tempApLed struct {
	Brightness *int  `json:"brightness,omitempty"`
	Enabled    *bool `json:"enabled,omitempty"`
}
