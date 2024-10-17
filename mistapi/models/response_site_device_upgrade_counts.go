package models

import (
	"encoding/json"
)

// ResponseSiteDeviceUpgradeCounts represents a ResponseSiteDeviceUpgradeCounts struct.
type ResponseSiteDeviceUpgradeCounts struct {
	// count of devices which cloud has requested to download firmware
	DownloadRequested *int `json:"download_requested,omitempty"`
	// count of ap’s which have the firmware downloaded
	Downloaded *int `json:"downloaded,omitempty"`
	// count of devices which have failed to upgrade
	Failed *int `json:"failed,omitempty"`
	// count of devices which are rebooting
	RebootInProgress *int `json:"reboot_in_progress,omitempty"`
	// count of devices which have rebooted successfully
	Rebooted *int `json:"rebooted,omitempty"`
	// count of devices which cloud has scheduled an upgrade for
	Scheduled *int `json:"scheduled,omitempty"`
	// count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped *int `json:"skipped,omitempty"`
	// count of devices part of this upgrade
	Total *int `json:"total,omitempty"`
	// count of devices which have upgraded successfully
	Upgraded             *int           `json:"upgraded,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSiteDeviceUpgradeCounts.
// It customizes the JSON marshaling process for ResponseSiteDeviceUpgradeCounts objects.
func (r ResponseSiteDeviceUpgradeCounts) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSiteDeviceUpgradeCounts object to a map representation for JSON marshaling.
func (r ResponseSiteDeviceUpgradeCounts) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSiteDeviceUpgradeCounts.
// It customizes the JSON unmarshaling process for ResponseSiteDeviceUpgradeCounts objects.
func (r *ResponseSiteDeviceUpgradeCounts) UnmarshalJSON(input []byte) error {
	var temp tempResponseSiteDeviceUpgradeCounts
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

// tempResponseSiteDeviceUpgradeCounts is a temporary struct used for validating the fields of ResponseSiteDeviceUpgradeCounts.
type tempResponseSiteDeviceUpgradeCounts struct {
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
