package models

import (
	"encoding/json"
)

// ApUsb represents a ApUsb struct.
// USB AP settings
// Note: if native imagotag is enabled, BLE will be disabled automatically
// Note: legacy, new config moved to ESL Config.
type ApUsb struct {
	// only if `type`==`imagotag`
	Cacert Optional[string] `json:"cacert"`
	// only if `type`==`imagotag`, channel selection, not needed by default, required for manual channel override only
	Channel *int `json:"channel,omitempty"`
	// whether to enable any usb config
	Enabled *bool `json:"enabled,omitempty"`
	// only if `type`==`imagotag`
	Host *string `json:"host,omitempty"`
	// only if `type`==`imagotag`
	Port *int `json:"port,omitempty"`
	// usb config type. enum: `hanshow`, `imagotag`, `solum`
	Type *ApUsbTypeEnum `json:"type,omitempty"`
	// only if `type`==`imagotag`, whether to turn on SSL verification
	VerifyCert *bool `json:"verify_cert,omitempty"`
	// only if `type`==`solum` or `type`==`hanshow`
	VlanId               *int           `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApUsb.
// It customizes the JSON marshaling process for ApUsb objects.
func (a ApUsb) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApUsb object to a map representation for JSON marshaling.
func (a ApUsb) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Cacert.IsValueSet() {
		if a.Cacert.Value() != nil {
			structMap["cacert"] = a.Cacert.Value()
		} else {
			structMap["cacert"] = nil
		}
	}
	if a.Channel != nil {
		structMap["channel"] = a.Channel
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Host != nil {
		structMap["host"] = a.Host
	}
	if a.Port != nil {
		structMap["port"] = a.Port
	}
	if a.Type != nil {
		structMap["type"] = a.Type
	}
	if a.VerifyCert != nil {
		structMap["verify_cert"] = a.VerifyCert
	}
	if a.VlanId != nil {
		structMap["vlan_id"] = a.VlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApUsb.
// It customizes the JSON unmarshaling process for ApUsb objects.
func (a *ApUsb) UnmarshalJSON(input []byte) error {
	var temp tempApUsb
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cacert", "channel", "enabled", "host", "port", "type", "verify_cert", "vlan_id")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Cacert = temp.Cacert
	a.Channel = temp.Channel
	a.Enabled = temp.Enabled
	a.Host = temp.Host
	a.Port = temp.Port
	a.Type = temp.Type
	a.VerifyCert = temp.VerifyCert
	a.VlanId = temp.VlanId
	return nil
}

// tempApUsb is a temporary struct used for validating the fields of ApUsb.
type tempApUsb struct {
	Cacert     Optional[string] `json:"cacert"`
	Channel    *int             `json:"channel,omitempty"`
	Enabled    *bool            `json:"enabled,omitempty"`
	Host       *string          `json:"host,omitempty"`
	Port       *int             `json:"port,omitempty"`
	Type       *ApUsbTypeEnum   `json:"type,omitempty"`
	VerifyCert *bool            `json:"verify_cert,omitempty"`
	VlanId     *int             `json:"vlan_id,omitempty"`
}
