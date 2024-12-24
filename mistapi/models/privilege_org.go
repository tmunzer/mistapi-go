package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// PrivilegeOrg represents a PrivilegeOrg struct.
// Privilieges settings
type PrivilegeOrg struct {
    // if `scope`==`org`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 PrivilegeOrgRoleEnum   `json:"role"`
    // enum: `org`, `site`, `sitegroup`
    Scope                PrivilegeOrgScopeEnum  `json:"scope"`
    // if `scope`==`site`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // if `scope`==`sitegroup`
    SitegroupId          *uuid.UUID             `json:"sitegroup_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PrivilegeOrg,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrivilegeOrg) String() string {
    return fmt.Sprintf(
    	"PrivilegeOrg[OrgId=%v, Role=%v, Scope=%v, SiteId=%v, SitegroupId=%v, AdditionalProperties=%v]",
    	p.OrgId, p.Role, p.Scope, p.SiteId, p.SitegroupId, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeOrg.
// It customizes the JSON marshaling process for PrivilegeOrg objects.
func (p PrivilegeOrg) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "org_id", "role", "scope", "site_id", "sitegroup_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrivilegeOrg object to a map representation for JSON marshaling.
func (p PrivilegeOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    structMap["role"] = p.Role
    structMap["scope"] = p.Scope
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    if p.SitegroupId != nil {
        structMap["sitegroup_id"] = p.SitegroupId
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "role", "scope", "site_id", "sitegroup_id")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.OrgId = temp.OrgId
    p.Role = *temp.Role
    p.Scope = *temp.Scope
    p.SiteId = temp.SiteId
    p.SitegroupId = temp.SitegroupId
    return nil
}

// tempPrivilegeOrg is a temporary struct used for validating the fields of PrivilegeOrg.
type tempPrivilegeOrg  struct {
    OrgId       *uuid.UUID             `json:"org_id,omitempty"`
    Role        *PrivilegeOrgRoleEnum  `json:"role"`
    Scope       *PrivilegeOrgScopeEnum `json:"scope"`
    SiteId      *uuid.UUID             `json:"site_id,omitempty"`
    SitegroupId *uuid.UUID             `json:"sitegroup_id,omitempty"`
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
