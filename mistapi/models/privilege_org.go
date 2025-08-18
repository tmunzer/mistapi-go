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

// PrivilegeOrg represents a PrivilegeOrg struct.
// Privileges settings
type PrivilegeOrg struct {
	// If `scope`==`org`
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
	Role PrivilegeOrgRoleEnum `json:"role"`
	// enum: `org`, `site`, `sitegroup`
	Scope PrivilegeOrgScopeEnum `json:"scope"`
	// If `scope`==`site`
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// If `scope`==`sitegroup`
	SitegroupId *uuid.UUID `json:"sitegroup_id,omitempty"`
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

// String implements the fmt.Stringer interface for PrivilegeOrg,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrivilegeOrg) String() string {
	return fmt.Sprintf(
		"PrivilegeOrg[OrgId=%v, Role=%v, Scope=%v, SiteId=%v, SitegroupId=%v, Views=%v, AdditionalProperties=%v]",
		p.OrgId, p.Role, p.Scope, p.SiteId, p.SitegroupId, p.Views, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrivilegeOrg.
// It customizes the JSON marshaling process for PrivilegeOrg objects.
func (p PrivilegeOrg) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"org_id", "role", "scope", "site_id", "sitegroup_id", "views"); err != nil {
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "role", "scope", "site_id", "sitegroup_id", "views")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.OrgId = temp.OrgId
	p.Role = *temp.Role
	p.Scope = *temp.Scope
	p.SiteId = temp.SiteId
	p.SitegroupId = temp.SitegroupId
	p.Views = temp.Views
	return nil
}

// tempPrivilegeOrg is a temporary struct used for validating the fields of PrivilegeOrg.
type tempPrivilegeOrg struct {
	OrgId       *uuid.UUID               `json:"org_id,omitempty"`
	Role        *PrivilegeOrgRoleEnum    `json:"role"`
	Scope       *PrivilegeOrgScopeEnum   `json:"scope"`
	SiteId      *uuid.UUID               `json:"site_id,omitempty"`
	SitegroupId *uuid.UUID               `json:"sitegroup_id,omitempty"`
	Views       []AdminPrivilegeViewEnum `json:"views,omitempty"`
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
	return errors.New(strings.Join(errs, "\n"))
}
