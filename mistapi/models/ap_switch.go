package models

import (
	"encoding/json"
)

// ApSwitch represents a ApSwitch struct.
// for people who want to fully control the vlans (advanced)
type ApSwitch struct {
	Enabled              *bool            `json:"enabled,omitempty"`
	Eth0                 *ApSwitchSetting `json:"eth0,omitempty"`
	Eth1                 *ApSwitchSetting `json:"eth1,omitempty"`
	Eth2                 *ApSwitchSetting `json:"eth2,omitempty"`
	Eth3                 *ApSwitchSetting `json:"eth3,omitempty"`
	Module               *ApSwitchSetting `json:"module,omitempty"`
	Wds                  *ApSwitchSetting `json:"wds,omitempty"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApSwitch.
// It customizes the JSON marshaling process for ApSwitch objects.
func (a ApSwitch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the ApSwitch object to a map representation for JSON marshaling.
func (a ApSwitch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Eth0 != nil {
		structMap["eth0"] = a.Eth0.toMap()
	}
	if a.Eth1 != nil {
		structMap["eth1"] = a.Eth1.toMap()
	}
	if a.Eth2 != nil {
		structMap["eth2"] = a.Eth2.toMap()
	}
	if a.Eth3 != nil {
		structMap["eth3"] = a.Eth3.toMap()
	}
	if a.Module != nil {
		structMap["module"] = a.Module.toMap()
	}
	if a.Wds != nil {
		structMap["wds"] = a.Wds.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApSwitch.
// It customizes the JSON unmarshaling process for ApSwitch objects.
func (a *ApSwitch) UnmarshalJSON(input []byte) error {
	var temp tempApSwitch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "eth0", "eth1", "eth2", "eth3", "module", "wds")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.Enabled = temp.Enabled
	a.Eth0 = temp.Eth0
	a.Eth1 = temp.Eth1
	a.Eth2 = temp.Eth2
	a.Eth3 = temp.Eth3
	a.Module = temp.Module
	a.Wds = temp.Wds
	return nil
}

// tempApSwitch is a temporary struct used for validating the fields of ApSwitch.
type tempApSwitch struct {
	Enabled *bool            `json:"enabled,omitempty"`
	Eth0    *ApSwitchSetting `json:"eth0,omitempty"`
	Eth1    *ApSwitchSetting `json:"eth1,omitempty"`
	Eth2    *ApSwitchSetting `json:"eth2,omitempty"`
	Eth3    *ApSwitchSetting `json:"eth3,omitempty"`
	Module  *ApSwitchSetting `json:"module,omitempty"`
	Wds     *ApSwitchSetting `json:"wds,omitempty"`
}
