package models

import (
	"encoding/json"
)

// DiscoveredSwitchAp represents a DiscoveredSwitchAp struct.
type DiscoveredSwitchAp struct {
	Hostname             *string        `json:"hostname,omitempty"`
	Mac                  *string        `json:"mac,omitempty"`
	PoeStatus            *bool          `json:"poe_status,omitempty"`
	Port                 *string        `json:"port,omitempty"`
	PortId               *string        `json:"port_id,omitempty"`
	PowerDraw            *float64       `json:"power_draw,omitempty"`
	When                 *string        `json:"when,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DiscoveredSwitchAp.
// It customizes the JSON marshaling process for DiscoveredSwitchAp objects.
func (d DiscoveredSwitchAp) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DiscoveredSwitchAp object to a map representation for JSON marshaling.
func (d DiscoveredSwitchAp) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Hostname != nil {
		structMap["hostname"] = d.Hostname
	}
	if d.Mac != nil {
		structMap["mac"] = d.Mac
	}
	if d.PoeStatus != nil {
		structMap["poe_status"] = d.PoeStatus
	}
	if d.Port != nil {
		structMap["port"] = d.Port
	}
	if d.PortId != nil {
		structMap["port_id"] = d.PortId
	}
	if d.PowerDraw != nil {
		structMap["power_draw"] = d.PowerDraw
	}
	if d.When != nil {
		structMap["when"] = d.When
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DiscoveredSwitchAp.
// It customizes the JSON unmarshaling process for DiscoveredSwitchAp objects.
func (d *DiscoveredSwitchAp) UnmarshalJSON(input []byte) error {
	var temp tempDiscoveredSwitchAp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "hostname", "mac", "poe_status", "port", "port_id", "power_draw", "when")
	if err != nil {
		return err
	}

	d.AdditionalProperties = additionalProperties
	d.Hostname = temp.Hostname
	d.Mac = temp.Mac
	d.PoeStatus = temp.PoeStatus
	d.Port = temp.Port
	d.PortId = temp.PortId
	d.PowerDraw = temp.PowerDraw
	d.When = temp.When
	return nil
}

// tempDiscoveredSwitchAp is a temporary struct used for validating the fields of DiscoveredSwitchAp.
type tempDiscoveredSwitchAp struct {
	Hostname  *string  `json:"hostname,omitempty"`
	Mac       *string  `json:"mac,omitempty"`
	PoeStatus *bool    `json:"poe_status,omitempty"`
	Port      *string  `json:"port,omitempty"`
	PortId    *string  `json:"port_id,omitempty"`
	PowerDraw *float64 `json:"power_draw,omitempty"`
	When      *string  `json:"when,omitempty"`
}
