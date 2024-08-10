package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientNac represents a ClientNac struct.
type ClientNac struct {
    Ap                   []string       `json:"ap,omitempty"`
    AuthType             *string        `json:"auth_type,omitempty"`
    CertCn               []string       `json:"cert_cn,omitempty"`
    CertIssuer           []string       `json:"cert_issuer,omitempty"`
    IdpId                *string        `json:"idp_id,omitempty"`
    IdpRole              []string       `json:"idp_role,omitempty"`
    LastAp               *string        `json:"last_ap,omitempty"`
    LastCertCn           *string        `json:"last_cert_cn,omitempty"`
    LastCertExpiry       *int           `json:"last_cert_expiry,omitempty"`
    LastCertIssuer       *string        `json:"last_cert_issuer,omitempty"`
    LastNacruleId        *string        `json:"last_nacrule_id,omitempty"`
    LastNacruleName      *string        `json:"last_nacrule_name,omitempty"`
    LastNasVendor        *string        `json:"last_nas_vendor,omitempty"`
    LastSsid             *string        `json:"last_ssid,omitempty"`
    LastStatus           *string        `json:"last_status,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    NacruleId            []string       `json:"nacrule_id,omitempty"`
    NacruleMatched       *bool          `json:"nacrule_matched,omitempty"`
    NacruleName          []string       `json:"nacrule_name,omitempty"`
    NasVendor            []string       `json:"nas_vendor,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    RandomMac            *bool          `json:"random_mac,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Ssid                 []string       `json:"ssid,omitempty"`
    Timestamp            *float64       `json:"timestamp,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientNac.
// It customizes the JSON marshaling process for ClientNac objects.
func (c ClientNac) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientNac object to a map representation for JSON marshaling.
func (c ClientNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ap != nil {
        structMap["ap"] = c.Ap
    }
    if c.AuthType != nil {
        structMap["auth_type"] = c.AuthType
    }
    if c.CertCn != nil {
        structMap["cert_cn"] = c.CertCn
    }
    if c.CertIssuer != nil {
        structMap["cert_issuer"] = c.CertIssuer
    }
    if c.IdpId != nil {
        structMap["idp_id"] = c.IdpId
    }
    if c.IdpRole != nil {
        structMap["idp_role"] = c.IdpRole
    }
    if c.LastAp != nil {
        structMap["last_ap"] = c.LastAp
    }
    if c.LastCertCn != nil {
        structMap["last_cert_cn"] = c.LastCertCn
    }
    if c.LastCertExpiry != nil {
        structMap["last_cert_expiry"] = c.LastCertExpiry
    }
    if c.LastCertIssuer != nil {
        structMap["last_cert_issuer"] = c.LastCertIssuer
    }
    if c.LastNacruleId != nil {
        structMap["last_nacrule_id"] = c.LastNacruleId
    }
    if c.LastNacruleName != nil {
        structMap["last_nacrule_name"] = c.LastNacruleName
    }
    if c.LastNasVendor != nil {
        structMap["last_nas_vendor"] = c.LastNasVendor
    }
    if c.LastSsid != nil {
        structMap["last_ssid"] = c.LastSsid
    }
    if c.LastStatus != nil {
        structMap["last_status"] = c.LastStatus
    }
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.NacruleId != nil {
        structMap["nacrule_id"] = c.NacruleId
    }
    if c.NacruleMatched != nil {
        structMap["nacrule_matched"] = c.NacruleMatched
    }
    if c.NacruleName != nil {
        structMap["nacrule_name"] = c.NacruleName
    }
    if c.NasVendor != nil {
        structMap["nas_vendor"] = c.NasVendor
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.RandomMac != nil {
        structMap["random_mac"] = c.RandomMac
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.Ssid != nil {
        structMap["ssid"] = c.Ssid
    }
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp
    }
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientNac.
// It customizes the JSON unmarshaling process for ClientNac objects.
func (c *ClientNac) UnmarshalJSON(input []byte) error {
    var temp tempClientNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "auth_type", "cert_cn", "cert_issuer", "idp_id", "idp_role", "last_ap", "last_cert_cn", "last_cert_expiry", "last_cert_issuer", "last_nacrule_id", "last_nacrule_name", "last_nas_vendor", "last_ssid", "last_status", "mac", "nacrule_id", "nacrule_matched", "nacrule_name", "nas_vendor", "org_id", "random_mac", "site_id", "ssid", "timestamp", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Ap = temp.Ap
    c.AuthType = temp.AuthType
    c.CertCn = temp.CertCn
    c.CertIssuer = temp.CertIssuer
    c.IdpId = temp.IdpId
    c.IdpRole = temp.IdpRole
    c.LastAp = temp.LastAp
    c.LastCertCn = temp.LastCertCn
    c.LastCertExpiry = temp.LastCertExpiry
    c.LastCertIssuer = temp.LastCertIssuer
    c.LastNacruleId = temp.LastNacruleId
    c.LastNacruleName = temp.LastNacruleName
    c.LastNasVendor = temp.LastNasVendor
    c.LastSsid = temp.LastSsid
    c.LastStatus = temp.LastStatus
    c.Mac = temp.Mac
    c.NacruleId = temp.NacruleId
    c.NacruleMatched = temp.NacruleMatched
    c.NacruleName = temp.NacruleName
    c.NasVendor = temp.NasVendor
    c.OrgId = temp.OrgId
    c.RandomMac = temp.RandomMac
    c.SiteId = temp.SiteId
    c.Ssid = temp.Ssid
    c.Timestamp = temp.Timestamp
    c.Type = temp.Type
    return nil
}

// tempClientNac is a temporary struct used for validating the fields of ClientNac.
type tempClientNac  struct {
    Ap              []string   `json:"ap,omitempty"`
    AuthType        *string    `json:"auth_type,omitempty"`
    CertCn          []string   `json:"cert_cn,omitempty"`
    CertIssuer      []string   `json:"cert_issuer,omitempty"`
    IdpId           *string    `json:"idp_id,omitempty"`
    IdpRole         []string   `json:"idp_role,omitempty"`
    LastAp          *string    `json:"last_ap,omitempty"`
    LastCertCn      *string    `json:"last_cert_cn,omitempty"`
    LastCertExpiry  *int       `json:"last_cert_expiry,omitempty"`
    LastCertIssuer  *string    `json:"last_cert_issuer,omitempty"`
    LastNacruleId   *string    `json:"last_nacrule_id,omitempty"`
    LastNacruleName *string    `json:"last_nacrule_name,omitempty"`
    LastNasVendor   *string    `json:"last_nas_vendor,omitempty"`
    LastSsid        *string    `json:"last_ssid,omitempty"`
    LastStatus      *string    `json:"last_status,omitempty"`
    Mac             *string    `json:"mac,omitempty"`
    NacruleId       []string   `json:"nacrule_id,omitempty"`
    NacruleMatched  *bool      `json:"nacrule_matched,omitempty"`
    NacruleName     []string   `json:"nacrule_name,omitempty"`
    NasVendor       []string   `json:"nas_vendor,omitempty"`
    OrgId           *uuid.UUID `json:"org_id,omitempty"`
    RandomMac       *bool      `json:"random_mac,omitempty"`
    SiteId          *uuid.UUID `json:"site_id,omitempty"`
    Ssid            []string   `json:"ssid,omitempty"`
    Timestamp       *float64   `json:"timestamp,omitempty"`
    Type            *string    `json:"type,omitempty"`
}
