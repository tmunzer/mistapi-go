package models

import (
	"encoding/json"
)

// OrgSettingDeviceCert represents a OrgSettingDeviceCert struct.
// common device cert, optional
type OrgSettingDeviceCert struct {
	Cert                 *string        `json:"cert,omitempty"`
	Key                  *string        `json:"key,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingDeviceCert.
// It customizes the JSON marshaling process for OrgSettingDeviceCert objects.
func (o OrgSettingDeviceCert) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingDeviceCert object to a map representation for JSON marshaling.
func (o OrgSettingDeviceCert) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Cert != nil {
		structMap["cert"] = o.Cert
	}
	if o.Key != nil {
		structMap["key"] = o.Key
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingDeviceCert.
// It customizes the JSON unmarshaling process for OrgSettingDeviceCert objects.
func (o *OrgSettingDeviceCert) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingDeviceCert
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cert", "key")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.Cert = temp.Cert
	o.Key = temp.Key
	return nil
}

// tempOrgSettingDeviceCert is a temporary struct used for validating the fields of OrgSettingDeviceCert.
type tempOrgSettingDeviceCert struct {
	Cert *string `json:"cert,omitempty"`
	Key  *string `json:"key,omitempty"`
}
