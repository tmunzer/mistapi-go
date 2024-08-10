package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Admin represents a Admin struct.
type Admin struct {
    AdminId              *uuid.UUID                 `json:"admin_id,omitempty"`
    // trade compliance status. enum: `blocked`, `restricted`
    ComplianceStatus     *AdminComplianceStatusEnum `json:"compliance_status,omitempty"`
    // if admin account is not an Org API Token
    Email                *string                    `json:"email,omitempty"`
    // if admin account is not an Org API Token
    EnableTwoFactor      *bool                      `json:"enable_two_factor,omitempty"`
    ExpireTime           *int                       `json:"expire_time,omitempty"`
    // if admin account is not an Org API Token
    // for an invite, this is the original first name used
    FirstName            *string                    `json:"first_name,omitempty"`
    // if admin account is not an Org API Token, how long the invite should be valid
    Hours                *int                       `json:"hours,omitempty"`
    // if admin account is not an Org API Token
    // for an invite, this is the original last name used
    LastName             *string                    `json:"last_name,omitempty"`
    // for Org API Token Only
    Name                 *string                    `json:"name,omitempty"`
    // when it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out
    NoTracking           *bool                      `json:"no_tracking,omitempty"`
    // if admin account is not an Org API Token
    OauthGoogle          *bool                      `json:"oauth_google,omitempty"`
    // if admin account is not an Org API Token
    // phone number (numbers only, including country code)
    Phone                *string                    `json:"phone,omitempty"`
    // if admin account is not an Org API Token
    // secondary phone number (numbers only, including country code)
    Phone2               *string                    `json:"phone2,omitempty"`
    // list of privileges the admin has
    Privileges           []PrivilegeSelf            `json:"privileges,omitempty"`
    SessionExpiry        *int64                     `json:"session_expiry,omitempty"`
    Tags                 []string                   `json:"tags,omitempty"`
    // if admin account is not an Org API Token
    // two factor status
    TwoFactorVerified    *bool                      `json:"two_factor_verified,omitempty"`
    // if admin account is not an Org API Token
    // an admin login via_sso is more restircted. (password and email
    // cannot be changed)
    ViaSso               *bool                      `json:"via_sso,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Admin.
// It customizes the JSON marshaling process for Admin objects.
func (a Admin) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the Admin object to a map representation for JSON marshaling.
func (a Admin) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AdminId != nil {
        structMap["admin_id"] = a.AdminId
    }
    if a.ComplianceStatus != nil {
        structMap["compliance_status"] = a.ComplianceStatus
    }
    if a.Email != nil {
        structMap["email"] = a.Email
    }
    if a.EnableTwoFactor != nil {
        structMap["enable_two_factor"] = a.EnableTwoFactor
    }
    if a.ExpireTime != nil {
        structMap["expire_time"] = a.ExpireTime
    }
    if a.FirstName != nil {
        structMap["first_name"] = a.FirstName
    }
    if a.Hours != nil {
        structMap["hours"] = a.Hours
    }
    if a.LastName != nil {
        structMap["last_name"] = a.LastName
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.NoTracking != nil {
        structMap["no_tracking"] = a.NoTracking
    }
    if a.OauthGoogle != nil {
        structMap["oauth_google"] = a.OauthGoogle
    }
    if a.Phone != nil {
        structMap["phone"] = a.Phone
    }
    if a.Phone2 != nil {
        structMap["phone2"] = a.Phone2
    }
    if a.Privileges != nil {
        structMap["privileges"] = a.Privileges
    }
    if a.SessionExpiry != nil {
        structMap["session_expiry"] = a.SessionExpiry
    }
    if a.Tags != nil {
        structMap["tags"] = a.Tags
    }
    if a.TwoFactorVerified != nil {
        structMap["two_factor_verified"] = a.TwoFactorVerified
    }
    if a.ViaSso != nil {
        structMap["via_sso"] = a.ViaSso
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Admin.
// It customizes the JSON unmarshaling process for Admin objects.
func (a *Admin) UnmarshalJSON(input []byte) error {
    var temp tempAdmin
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_id", "compliance_status", "email", "enable_two_factor", "expire_time", "first_name", "hours", "last_name", "name", "no_tracking", "oauth_google", "phone", "phone2", "privileges", "session_expiry", "tags", "two_factor_verified", "via_sso")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AdminId = temp.AdminId
    a.ComplianceStatus = temp.ComplianceStatus
    a.Email = temp.Email
    a.EnableTwoFactor = temp.EnableTwoFactor
    a.ExpireTime = temp.ExpireTime
    a.FirstName = temp.FirstName
    a.Hours = temp.Hours
    a.LastName = temp.LastName
    a.Name = temp.Name
    a.NoTracking = temp.NoTracking
    a.OauthGoogle = temp.OauthGoogle
    a.Phone = temp.Phone
    a.Phone2 = temp.Phone2
    a.Privileges = temp.Privileges
    a.SessionExpiry = temp.SessionExpiry
    a.Tags = temp.Tags
    a.TwoFactorVerified = temp.TwoFactorVerified
    a.ViaSso = temp.ViaSso
    return nil
}

// tempAdmin is a temporary struct used for validating the fields of Admin.
type tempAdmin  struct {
    AdminId           *uuid.UUID                 `json:"admin_id,omitempty"`
    ComplianceStatus  *AdminComplianceStatusEnum `json:"compliance_status,omitempty"`
    Email             *string                    `json:"email,omitempty"`
    EnableTwoFactor   *bool                      `json:"enable_two_factor,omitempty"`
    ExpireTime        *int                       `json:"expire_time,omitempty"`
    FirstName         *string                    `json:"first_name,omitempty"`
    Hours             *int                       `json:"hours,omitempty"`
    LastName          *string                    `json:"last_name,omitempty"`
    Name              *string                    `json:"name,omitempty"`
    NoTracking        *bool                      `json:"no_tracking,omitempty"`
    OauthGoogle       *bool                      `json:"oauth_google,omitempty"`
    Phone             *string                    `json:"phone,omitempty"`
    Phone2            *string                    `json:"phone2,omitempty"`
    Privileges        []PrivilegeSelf            `json:"privileges,omitempty"`
    SessionExpiry     *int64                     `json:"session_expiry,omitempty"`
    Tags              []string                   `json:"tags,omitempty"`
    TwoFactorVerified *bool                      `json:"two_factor_verified,omitempty"`
    ViaSso            *bool                      `json:"via_sso,omitempty"`
}
