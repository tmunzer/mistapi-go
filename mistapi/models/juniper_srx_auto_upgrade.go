// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JuniperSrxAutoUpgrade represents a JuniperSrxAutoUpgrade struct.
// auto_upgrade device first time it is onboarded
type JuniperSrxAutoUpgrade struct {
	// Property key is the SRX Hardware model (e.g. "SRX4600")
	CustomVersions       map[string]string      `json:"custom_versions,omitempty"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	Snapshot             *bool                  `json:"snapshot,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JuniperSrxAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JuniperSrxAutoUpgrade) String() string {
	return fmt.Sprintf(
		"JuniperSrxAutoUpgrade[CustomVersions=%v, Enabled=%v, Snapshot=%v, AdditionalProperties=%v]",
		j.CustomVersions, j.Enabled, j.Snapshot, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JuniperSrxAutoUpgrade.
// It customizes the JSON marshaling process for JuniperSrxAutoUpgrade objects.
func (j JuniperSrxAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"custom_versions", "enabled", "snapshot"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JuniperSrxAutoUpgrade object to a map representation for JSON marshaling.
func (j JuniperSrxAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.CustomVersions != nil {
		structMap["custom_versions"] = j.CustomVersions
	}
	if j.Enabled != nil {
		structMap["enabled"] = j.Enabled
	}
	if j.Snapshot != nil {
		structMap["snapshot"] = j.Snapshot
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JuniperSrxAutoUpgrade.
// It customizes the JSON unmarshaling process for JuniperSrxAutoUpgrade objects.
func (j *JuniperSrxAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempJuniperSrxAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_versions", "enabled", "snapshot")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.CustomVersions = temp.CustomVersions
	j.Enabled = temp.Enabled
	j.Snapshot = temp.Snapshot
	return nil
}

// tempJuniperSrxAutoUpgrade is a temporary struct used for validating the fields of JuniperSrxAutoUpgrade.
type tempJuniperSrxAutoUpgrade struct {
	CustomVersions map[string]string `json:"custom_versions,omitempty"`
	Enabled        *bool             `json:"enabled,omitempty"`
	Snapshot       *bool             `json:"snapshot,omitempty"`
}
