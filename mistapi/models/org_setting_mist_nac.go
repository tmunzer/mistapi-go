// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingMistNac represents a OrgSettingMistNac struct.
type OrgSettingMistNac struct {
	// List of PEM-encoded ca certs
	Cacerts []string `json:"cacerts,omitempty"`
	// use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm.
	DefaultIdpId *string `json:"default_idp_id,omitempty"`
	// to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html
	DisableRsaeAlgorithms *bool `json:"disable_rsae_algorithms,omitempty"`
	// eap ssl security level, see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR
	EapSslSecurityLevel *int `json:"eap_ssl_security_level,omitempty"`
	// By default, NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliance NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled
	EuOnly *bool `json:"eu_only,omitempty"`
	// allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`
	IdpMachineCertLookupField *IdpMachineCertLookupFieldEnum `json:"idp_machine_cert_lookup_field,omitempty"`
	// allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`
	IdpUserCertLookupField *IdpUserCertLookupFieldEnum `json:"idp_user_cert_lookup_field,omitempty"`
	Idps                   []OrgSettingMistNacIdp      `json:"idps,omitempty"`
	// radius server cert to be presented in EAP TLS
	ServerCert *OrgSettingMistNacServerCert `json:"server_cert,omitempty"`
	// by default, NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4. enum: `v4`, `v6`
	UseIpVersion *OrgSettingMistNacIpVersionEnum `json:"use_ip_version,omitempty"`
	// By default, NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(RadSec) to reach mist-nac. Set `use_ssl_port`==`true` to override that port with TCP43 (ssl), This is an org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled
	UseSslPort           *bool                  `json:"use_ssl_port,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMistNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMistNac) String() string {
	return fmt.Sprintf(
		"OrgSettingMistNac[Cacerts=%v, DefaultIdpId=%v, DisableRsaeAlgorithms=%v, EapSslSecurityLevel=%v, EuOnly=%v, IdpMachineCertLookupField=%v, IdpUserCertLookupField=%v, Idps=%v, ServerCert=%v, UseIpVersion=%v, UseSslPort=%v, AdditionalProperties=%v]",
		o.Cacerts, o.DefaultIdpId, o.DisableRsaeAlgorithms, o.EapSslSecurityLevel, o.EuOnly, o.IdpMachineCertLookupField, o.IdpUserCertLookupField, o.Idps, o.ServerCert, o.UseIpVersion, o.UseSslPort, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNac.
// It customizes the JSON marshaling process for OrgSettingMistNac objects.
func (o OrgSettingMistNac) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"cacerts", "default_idp_id", "disable_rsae_algorithms", "eap_ssl_security_level", "eu_only", "idp_machine_cert_lookup_field", "idp_user_cert_lookup_field", "idps", "server_cert", "use_ip_version", "use_ssl_port"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNac object to a map representation for JSON marshaling.
func (o OrgSettingMistNac) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Cacerts != nil {
		structMap["cacerts"] = o.Cacerts
	}
	if o.DefaultIdpId != nil {
		structMap["default_idp_id"] = o.DefaultIdpId
	}
	if o.DisableRsaeAlgorithms != nil {
		structMap["disable_rsae_algorithms"] = o.DisableRsaeAlgorithms
	}
	if o.EapSslSecurityLevel != nil {
		structMap["eap_ssl_security_level"] = o.EapSslSecurityLevel
	}
	if o.EuOnly != nil {
		structMap["eu_only"] = o.EuOnly
	}
	if o.IdpMachineCertLookupField != nil {
		structMap["idp_machine_cert_lookup_field"] = o.IdpMachineCertLookupField
	}
	if o.IdpUserCertLookupField != nil {
		structMap["idp_user_cert_lookup_field"] = o.IdpUserCertLookupField
	}
	if o.Idps != nil {
		structMap["idps"] = o.Idps
	}
	if o.ServerCert != nil {
		structMap["server_cert"] = o.ServerCert.toMap()
	}
	if o.UseIpVersion != nil {
		structMap["use_ip_version"] = o.UseIpVersion
	}
	if o.UseSslPort != nil {
		structMap["use_ssl_port"] = o.UseSslPort
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMistNac.
// It customizes the JSON unmarshaling process for OrgSettingMistNac objects.
func (o *OrgSettingMistNac) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingMistNac
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cacerts", "default_idp_id", "disable_rsae_algorithms", "eap_ssl_security_level", "eu_only", "idp_machine_cert_lookup_field", "idp_user_cert_lookup_field", "idps", "server_cert", "use_ip_version", "use_ssl_port")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Cacerts = temp.Cacerts
	o.DefaultIdpId = temp.DefaultIdpId
	o.DisableRsaeAlgorithms = temp.DisableRsaeAlgorithms
	o.EapSslSecurityLevel = temp.EapSslSecurityLevel
	o.EuOnly = temp.EuOnly
	o.IdpMachineCertLookupField = temp.IdpMachineCertLookupField
	o.IdpUserCertLookupField = temp.IdpUserCertLookupField
	o.Idps = temp.Idps
	o.ServerCert = temp.ServerCert
	o.UseIpVersion = temp.UseIpVersion
	o.UseSslPort = temp.UseSslPort
	return nil
}

// tempOrgSettingMistNac is a temporary struct used for validating the fields of OrgSettingMistNac.
type tempOrgSettingMistNac struct {
	Cacerts                   []string                        `json:"cacerts,omitempty"`
	DefaultIdpId              *string                         `json:"default_idp_id,omitempty"`
	DisableRsaeAlgorithms     *bool                           `json:"disable_rsae_algorithms,omitempty"`
	EapSslSecurityLevel       *int                            `json:"eap_ssl_security_level,omitempty"`
	EuOnly                    *bool                           `json:"eu_only,omitempty"`
	IdpMachineCertLookupField *IdpMachineCertLookupFieldEnum  `json:"idp_machine_cert_lookup_field,omitempty"`
	IdpUserCertLookupField    *IdpUserCertLookupFieldEnum     `json:"idp_user_cert_lookup_field,omitempty"`
	Idps                      []OrgSettingMistNacIdp          `json:"idps,omitempty"`
	ServerCert                *OrgSettingMistNacServerCert    `json:"server_cert,omitempty"`
	UseIpVersion              *OrgSettingMistNacIpVersionEnum `json:"use_ip_version,omitempty"`
	UseSslPort                *bool                           `json:"use_ssl_port,omitempty"`
}
