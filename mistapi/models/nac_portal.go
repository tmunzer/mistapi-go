package models

import (
    "encoding/json"
)

// NacPortal represents a NacPortal struct.
type NacPortal struct {
    // enum: `wireless`, `wireless+wired`
    AccessType             *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
    // background image
    BgImageUrl             *string                  `json:"bg_image_url,omitempty"`
    // in days
    CertExpireTime         *int                     `json:"cert_expire_time,omitempty"`
    // model, version, fingering, events (connecting, disconnect, roaming), which ap
    EnableTelemetry        *bool                    `json:"enable_telemetry,omitempty"`
    // in days
    ExpiryNotificationTime *int                     `json:"expiry_notification_time,omitempty"`
    Name                   *string                  `json:"name,omitempty"`
    // phase 2
    NotifyExpiry           *bool                    `json:"notify_expiry,omitempty"`
    Ssid                   *string                  `json:"ssid,omitempty"`
    Sso                    *NacPortalSso            `json:"sso,omitempty"`
    TemplateUrl            *string                  `json:"template_url,omitempty"`
    ThumbnailUrl           *string                  `json:"thumbnail_url,omitempty"`
    Tos                    *string                  `json:"tos,omitempty"`
    AdditionalProperties   map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacPortal.
// It customizes the JSON marshaling process for NacPortal objects.
func (n NacPortal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NacPortal object to a map representation for JSON marshaling.
func (n NacPortal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AccessType != nil {
        structMap["access_type"] = n.AccessType
    }
    if n.BgImageUrl != nil {
        structMap["bg_image_url"] = n.BgImageUrl
    }
    if n.CertExpireTime != nil {
        structMap["cert_expire_time"] = n.CertExpireTime
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortal.
// It customizes the JSON unmarshaling process for NacPortal objects.
func (n *NacPortal) UnmarshalJSON(input []byte) error {
    var temp nacPortal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "access_type", "bg_image_url", "cert_expire_time", "enable_telemetry", "expiry_notification_time", "name", "notify_expiry", "ssid", "sso", "template_url", "thumbnail_url", "tos")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.AccessType = temp.AccessType
    n.BgImageUrl = temp.BgImageUrl
    n.CertExpireTime = temp.CertExpireTime
    n.EnableTelemetry = temp.EnableTelemetry
    n.ExpiryNotificationTime = temp.ExpiryNotificationTime
    n.Name = temp.Name
    n.NotifyExpiry = temp.NotifyExpiry
    n.Ssid = temp.Ssid
    n.Sso = temp.Sso
    n.TemplateUrl = temp.TemplateUrl
    n.ThumbnailUrl = temp.ThumbnailUrl
    n.Tos = temp.Tos
    return nil
}

// nacPortal is a temporary struct used for validating the fields of NacPortal.
type nacPortal  struct {
    AccessType             *NacPortalAccessTypeEnum `json:"access_type,omitempty"`
    BgImageUrl             *string                  `json:"bg_image_url,omitempty"`
    CertExpireTime         *int                     `json:"cert_expire_time,omitempty"`
    EnableTelemetry        *bool                    `json:"enable_telemetry,omitempty"`
    ExpiryNotificationTime *int                     `json:"expiry_notification_time,omitempty"`
    Name                   *string                  `json:"name,omitempty"`
    NotifyExpiry           *bool                    `json:"notify_expiry,omitempty"`
    Ssid                   *string                  `json:"ssid,omitempty"`
    Sso                    *NacPortalSso            `json:"sso,omitempty"`
    TemplateUrl            *string                  `json:"template_url,omitempty"`
    ThumbnailUrl           *string                  `json:"thumbnail_url,omitempty"`
    Tos                    *string                  `json:"tos,omitempty"`
}
