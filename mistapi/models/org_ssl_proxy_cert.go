package models

import (
	"encoding/json"
)

// OrgSslProxyCert represents a OrgSslProxyCert struct.
type OrgSslProxyCert struct {
	Cert                 *string        `json:"cert,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSslProxyCert.
// It customizes the JSON marshaling process for OrgSslProxyCert objects.
func (o OrgSslProxyCert) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSslProxyCert object to a map representation for JSON marshaling.
func (o OrgSslProxyCert) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Cert != nil {
		structMap["cert"] = o.Cert
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSslProxyCert.
// It customizes the JSON unmarshaling process for OrgSslProxyCert objects.
func (o *OrgSslProxyCert) UnmarshalJSON(input []byte) error {
	var temp tempOrgSslProxyCert
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cert")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.Cert = temp.Cert
	return nil
}

// tempOrgSslProxyCert is a temporary struct used for validating the fields of OrgSslProxyCert.
type tempOrgSslProxyCert struct {
	Cert *string `json:"cert,omitempty"`
}
