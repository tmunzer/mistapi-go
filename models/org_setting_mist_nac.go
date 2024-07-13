package models

import (
    "encoding/json"
)

// OrgSettingMistNac represents a OrgSettingMistNac struct.
type OrgSettingMistNac struct {
    // the CA certs we use to verify client certs
    Cacerts              *string                         `json:"cacerts,omitempty"`
    // use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm.
    DefaultIdpId         *string                         `json:"default_idp_id,omitempty"`
    // eap ssl security level
    // see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR
    EapSslSecurityLevel  *int                            `json:"eap_ssl_security_level,omitempty"`
    // By default NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site.
    // For strict GDPR compliancy NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled
    EuOnly               *bool                           `json:"eu_only,omitempty"`
    Idps                 []OrgSettingMistNacIdp          `json:"idps,omitempty"`
    // radius server cert to be presented in EAP TLS
    ServerCert           *OrgSettingMistNacServerCert    `json:"server_cert,omitempty"`
    // by default NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4
    UseIpVersion         *OrgSettingMistNacIpVersionEnum `json:"use_ip_version,omitempty"`
    // By default NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(radsec) to reach mist-nac. 
    // Set `use_ssl_port`==`true` to override that port with TCP43 (ssl), 
    // This is a org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled
    UseSslPort           *bool                           `json:"use_ssl_port,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMistNac.
// It customizes the JSON marshaling process for OrgSettingMistNac objects.
func (o OrgSettingMistNac) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMistNac object to a map representation for JSON marshaling.
func (o OrgSettingMistNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Cacerts != nil {
        structMap["cacerts"] = o.Cacerts
    }
    if o.DefaultIdpId != nil {
        structMap["default_idp_id"] = o.DefaultIdpId
    }
    if o.EapSslSecurityLevel != nil {
        structMap["eap_ssl_security_level"] = o.EapSslSecurityLevel
    }
    if o.EuOnly != nil {
        structMap["eu_only"] = o.EuOnly
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
    var temp orgSettingMistNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cacerts", "default_idp_id", "eap_ssl_security_level", "eu_only", "idps", "server_cert", "use_ip_version", "use_ssl_port")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Cacerts = temp.Cacerts
    o.DefaultIdpId = temp.DefaultIdpId
    o.EapSslSecurityLevel = temp.EapSslSecurityLevel
    o.EuOnly = temp.EuOnly
    o.Idps = temp.Idps
    o.ServerCert = temp.ServerCert
    o.UseIpVersion = temp.UseIpVersion
    o.UseSslPort = temp.UseSslPort
    return nil
}

// orgSettingMistNac is a temporary struct used for validating the fields of OrgSettingMistNac.
type orgSettingMistNac  struct {
    Cacerts             *string                         `json:"cacerts,omitempty"`
    DefaultIdpId        *string                         `json:"default_idp_id,omitempty"`
    EapSslSecurityLevel *int                            `json:"eap_ssl_security_level,omitempty"`
    EuOnly              *bool                           `json:"eu_only,omitempty"`
    Idps                []OrgSettingMistNacIdp          `json:"idps,omitempty"`
    ServerCert          *OrgSettingMistNacServerCert    `json:"server_cert,omitempty"`
    UseIpVersion        *OrgSettingMistNacIpVersionEnum `json:"use_ip_version,omitempty"`
    UseSslPort          *bool                           `json:"use_ssl_port,omitempty"`
}