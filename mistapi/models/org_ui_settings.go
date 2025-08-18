// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OrgUiSettings represents a OrgUiSettings struct.
type OrgUiSettings struct {
    // When the object has been created, in epoch
    CreatedTime          *float64                  `json:"created_time,omitempty"`
    Description          *string                   `json:"description,omitempty"`
    ForSite              *bool                     `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                `json:"id,omitempty"`
    // Whether this is a custom databoard or not
    IsCustomDataboard    *bool                     `json:"isCustomDataboard,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                  `json:"modified_time,omitempty"`
    // Name of the databoard
    Name                 *string                   `json:"name,omitempty"`
    OrgId                *uuid.UUID                `json:"org_id,omitempty"`
    // enum: `marvisdashboard`
    Purpose              *OrgUiSettingsPurposeEnum `json:"purpose,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    // List of tiles in the databoard
    Tiles                []OrgUiSettingsTile       `json:"tiles,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for OrgUiSettings,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgUiSettings) String() string {
    return fmt.Sprintf(
    	"OrgUiSettings[CreatedTime=%v, Description=%v, ForSite=%v, Id=%v, IsCustomDataboard=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Purpose=%v, SiteId=%v, Tiles=%v, AdditionalProperties=%v]",
    	o.CreatedTime, o.Description, o.ForSite, o.Id, o.IsCustomDataboard, o.ModifiedTime, o.Name, o.OrgId, o.Purpose, o.SiteId, o.Tiles, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgUiSettings.
// It customizes the JSON marshaling process for OrgUiSettings objects.
func (o OrgUiSettings) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "created_time", "description", "for_site", "id", "isCustomDataboard", "modified_time", "name", "org_id", "purpose", "site_id", "tiles"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgUiSettings object to a map representation for JSON marshaling.
func (o OrgUiSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Description != nil {
        structMap["description"] = o.Description
    }
    if o.ForSite != nil {
        structMap["for_site"] = o.ForSite
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.IsCustomDataboard != nil {
        structMap["isCustomDataboard"] = o.IsCustomDataboard
    }
    if o.ModifiedTime != nil {
        structMap["modified_time"] = o.ModifiedTime
    }
    if o.Name != nil {
        structMap["name"] = o.Name
    }
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.Purpose != nil {
        structMap["purpose"] = o.Purpose
    }
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    if o.Tiles != nil {
        structMap["tiles"] = o.Tiles
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgUiSettings.
// It customizes the JSON unmarshaling process for OrgUiSettings objects.
func (o *OrgUiSettings) UnmarshalJSON(input []byte) error {
    var temp tempOrgUiSettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "description", "for_site", "id", "isCustomDataboard", "modified_time", "name", "org_id", "purpose", "site_id", "tiles")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.CreatedTime = temp.CreatedTime
    o.Description = temp.Description
    o.ForSite = temp.ForSite
    o.Id = temp.Id
    o.IsCustomDataboard = temp.IsCustomDataboard
    o.ModifiedTime = temp.ModifiedTime
    o.Name = temp.Name
    o.OrgId = temp.OrgId
    o.Purpose = temp.Purpose
    o.SiteId = temp.SiteId
    o.Tiles = temp.Tiles
    return nil
}

// tempOrgUiSettings is a temporary struct used for validating the fields of OrgUiSettings.
type tempOrgUiSettings  struct {
    CreatedTime       *float64                  `json:"created_time,omitempty"`
    Description       *string                   `json:"description,omitempty"`
    ForSite           *bool                     `json:"for_site,omitempty"`
    Id                *uuid.UUID                `json:"id,omitempty"`
    IsCustomDataboard *bool                     `json:"isCustomDataboard,omitempty"`
    ModifiedTime      *float64                  `json:"modified_time,omitempty"`
    Name              *string                   `json:"name,omitempty"`
    OrgId             *uuid.UUID                `json:"org_id,omitempty"`
    Purpose           *OrgUiSettingsPurposeEnum `json:"purpose,omitempty"`
    SiteId            *uuid.UUID                `json:"site_id,omitempty"`
    Tiles             []OrgUiSettingsTile       `json:"tiles,omitempty"`
}
