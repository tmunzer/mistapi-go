package models

import (
    "encoding/json"
)

// UpgradeOrgDeviceTargets represents a UpgradeOrgDeviceTargets struct.
type UpgradeOrgDeviceTargets struct {
    // list of devices MAC Addresses which cloud has requested to download firmware
    DownloadRequested    []string       `json:"download_requested,omitempty"`
    // list of devices MAC Addresses which have the firmware downloaded
    Downloaded           []string       `json:"downloaded,omitempty"`
    // list of devices MAC Addresses which have failed to upgrade
    Failed               []string       `json:"failed,omitempty"`
    // list of devices MAC Addresses which are rebooting
    RebootInProgress     []string       `json:"reboot_in_progress,omitempty"`
    // list of devices MAC Addresses which have rebooted successfully
    Rebooted             []string       `json:"rebooted,omitempty"`
    // list of devices MAC Addresses which cloud has scheduled an upgrade for
    Scheduled            []string       `json:"scheduled,omitempty"`
    // list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade
    Skipped              []string       `json:"skipped,omitempty"`
    // count of devices part of this upgrade
    Total                *int           `json:"total,omitempty"`
    // count of devices which have upgraded successfully
    Upgraded             []string       `json:"upgraded,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDeviceTargets.
// It customizes the JSON marshaling process for UpgradeOrgDeviceTargets objects.
func (u UpgradeOrgDeviceTargets) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDeviceTargets object to a map representation for JSON marshaling.
func (u UpgradeOrgDeviceTargets) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDeviceTargets.
// It customizes the JSON unmarshaling process for UpgradeOrgDeviceTargets objects.
func (u *UpgradeOrgDeviceTargets) UnmarshalJSON(input []byte) error {
    var temp upgradeOrgDeviceTargets
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "download_requested", "downloaded", "failed", "reboot_in_progress", "rebooted", "scheduled", "skipped", "total", "upgraded")
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

// upgradeOrgDeviceTargets is a temporary struct used for validating the fields of UpgradeOrgDeviceTargets.
type upgradeOrgDeviceTargets  struct {
    DownloadRequested []string `json:"download_requested,omitempty"`
    Downloaded        []string `json:"downloaded,omitempty"`
    Failed            []string `json:"failed,omitempty"`
    RebootInProgress  []string `json:"reboot_in_progress,omitempty"`
    Rebooted          []string `json:"rebooted,omitempty"`
    Scheduled         []string `json:"scheduled,omitempty"`
    Skipped           []string `json:"skipped,omitempty"`
    Total             *int     `json:"total,omitempty"`
    Upgraded          []string `json:"upgraded,omitempty"`
}
