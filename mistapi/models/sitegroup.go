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

// Sitegroup represents a Sitegroup struct.
// Sites Group
type Sitegroup struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 string                 `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteIds              []uuid.UUID            `json:"site_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Sitegroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Sitegroup) String() string {
    return fmt.Sprintf(
    	"Sitegroup[CreatedTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SiteIds=%v, AdditionalProperties=%v]",
    	s.CreatedTime, s.Id, s.ModifiedTime, s.Name, s.OrgId, s.SiteIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Sitegroup.
// It customizes the JSON marshaling process for Sitegroup objects.
func (s Sitegroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "created_time", "id", "modified_time", "name", "org_id", "site_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Sitegroup object to a map representation for JSON marshaling.
func (s Sitegroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
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
    if s.SiteIds != nil {
        structMap["site_ids"] = s.SiteIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Sitegroup.
// It customizes the JSON unmarshaling process for Sitegroup objects.
func (s *Sitegroup) UnmarshalJSON(input []byte) error {
    var temp tempSitegroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "name", "org_id", "site_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CreatedTime = temp.CreatedTime
    s.Id = temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.Name = *temp.Name
    s.OrgId = temp.OrgId
    s.SiteIds = temp.SiteIds
    return nil
}

// tempSitegroup is a temporary struct used for validating the fields of Sitegroup.
type tempSitegroup  struct {
    CreatedTime  *float64    `json:"created_time,omitempty"`
    Id           *uuid.UUID  `json:"id,omitempty"`
    ModifiedTime *float64    `json:"modified_time,omitempty"`
    Name         *string     `json:"name"`
    OrgId        *uuid.UUID  `json:"org_id,omitempty"`
    SiteIds      []uuid.UUID `json:"site_ids,omitempty"`
}

func (s *tempSitegroup) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sitegroup`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
