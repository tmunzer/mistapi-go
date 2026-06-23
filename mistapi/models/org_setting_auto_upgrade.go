// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingAutoUpgrade represents a OrgSettingAutoUpgrade struct.
// Organization-wide AP automatic firmware upgrade policy
type OrgSettingAutoUpgrade struct {
	// Per-AP-model firmware versions or channels used for auto-upgrade. Property key is the AP model name (e.g. "AP41"), value is the firmware version or release channel (e.g. `stable`)
	CustomVersions map[string]string `json:"custom_versions,omitempty"`
	// enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
	DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
	// Whether AP auto-upgrade is enabled. Note that Mist may auto-upgrade APs if the running version is no longer supported.
	Enabled *bool `json:"enabled,omitempty"`
	// `any` or HH:MM (24-hour format). Upgrade will happen within up to 1 hour from this time.
	TimeOfDay *string `json:"time_of_day,omitempty"`
	// desired version. enum: `beta`, `custom`, `stable`
	Version              *SiteAutoUpgradeVersionEnum `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingAutoUpgrade) String() string {
	return fmt.Sprintf(
		"OrgSettingAutoUpgrade[CustomVersions=%v, DayOfWeek=%v, Enabled=%v, TimeOfDay=%v, Version=%v, AdditionalProperties=%v]",
		o.CustomVersions, o.DayOfWeek, o.Enabled, o.TimeOfDay, o.Version, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoUpgrade.
// It customizes the JSON marshaling process for OrgSettingAutoUpgrade objects.
func (o OrgSettingAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"custom_versions", "day_of_week", "enabled", "time_of_day", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoUpgrade object to a map representation for JSON marshaling.
func (o OrgSettingAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.CustomVersions != nil {
		structMap["custom_versions"] = o.CustomVersions
	}
	if o.DayOfWeek != nil {
		structMap["day_of_week"] = o.DayOfWeek
	}
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	if o.TimeOfDay != nil {
		structMap["time_of_day"] = o.TimeOfDay
	}
	if o.Version != nil {
		structMap["version"] = o.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingAutoUpgrade.
// It customizes the JSON unmarshaling process for OrgSettingAutoUpgrade objects.
func (o *OrgSettingAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_versions", "day_of_week", "enabled", "time_of_day", "version")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.CustomVersions = temp.CustomVersions
	o.DayOfWeek = temp.DayOfWeek
	o.Enabled = temp.Enabled
	o.TimeOfDay = temp.TimeOfDay
	o.Version = temp.Version
	return nil
}

// tempOrgSettingAutoUpgrade is a temporary struct used for validating the fields of OrgSettingAutoUpgrade.
type tempOrgSettingAutoUpgrade struct {
	CustomVersions map[string]string           `json:"custom_versions,omitempty"`
	DayOfWeek      *DayOfWeekEnum              `json:"day_of_week,omitempty"`
	Enabled        *bool                       `json:"enabled,omitempty"`
	TimeOfDay      *string                     `json:"time_of_day,omitempty"`
	Version        *SiteAutoUpgradeVersionEnum `json:"version,omitempty"`
}
