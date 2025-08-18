// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ConstNacEvent represents a ConstNacEvent struct.
type ConstNacEvent struct {
	Ap          *string    `json:"ap,omitempty"`
	Bssid       *string    `json:"bssid,omitempty"`
	CertCn      *string    `json:"cert_cn,omitempty"`
	CertExpiry  *int       `json:"cert_expiry,omitempty"`
	CertIssuer  *string    `json:"cert_issuer,omitempty"`
	CertSanUpn  []string   `json:"cert_san_upn,omitempty"`
	CertSerial  *string    `json:"cert_serial,omitempty"`
	CertSubject *string    `json:"cert_subject,omitempty"`
	EapType     *string    `json:"eap_type,omitempty"`
	NasVendor   *string    `json:"nas_vendor,omitempty"`
	OrgId       *uuid.UUID `json:"org_id,omitempty"`
	RandomMac   *bool      `json:"random_mac,omitempty"`
	SiteId      *uuid.UUID `json:"site_id,omitempty"`
	Ssid        *string    `json:"ssid,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	Type                 *string                `json:"type,omitempty"`
	Username             *string                `json:"username,omitempty"`
	Wcid                 *uuid.UUID             `json:"wcid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstNacEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstNacEvent) String() string {
	return fmt.Sprintf(
		"ConstNacEvent[Ap=%v, Bssid=%v, CertCn=%v, CertExpiry=%v, CertIssuer=%v, CertSanUpn=%v, CertSerial=%v, CertSubject=%v, EapType=%v, NasVendor=%v, OrgId=%v, RandomMac=%v, SiteId=%v, Ssid=%v, Timestamp=%v, Type=%v, Username=%v, Wcid=%v, AdditionalProperties=%v]",
		c.Ap, c.Bssid, c.CertCn, c.CertExpiry, c.CertIssuer, c.CertSanUpn, c.CertSerial, c.CertSubject, c.EapType, c.NasVendor, c.OrgId, c.RandomMac, c.SiteId, c.Ssid, c.Timestamp, c.Type, c.Username, c.Wcid, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstNacEvent.
// It customizes the JSON marshaling process for ConstNacEvent objects.
func (c ConstNacEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"ap", "bssid", "cert_cn", "cert_expiry", "cert_issuer", "cert_san_upn", "cert_serial", "cert_subject", "eap_type", "nas_vendor", "org_id", "random_mac", "site_id", "ssid", "timestamp", "type", "username", "wcid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstNacEvent object to a map representation for JSON marshaling.
func (c ConstNacEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Ap != nil {
		structMap["ap"] = c.Ap
	}
	if c.Bssid != nil {
		structMap["bssid"] = c.Bssid
	}
	if c.CertCn != nil {
		structMap["cert_cn"] = c.CertCn
	}
	if c.CertExpiry != nil {
		structMap["cert_expiry"] = c.CertExpiry
	}
	if c.CertIssuer != nil {
		structMap["cert_issuer"] = c.CertIssuer
	}
	if c.CertSanUpn != nil {
		structMap["cert_san_upn"] = c.CertSanUpn
	}
	if c.CertSerial != nil {
		structMap["cert_serial"] = c.CertSerial
	}
	if c.CertSubject != nil {
		structMap["cert_subject"] = c.CertSubject
	}
	if c.EapType != nil {
		structMap["eap_type"] = c.EapType
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
	if c.Username != nil {
		structMap["username"] = c.Username
	}
	if c.Wcid != nil {
		structMap["wcid"] = c.Wcid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstNacEvent.
// It customizes the JSON unmarshaling process for ConstNacEvent objects.
func (c *ConstNacEvent) UnmarshalJSON(input []byte) error {
	var temp tempConstNacEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "bssid", "cert_cn", "cert_expiry", "cert_issuer", "cert_san_upn", "cert_serial", "cert_subject", "eap_type", "nas_vendor", "org_id", "random_mac", "site_id", "ssid", "timestamp", "type", "username", "wcid")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Ap = temp.Ap
	c.Bssid = temp.Bssid
	c.CertCn = temp.CertCn
	c.CertExpiry = temp.CertExpiry
	c.CertIssuer = temp.CertIssuer
	c.CertSanUpn = temp.CertSanUpn
	c.CertSerial = temp.CertSerial
	c.CertSubject = temp.CertSubject
	c.EapType = temp.EapType
	c.NasVendor = temp.NasVendor
	c.OrgId = temp.OrgId
	c.RandomMac = temp.RandomMac
	c.SiteId = temp.SiteId
	c.Ssid = temp.Ssid
	c.Timestamp = temp.Timestamp
	c.Type = temp.Type
	c.Username = temp.Username
	c.Wcid = temp.Wcid
	return nil
}

// tempConstNacEvent is a temporary struct used for validating the fields of ConstNacEvent.
type tempConstNacEvent struct {
	Ap          *string    `json:"ap,omitempty"`
	Bssid       *string    `json:"bssid,omitempty"`
	CertCn      *string    `json:"cert_cn,omitempty"`
	CertExpiry  *int       `json:"cert_expiry,omitempty"`
	CertIssuer  *string    `json:"cert_issuer,omitempty"`
	CertSanUpn  []string   `json:"cert_san_upn,omitempty"`
	CertSerial  *string    `json:"cert_serial,omitempty"`
	CertSubject *string    `json:"cert_subject,omitempty"`
	EapType     *string    `json:"eap_type,omitempty"`
	NasVendor   *string    `json:"nas_vendor,omitempty"`
	OrgId       *uuid.UUID `json:"org_id,omitempty"`
	RandomMac   *bool      `json:"random_mac,omitempty"`
	SiteId      *uuid.UUID `json:"site_id,omitempty"`
	Ssid        *string    `json:"ssid,omitempty"`
	Timestamp   *float64   `json:"timestamp,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Username    *string    `json:"username,omitempty"`
	Wcid        *uuid.UUID `json:"wcid,omitempty"`
}
