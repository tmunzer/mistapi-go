// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WlanBonjour represents a WlanBonjour struct.
// Bonjour gateway wlan settings
type WlanBonjour struct {
	// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
	AdditionalVlanIds *AdditionalVlanIds `json:"additional_vlan_ids,omitempty"`
	// Whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false
	Enabled *bool `json:"enabled,omitempty"`
	// What services are allowed.
	// Property key is the service name
	Services             map[string]WlanBonjourServiceProperties `json:"services,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for WlanBonjour,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanBonjour) String() string {
	return fmt.Sprintf(
		"WlanBonjour[AdditionalVlanIds=%v, Enabled=%v, Services=%v, AdditionalProperties=%v]",
		w.AdditionalVlanIds, w.Enabled, w.Services, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanBonjour.
// It customizes the JSON marshaling process for WlanBonjour objects.
func (w WlanBonjour) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"additional_vlan_ids", "enabled", "services"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WlanBonjour object to a map representation for JSON marshaling.
func (w WlanBonjour) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.AdditionalVlanIds != nil {
		structMap["additional_vlan_ids"] = w.AdditionalVlanIds.toMap()
	}
	if w.Enabled != nil {
		structMap["enabled"] = w.Enabled
	}
	if w.Services != nil {
		structMap["services"] = w.Services
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanBonjour.
// It customizes the JSON unmarshaling process for WlanBonjour objects.
func (w *WlanBonjour) UnmarshalJSON(input []byte) error {
	var temp tempWlanBonjour
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_vlan_ids", "enabled", "services")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.AdditionalVlanIds = temp.AdditionalVlanIds
	w.Enabled = temp.Enabled
	w.Services = temp.Services
	return nil
}

// tempWlanBonjour is a temporary struct used for validating the fields of WlanBonjour.
type tempWlanBonjour struct {
	AdditionalVlanIds *AdditionalVlanIds                      `json:"additional_vlan_ids,omitempty"`
	Enabled           *bool                                   `json:"enabled,omitempty"`
	Services          map[string]WlanBonjourServiceProperties `json:"services,omitempty"`
}
