// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// OrgSettingCacertsConfig represents a OrgSettingCacertsConfig struct.
// Per-issuer CA certificate configuration used to verify client certificates
type OrgSettingCacertsConfig struct {
	// PEM-encoded CA certificate
	Cert string `json:"cert"`
	// Whether CRL checks are enabled. When true, CRL from AIA is used if available unless `crl_url` is set.
	CrlEnabled *bool `json:"crl_enabled,omitempty"`
	// Optional override URL for the certificate CRL distribution point
	CrlUrl *string `json:"crl_url,omitempty"`
	// Optional user-friendly label for the CA issuer configuration
	Name *string `json:"name,omitempty"`
	// Whether OCSP checks are enabled. When true, OCSP responder from AIA is used if available unless `ocsp_url` is set.
	OcspEnabled *bool `json:"ocsp_enabled,omitempty"`
	// Optional override URL for the OCSP responder
	OcspUrl              *string                `json:"ocsp_url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingCacertsConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingCacertsConfig) String() string {
	return fmt.Sprintf(
		"OrgSettingCacertsConfig[Cert=%v, CrlEnabled=%v, CrlUrl=%v, Name=%v, OcspEnabled=%v, OcspUrl=%v, AdditionalProperties=%v]",
		o.Cert, o.CrlEnabled, o.CrlUrl, o.Name, o.OcspEnabled, o.OcspUrl, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingCacertsConfig.
// It customizes the JSON marshaling process for OrgSettingCacertsConfig objects.
func (o OrgSettingCacertsConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"cert", "crl_enabled", "crl_url", "name", "ocsp_enabled", "ocsp_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingCacertsConfig object to a map representation for JSON marshaling.
func (o OrgSettingCacertsConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	structMap["cert"] = o.Cert
	if o.CrlEnabled != nil {
		structMap["crl_enabled"] = o.CrlEnabled
	}
	if o.CrlUrl != nil {
		structMap["crl_url"] = o.CrlUrl
	}
	if o.Name != nil {
		structMap["name"] = o.Name
	}
	if o.OcspEnabled != nil {
		structMap["ocsp_enabled"] = o.OcspEnabled
	}
	if o.OcspUrl != nil {
		structMap["ocsp_url"] = o.OcspUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingCacertsConfig.
// It customizes the JSON unmarshaling process for OrgSettingCacertsConfig objects.
func (o *OrgSettingCacertsConfig) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingCacertsConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert", "crl_enabled", "crl_url", "name", "ocsp_enabled", "ocsp_url")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Cert = *temp.Cert
	o.CrlEnabled = temp.CrlEnabled
	o.CrlUrl = temp.CrlUrl
	o.Name = temp.Name
	o.OcspEnabled = temp.OcspEnabled
	o.OcspUrl = temp.OcspUrl
	return nil
}

// tempOrgSettingCacertsConfig is a temporary struct used for validating the fields of OrgSettingCacertsConfig.
type tempOrgSettingCacertsConfig struct {
	Cert        *string `json:"cert"`
	CrlEnabled  *bool   `json:"crl_enabled,omitempty"`
	CrlUrl      *string `json:"crl_url,omitempty"`
	Name        *string `json:"name,omitempty"`
	OcspEnabled *bool   `json:"ocsp_enabled,omitempty"`
	OcspUrl     *string `json:"ocsp_url,omitempty"`
}

func (o *tempOrgSettingCacertsConfig) validate() error {
	var errs []string
	if o.Cert == nil {
		errs = append(errs, "required field `cert` is missing for type `org_setting_cacerts_config`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
