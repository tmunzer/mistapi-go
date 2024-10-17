package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// OrgSettingInstaller represents a OrgSettingInstaller struct.
type OrgSettingInstaller struct {
	AllowAllDevices      *bool          `json:"allow_all_devices,omitempty"`
	AllowAllSites        *bool          `json:"allow_all_sites,omitempty"`
	ExtraSiteIds         []uuid.UUID    `json:"extra_site_ids,omitempty"`
	GracePeriod          *int           `json:"grace_period,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingInstaller.
// It customizes the JSON marshaling process for OrgSettingInstaller objects.
func (o OrgSettingInstaller) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingInstaller object to a map representation for JSON marshaling.
func (o OrgSettingInstaller) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.AllowAllDevices != nil {
		structMap["allow_all_devices"] = o.AllowAllDevices
	}
	if o.AllowAllSites != nil {
		structMap["allow_all_sites"] = o.AllowAllSites
	}
	if o.ExtraSiteIds != nil {
		structMap["extra_site_ids"] = o.ExtraSiteIds
	}
	if o.GracePeriod != nil {
		structMap["grace_period"] = o.GracePeriod
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingInstaller.
// It customizes the JSON unmarshaling process for OrgSettingInstaller objects.
func (o *OrgSettingInstaller) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingInstaller
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_all_devices", "allow_all_sites", "extra_site_ids", "grace_period")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.AllowAllDevices = temp.AllowAllDevices
	o.AllowAllSites = temp.AllowAllSites
	o.ExtraSiteIds = temp.ExtraSiteIds
	o.GracePeriod = temp.GracePeriod
	return nil
}

// tempOrgSettingInstaller is a temporary struct used for validating the fields of OrgSettingInstaller.
type tempOrgSettingInstaller struct {
	AllowAllDevices *bool       `json:"allow_all_devices,omitempty"`
	AllowAllSites   *bool       `json:"allow_all_sites,omitempty"`
	ExtraSiteIds    []uuid.UUID `json:"extra_site_ids,omitempty"`
	GracePeriod     *int        `json:"grace_period,omitempty"`
}
