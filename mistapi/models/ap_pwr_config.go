package models

import (
	"encoding/json"
)

// ApPwrConfig represents a ApPwrConfig struct.
// power related configs
type ApPwrConfig struct {
	// additional power to request during negotiating with PSE over PoE, in mW
	Base *int `json:"base,omitempty"`
	// whether to enable power out to peripheral, meanwhile will reduce power to wifi (only for AP45 at power mode)
	PreferUsbOverWifi    *bool          `json:"prefer_usb_over_wifi,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApPwrConfig.
// It customizes the JSON marshaling process for ApPwrConfig objects.
func (a ApPwrConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApPwrConfig object to a map representation for JSON marshaling.
func (a ApPwrConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Base != nil {
		structMap["base"] = a.Base
	}
	if a.PreferUsbOverWifi != nil {
		structMap["prefer_usb_over_wifi"] = a.PreferUsbOverWifi
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApPwrConfig.
// It customizes the JSON unmarshaling process for ApPwrConfig objects.
func (a *ApPwrConfig) UnmarshalJSON(input []byte) error {
	var temp tempApPwrConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "base", "prefer_usb_over_wifi")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Base = temp.Base
	a.PreferUsbOverWifi = temp.PreferUsbOverWifi
	return nil
}

// tempApPwrConfig is a temporary struct used for validating the fields of ApPwrConfig.
type tempApPwrConfig struct {
	Base              *int  `json:"base,omitempty"`
	PreferUsbOverWifi *bool `json:"prefer_usb_over_wifi,omitempty"`
}
