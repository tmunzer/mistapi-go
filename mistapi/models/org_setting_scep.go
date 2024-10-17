package models

import (
	"encoding/json"
)

// OrgSettingScep represents a OrgSettingScep struct.
type OrgSettingScep struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingScep.
// It customizes the JSON marshaling process for OrgSettingScep objects.
func (o OrgSettingScep) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingScep object to a map representation for JSON marshaling.
func (o OrgSettingScep) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingScep.
// It customizes the JSON unmarshaling process for OrgSettingScep objects.
func (o *OrgSettingScep) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingScep
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.Enabled = temp.Enabled
	return nil
}

// tempOrgSettingScep is a temporary struct used for validating the fields of OrgSettingScep.
type tempOrgSettingScep struct {
	Enabled *bool `json:"enabled,omitempty"`
}
