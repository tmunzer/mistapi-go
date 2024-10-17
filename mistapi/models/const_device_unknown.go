package models

import (
	"encoding/json"
)

// ConstDeviceUnknown represents a ConstDeviceUnknown struct.
type ConstDeviceUnknown struct {
	ApType               *string        `json:"ap_type,omitempty"`
	Description          *string        `json:"description,omitempty"`
	Display              *string        `json:"display,omitempty"`
	HasExtio             *bool          `json:"has_extio,omitempty"`
	HasVble              *bool          `json:"has_vble,omitempty"`
	HasWifiBand24        *bool          `json:"has_wifi_band24,omitempty"`
	HasWifiBand5         *bool          `json:"has_wifi_band5,omitempty"`
	Model                *string        `json:"model,omitempty"`
	Unmanaged            *bool          `json:"unmanaged,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceUnknown.
// It customizes the JSON marshaling process for ConstDeviceUnknown objects.
func (c ConstDeviceUnknown) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceUnknown object to a map representation for JSON marshaling.
func (c ConstDeviceUnknown) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.ApType != nil {
		structMap["ap_type"] = c.ApType
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	if c.Display != nil {
		structMap["display"] = c.Display
	}
	if c.HasExtio != nil {
		structMap["has_extio"] = c.HasExtio
	}
	if c.HasVble != nil {
		structMap["has_vble"] = c.HasVble
	}
	if c.HasWifiBand24 != nil {
		structMap["has_wifi_band24"] = c.HasWifiBand24
	}
	if c.HasWifiBand5 != nil {
		structMap["has_wifi_band5"] = c.HasWifiBand5
	}
	if c.Model != nil {
		structMap["model"] = c.Model
	}
	if c.Unmanaged != nil {
		structMap["unmanaged"] = c.Unmanaged
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceUnknown.
// It customizes the JSON unmarshaling process for ConstDeviceUnknown objects.
func (c *ConstDeviceUnknown) UnmarshalJSON(input []byte) error {
	var temp tempConstDeviceUnknown
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_type", "description", "display", "has_extio", "has_vble", "has_wifi_band24", "has_wifi_band5", "model", "unmanaged")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.ApType = temp.ApType
	c.Description = temp.Description
	c.Display = temp.Display
	c.HasExtio = temp.HasExtio
	c.HasVble = temp.HasVble
	c.HasWifiBand24 = temp.HasWifiBand24
	c.HasWifiBand5 = temp.HasWifiBand5
	c.Model = temp.Model
	c.Unmanaged = temp.Unmanaged
	return nil
}

// tempConstDeviceUnknown is a temporary struct used for validating the fields of ConstDeviceUnknown.
type tempConstDeviceUnknown struct {
	ApType        *string `json:"ap_type,omitempty"`
	Description   *string `json:"description,omitempty"`
	Display       *string `json:"display,omitempty"`
	HasExtio      *bool   `json:"has_extio,omitempty"`
	HasVble       *bool   `json:"has_vble,omitempty"`
	HasWifiBand24 *bool   `json:"has_wifi_band24,omitempty"`
	HasWifiBand5  *bool   `json:"has_wifi_band5,omitempty"`
	Model         *string `json:"model,omitempty"`
	Unmanaged     *bool   `json:"unmanaged,omitempty"`
}
