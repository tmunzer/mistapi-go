package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// AdminPrivilege represents a AdminPrivilege struct.
// Privilieges settings
type AdminPrivilege struct {
    // required if `scope`==`msp`
    MspId                *uuid.UUID              `json:"msp_id,omitempty"`
    // logo of the MSP (if the MSP belongs to an Advanced tier)
    MspLogoUrl           *string                 `json:"msp_logo_url,omitempty"`
    // name of the MSP (if the org belongs to an MSP)
    MspName              Optional[string]        `json:"msp_name"`
    // custom url of the MSP (if the MSP belongs to an Advanced tier)
    MspUrl               *string                 `json:"msp_url,omitempty"`
    // name of the org/site/MSP depending on object scope
    Name                 *string                 `json:"name,omitempty"`
    // required if `scope`==`org`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // name of the org (for a site belonging to org)
    OrgName              *string                 `json:"org_name,omitempty"`
    // if `scope`==`orggroup`
    OrggroupIds          []uuid.UUID             `json:"orggroup_ids,omitempty"`
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 AdminPrivilegeRoleEnum  `json:"role"`
    // enum: `msp`, `org`, `orggroup`, `site`, `sitegroup`
    Scope                AdminPrivilegeScopeEnum `json:"scope"`
    // required if `scope`==`site`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    SitegroupIds         []uuid.UUID             `json:"sitegroup_ids,omitempty"`
    // Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
    // You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.
    // Below are the list of supported UI views. Note that this is UI only feature.
    // | UI View | Required Role | Description |
    // | --- | --- | --- |
    // | `reporting` | `read` | full access to all analytics tools |
    // | `marketing` | `read` | can view analytics and location maps |
    // | `super_observer` | `read` | can view all the organization except the subscription page |
    // | `location` | `write` | can view and manage location maps, can view analytics |
    // | `security` | `write` | can view and manage site labels, policies and security |
    // | `switch_admin` | `helpdesk` | can view and manage Switch ports, can view wired clients |
    // | `mxedge_admin` | `admin` | can view and manage Mist edges and Mist tunnels |
    // | `lobby_admin` | `admin` | full access to Org and Site Pre-shared keys |
    Views                *AdminPrivilegeViewEnum `json:"views,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for AdminPrivilege,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AdminPrivilege) String() string {
    return fmt.Sprintf(
    	"AdminPrivilege[MspId=%v, MspLogoUrl=%v, MspName=%v, MspUrl=%v, Name=%v, OrgId=%v, OrgName=%v, OrggroupIds=%v, Role=%v, Scope=%v, SiteId=%v, SitegroupIds=%v, Views=%v, AdditionalProperties=%v]",
    	a.MspId, a.MspLogoUrl, a.MspName, a.MspUrl, a.Name, a.OrgId, a.OrgName, a.OrggroupIds, a.Role, a.Scope, a.SiteId, a.SitegroupIds, a.Views, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AdminPrivilege.
// It customizes the JSON marshaling process for AdminPrivilege objects.
func (a AdminPrivilege) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "msp_id", "msp_logo_url", "msp_name", "msp_url", "name", "org_id", "org_name", "orggroup_ids", "role", "scope", "site_id", "sitegroup_ids", "views"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AdminPrivilege object to a map representation for JSON marshaling.
func (a AdminPrivilege) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.MspId != nil {
        structMap["msp_id"] = a.MspId
    }
    if a.MspLogoUrl != nil {
        structMap["msp_logo_url"] = a.MspLogoUrl
    }
    if a.MspName.IsValueSet() {
        if a.MspName.Value() != nil {
            structMap["msp_name"] = a.MspName.Value()
        } else {
            structMap["msp_name"] = nil
        }
    }
    if a.MspUrl != nil {
        structMap["msp_url"] = a.MspUrl
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.OrgName != nil {
        structMap["org_name"] = a.OrgName
    }
    if a.OrggroupIds != nil {
        structMap["orggroup_ids"] = a.OrggroupIds
    }
    structMap["role"] = a.Role
    structMap["scope"] = a.Scope
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.SitegroupIds != nil {
        structMap["sitegroup_ids"] = a.SitegroupIds
    }
    if a.Views != nil {
        structMap["views"] = a.Views
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdminPrivilege.
// It customizes the JSON unmarshaling process for AdminPrivilege objects.
func (a *AdminPrivilege) UnmarshalJSON(input []byte) error {
    var temp tempAdminPrivilege
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "msp_id", "msp_logo_url", "msp_name", "msp_url", "name", "org_id", "org_name", "orggroup_ids", "role", "scope", "site_id", "sitegroup_ids", "views")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.MspId = temp.MspId
    a.MspLogoUrl = temp.MspLogoUrl
    a.MspName = temp.MspName
    a.MspUrl = temp.MspUrl
    a.Name = temp.Name
    a.OrgId = temp.OrgId
    a.OrgName = temp.OrgName
    a.OrggroupIds = temp.OrggroupIds
    a.Role = *temp.Role
    a.Scope = *temp.Scope
    a.SiteId = temp.SiteId
    a.SitegroupIds = temp.SitegroupIds
    a.Views = temp.Views
    return nil
}

// tempAdminPrivilege is a temporary struct used for validating the fields of AdminPrivilege.
type tempAdminPrivilege  struct {
    MspId        *uuid.UUID               `json:"msp_id,omitempty"`
    MspLogoUrl   *string                  `json:"msp_logo_url,omitempty"`
    MspName      Optional[string]         `json:"msp_name"`
    MspUrl       *string                  `json:"msp_url,omitempty"`
    Name         *string                  `json:"name,omitempty"`
    OrgId        *uuid.UUID               `json:"org_id,omitempty"`
    OrgName      *string                  `json:"org_name,omitempty"`
    OrggroupIds  []uuid.UUID              `json:"orggroup_ids,omitempty"`
    Role         *AdminPrivilegeRoleEnum  `json:"role"`
    Scope        *AdminPrivilegeScopeEnum `json:"scope"`
    SiteId       *uuid.UUID               `json:"site_id,omitempty"`
    SitegroupIds []uuid.UUID              `json:"sitegroup_ids,omitempty"`
    Views        *AdminPrivilegeViewEnum  `json:"views,omitempty"`
}

func (a *tempAdminPrivilege) validate() error {
    var errs []string
    if a.Role == nil {
        errs = append(errs, "required field `role` is missing for type `admin_privilege`")
    }
    if a.Scope == nil {
        errs = append(errs, "required field `scope` is missing for type `admin_privilege`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
