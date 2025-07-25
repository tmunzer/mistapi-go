// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// PskPortalSso represents a PskPortalSso struct.
// If `auth`==`sso`
type PskPortalSso struct {
    // // allowed roles for accessing psk portal, if none, any role is permitted
    AllowedRoles         []string                     `json:"allowed_roles,omitempty"`
    IdpCert              *string                      `json:"idp_cert,omitempty"`
    // Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`. enum: `sha1`, `sha256`, `sha384`, `sha512`
    IdpSignAlgo          *PskPortalSsoIdpSignAlgoEnum `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl            *string                      `json:"idp_sso_url,omitempty"`
    Issuer               *string                      `json:"issuer,omitempty"`
    NameidFormat         *string                      `json:"nameid_format,omitempty"`
    // Property key is the role name, property value is the SSO Attribute
    RoleMapping          map[string]string            `json:"role_mapping,omitempty"`
    // If enabled, the `role` above will be ignored
    UseSsoRoleForPskRole *bool                        `json:"use_sso_role_for_psk_role,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalSso,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalSso) String() string {
    return fmt.Sprintf(
    	"PskPortalSso[AllowedRoles=%v, IdpCert=%v, IdpSignAlgo=%v, IdpSsoUrl=%v, Issuer=%v, NameidFormat=%v, RoleMapping=%v, UseSsoRoleForPskRole=%v, AdditionalProperties=%v]",
    	p.AllowedRoles, p.IdpCert, p.IdpSignAlgo, p.IdpSsoUrl, p.Issuer, p.NameidFormat, p.RoleMapping, p.UseSsoRoleForPskRole, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalSso.
// It customizes the JSON marshaling process for PskPortalSso objects.
func (p PskPortalSso) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "allowed_roles", "idp_cert", "idp_sign_algo", "idp_sso_url", "issuer", "nameid_format", "role_mapping", "use_sso_role_for_psk_role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalSso object to a map representation for JSON marshaling.
func (p PskPortalSso) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AllowedRoles != nil {
        structMap["allowed_roles"] = p.AllowedRoles
    }
    if p.IdpCert != nil {
        structMap["idp_cert"] = p.IdpCert
    }
    if p.IdpSignAlgo != nil {
        structMap["idp_sign_algo"] = p.IdpSignAlgo
    }
    if p.IdpSsoUrl != nil {
        structMap["idp_sso_url"] = p.IdpSsoUrl
    }
    if p.Issuer != nil {
        structMap["issuer"] = p.Issuer
    }
    if p.NameidFormat != nil {
        structMap["nameid_format"] = p.NameidFormat
    }
    if p.RoleMapping != nil {
        structMap["role_mapping"] = p.RoleMapping
    }
    if p.UseSsoRoleForPskRole != nil {
        structMap["use_sso_role_for_psk_role"] = p.UseSsoRoleForPskRole
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalSso.
// It customizes the JSON unmarshaling process for PskPortalSso objects.
func (p *PskPortalSso) UnmarshalJSON(input []byte) error {
    var temp tempPskPortalSso
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allowed_roles", "idp_cert", "idp_sign_algo", "idp_sso_url", "issuer", "nameid_format", "role_mapping", "use_sso_role_for_psk_role")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.AllowedRoles = temp.AllowedRoles
    p.IdpCert = temp.IdpCert
    p.IdpSignAlgo = temp.IdpSignAlgo
    p.IdpSsoUrl = temp.IdpSsoUrl
    p.Issuer = temp.Issuer
    p.NameidFormat = temp.NameidFormat
    p.RoleMapping = temp.RoleMapping
    p.UseSsoRoleForPskRole = temp.UseSsoRoleForPskRole
    return nil
}

// tempPskPortalSso is a temporary struct used for validating the fields of PskPortalSso.
type tempPskPortalSso  struct {
    AllowedRoles         []string                     `json:"allowed_roles,omitempty"`
    IdpCert              *string                      `json:"idp_cert,omitempty"`
    IdpSignAlgo          *PskPortalSsoIdpSignAlgoEnum `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl            *string                      `json:"idp_sso_url,omitempty"`
    Issuer               *string                      `json:"issuer,omitempty"`
    NameidFormat         *string                      `json:"nameid_format,omitempty"`
    RoleMapping          map[string]string            `json:"role_mapping,omitempty"`
    UseSsoRoleForPskRole *bool                        `json:"use_sso_role_for_psk_role,omitempty"`
}
