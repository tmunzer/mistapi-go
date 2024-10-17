package models

import (
	"encoding/json"
)

// OrgSettingCelona represents a OrgSettingCelona struct.
type OrgSettingCelona struct {
	ApiKey               *string        `json:"api_key,omitempty"`
	ApiPrefix            *string        `json:"api_prefix,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingCelona.
// It customizes the JSON marshaling process for OrgSettingCelona objects.
func (o OrgSettingCelona) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingCelona object to a map representation for JSON marshaling.
func (o OrgSettingCelona) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.ApiKey != nil {
		structMap["api_key"] = o.ApiKey
	}
	if o.ApiPrefix != nil {
		structMap["api_prefix"] = o.ApiPrefix
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingCelona.
// It customizes the JSON unmarshaling process for OrgSettingCelona objects.
func (o *OrgSettingCelona) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingCelona
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "api_key", "api_prefix")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.ApiKey = temp.ApiKey
	o.ApiPrefix = temp.ApiPrefix
	return nil
}

// tempOrgSettingCelona is a temporary struct used for validating the fields of OrgSettingCelona.
type tempOrgSettingCelona struct {
	ApiKey    *string `json:"api_key,omitempty"`
	ApiPrefix *string `json:"api_prefix,omitempty"`
}
