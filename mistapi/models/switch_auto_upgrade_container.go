// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchAutoUpgradeContainer represents a SwitchAutoUpgradeContainer struct.
type SwitchAutoUpgradeContainer struct {
	AutoUpgrade          *SwitchAutoUpgrade     `json:"auto_upgrade,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchAutoUpgradeContainer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchAutoUpgradeContainer) String() string {
	return fmt.Sprintf(
		"SwitchAutoUpgradeContainer[AutoUpgrade=%v, AdditionalProperties=%v]",
		s.AutoUpgrade, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchAutoUpgradeContainer.
// It customizes the JSON marshaling process for SwitchAutoUpgradeContainer objects.
func (s SwitchAutoUpgradeContainer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auto_upgrade"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchAutoUpgradeContainer object to a map representation for JSON marshaling.
func (s SwitchAutoUpgradeContainer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AutoUpgrade != nil {
		structMap["auto_upgrade"] = s.AutoUpgrade.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchAutoUpgradeContainer.
// It customizes the JSON unmarshaling process for SwitchAutoUpgradeContainer objects.
func (s *SwitchAutoUpgradeContainer) UnmarshalJSON(input []byte) error {
	var temp tempSwitchAutoUpgradeContainer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AutoUpgrade = temp.AutoUpgrade
	return nil
}

// tempSwitchAutoUpgradeContainer is a temporary struct used for validating the fields of SwitchAutoUpgradeContainer.
type tempSwitchAutoUpgradeContainer struct {
	AutoUpgrade *SwitchAutoUpgrade `json:"auto_upgrade,omitempty"`
}
