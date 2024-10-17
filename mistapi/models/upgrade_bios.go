package models

import (
	"encoding/json"
)

// UpgradeBios represents a UpgradeBios struct.
type UpgradeBios struct {
	// Reboot device immediately after upgrade is completed
	Reboot *bool `json:"reboot,omitempty"`
	// specific bios version
	Version              *string        `json:"version,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpgradeBios.
// It customizes the JSON marshaling process for UpgradeBios objects.
func (u UpgradeBios) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpgradeBios object to a map representation for JSON marshaling.
func (u UpgradeBios) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Reboot != nil {
		structMap["reboot"] = u.Reboot
	}
	if u.Version != nil {
		structMap["version"] = u.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeBios.
// It customizes the JSON unmarshaling process for UpgradeBios objects.
func (u *UpgradeBios) UnmarshalJSON(input []byte) error {
	var temp tempUpgradeBios
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "reboot", "version")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.Reboot = temp.Reboot
	u.Version = temp.Version
	return nil
}

// tempUpgradeBios is a temporary struct used for validating the fields of UpgradeBios.
type tempUpgradeBios struct {
	Reboot  *bool   `json:"reboot,omitempty"`
	Version *string `json:"version,omitempty"`
}
