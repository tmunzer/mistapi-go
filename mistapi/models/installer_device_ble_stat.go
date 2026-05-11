// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// InstallerDeviceBleStat represents a InstallerDeviceBleStat struct.
// BLE statistics for the device
type InstallerDeviceBleStat struct {
	Major                *int                   `json:"major,omitempty"`
	Minors               []int                  `json:"minors,omitempty"`
	Uuid                 *uuid.UUID             `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InstallerDeviceBleStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InstallerDeviceBleStat) String() string {
	return fmt.Sprintf(
		"InstallerDeviceBleStat[Major=%v, Minors=%v, Uuid=%v, AdditionalProperties=%v]",
		i.Major, i.Minors, i.Uuid, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InstallerDeviceBleStat.
// It customizes the JSON marshaling process for InstallerDeviceBleStat objects.
func (i InstallerDeviceBleStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"major", "minors", "uuid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InstallerDeviceBleStat object to a map representation for JSON marshaling.
func (i InstallerDeviceBleStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Major != nil {
		structMap["major"] = i.Major
	}
	if i.Minors != nil {
		structMap["minors"] = i.Minors
	}
	if i.Uuid != nil {
		structMap["uuid"] = i.Uuid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InstallerDeviceBleStat.
// It customizes the JSON unmarshaling process for InstallerDeviceBleStat objects.
func (i *InstallerDeviceBleStat) UnmarshalJSON(input []byte) error {
	var temp tempInstallerDeviceBleStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "major", "minors", "uuid")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.Major = temp.Major
	i.Minors = temp.Minors
	i.Uuid = temp.Uuid
	return nil
}

// tempInstallerDeviceBleStat is a temporary struct used for validating the fields of InstallerDeviceBleStat.
type tempInstallerDeviceBleStat struct {
	Major  *int       `json:"major,omitempty"`
	Minors []int      `json:"minors,omitempty"`
	Uuid   *uuid.UUID `json:"uuid,omitempty"`
}
