package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// PrivilegeSelf represents a PrivilegeSelf struct.
// Privilieges settings
type PrivilegeSelf struct {
    MspId                *uuid.UUID              `json:"msp_id,omitempty"`
    // logo of the MSP (if the MSP belongs to an Advanced tier)
    MspLogoUrl           *string                 `json:"msp_logo_url,omitempty"`
    // name of the MSP (if the org belongs to an MSP)
    MspName              Optional[string]        `json:"msp_name"`
    // custom url of the MSP (if the MSP belongs to an Advanced tier)
    MspUrl               *string                 `json:"msp_url,omitempty"`
    // name of the org/site/MSP depending on object scope
    Name                 *string                 `json:"name,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // name of the org (for a site belonging to org)
    OrgName              *string                 `json:"org_name,omitempty"`
    // if `scope`==`orggroup`
    OrggroupIds          []uuid.UUID             `json:"orggroup_ids,omitempty"`
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 PrivilegeSelfRoleEnum   `json:"role"`
    // enum: `msp`, `org`, `orggroup`, `site`, `sitegroup`
    Scope                PrivilegeSelfScopeEnum  `json:"scope"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    SitegroupIds         []uuid.UUID             `json:"sitegroup_ids,omitempty"`
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
    Views                *PrivilegeSelfViewsEnum `json:"views,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeSelf.
// It customizes the JSON marshaling process for PrivilegeSelf objects.
func (p PrivilegeSelf) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrivilegeSelf object to a map representation for JSON marshaling.
func (p PrivilegeSelf) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.MspId != nil {
        structMap["msp_id"] = p.MspId
    }
    if p.MspLogoUrl != nil {
        structMap["msp_logo_url"] = p.MspLogoUrl
    }
    if p.MspName.IsValueSet() {
        if p.MspName.Value() != nil {
            structMap["msp_name"] = p.MspName.Value()
        } else {
            structMap["msp_name"] = nil
        }
    }
    if p.MspUrl != nil {
        structMap["msp_url"] = p.MspUrl
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    if p.OrgName != nil {
        structMap["org_name"] = p.OrgName
    }
    if p.OrggroupIds != nil {
        structMap["orggroup_ids"] = p.OrggroupIds
    }
    structMap["role"] = p.Role
    structMap["scope"] = p.Scope
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    if p.SitegroupIds != nil {
        structMap["sitegroup_ids"] = p.SitegroupIds
    }
    if p.Views != nil {
        structMap["views"] = p.Views
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrivilegeSelf.
// It customizes the JSON unmarshaling process for PrivilegeSelf objects.
func (p *PrivilegeSelf) UnmarshalJSON(input []byte) error {
    var temp tempPrivilegeSelf
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "msp_id", "msp_logo_url", "msp_name", "msp_url", "name", "org_id", "org_name", "orggroup_ids", "role", "scope", "site_id", "sitegroup_ids", "views")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.MspId = temp.MspId
    p.MspLogoUrl = temp.MspLogoUrl
    p.MspName = temp.MspName
    p.MspUrl = temp.MspUrl
    p.Name = temp.Name
    p.OrgId = temp.OrgId
    p.OrgName = temp.OrgName
    p.OrggroupIds = temp.OrggroupIds
    p.Role = *temp.Role
    p.Scope = *temp.Scope
    p.SiteId = temp.SiteId
    p.SitegroupIds = temp.SitegroupIds
    p.Views = temp.Views
    return nil
}

// tempPrivilegeSelf is a temporary struct used for validating the fields of PrivilegeSelf.
type tempPrivilegeSelf  struct {
    MspId        *uuid.UUID              `json:"msp_id,omitempty"`
    MspLogoUrl   *string                 `json:"msp_logo_url,omitempty"`
    MspName      Optional[string]        `json:"msp_name"`
    MspUrl       *string                 `json:"msp_url,omitempty"`
    Name         *string                 `json:"name,omitempty"`
    OrgId        *uuid.UUID              `json:"org_id,omitempty"`
    OrgName      *string                 `json:"org_name,omitempty"`
    OrggroupIds  []uuid.UUID             `json:"orggroup_ids,omitempty"`
    Role         *PrivilegeSelfRoleEnum  `json:"role"`
    Scope        *PrivilegeSelfScopeEnum `json:"scope"`
    SiteId       *uuid.UUID              `json:"site_id,omitempty"`
    SitegroupIds []uuid.UUID             `json:"sitegroup_ids,omitempty"`
    Views        *PrivilegeSelfViewsEnum `json:"views,omitempty"`
}

func (p *tempPrivilegeSelf) validate() error {
    var errs []string
    if p.Role == nil {
        errs = append(errs, "required field `role` is missing for type `privilege_self`")
    }
    if p.Scope == nil {
        errs = append(errs, "required field `scope` is missing for type `privilege_self`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
