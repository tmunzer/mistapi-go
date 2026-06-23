// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApZigbee represents a ApZigbee struct.
// Zigbee radio and network settings applied to an AP or AP profile
type ApZigbee struct {
	// Controls whether new Zigbee devices are allowed to join the network. enum: `always`, `manual`
	AllowJoin *ApZigbeeAllowJoinEnum `json:"allow_join,omitempty"`
	// Zigbee channel (2.4 GHz). `0` means auto; valid fixed values are 11–26
	Channel *int `json:"channel,omitempty"`
	// Whether to enable Zigbee on this AP
	Enabled *bool `json:"enabled,omitempty"`
	// Extended PAN ID in hex string format; only applicable when `pan_id` is also specified
	ExtendedPanId Optional[string] `json:"extended_pan_id"`
	// PAN ID in hex string format; if not specified, assigned automatically
	PanId                Optional[string]       `json:"pan_id"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApZigbee,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApZigbee) String() string {
	return fmt.Sprintf(
		"ApZigbee[AllowJoin=%v, Channel=%v, Enabled=%v, ExtendedPanId=%v, PanId=%v, AdditionalProperties=%v]",
		a.AllowJoin, a.Channel, a.Enabled, a.ExtendedPanId, a.PanId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApZigbee.
// It customizes the JSON marshaling process for ApZigbee objects.
func (a ApZigbee) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"allow_join", "channel", "enabled", "extended_pan_id", "pan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApZigbee object to a map representation for JSON marshaling.
func (a ApZigbee) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AllowJoin != nil {
		structMap["allow_join"] = a.AllowJoin
	}
	if a.Channel != nil {
		structMap["channel"] = a.Channel
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.ExtendedPanId.IsValueSet() {
		if a.ExtendedPanId.Value() != nil {
			structMap["extended_pan_id"] = a.ExtendedPanId.Value()
		} else {
			structMap["extended_pan_id"] = nil
		}
	}
	if a.PanId.IsValueSet() {
		if a.PanId.Value() != nil {
			structMap["pan_id"] = a.PanId.Value()
		} else {
			structMap["pan_id"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApZigbee.
// It customizes the JSON unmarshaling process for ApZigbee objects.
func (a *ApZigbee) UnmarshalJSON(input []byte) error {
	var temp tempApZigbee
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_join", "channel", "enabled", "extended_pan_id", "pan_id")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.AllowJoin = temp.AllowJoin
	a.Channel = temp.Channel
	a.Enabled = temp.Enabled
	a.ExtendedPanId = temp.ExtendedPanId
	a.PanId = temp.PanId
	return nil
}

// tempApZigbee is a temporary struct used for validating the fields of ApZigbee.
type tempApZigbee struct {
	AllowJoin     *ApZigbeeAllowJoinEnum `json:"allow_join,omitempty"`
	Channel       *int                   `json:"channel,omitempty"`
	Enabled       *bool                  `json:"enabled,omitempty"`
	ExtendedPanId Optional[string]       `json:"extended_pan_id"`
	PanId         Optional[string]       `json:"pan_id"`
}
