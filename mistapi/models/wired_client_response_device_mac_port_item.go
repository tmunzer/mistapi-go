package models

import (
	"encoding/json"
)

// WiredClientResponseDeviceMacPortItem represents a WiredClientResponseDeviceMacPortItem struct.
type WiredClientResponseDeviceMacPortItem struct {
	DeviceMac            *string        `json:"device_mac,omitempty"`
	Ip                   *string        `json:"ip,omitempty"`
	PortId               *string        `json:"port_id,omitempty"`
	PortParent           *string        `json:"port_parent,omitempty"`
	Start                *string        `json:"start,omitempty"`
	Vlan                 *int           `json:"vlan,omitempty"`
	When                 *string        `json:"when,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WiredClientResponseDeviceMacPortItem.
// It customizes the JSON marshaling process for WiredClientResponseDeviceMacPortItem objects.
func (w WiredClientResponseDeviceMacPortItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WiredClientResponseDeviceMacPortItem object to a map representation for JSON marshaling.
func (w WiredClientResponseDeviceMacPortItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.DeviceMac != nil {
		structMap["device_mac"] = w.DeviceMac
	}
	if w.Ip != nil {
		structMap["ip"] = w.Ip
	}
	if w.PortId != nil {
		structMap["port_id"] = w.PortId
	}
	if w.PortParent != nil {
		structMap["port_parent"] = w.PortParent
	}
	if w.Start != nil {
		structMap["start"] = w.Start
	}
	if w.Vlan != nil {
		structMap["vlan"] = w.Vlan
	}
	if w.When != nil {
		structMap["when"] = w.When
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WiredClientResponseDeviceMacPortItem.
// It customizes the JSON unmarshaling process for WiredClientResponseDeviceMacPortItem objects.
func (w *WiredClientResponseDeviceMacPortItem) UnmarshalJSON(input []byte) error {
	var temp tempWiredClientResponseDeviceMacPortItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "ip", "port_id", "port_parent", "start", "vlan", "when")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.DeviceMac = temp.DeviceMac
	w.Ip = temp.Ip
	w.PortId = temp.PortId
	w.PortParent = temp.PortParent
	w.Start = temp.Start
	w.Vlan = temp.Vlan
	w.When = temp.When
	return nil
}

// tempWiredClientResponseDeviceMacPortItem is a temporary struct used for validating the fields of WiredClientResponseDeviceMacPortItem.
type tempWiredClientResponseDeviceMacPortItem struct {
	DeviceMac  *string `json:"device_mac,omitempty"`
	Ip         *string `json:"ip,omitempty"`
	PortId     *string `json:"port_id,omitempty"`
	PortParent *string `json:"port_parent,omitempty"`
	Start      *string `json:"start,omitempty"`
	Vlan       *int    `json:"vlan,omitempty"`
	When       *string `json:"when,omitempty"`
}
