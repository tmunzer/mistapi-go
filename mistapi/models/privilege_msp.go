// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// PrivilegeMsp represents a PrivilegeMsp struct.
// Privileges settings
type PrivilegeMsp struct {
    // If `scope`==`org`
    OrgId                *uuid.UUID               `json:"org_id,omitempty"`
    // Name of the org (for a site belonging to org)
    OrgName              *string                  `json:"org_name,omitempty"`
    // If `scope`==`orggroup`
    OrggroupId           *uuid.UUID               `json:"orggroup_id,omitempty"`
    // access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
    Role                 PrivilegeMspRoleEnum     `json:"role"`
    // enum: `msp`, `org`, `orggroup`
    Scope                PrivilegeMspScopeEnum    `json:"scope"`
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
    Views                []AdminPrivilegeViewEnum `json:"views,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for PrivilegeMsp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrivilegeMsp) String() string {
    return fmt.Sprintf(
    	"PrivilegeMsp[OrgId=%v, OrgName=%v, OrggroupId=%v, Role=%v, Scope=%v, Views=%v, AdditionalProperties=%v]",
    	p.OrgId, p.OrgName, p.OrggroupId, p.Role, p.Scope, p.Views, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeMsp.
// It customizes the JSON marshaling process for PrivilegeMsp objects.
func (p PrivilegeMsp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "org_id", "org_name", "orggroup_id", "role", "scope", "views"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrivilegeMsp object to a map representation for JSON marshaling.
func (p PrivilegeMsp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "org_name", "orggroup_id", "role", "scope", "views")
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
    OrgId      *uuid.UUID               `json:"org_id,omitempty"`
    OrgName    *string                  `json:"org_name,omitempty"`
    OrggroupId *uuid.UUID               `json:"orggroup_id,omitempty"`
    Role       *PrivilegeMspRoleEnum    `json:"role"`
    Scope      *PrivilegeMspScopeEnum   `json:"scope"`
    Views      []AdminPrivilegeViewEnum `json:"views,omitempty"`
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
    return errors.New(strings.Join (errs, "\n"))
}
