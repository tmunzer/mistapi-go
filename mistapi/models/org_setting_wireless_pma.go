package models

import (
	"encoding/json"
)

// OrgSettingWirelessPma represents a OrgSettingWirelessPma struct.
type OrgSettingWirelessPma struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingWirelessPma.
// It customizes the JSON marshaling process for OrgSettingWirelessPma objects.
func (o OrgSettingWirelessPma) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingWirelessPma object to a map representation for JSON marshaling.
func (o OrgSettingWirelessPma) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingWirelessPma.
// It customizes the JSON unmarshaling process for OrgSettingWirelessPma objects.
func (o *OrgSettingWirelessPma) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingWirelessPma
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

// tempOrgSettingWirelessPma is a temporary struct used for validating the fields of OrgSettingWirelessPma.
type tempOrgSettingWirelessPma struct {
	Enabled *bool `json:"enabled,omitempty"`
}
