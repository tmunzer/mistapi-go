package models

import (
	"encoding/json"
)

// JseDevice represents a JseDevice struct.
type JseDevice struct {
	// when available
	ExtIp                *string        `json:"ext_ip,omitempty"`
	LastSeen             *float64       `json:"last_seen,omitempty"`
	Mac                  *string        `json:"mac,omitempty"`
	Model                *string        `json:"model,omitempty"`
	Serial               *string        `json:"serial,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for JseDevice.
// It customizes the JSON marshaling process for JseDevice objects.
func (j JseDevice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(j.toMap())
}

// toMap converts the JseDevice object to a map representation for JSON marshaling.
func (j JseDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, j.AdditionalProperties)
	if j.ExtIp != nil {
		structMap["ext_ip"] = j.ExtIp
	}
	if j.LastSeen != nil {
		structMap["last_seen"] = j.LastSeen
	}
	if j.Mac != nil {
		structMap["mac"] = j.Mac
	}
	if j.Model != nil {
		structMap["model"] = j.Model
	}
	if j.Serial != nil {
		structMap["serial"] = j.Serial
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JseDevice.
// It customizes the JSON unmarshaling process for JseDevice objects.
func (j *JseDevice) UnmarshalJSON(input []byte) error {
	var temp tempJseDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ext_ip", "last_seen", "mac", "model", "serial")
	if err != nil {
		return err
	}

	j.AdditionalProperties = additionalProperties
	j.ExtIp = temp.ExtIp
	j.LastSeen = temp.LastSeen
	j.Mac = temp.Mac
	j.Model = temp.Model
	j.Serial = temp.Serial
	return nil
}

// tempJseDevice is a temporary struct used for validating the fields of JseDevice.
type tempJseDevice struct {
	ExtIp    *string  `json:"ext_ip,omitempty"`
	LastSeen *float64 `json:"last_seen,omitempty"`
	Mac      *string  `json:"mac,omitempty"`
	Model    *string  `json:"model,omitempty"`
	Serial   *string  `json:"serial,omitempty"`
}
