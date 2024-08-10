package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// PrivilegeMsp represents a PrivilegeMsp struct.
// Privilieges settings
type PrivilegeMsp struct {
    // if `scope`==`org`
    OrgId                *uuid.UUID            `json:"org_id,omitempty"`
    // name of the org (for a site belonging to org)
    OrgName              *string               `json:"org_name,omitempty"`
    // if `scope`==`orggroup`
    OrggroupId           *uuid.UUID            `json:"orggroup_id,omitempty"`
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 PrivilegeMspRoleEnum  `json:"role"`
    // enum: `msp`, `org`, `orggroup`
    Scope                PrivilegeMspScopeEnum `json:"scope"`
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
    Views                *PrivilegeMspViewEnum `json:"views,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeMsp.
// It customizes the JSON marshaling process for PrivilegeMsp objects.
func (p PrivilegeMsp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrivilegeMsp object to a map representation for JSON marshaling.
func (p PrivilegeMsp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.OrgId != nil {
        structMap["org_id"] = p.OrgId
    }
    if p.OrgName != nil {
        structMap["org_name"] = p.OrgName
    }
    if p.OrggroupId != nil {
        structMap["orggroup_id"] = p.OrggroupId
    }
    structMap["role"] = p.Role
    structMap["scope"] = p.Scope
    if p.Views != nil {
        structMap["views"] = p.Views
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrivilegeMsp.
// It customizes the JSON unmarshaling process for PrivilegeMsp objects.
func (p *PrivilegeMsp) UnmarshalJSON(input []byte) error {
    var temp tempPrivilegeMsp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "org_id", "org_name", "orggroup_id", "role", "scope", "views")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.OrgId = temp.OrgId
    p.OrgName = temp.OrgName
    p.OrggroupId = temp.OrggroupId
    p.Role = *temp.Role
    p.Scope = *temp.Scope
    p.Views = temp.Views
    return nil
}

// tempPrivilegeMsp is a temporary struct used for validating the fields of PrivilegeMsp.
type tempPrivilegeMsp  struct {
    OrgId      *uuid.UUID             `json:"org_id,omitempty"`
    OrgName    *string                `json:"org_name,omitempty"`
    OrggroupId *uuid.UUID             `json:"orggroup_id,omitempty"`
    Role       *PrivilegeMspRoleEnum  `json:"role"`
    Scope      *PrivilegeMspScopeEnum `json:"scope"`
    Views      *PrivilegeMspViewEnum  `json:"views,omitempty"`
}

func (p *tempPrivilegeMsp) validate() error {
    var errs []string
    if p.Role == nil {
        errs = append(errs, "required field `role` is missing for type `privilege_msp`")
    }
    if p.Scope == nil {
        errs = append(errs, "required field `scope` is missing for type `privilege_msp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
