package models

import (
    "encoding/json"
    "fmt"
)

// NacPortal represents a NacPortal struct.
type NacPortal struct {
    // if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`
    AccessType             *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
    // Background image
    BgImageUrl             *string                  `json:"bg_image_url,omitempty"`
    // In days
    CertExpireTime         *int                     `json:"cert_expire_time,omitempty"`
    // enum: `wpa2`, `wpa3`
    EapType                *NacPortalEapTypeEnum    `json:"eap_type,omitempty"`
    // Model, version, fingering, events (connecting, disconnect, roaming), which ap
    EnableTelemetry        *bool                    `json:"enable_telemetry,omitempty"`
    // In days
    ExpiryNotificationTime *int                     `json:"expiry_notification_time,omitempty"`
    // Portal wlan settings
    GuestPortalConfig      *WlanPortal              `json:"guest_portal_config,omitempty"`
    Name                   *string                  `json:"name,omitempty"`
    // phase 2
    NotifyExpiry           *bool                    `json:"notify_expiry,omitempty"`
    Ssid                   *string                  `json:"ssid,omitempty"`
    Sso                    *NacPortalSso            `json:"sso,omitempty"`
    TemplateUrl            *string                  `json:"template_url,omitempty"`
    ThumbnailUrl           *string                  `json:"thumbnail_url,omitempty"`
    Tos                    *string                  `json:"tos,omitempty"`
    // enum: `guest`, `marvis_client`
    Type                   *NacPortalTypeEnum       `json:"type,omitempty"`
    AdditionalProperties   map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for NacPortal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacPortal) String() string {
    return fmt.Sprintf(
    	"NacPortal[AccessType=%v, BgImageUrl=%v, CertExpireTime=%v, EapType=%v, EnableTelemetry=%v, ExpiryNotificationTime=%v, GuestPortalConfig=%v, Name=%v, NotifyExpiry=%v, Ssid=%v, Sso=%v, TemplateUrl=%v, ThumbnailUrl=%v, Tos=%v, Type=%v, AdditionalProperties=%v]",
    	n.AccessType, n.BgImageUrl, n.CertExpireTime, n.EapType, n.EnableTelemetry, n.ExpiryNotificationTime, n.GuestPortalConfig, n.Name, n.NotifyExpiry, n.Ssid, n.Sso, n.TemplateUrl, n.ThumbnailUrl, n.Tos, n.Type, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacPortal.
// It customizes the JSON marshaling process for NacPortal objects.
func (n NacPortal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "access_type", "bg_image_url", "cert_expire_time", "eap_type", "enable_telemetry", "expiry_notification_time", "guest_portal_config", "name", "notify_expiry", "ssid", "sso", "template_url", "thumbnail_url", "tos", "type"); err != nil {
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
    if n.GuestPortalConfig != nil {
        structMap["guest_portal_config"] = n.GuestPortalConfig.toMap()
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.NotifyExpiry != nil {
        structMap["notify_expiry"] = n.NotifyExpiry
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "access_type", "bg_image_url", "cert_expire_time", "eap_type", "enable_telemetry", "expiry_notification_time", "guest_portal_config", "name", "notify_expiry", "ssid", "sso", "template_url", "thumbnail_url", "tos", "type")
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
    n.GuestPortalConfig = temp.GuestPortalConfig
    n.Name = temp.Name
    n.NotifyExpiry = temp.NotifyExpiry
    n.Ssid = temp.Ssid
    n.Sso = temp.Sso
    n.TemplateUrl = temp.TemplateUrl
    n.ThumbnailUrl = temp.ThumbnailUrl
    n.Tos = temp.Tos
    n.Type = temp.Type
    return nil
}

// tempNacPortal is a temporary struct used for validating the fields of NacPortal.
type tempNacPortal  struct {
    AccessType             *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
    BgImageUrl             *string                  `json:"bg_image_url,omitempty"`
    CertExpireTime         *int                     `json:"cert_expire_time,omitempty"`
    EapType                *NacPortalEapTypeEnum    `json:"eap_type,omitempty"`
    EnableTelemetry        *bool                    `json:"enable_telemetry,omitempty"`
    ExpiryNotificationTime *int                     `json:"expiry_notification_time,omitempty"`
    GuestPortalConfig      *WlanPortal              `json:"guest_portal_config,omitempty"`
    Name                   *string                  `json:"name,omitempty"`
    NotifyExpiry           *bool                    `json:"notify_expiry,omitempty"`
    Ssid                   *string                  `json:"ssid,omitempty"`
    Sso                    *NacPortalSso            `json:"sso,omitempty"`
    TemplateUrl            *string                  `json:"template_url,omitempty"`
    ThumbnailUrl           *string                  `json:"thumbnail_url,omitempty"`
    Tos                    *string                  `json:"tos,omitempty"`
    Type                   *NacPortalTypeEnum       `json:"type,omitempty"`
}
