package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// PrivilegeOrg represents a PrivilegeOrg struct.
// Privilieges settings
type PrivilegeOrg struct {
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 PrivilegeOrgRoleEnum   `json:"role"`
    // enum: `org`, `site`, `sitegroup`
    Scope                PrivilegeOrgScopeEnum  `json:"scope"`
    // if `scope`==`site`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // if `scope`==`sitegroup`
    SitegroupId          *uuid.UUID             `json:"sitegroup_id,omitempty"`
    // Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
    // You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.
    // Below are the list of supported UI views. Note that this is UI only feature
    // Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
    // You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.
    // Below are the list of supported UI views. Note that this is UI only feature
    // | UI View | Description |
    // | --- | --- |
    // | `reporting` | full access to all analytics tools |
    // | `marketing` | can view analytics and location maps |
    // | `location` | can view and manage location maps |
    // | `security` | can view and manage WLAN, rogues and authentication |
    // | `switch_admin` | can view and manage Switch ports |
    // | `mxedge_admin` | can view and manage Mist edges and Mist tunnels |
    // | `lobby_admin` | full access to Org and Site Pre-shared keys |
    Views                *PrivilegeOrgViewsEnum `json:"views,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeOrg.
// It customizes the JSON marshaling process for PrivilegeOrg objects.
func (p PrivilegeOrg) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrivilegeOrg object to a map representation for JSON marshaling.
func (p PrivilegeOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["role"] = p.Role
    structMap["scope"] = p.Scope
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    if p.SitegroupId != nil {
        structMap["sitegroup_id"] = p.SitegroupId
    }
    if p.Views != nil {
        structMap["views"] = p.Views
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrivilegeOrg.
// It customizes the JSON unmarshaling process for PrivilegeOrg objects.
func (p *PrivilegeOrg) UnmarshalJSON(input []byte) error {
    var temp tempPrivilegeOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "role", "scope", "site_id", "sitegroup_id", "views")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Role = *temp.Role
    p.Scope = *temp.Scope
    p.SiteId = temp.SiteId
    p.SitegroupId = temp.SitegroupId
    p.Views = temp.Views
    return nil
}

// tempPrivilegeOrg is a temporary struct used for validating the fields of PrivilegeOrg.
type tempPrivilegeOrg  struct {
    Role        *PrivilegeOrgRoleEnum  `json:"role"`
    Scope       *PrivilegeOrgScopeEnum `json:"scope"`
    SiteId      *uuid.UUID             `json:"site_id,omitempty"`
    SitegroupId *uuid.UUID             `json:"sitegroup_id,omitempty"`
    Views       *PrivilegeOrgViewsEnum `json:"views,omitempty"`
}

func (p *tempPrivilegeOrg) validate() error {
    var errs []string
    if p.Role == nil {
        errs = append(errs, "required field `role` is missing for type `privilege_org`")
    }
    if p.Scope == nil {
        errs = append(errs, "required field `scope` is missing for type `privilege_org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
