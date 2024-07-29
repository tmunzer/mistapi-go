package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Admin represents a Admin struct.
type Admin struct {
    AdminId              *uuid.UUID                 `json:"admin_id,omitempty"`
    // trade compliance status. enum: `blocked`, `restricted`
    ComplianceStatus     *AdminComplianceStatusEnum `json:"compliance_status,omitempty"`
    Email                string                     `json:"email"`
    EnableTwoFactor      *bool                      `json:"enable_two_factor,omitempty"`
    ExpireTime           *int                       `json:"expire_time,omitempty"`
    // for an invite, this is the original first name used
    FirstName            string                     `json:"first_name"`
    // how long the invite should be valid
    Hours                *int                       `json:"hours,omitempty"`
    // for an invite, this is the original last name used
    LastName             string                     `json:"last_name"`
    OauthGoogle          *bool                      `json:"oauth_google,omitempty"`
    // phone number (numbers only, including country code)
    Phone                *string                    `json:"phone,omitempty"`
    // secondary phone number (numbers only, including country code)
    Phone2               *string                    `json:"phone2,omitempty"`
    // list of privileges the admin has
    Privileges           []PrivilegeSelf            `json:"privileges,omitempty"`
    SessionExpiry        *int64                     `json:"session_expiry,omitempty"`
    Tags                 []string                   `json:"tags,omitempty"`
    // two factor status
    TwoFactorVerified    *bool                      `json:"two_factor_verified,omitempty"`
    // an admin alogin via_sso is more restircted. (password and email cannot be changed)
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
    structMap["email"] = a.Email
    if a.EnableTwoFactor != nil {
        structMap["enable_two_factor"] = a.EnableTwoFactor
    }
    if a.ExpireTime != nil {
        structMap["expire_time"] = a.ExpireTime
    }
    structMap["first_name"] = a.FirstName
    if a.Hours != nil {
        structMap["hours"] = a.Hours
    }
    structMap["last_name"] = a.LastName
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
    var temp admin
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_id", "compliance_status", "email", "enable_two_factor", "expire_time", "first_name", "hours", "last_name", "oauth_google", "phone", "phone2", "privileges", "session_expiry", "tags", "two_factor_verified", "via_sso")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AdminId = temp.AdminId
    a.ComplianceStatus = temp.ComplianceStatus
    a.Email = *temp.Email
    a.EnableTwoFactor = temp.EnableTwoFactor
    a.ExpireTime = temp.ExpireTime
    a.FirstName = *temp.FirstName
    a.Hours = temp.Hours
    a.LastName = *temp.LastName
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

// admin is a temporary struct used for validating the fields of Admin.
type admin  struct {
    AdminId           *uuid.UUID                 `json:"admin_id,omitempty"`
    ComplianceStatus  *AdminComplianceStatusEnum `json:"compliance_status,omitempty"`
    Email             *string                    `json:"email"`
    EnableTwoFactor   *bool                      `json:"enable_two_factor,omitempty"`
    ExpireTime        *int                       `json:"expire_time,omitempty"`
    FirstName         *string                    `json:"first_name"`
    Hours             *int                       `json:"hours,omitempty"`
    LastName          *string                    `json:"last_name"`
    OauthGoogle       *bool                      `json:"oauth_google,omitempty"`
    Phone             *string                    `json:"phone,omitempty"`
    Phone2            *string                    `json:"phone2,omitempty"`
    Privileges        []PrivilegeSelf            `json:"privileges,omitempty"`
    SessionExpiry     *int64                     `json:"session_expiry,omitempty"`
    Tags              []string                   `json:"tags,omitempty"`
    TwoFactorVerified *bool                      `json:"two_factor_verified,omitempty"`
    ViaSso            *bool                      `json:"via_sso,omitempty"`
}

func (a *admin) validate() error {
    var errs []string
    if a.Email == nil {
        errs = append(errs, "required field `email` is missing for type `Admin`")
    }
    if a.FirstName == nil {
        errs = append(errs, "required field `first_name` is missing for type `Admin`")
    }
    if a.LastName == nil {
        errs = append(errs, "required field `last_name` is missing for type `Admin`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
