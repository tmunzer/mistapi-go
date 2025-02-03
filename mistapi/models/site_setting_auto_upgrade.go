package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingAutoUpgrade represents a SiteSettingAutoUpgrade struct.
// Auto Upgrade Settings
type SiteSettingAutoUpgrade struct {
    // Custom versions for different models. Property key is the model name (e.g. "AP41")
    CustomVersions       map[string]string           `json:"custom_versions,omitempty"`
    // enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
    DayOfWeek            *DayOfWeekEnum              `json:"day_of_week,omitempty"`
    // Whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported)
    Enabled              *bool                       `json:"enabled,omitempty"`
    // `any` / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time
    TimeOfDay            *string                     `json:"time_of_day,omitempty"`
    // desired version. enum: `beta`, `custom`, `stable`
    Version              *SiteAutoUpgradeVersionEnum `json:"version,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingAutoUpgrade) String() string {
    return fmt.Sprintf(
    	"SiteSettingAutoUpgrade[CustomVersions=%v, DayOfWeek=%v, Enabled=%v, TimeOfDay=%v, Version=%v, AdditionalProperties=%v]",
    	s.CustomVersions, s.DayOfWeek, s.Enabled, s.TimeOfDay, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingAutoUpgrade.
// It customizes the JSON marshaling process for SiteSettingAutoUpgrade objects.
func (s SiteSettingAutoUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "custom_versions", "day_of_week", "enabled", "time_of_day", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingAutoUpgrade object to a map representation for JSON marshaling.
func (s SiteSettingAutoUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomVersions != nil {
        structMap["custom_versions"] = s.CustomVersions
    }
    if s.DayOfWeek != nil {
        structMap["day_of_week"] = s.DayOfWeek
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.TimeOfDay != nil {
        structMap["time_of_day"] = s.TimeOfDay
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingAutoUpgrade.
// It customizes the JSON unmarshaling process for SiteSettingAutoUpgrade objects.
func (s *SiteSettingAutoUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingAutoUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_versions", "day_of_week", "enabled", "time_of_day", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CustomVersions = temp.CustomVersions
    s.DayOfWeek = temp.DayOfWeek
    s.Enabled = temp.Enabled
    s.TimeOfDay = temp.TimeOfDay
    s.Version = temp.Version
    return nil
}

// tempSiteSettingAutoUpgrade is a temporary struct used for validating the fields of SiteSettingAutoUpgrade.
type tempSiteSettingAutoUpgrade  struct {
    CustomVersions map[string]string           `json:"custom_versions,omitempty"`
    DayOfWeek      *DayOfWeekEnum              `json:"day_of_week,omitempty"`
    Enabled        *bool                       `json:"enabled,omitempty"`
    TimeOfDay      *string                     `json:"time_of_day,omitempty"`
    Version        *SiteAutoUpgradeVersionEnum `json:"version,omitempty"`
}
