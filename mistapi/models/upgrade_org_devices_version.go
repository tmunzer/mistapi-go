// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UpgradeOrgDevicesVersion represents a UpgradeOrgDevicesVersion struct.
type UpgradeOrgDevicesVersion struct {
	// enum: `ap`, `junos`
	FirmwareType *UpgradeOrgDevicesVersionFirmwareTypeEnum `json:"firmware_type,omitempty"`
	// If `firmware_type`==`ap`, set to `true` if upgrade is needed when target version <= running version
	Force *bool `json:"force,omitempty"`
	// If `firmware_type`==`junos`, used to select different versions for different models (Overrides `version` for the specified models). Property key is the hadware model (e.g. `EX4400-24MP`), Property value is the firmware version (e.g. `23.4R1.9`)
	ModelVersion map[string]string `json:"model_version,omitempty"`
	// version of the firmware to deploy
	Version              *string                `json:"version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevicesVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevicesVersion) String() string {
	return fmt.Sprintf(
		"UpgradeOrgDevicesVersion[FirmwareType=%v, Force=%v, ModelVersion=%v, Version=%v, AdditionalProperties=%v]",
		u.FirmwareType, u.Force, u.ModelVersion, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevicesVersion.
// It customizes the JSON marshaling process for UpgradeOrgDevicesVersion objects.
func (u UpgradeOrgDevicesVersion) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"firmware_type", "force", "model_version", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevicesVersion object to a map representation for JSON marshaling.
func (u UpgradeOrgDevicesVersion) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.FirmwareType != nil {
		structMap["firmware_type"] = u.FirmwareType
	}
	if u.Force != nil {
		structMap["force"] = u.Force
	}
	if u.ModelVersion != nil {
		structMap["model_version"] = u.ModelVersion
	}
	if u.Version != nil {
		structMap["version"] = u.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevicesVersion.
// It customizes the JSON unmarshaling process for UpgradeOrgDevicesVersion objects.
func (u *UpgradeOrgDevicesVersion) UnmarshalJSON(input []byte) error {
	var temp tempUpgradeOrgDevicesVersion
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "firmware_type", "force", "model_version", "version")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.FirmwareType = temp.FirmwareType
	u.Force = temp.Force
	u.ModelVersion = temp.ModelVersion
	u.Version = temp.Version
	return nil
}

// tempUpgradeOrgDevicesVersion is a temporary struct used for validating the fields of UpgradeOrgDevicesVersion.
type tempUpgradeOrgDevicesVersion struct {
	FirmwareType *UpgradeOrgDevicesVersionFirmwareTypeEnum `json:"firmware_type,omitempty"`
	Force        *bool                                     `json:"force,omitempty"`
	ModelVersion map[string]string                         `json:"model_version,omitempty"`
	Version      *string                                   `json:"version,omitempty"`
}
