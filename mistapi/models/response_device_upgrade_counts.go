package models

import (
	"encoding/json"
)

// ResponseDeviceUpgradeCounts represents a ResponseDeviceUpgradeCounts struct.
type ResponseDeviceUpgradeCounts struct {
	// list of devices MAC Addresses which cloud has requested to download firmware
	DownloadRequested []string `json:"download_requested,omitempty"`
	// list of devices MAC Addresses which have the firmware downloaded
	Downloaded []string `json:"downloaded,omitempty"`
	// list of devices MAC Addresses which have failed to upgrade
	Failed []string `json:"failed,omitempty"`
	// list of devices MAC Addresses which are rebooting
	RebootInProgress []string `json:"reboot_in_progress,omitempty"`
	// list of devices MAC Addresses which have rebooted successfully
	Rebooted []string `json:"rebooted,omitempty"`
	// list of devices MAC Addresses which cloud has scheduled an upgrade for
	Scheduled []string `json:"scheduled,omitempty"`
	// list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped []string `json:"skipped,omitempty"`
	// count of devices part of this upgrade
	Total *int `json:"total,omitempty"`
	// count of devices which have upgraded successfully
	Upgraded             []string       `json:"upgraded,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceUpgradeCounts.
// It customizes the JSON marshaling process for ResponseDeviceUpgradeCounts objects.
func (r ResponseDeviceUpgradeCounts) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceUpgradeCounts object to a map representation for JSON marshaling.
func (r ResponseDeviceUpgradeCounts) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.DownloadRequested != nil {
		structMap["download_requested"] = r.DownloadRequested
	}
	if r.Downloaded != nil {
		structMap["downloaded"] = r.Downloaded
	}
	if r.Failed != nil {
		structMap["failed"] = r.Failed
	}
	if r.RebootInProgress != nil {
		structMap["reboot_in_progress"] = r.RebootInProgress
	}
	if r.Rebooted != nil {
		structMap["rebooted"] = r.Rebooted
	}
	if r.Scheduled != nil {
		structMap["scheduled"] = r.Scheduled
	}
	if r.Skipped != nil {
		structMap["skipped"] = r.Skipped
	}
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	if r.Upgraded != nil {
		structMap["upgraded"] = r.Upgraded
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceUpgradeCounts.
// It customizes the JSON unmarshaling process for ResponseDeviceUpgradeCounts objects.
func (r *ResponseDeviceUpgradeCounts) UnmarshalJSON(input []byte) error {
	var temp tempResponseDeviceUpgradeCounts
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "download_requested", "downloaded", "failed", "reboot_in_progress", "rebooted", "scheduled", "skipped", "total", "upgraded")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.DownloadRequested = temp.DownloadRequested
	r.Downloaded = temp.Downloaded
	r.Failed = temp.Failed
	r.RebootInProgress = temp.RebootInProgress
	r.Rebooted = temp.Rebooted
	r.Scheduled = temp.Scheduled
	r.Skipped = temp.Skipped
	r.Total = temp.Total
	r.Upgraded = temp.Upgraded
	return nil
}

// tempResponseDeviceUpgradeCounts is a temporary struct used for validating the fields of ResponseDeviceUpgradeCounts.
type tempResponseDeviceUpgradeCounts struct {
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
