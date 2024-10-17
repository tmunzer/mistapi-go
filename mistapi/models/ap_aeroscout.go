package models

import (
	"encoding/json"
)

// ApAeroscout represents a ApAeroscout struct.
// Aeroscout AP settings
type ApAeroscout struct {
	// whether to enable aeroscout config
	Enabled *bool `json:"enabled,omitempty"`
	// required if enabled, aeroscout server host
	Host Optional[string] `json:"host"`
	// whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation
	LocateConnected      *bool          `json:"locate_connected,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApAeroscout.
// It customizes the JSON marshaling process for ApAeroscout objects.
func (a ApAeroscout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApAeroscout object to a map representation for JSON marshaling.
func (a ApAeroscout) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Host.IsValueSet() {
		if a.Host.Value() != nil {
			structMap["host"] = a.Host.Value()
		} else {
			structMap["host"] = nil
		}
	}
	if a.LocateConnected != nil {
		structMap["locate_connected"] = a.LocateConnected
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApAeroscout.
// It customizes the JSON unmarshaling process for ApAeroscout objects.
func (a *ApAeroscout) UnmarshalJSON(input []byte) error {
	var temp tempApAeroscout
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "host", "locate_connected")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Enabled = temp.Enabled
	a.Host = temp.Host
	a.LocateConnected = temp.LocateConnected
	return nil
}

// tempApAeroscout is a temporary struct used for validating the fields of ApAeroscout.
type tempApAeroscout struct {
	Enabled         *bool            `json:"enabled,omitempty"`
	Host            Optional[string] `json:"host"`
	LocateConnected *bool            `json:"locate_connected,omitempty"`
}
