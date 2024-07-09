package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SsoRoleOrg represents a SsoRoleOrg struct.
// SSO Role response
type SsoRoleOrg struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    ForSite              *bool          `json:"for_site,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    MspId                *uuid.UUID     `json:"msp_id,omitempty"`
    Name                 string         `json:"name"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    Privileges           []PrivilegeOrg `json:"privileges"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsoRoleOrg.
// It customizes the JSON marshaling process for SsoRoleOrg objects.
func (s SsoRoleOrg) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SsoRoleOrg object to a map representation for JSON marshaling.
func (s SsoRoleOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.MspId != nil {
        structMap["msp_id"] = s.MspId
    }
    structMap["name"] = s.Name
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    structMap["privileges"] = s.Privileges
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoRoleOrg.
// It customizes the JSON unmarshaling process for SsoRoleOrg objects.
func (s *SsoRoleOrg) UnmarshalJSON(input []byte) error {
    var temp ssoRoleOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "for_site", "id", "modified_time", "msp_id", "name", "org_id", "privileges", "site_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CreatedTime = temp.CreatedTime
    s.ForSite = temp.ForSite
    s.Id = temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.MspId = temp.MspId
    s.Name = *temp.Name
    s.OrgId = temp.OrgId
    s.Privileges = *temp.Privileges
    s.SiteId = temp.SiteId
    return nil
}

// ssoRoleOrg is a temporary struct used for validating the fields of SsoRoleOrg.
type ssoRoleOrg  struct {
    CreatedTime  *float64        `json:"created_time,omitempty"`
    ForSite      *bool           `json:"for_site,omitempty"`
    Id           *uuid.UUID      `json:"id,omitempty"`
    ModifiedTime *float64        `json:"modified_time,omitempty"`
    MspId        *uuid.UUID      `json:"msp_id,omitempty"`
    Name         *string         `json:"name"`
    OrgId        *uuid.UUID      `json:"org_id,omitempty"`
    Privileges   *[]PrivilegeOrg `json:"privileges"`
    SiteId       *uuid.UUID      `json:"site_id,omitempty"`
}

func (s *ssoRoleOrg) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sso_Role_Org`")
    }
    if s.Privileges == nil {
        errs = append(errs, "required field `privileges` is missing for type `Sso_Role_Org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
