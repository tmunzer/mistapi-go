package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Sdkinvite represents a Sdkinvite struct.
// SDK invite
type Sdkinvite struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    ExpireTime           *int                   `json:"expire_time,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // Name, will show up in mobile
    Name                 string                 `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Number of time this invite can be used
    Quota                *int                   `json:"quota,omitempty"`
    // Whether quota limiting is enabled
    QuotaLimited         *bool                  `json:"quota_limited,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Sdkinvite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Sdkinvite) String() string {
    return fmt.Sprintf(
    	"Sdkinvite[CreatedTime=%v, Enabled=%v, ExpireTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Quota=%v, QuotaLimited=%v, SiteId=%v, AdditionalProperties=%v]",
    	s.CreatedTime, s.Enabled, s.ExpireTime, s.Id, s.ModifiedTime, s.Name, s.OrgId, s.Quota, s.QuotaLimited, s.SiteId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Sdkinvite.
// It customizes the JSON marshaling process for Sdkinvite objects.
func (s Sdkinvite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "created_time", "enabled", "expire_time", "id", "modified_time", "name", "org_id", "quota", "quota_limited", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Sdkinvite object to a map representation for JSON marshaling.
func (s Sdkinvite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSdkinvite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "enabled", "expire_time", "id", "modified_time", "name", "org_id", "quota", "quota_limited", "site_id")
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

// tempSdkinvite is a temporary struct used for validating the fields of Sdkinvite.
type tempSdkinvite  struct {
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

func (s *tempSdkinvite) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sdkinvite`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
