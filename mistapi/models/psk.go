package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Psk represents a Psk struct.
// PSK
type Psk struct {
    // sso id for psk created from psk portal
    AdminSsoId             *string        `json:"admin_sso_id,omitempty"`
    CreatedTime            *float64       `json:"created_time,omitempty"`
    // email to send psk expiring notifications to
    Email                  *string        `json:"email,omitempty"`
    // Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)
    ExpireTime             Optional[int]  `json:"expire_time"`
    // Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire
    ExpiryNotificationTime *int           `json:"expiry_notification_time,omitempty"`
    Id                     *uuid.UUID     `json:"id,omitempty"`
    // if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`
    Mac                    *string        `json:"mac,omitempty"`
    // if `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000
    Macs                   []string       `json:"macs,omitempty"`
    // For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited)
    MaxUsage               *int           `json:"max_usage,omitempty"`
    ModifiedTime           *float64       `json:"modified_time,omitempty"`
    Name                   string         `json:"name"`
    Note                   *string        `json:"note,omitempty"`
    // If set to true, reminder notification will be sent when psk is about to expire
    NotifyExpiry           *bool          `json:"notify_expiry,omitempty"`
    // If set to true, notification will be sent when psk is created or edited
    NotifyOnCreateOrEdit   *bool          `json:"notify_on_create_or_edit,omitempty"`
    // previous passphrase of the PSK if it has been rotated
    OldPassphrase          *string        `json:"old_passphrase,omitempty"`
    OrgId                  *uuid.UUID     `json:"org_id,omitempty"`
    // passphrase of the PSK (8-63 character or 64 in hex)
    Passphrase             string         `json:"passphrase"`
    Role                   *string        `json:"role,omitempty"`
    SiteId                 *uuid.UUID     `json:"site_id,omitempty"`
    // SSID this PSK should be applicable to
    Ssid                   string         `json:"ssid"`
    // enum: `macs`, `multi`, `single`
    Usage                  *PskUsageEnum  `json:"usage,omitempty"`
    // VLAN for this PSK key
    VlanId                 *PskVlanId     `json:"vlan_id,omitempty"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Psk.
// It customizes the JSON marshaling process for Psk objects.
func (p Psk) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the Psk object to a map representation for JSON marshaling.
func (p Psk) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AdminSsoId != nil {
        structMap["admin_sso_id"] = p.AdminSsoId
    }
    if p.CreatedTime != nil {
        structMap["created_time"] = p.CreatedTime
    }
    if p.Email != nil {
        structMap["email"] = p.Email
    }
    if p.ExpireTime.IsValueSet() {
        if p.ExpireTime.Value() != nil {
            structMap["expire_time"] = p.ExpireTime.Value()
        } else {
            structMap["expire_time"] = nil
        }
    }
    if p.ExpiryNotificationTime != nil {
        structMap["expiry_notification_time"] = p.ExpiryNotificationTime
    }
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Mac != nil {
        structMap["mac"] = p.Mac
    }
    if p.Macs != nil {
        structMap["macs"] = p.Macs
    }
    if p.MaxUsage != nil {
        structMap["max_usage"] = p.MaxUsage
    }
    if p.ModifiedTime != nil {
        structMap["modified_time"] = p.ModifiedTime
    }
    structMap["name"] = p.Name
    if p.Note != nil {
        structMap["note"] = p.Note
    }
    if p.NotifyExpiry != nil {
        structMap["notify_expiry"] = p.NotifyExpiry
    }
    if p.NotifyOnCreateOrEdit != nil {
        structMap["notify_on_create_or_edit"] = p.NotifyOnCreateOrEdit
    }
    if p.OldPassphrase != nil {
        structMap["old_passphrase"] = p.OldPassphrase
    }
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    structMap["passphrase"] = p.Passphrase
    if p.Role != nil {
        structMap["role"] = p.Role
    }
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    structMap["ssid"] = p.Ssid
    if p.Usage != nil {
        structMap["usage"] = p.Usage
    }
    if p.VlanId != nil {
        structMap["vlan_id"] = p.VlanId.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Psk.
// It customizes the JSON unmarshaling process for Psk objects.
func (p *Psk) UnmarshalJSON(input []byte) error {
    var temp psk
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_sso_id", "created_time", "email", "expire_time", "expiry_notification_time", "id", "mac", "macs", "max_usage", "modified_time", "name", "note", "notify_expiry", "notify_on_create_or_edit", "old_passphrase", "org_id", "passphrase", "role", "site_id", "ssid", "usage", "vlan_id")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.AdminSsoId = temp.AdminSsoId
    p.CreatedTime = temp.CreatedTime
    p.Email = temp.Email
    p.ExpireTime = temp.ExpireTime
    p.ExpiryNotificationTime = temp.ExpiryNotificationTime
    p.Id = temp.Id
    p.Mac = temp.Mac
    p.Macs = temp.Macs
    p.MaxUsage = temp.MaxUsage
    p.ModifiedTime = temp.ModifiedTime
    p.Name = *temp.Name
    p.Note = temp.Note
    p.NotifyExpiry = temp.NotifyExpiry
    p.NotifyOnCreateOrEdit = temp.NotifyOnCreateOrEdit
    p.OldPassphrase = temp.OldPassphrase
    p.OrgId = temp.OrgId
    p.Passphrase = *temp.Passphrase
    p.Role = temp.Role
    p.SiteId = temp.SiteId
    p.Ssid = *temp.Ssid
    p.Usage = temp.Usage
    p.VlanId = temp.VlanId
    return nil
}

// psk is a temporary struct used for validating the fields of Psk.
type psk  struct {
    AdminSsoId             *string       `json:"admin_sso_id,omitempty"`
    CreatedTime            *float64      `json:"created_time,omitempty"`
    Email                  *string       `json:"email,omitempty"`
    ExpireTime             Optional[int] `json:"expire_time"`
    ExpiryNotificationTime *int          `json:"expiry_notification_time,omitempty"`
    Id                     *uuid.UUID    `json:"id,omitempty"`
    Mac                    *string       `json:"mac,omitempty"`
    Macs                   []string      `json:"macs,omitempty"`
    MaxUsage               *int          `json:"max_usage,omitempty"`
    ModifiedTime           *float64      `json:"modified_time,omitempty"`
    Name                   *string       `json:"name"`
    Note                   *string       `json:"note,omitempty"`
    NotifyExpiry           *bool         `json:"notify_expiry,omitempty"`
    NotifyOnCreateOrEdit   *bool         `json:"notify_on_create_or_edit,omitempty"`
    OldPassphrase          *string       `json:"old_passphrase,omitempty"`
    OrgId                  *uuid.UUID    `json:"org_id,omitempty"`
    Passphrase             *string       `json:"passphrase"`
    Role                   *string       `json:"role,omitempty"`
    SiteId                 *uuid.UUID    `json:"site_id,omitempty"`
    Ssid                   *string       `json:"ssid"`
    Usage                  *PskUsageEnum `json:"usage,omitempty"`
    VlanId                 *PskVlanId    `json:"vlan_id,omitempty"`
}

func (p *psk) validate() error {
    var errs []string
    if p.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Psk`")
    }
    if p.Passphrase == nil {
        errs = append(errs, "required field `passphrase` is missing for type `Psk`")
    }
    if p.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `Psk`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
