package models

import (
    "encoding/json"
    "fmt"
)

// NacPortalSso represents a NacPortalSso struct.
type NacPortalSso struct {
    IdpCert              *string                      `json:"idp_cert,omitempty"`
    // Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`.
    IdpSignAlgo          *NacPortalSsoIdpSignAlgoEnum `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl            *string                      `json:"idp_sso_url,omitempty"`
    Issuer               *string                      `json:"issuer,omitempty"`
    NameidFormat         *string                      `json:"nameid_format,omitempty"`
    SsoRoleMatching      []NacPortalSsoRoleMatching   `json:"sso_role_matching,omitempty"`
    // if it's desired to inject a role into Cert's Subject (so it can be used later on in policy)
    UseSsoRoleForCert    *bool                        `json:"use_sso_role_for_cert,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for NacPortalSso,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacPortalSso) String() string {
    return fmt.Sprintf(
    	"NacPortalSso[IdpCert=%v, IdpSignAlgo=%v, IdpSsoUrl=%v, Issuer=%v, NameidFormat=%v, SsoRoleMatching=%v, UseSsoRoleForCert=%v, AdditionalProperties=%v]",
    	n.IdpCert, n.IdpSignAlgo, n.IdpSsoUrl, n.Issuer, n.NameidFormat, n.SsoRoleMatching, n.UseSsoRoleForCert, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacPortalSso.
// It customizes the JSON marshaling process for NacPortalSso objects.
func (n NacPortalSso) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "idp_cert", "idp_sign_algo", "idp_sso_url", "issuer", "nameid_format", "sso_role_matching", "use_sso_role_for_cert"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NacPortalSso object to a map representation for JSON marshaling.
func (n NacPortalSso) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.IdpCert != nil {
        structMap["idp_cert"] = n.IdpCert
    }
    if n.IdpSignAlgo != nil {
        structMap["idp_sign_algo"] = n.IdpSignAlgo
    }
    if n.IdpSsoUrl != nil {
        structMap["idp_sso_url"] = n.IdpSsoUrl
    }
    if n.Issuer != nil {
        structMap["issuer"] = n.Issuer
    }
    if n.NameidFormat != nil {
        structMap["nameid_format"] = n.NameidFormat
    }
    if n.SsoRoleMatching != nil {
        structMap["sso_role_matching"] = n.SsoRoleMatching
    }
    if n.UseSsoRoleForCert != nil {
        structMap["use_sso_role_for_cert"] = n.UseSsoRoleForCert
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortalSso.
// It customizes the JSON unmarshaling process for NacPortalSso objects.
func (n *NacPortalSso) UnmarshalJSON(input []byte) error {
    var temp tempNacPortalSso
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "idp_cert", "idp_sign_algo", "idp_sso_url", "issuer", "nameid_format", "sso_role_matching", "use_sso_role_for_cert")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.IdpCert = temp.IdpCert
    n.IdpSignAlgo = temp.IdpSignAlgo
    n.IdpSsoUrl = temp.IdpSsoUrl
    n.Issuer = temp.Issuer
    n.NameidFormat = temp.NameidFormat
    n.SsoRoleMatching = temp.SsoRoleMatching
    n.UseSsoRoleForCert = temp.UseSsoRoleForCert
    return nil
}

// tempNacPortalSso is a temporary struct used for validating the fields of NacPortalSso.
type tempNacPortalSso  struct {
    IdpCert           *string                      `json:"idp_cert,omitempty"`
    IdpSignAlgo       *NacPortalSsoIdpSignAlgoEnum `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl         *string                      `json:"idp_sso_url,omitempty"`
    Issuer            *string                      `json:"issuer,omitempty"`
    NameidFormat      *string                      `json:"nameid_format,omitempty"`
    SsoRoleMatching   []NacPortalSsoRoleMatching   `json:"sso_role_matching,omitempty"`
    UseSsoRoleForCert *bool                        `json:"use_sso_role_for_cert,omitempty"`
}
