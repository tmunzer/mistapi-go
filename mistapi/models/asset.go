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

// Asset represents a Asset struct.
// Asset
type Asset struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // Bluetooth MAC
    Mac                  string                 `json:"mac"`
    MapId                *uuid.UUID             `json:"map_id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // Name / label of the device
    Name                 string                 `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    TagId                *uuid.UUID             `json:"tag_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Asset,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a Asset) String() string {
    return fmt.Sprintf(
    	"Asset[CreatedTime=%v, ForSite=%v, Id=%v, Mac=%v, MapId=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SiteId=%v, TagId=%v, AdditionalProperties=%v]",
    	a.CreatedTime, a.ForSite, a.Id, a.Mac, a.MapId, a.ModifiedTime, a.Name, a.OrgId, a.SiteId, a.TagId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Asset.
// It customizes the JSON marshaling process for Asset objects.
func (a Asset) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "created_time", "for_site", "id", "mac", "map_id", "modified_time", "name", "org_id", "site_id", "tag_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the Asset object to a map representation for JSON marshaling.
func (a Asset) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    if a.ForSite != nil {
        structMap["for_site"] = a.ForSite
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    structMap["mac"] = a.Mac
    if a.MapId != nil {
        structMap["map_id"] = a.MapId
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    structMap["name"] = a.Name
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.TagId != nil {
        structMap["tag_id"] = a.TagId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Asset.
// It customizes the JSON unmarshaling process for Asset objects.
func (a *Asset) UnmarshalJSON(input []byte) error {
    var temp tempAsset
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "for_site", "id", "mac", "map_id", "modified_time", "name", "org_id", "site_id", "tag_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CreatedTime = temp.CreatedTime
    a.ForSite = temp.ForSite
    a.Id = temp.Id
    a.Mac = *temp.Mac
    a.MapId = temp.MapId
    a.ModifiedTime = temp.ModifiedTime
    a.Name = *temp.Name
    a.OrgId = temp.OrgId
    a.SiteId = temp.SiteId
    a.TagId = temp.TagId
    return nil
}

// tempAsset is a temporary struct used for validating the fields of Asset.
type tempAsset  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    ForSite      *bool      `json:"for_site,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    Mac          *string    `json:"mac"`
    MapId        *uuid.UUID `json:"map_id,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    Name         *string    `json:"name"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    TagId        *uuid.UUID `json:"tag_id,omitempty"`
}

func (a *tempAsset) validate() error {
    var errs []string
    if a.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `asset`")
    }
    if a.Name == nil {
        errs = append(errs, "required field `name` is missing for type `asset`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
