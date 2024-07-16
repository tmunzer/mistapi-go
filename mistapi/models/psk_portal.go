package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// PskPortal represents a PskPortal struct.
type PskPortal struct {
    // Note: `sponsor` not yet available
    Auth                         *PskPortalAuthEnum        `json:"auth,omitempty"`
    BgImageUrl                   *string                   `json:"bg_image_url,omitempty"`
    // used to cleanup exited psk when portal delete or ssid changed
    CleanupPsk                   *bool                     `json:"cleanup_psk,omitempty"`
    CreatedTime                  *float64                  `json:"created_time,omitempty"`
    // unit min
    ExpireTime                   *int                      `json:"expire_time,omitempty"`
    // Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire
    ExpiryNotificationTime       *int                      `json:"expiry_notification_time,omitempty"`
    // only if `type`==`admin`
    HidePsksCreatedByOtherAdmins *bool                     `json:"hide_psks_created_by_other_admins,omitempty"`
    Id                           *string                   `json:"id,omitempty"`
    // `max_usage`==`0` means unlimited
    MaxUsage                     *int                      `json:"max_usage,omitempty"`
    ModifiedTime                 *float64                  `json:"modified_time,omitempty"`
    Name                         string                    `json:"name"`
    // optional, will include the link in the notification email the customer can either provide their own url or use the one generate from mist, or do a url shorterner against either
    NotificationRenewUrl         *string                   `json:"notification_renew_url,omitempty"`
    // If set to true, reminder notification will be sent when psk is about to expire
    NotifyExpiry                 *bool                     `json:"notify_expiry,omitempty"`
    NotifyOnCreateOrEdit         *bool                     `json:"notify_on_create_or_edit,omitempty"`
    OrgId                        *uuid.UUID                `json:"org_id,omitempty"`
    PassphraseRules              *PskPortalPassphraseRules `json:"passphrase_rules,omitempty"`
    // what information to ask for (email is required by default)
    RequiredFields               []string                  `json:"required_fields,omitempty"`
    Role                         *string                   `json:"role,omitempty"`
    // intended SSID
    Ssid                         string                    `json:"ssid"`
    // if `auth`==`sso`
    Sso                          *PskPortalSso             `json:"sso,omitempty"`
    // UI customization
    TemplateUrl                  *string                   `json:"template_url,omitempty"`
    ThumbnailUrl                 *string                   `json:"thumbnail_url,omitempty"`
    // for personal psk portal
    Type                         *PskPortalTypeEnum        `json:"type,omitempty"`
    VlanId                       *int                      `json:"vlan_id,omitempty"`
    AdditionalProperties         map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PskPortal.
// It customizes the JSON marshaling process for PskPortal objects.
func (p PskPortal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortal object to a map representation for JSON marshaling.
func (p PskPortal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Auth != nil {
        structMap["auth"] = p.Auth
    }
    if p.BgImageUrl != nil {
        structMap["bg_image_url"] = p.BgImageUrl
    }
    if p.CleanupPsk != nil {
        structMap["cleanup_psk"] = p.CleanupPsk
    }
    if p.CreatedTime != nil {
        structMap["created_time"] = p.CreatedTime
    }
    if p.ExpireTime != nil {
        structMap["expire_time"] = p.ExpireTime
    }
    if p.ExpiryNotificationTime != nil {
        structMap["expiry_notification_time"] = p.ExpiryNotificationTime
    }
    if p.HidePsksCreatedByOtherAdmins != nil {
        structMap["hide_psks_created_by_other_admins"] = p.HidePsksCreatedByOtherAdmins
    }
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.MaxUsage != nil {
        structMap["max_usage"] = p.MaxUsage
    }
    if p.ModifiedTime != nil {
        structMap["modified_time"] = p.ModifiedTime
    }
    structMap["name"] = p.Name
    if p.NotificationRenewUrl != nil {
        structMap["notification_renew_url"] = p.NotificationRenewUrl
    }
    if p.NotifyExpiry != nil {
        structMap["notify_expiry"] = p.NotifyExpiry
    }
    if p.NotifyOnCreateOrEdit != nil {
        structMap["notify_on_create_or_edit"] = p.NotifyOnCreateOrEdit
    }
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    if p.PassphraseRules != nil {
        structMap["passphrase_rules"] = p.PassphraseRules.toMap()
    }
    if p.RequiredFields != nil {
        structMap["required_fields"] = p.RequiredFields
    }
    if p.Role != nil {
        structMap["role"] = p.Role
    }
    structMap["ssid"] = p.Ssid
    if p.Sso != nil {
        structMap["sso"] = p.Sso.toMap()
    }
    if p.TemplateUrl != nil {
        structMap["template_url"] = p.TemplateUrl
    }
    if p.ThumbnailUrl != nil {
        structMap["thumbnail_url"] = p.ThumbnailUrl
    }
    if p.Type != nil {
        structMap["type"] = p.Type
    }
    if p.VlanId != nil {
        structMap["vlan_id"] = p.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortal.
// It customizes the JSON unmarshaling process for PskPortal objects.
func (p *PskPortal) UnmarshalJSON(input []byte) error {
    var temp pskPortal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth", "bg_image_url", "cleanup_psk", "created_time", "expire_time", "expiry_notification_time", "hide_psks_created_by_other_admins", "id", "max_usage", "modified_time", "name", "notification_renew_url", "notify_expiry", "notify_on_create_or_edit", "org_id", "passphrase_rules", "required_fields", "role", "ssid", "sso", "template_url", "thumbnail_url", "type", "vlan_id")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Auth = temp.Auth
    p.BgImageUrl = temp.BgImageUrl
    p.CleanupPsk = temp.CleanupPsk
    p.CreatedTime = temp.CreatedTime
    p.ExpireTime = temp.ExpireTime
    p.ExpiryNotificationTime = temp.ExpiryNotificationTime
    p.HidePsksCreatedByOtherAdmins = temp.HidePsksCreatedByOtherAdmins
    p.Id = temp.Id
    p.MaxUsage = temp.MaxUsage
    p.ModifiedTime = temp.ModifiedTime
    p.Name = *temp.Name
    p.NotificationRenewUrl = temp.NotificationRenewUrl
    p.NotifyExpiry = temp.NotifyExpiry
    p.NotifyOnCreateOrEdit = temp.NotifyOnCreateOrEdit
    p.OrgId = temp.OrgId
    p.PassphraseRules = temp.PassphraseRules
    p.RequiredFields = temp.RequiredFields
    p.Role = temp.Role
    p.Ssid = *temp.Ssid
    p.Sso = temp.Sso
    p.TemplateUrl = temp.TemplateUrl
    p.ThumbnailUrl = temp.ThumbnailUrl
    p.Type = temp.Type
    p.VlanId = temp.VlanId
    return nil
}

// pskPortal is a temporary struct used for validating the fields of PskPortal.
type pskPortal  struct {
    Auth                         *PskPortalAuthEnum        `json:"auth,omitempty"`
    BgImageUrl                   *string                   `json:"bg_image_url,omitempty"`
    CleanupPsk                   *bool                     `json:"cleanup_psk,omitempty"`
    CreatedTime                  *float64                  `json:"created_time,omitempty"`
    ExpireTime                   *int                      `json:"expire_time,omitempty"`
    ExpiryNotificationTime       *int                      `json:"expiry_notification_time,omitempty"`
    HidePsksCreatedByOtherAdmins *bool                     `json:"hide_psks_created_by_other_admins,omitempty"`
    Id                           *string                   `json:"id,omitempty"`
    MaxUsage                     *int                      `json:"max_usage,omitempty"`
    ModifiedTime                 *float64                  `json:"modified_time,omitempty"`
    Name                         *string                   `json:"name"`
    NotificationRenewUrl         *string                   `json:"notification_renew_url,omitempty"`
    NotifyExpiry                 *bool                     `json:"notify_expiry,omitempty"`
    NotifyOnCreateOrEdit         *bool                     `json:"notify_on_create_or_edit,omitempty"`
    OrgId                        *uuid.UUID                `json:"org_id,omitempty"`
    PassphraseRules              *PskPortalPassphraseRules `json:"passphrase_rules,omitempty"`
    RequiredFields               []string                  `json:"required_fields,omitempty"`
    Role                         *string                   `json:"role,omitempty"`
    Ssid                         *string                   `json:"ssid"`
    Sso                          *PskPortalSso             `json:"sso,omitempty"`
    TemplateUrl                  *string                   `json:"template_url,omitempty"`
    ThumbnailUrl                 *string                   `json:"thumbnail_url,omitempty"`
    Type                         *PskPortalTypeEnum        `json:"type,omitempty"`
    VlanId                       *int                      `json:"vlan_id,omitempty"`
}

func (p *pskPortal) validate() error {
    var errs []string
    if p.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Psk_Portal`")
    }
    if p.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `Psk_Portal`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
