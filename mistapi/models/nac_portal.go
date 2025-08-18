// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NacPortal represents a NacPortal struct.
type NacPortal struct {
	// if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`
	AccessType *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
	// Background image
	BgImageUrl *string `json:"bg_image_url,omitempty"`
	// In days
	CertExpireTime *int `json:"cert_expire_time,omitempty"`
	// enum: `wpa2`, `wpa3`
	EapType *NacPortalEapTypeEnum `json:"eap_type,omitempty"`
	// Model, version, fingering, events (connecting, disconnect, roaming), which ap
	EnableTelemetry *bool `json:"enable_telemetry,omitempty"`
	// In days
	ExpiryNotificationTime *int    `json:"expiry_notification_time,omitempty"`
	Name                   *string `json:"name,omitempty"`
	// phase 2
	NotifyExpiry *bool `json:"notify_expiry,omitempty"`
	// Guest portal configuration when `type`==`guest_portal`. If
	// * `auth`==`none`, the user is presented with a terms of service and can click and continue.
	// * `auth`==`external`, the user is redirected to an external URL for authentication.
	// * `auth`==`multi`, the user is presented with a choice of authentication methods:
	// - social logins: facebook / google / amazon / microsoft / azure
	// - sponsor
	// - sms: supported provider: twillio
	// - email
	// - sso
	// - userpass: pre created guest list
	Portal *NacPortalGuestPortal `json:"portal,omitempty"`
	// If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_jwt_secret` will be generated
	PortalAuthorizeJwtSecret *string `json:"portal_authorize_jwt_secret,omitempty"`
	// If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_url` will be generated
	PortalAuthorizeUrl *string `json:"portal_authorize_url,omitempty"`
	// If `type`==`guest_portal` or `type`==`guest_admin` and ans SSO is enabled, the `portal_sso_url` will be generated (which needs to be configured in your IDP
	PortalSsoUrl *string       `json:"portal_sso_url,omitempty"`
	Ssid         *string       `json:"ssid,omitempty"`
	Sso          *NacPortalSso `json:"sso,omitempty"`
	TemplateUrl  *string       `json:"template_url,omitempty"`
	ThumbnailUrl *string       `json:"thumbnail_url,omitempty"`
	Tos          *string       `json:"tos,omitempty"`
	// enum:
	// * `guest_admin`: NAC-Based Portal Admin for Pre Created Guest Authentication
	// * `guest_portal`: NAC-Based Guest Portal
	// * `marvis_client`
	Type *NacPortalTypeEnum `json:"type,omitempty"`
	// If `auth`==`guest_admin`, the URL to the guest admin portal
	UiUrl                *string                `json:"ui_url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacPortal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacPortal) String() string {
	return fmt.Sprintf(
		"NacPortal[AccessType=%v, BgImageUrl=%v, CertExpireTime=%v, EapType=%v, EnableTelemetry=%v, ExpiryNotificationTime=%v, Name=%v, NotifyExpiry=%v, Portal=%v, PortalAuthorizeJwtSecret=%v, PortalAuthorizeUrl=%v, PortalSsoUrl=%v, Ssid=%v, Sso=%v, TemplateUrl=%v, ThumbnailUrl=%v, Tos=%v, Type=%v, UiUrl=%v, AdditionalProperties=%v]",
		n.AccessType, n.BgImageUrl, n.CertExpireTime, n.EapType, n.EnableTelemetry, n.ExpiryNotificationTime, n.Name, n.NotifyExpiry, n.Portal, n.PortalAuthorizeJwtSecret, n.PortalAuthorizeUrl, n.PortalSsoUrl, n.Ssid, n.Sso, n.TemplateUrl, n.ThumbnailUrl, n.Tos, n.Type, n.UiUrl, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacPortal.
// It customizes the JSON marshaling process for NacPortal objects.
func (n NacPortal) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"access_type", "bg_image_url", "cert_expire_time", "eap_type", "enable_telemetry", "expiry_notification_time", "name", "notify_expiry", "portal", "portal_authorize_jwt_secret", "portal_authorize_url", "portal_sso_url", "ssid", "sso", "template_url", "thumbnail_url", "tos", "type", "ui_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacPortal object to a map representation for JSON marshaling.
func (n NacPortal) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.AccessType != nil {
		structMap["access_type"] = n.AccessType
	}
	if n.BgImageUrl != nil {
		structMap["bg_image_url"] = n.BgImageUrl
	}
	if n.CertExpireTime != nil {
		structMap["cert_expire_time"] = n.CertExpireTime
	}
	if n.EapType != nil {
		structMap["eap_type"] = n.EapType
	}
	if n.EnableTelemetry != nil {
		structMap["enable_telemetry"] = n.EnableTelemetry
	}
	if n.ExpiryNotificationTime != nil {
		structMap["expiry_notification_time"] = n.ExpiryNotificationTime
	}
	if n.Name != nil {
		structMap["name"] = n.Name
	}
	if n.NotifyExpiry != nil {
		structMap["notify_expiry"] = n.NotifyExpiry
	}
	if n.Portal != nil {
		structMap["portal"] = n.Portal.toMap()
	}
	if n.PortalAuthorizeJwtSecret != nil {
		structMap["portal_authorize_jwt_secret"] = n.PortalAuthorizeJwtSecret
	}
	if n.PortalAuthorizeUrl != nil {
		structMap["portal_authorize_url"] = n.PortalAuthorizeUrl
	}
	if n.PortalSsoUrl != nil {
		structMap["portal_sso_url"] = n.PortalSsoUrl
	}
	if n.Ssid != nil {
		structMap["ssid"] = n.Ssid
	}
	if n.Sso != nil {
		structMap["sso"] = n.Sso.toMap()
	}
	if n.TemplateUrl != nil {
		structMap["template_url"] = n.TemplateUrl
	}
	if n.ThumbnailUrl != nil {
		structMap["thumbnail_url"] = n.ThumbnailUrl
	}
	if n.Tos != nil {
		structMap["tos"] = n.Tos
	}
	if n.Type != nil {
		structMap["type"] = n.Type
	}
	if n.UiUrl != nil {
		structMap["ui_url"] = n.UiUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortal.
// It customizes the JSON unmarshaling process for NacPortal objects.
func (n *NacPortal) UnmarshalJSON(input []byte) error {
	var temp tempNacPortal
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "access_type", "bg_image_url", "cert_expire_time", "eap_type", "enable_telemetry", "expiry_notification_time", "name", "notify_expiry", "portal", "portal_authorize_jwt_secret", "portal_authorize_url", "portal_sso_url", "ssid", "sso", "template_url", "thumbnail_url", "tos", "type", "ui_url")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.AccessType = temp.AccessType
	n.BgImageUrl = temp.BgImageUrl
	n.CertExpireTime = temp.CertExpireTime
	n.EapType = temp.EapType
	n.EnableTelemetry = temp.EnableTelemetry
	n.ExpiryNotificationTime = temp.ExpiryNotificationTime
	n.Name = temp.Name
	n.NotifyExpiry = temp.NotifyExpiry
	n.Portal = temp.Portal
	n.PortalAuthorizeJwtSecret = temp.PortalAuthorizeJwtSecret
	n.PortalAuthorizeUrl = temp.PortalAuthorizeUrl
	n.PortalSsoUrl = temp.PortalSsoUrl
	n.Ssid = temp.Ssid
	n.Sso = temp.Sso
	n.TemplateUrl = temp.TemplateUrl
	n.ThumbnailUrl = temp.ThumbnailUrl
	n.Tos = temp.Tos
	n.Type = temp.Type
	n.UiUrl = temp.UiUrl
	return nil
}

// tempNacPortal is a temporary struct used for validating the fields of NacPortal.
type tempNacPortal struct {
	AccessType               *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
	BgImageUrl               *string                  `json:"bg_image_url,omitempty"`
	CertExpireTime           *int                     `json:"cert_expire_time,omitempty"`
	EapType                  *NacPortalEapTypeEnum    `json:"eap_type,omitempty"`
	EnableTelemetry          *bool                    `json:"enable_telemetry,omitempty"`
	ExpiryNotificationTime   *int                     `json:"expiry_notification_time,omitempty"`
	Name                     *string                  `json:"name,omitempty"`
	NotifyExpiry             *bool                    `json:"notify_expiry,omitempty"`
	Portal                   *NacPortalGuestPortal    `json:"portal,omitempty"`
	PortalAuthorizeJwtSecret *string                  `json:"portal_authorize_jwt_secret,omitempty"`
	PortalAuthorizeUrl       *string                  `json:"portal_authorize_url,omitempty"`
	PortalSsoUrl             *string                  `json:"portal_sso_url,omitempty"`
	Ssid                     *string                  `json:"ssid,omitempty"`
	Sso                      *NacPortalSso            `json:"sso,omitempty"`
	TemplateUrl              *string                  `json:"template_url,omitempty"`
	ThumbnailUrl             *string                  `json:"thumbnail_url,omitempty"`
	Tos                      *string                  `json:"tos,omitempty"`
	Type                     *NacPortalTypeEnum       `json:"type,omitempty"`
	UiUrl                    *string                  `json:"ui_url,omitempty"`
}
