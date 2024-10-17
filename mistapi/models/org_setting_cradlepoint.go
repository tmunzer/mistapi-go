package models

import (
	"encoding/json"
)

// OrgSettingCradlepoint represents a OrgSettingCradlepoint struct.
type OrgSettingCradlepoint struct {
	CpApiId              *string        `json:"cp_api_id,omitempty"`
	CpApiKey             *string        `json:"cp_api_key,omitempty"`
	EcmApiId             *string        `json:"ecm_api_id,omitempty"`
	EcmApiKey            *string        `json:"ecm_api_key,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingCradlepoint.
// It customizes the JSON marshaling process for OrgSettingCradlepoint objects.
func (o OrgSettingCradlepoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingCradlepoint object to a map representation for JSON marshaling.
func (o OrgSettingCradlepoint) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.CpApiId != nil {
		structMap["cp_api_id"] = o.CpApiId
	}
	if o.CpApiKey != nil {
		structMap["cp_api_key"] = o.CpApiKey
	}
	if o.EcmApiId != nil {
		structMap["ecm_api_id"] = o.EcmApiId
	}
	if o.EcmApiKey != nil {
		structMap["ecm_api_key"] = o.EcmApiKey
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingCradlepoint.
// It customizes the JSON unmarshaling process for OrgSettingCradlepoint objects.
func (o *OrgSettingCradlepoint) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingCradlepoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.CpApiId = temp.CpApiId
	o.CpApiKey = temp.CpApiKey
	o.EcmApiId = temp.EcmApiId
	o.EcmApiKey = temp.EcmApiKey
	return nil
}

// tempOrgSettingCradlepoint is a temporary struct used for validating the fields of OrgSettingCradlepoint.
type tempOrgSettingCradlepoint struct {
	CpApiId   *string `json:"cp_api_id,omitempty"`
	CpApiKey  *string `json:"cp_api_key,omitempty"`
	EcmApiId  *string `json:"ecm_api_id,omitempty"`
	EcmApiKey *string `json:"ecm_api_key,omitempty"`
}
