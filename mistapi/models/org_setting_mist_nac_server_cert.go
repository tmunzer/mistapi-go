package models

import (
	"encoding/json"
)

// OrgSettingMistNacServerCert represents a OrgSettingMistNacServerCert struct.
// radius server cert to be presented in EAP TLS
type OrgSettingMistNacServerCert struct {
	Cert *string `json:"cert,omitempty"`
	Key  *string `json:"key,omitempty"`
	// private key password (optional)
	Password             *string        `json:"password,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNacServerCert.
// It customizes the JSON marshaling process for OrgSettingMistNacServerCert objects.
func (o OrgSettingMistNacServerCert) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNacServerCert object to a map representation for JSON marshaling.
func (o OrgSettingMistNacServerCert) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Cert != nil {
		structMap["cert"] = o.Cert
	}
	if o.Key != nil {
		structMap["key"] = o.Key
	}
	if o.Password != nil {
		structMap["password"] = o.Password
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNacServerCert.
// It customizes the JSON unmarshaling process for OrgSettingMistNacServerCert objects.
func (o *OrgSettingMistNacServerCert) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingMistNacServerCert
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cert", "key", "password")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.Cert = temp.Cert
	o.Key = temp.Key
	o.Password = temp.Password
	return nil
}

// tempOrgSettingMistNacServerCert is a temporary struct used for validating the fields of OrgSettingMistNacServerCert.
type tempOrgSettingMistNacServerCert struct {
	Cert     *string `json:"cert,omitempty"`
	Key      *string `json:"key,omitempty"`
	Password *string `json:"password,omitempty"`
}
