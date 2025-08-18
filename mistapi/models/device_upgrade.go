// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DeviceUpgrade represents a DeviceUpgrade struct.
type DeviceUpgrade struct {
	// For Switches and Gateways only (APs are automatically rebooted). Reboot device immediately after upgrade is completed
	Reboot *bool `json:"reboot,omitempty"`
	// For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time`
	RebootAt *int `json:"reboot_at,omitempty"`
	// For Junos devices only. Perform recovery snapshot after device is rebooted
	Snapshot *bool `json:"snapshot,omitempty"`
	// Firmware download start time in epoch
	StartTime *int `json:"start_time,omitempty"`
	// Specific version / `stable`, default is to use the latest
	Version              string                 `json:"version"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceUpgrade) String() string {
	return fmt.Sprintf(
		"DeviceUpgrade[Reboot=%v, RebootAt=%v, Snapshot=%v, StartTime=%v, Version=%v, AdditionalProperties=%v]",
		d.Reboot, d.RebootAt, d.Snapshot, d.StartTime, d.Version, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceUpgrade.
// It customizes the JSON marshaling process for DeviceUpgrade objects.
func (d DeviceUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"reboot", "reboot_at", "snapshot", "start_time", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceUpgrade object to a map representation for JSON marshaling.
func (d DeviceUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Reboot != nil {
		structMap["reboot"] = d.Reboot
	}
	if d.RebootAt != nil {
		structMap["reboot_at"] = d.RebootAt
	}
	if d.Snapshot != nil {
		structMap["snapshot"] = d.Snapshot
	}
	if d.StartTime != nil {
		structMap["start_time"] = d.StartTime
	}
	structMap["version"] = d.Version
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceUpgrade.
// It customizes the JSON unmarshaling process for DeviceUpgrade objects.
func (d *DeviceUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempDeviceUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reboot", "reboot_at", "snapshot", "start_time", "version")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Reboot = temp.Reboot
	d.RebootAt = temp.RebootAt
	d.Snapshot = temp.Snapshot
	d.StartTime = temp.StartTime
	d.Version = *temp.Version
	return nil
}

// tempDeviceUpgrade is a temporary struct used for validating the fields of DeviceUpgrade.
type tempDeviceUpgrade struct {
	Reboot    *bool   `json:"reboot,omitempty"`
	RebootAt  *int    `json:"reboot_at,omitempty"`
	Snapshot  *bool   `json:"snapshot,omitempty"`
	StartTime *int    `json:"start_time,omitempty"`
	Version   *string `json:"version"`
}

func (d *tempDeviceUpgrade) validate() error {
	var errs []string
	if d.Version == nil {
		errs = append(errs, "required field `version` is missing for type `device_upgrade`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
