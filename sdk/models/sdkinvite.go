package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Sdkinvite represents a Sdkinvite struct.
// SDK invite
type Sdkinvite struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    Enabled              *bool          `json:"enabled,omitempty"`
    ExpireTime           *int           `json:"expire_time,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    // name, will show up in mobile
    Name                 string         `json:"name"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    // number of time this invite can be used
    Quota                *int           `json:"quota,omitempty"`
    // whether quota limiting is enabled
    QuotaLimited         *bool          `json:"quota_limited,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Sdkinvite.
// It customizes the JSON marshaling process for Sdkinvite objects.
func (s Sdkinvite) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Sdkinvite object to a map representation for JSON marshaling.
func (s Sdkinvite) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.ExpireTime != nil {
        structMap["expire_time"] = s.ExpireTime
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    structMap["name"] = s.Name
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.Quota != nil {
        structMap["quota"] = s.Quota
    }
    if s.QuotaLimited != nil {
        structMap["quota_limited"] = s.QuotaLimited
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Sdkinvite.
// It customizes the JSON unmarshaling process for Sdkinvite objects.
func (s *Sdkinvite) UnmarshalJSON(input []byte) error {
    var temp sdkinvite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "enabled", "expire_time", "id", "modified_time", "name", "org_id", "quota", "quota_limited", "site_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CreatedTime = temp.CreatedTime
    s.Enabled = temp.Enabled
    s.ExpireTime = temp.ExpireTime
    s.Id = temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.Name = *temp.Name
    s.OrgId = temp.OrgId
    s.Quota = temp.Quota
    s.QuotaLimited = temp.QuotaLimited
    s.SiteId = temp.SiteId
    return nil
}

// sdkinvite is a temporary struct used for validating the fields of Sdkinvite.
type sdkinvite  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    Enabled      *bool      `json:"enabled,omitempty"`
    ExpireTime   *int       `json:"expire_time,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    Name         *string    `json:"name"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    Quota        *int       `json:"quota,omitempty"`
    QuotaLimited *bool      `json:"quota_limited,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
}

func (s *sdkinvite) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sdkinvite`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
