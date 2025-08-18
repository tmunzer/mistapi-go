// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UpgradeSiteDevicesCounts represents a UpgradeSiteDevicesCounts struct.
type UpgradeSiteDevicesCounts struct {
	// Count of devices which cloud has requested to download firmware
	DownloadRequested *int `json:"download_requested,omitempty"`
	// Count of ap's which have the firmware downloaded
	Downloaded *int `json:"downloaded,omitempty"`
	// Count of devices which have failed to upgrade
	Failed *int `json:"failed,omitempty"`
	// Count of devices which are rebooting
	RebootInProgress *int `json:"reboot_in_progress,omitempty"`
	// Count of devices which have rebooted successfully
	Rebooted *int `json:"rebooted,omitempty"`
	// Count of devices which cloud has scheduled an upgrade for
	Scheduled *int `json:"scheduled,omitempty"`
	// Count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped *int `json:"skipped,omitempty"`
	// Count of devices part of this upgrade
	Total *int `json:"total,omitempty"`
	// Count of devices which have upgraded successfully
	Upgraded             *int                   `json:"upgraded,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeSiteDevicesCounts,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeSiteDevicesCounts) String() string {
	return fmt.Sprintf(
		"UpgradeSiteDevicesCounts[DownloadRequested=%v, Downloaded=%v, Failed=%v, RebootInProgress=%v, Rebooted=%v, Scheduled=%v, Skipped=%v, Total=%v, Upgraded=%v, AdditionalProperties=%v]",
		u.DownloadRequested, u.Downloaded, u.Failed, u.RebootInProgress, u.Rebooted, u.Scheduled, u.Skipped, u.Total, u.Upgraded, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeSiteDevicesCounts.
// It customizes the JSON marshaling process for UpgradeSiteDevicesCounts objects.
func (u UpgradeSiteDevicesCounts) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"download_requested", "downloaded", "failed", "reboot_in_progress", "rebooted", "scheduled", "skipped", "total", "upgraded"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpgradeSiteDevicesCounts object to a map representation for JSON marshaling.
func (u UpgradeSiteDevicesCounts) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.DownloadRequested != nil {
		structMap["download_requested"] = u.DownloadRequested
	}
	if u.Downloaded != nil {
		structMap["downloaded"] = u.Downloaded
	}
	if u.Failed != nil {
		structMap["failed"] = u.Failed
	}
	if u.RebootInProgress != nil {
		structMap["reboot_in_progress"] = u.RebootInProgress
	}
	if u.Rebooted != nil {
		structMap["rebooted"] = u.Rebooted
	}
	if u.Scheduled != nil {
		structMap["scheduled"] = u.Scheduled
	}
	if u.Skipped != nil {
		structMap["skipped"] = u.Skipped
	}
	if u.Total != nil {
		structMap["total"] = u.Total
	}
	if u.Upgraded != nil {
		structMap["upgraded"] = u.Upgraded
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeSiteDevicesCounts.
// It customizes the JSON unmarshaling process for UpgradeSiteDevicesCounts objects.
func (u *UpgradeSiteDevicesCounts) UnmarshalJSON(input []byte) error {
	var temp tempUpgradeSiteDevicesCounts
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "download_requested", "downloaded", "failed", "reboot_in_progress", "rebooted", "scheduled", "skipped", "total", "upgraded")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.DownloadRequested = temp.DownloadRequested
	u.Downloaded = temp.Downloaded
	u.Failed = temp.Failed
	u.RebootInProgress = temp.RebootInProgress
	u.Rebooted = temp.Rebooted
	u.Scheduled = temp.Scheduled
	u.Skipped = temp.Skipped
	u.Total = temp.Total
	u.Upgraded = temp.Upgraded
	return nil
}

// tempUpgradeSiteDevicesCounts is a temporary struct used for validating the fields of UpgradeSiteDevicesCounts.
type tempUpgradeSiteDevicesCounts struct {
	DownloadRequested *int `json:"download_requested,omitempty"`
	Downloaded        *int `json:"downloaded,omitempty"`
	Failed            *int `json:"failed,omitempty"`
	RebootInProgress  *int `json:"reboot_in_progress,omitempty"`
	Rebooted          *int `json:"rebooted,omitempty"`
	Scheduled         *int `json:"scheduled,omitempty"`
	Skipped           *int `json:"skipped,omitempty"`
	Total             *int `json:"total,omitempty"`
	Upgraded          *int `json:"upgraded,omitempty"`
}
