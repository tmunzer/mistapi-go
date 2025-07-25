// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// AamwProfile represents a AamwProfile struct.
type AamwProfile struct {
    Categories           []AamwProfileCategory  `json:"categories,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // enum: `block`, `permit`
    FallbackAction       *AamwProfileActionEnum `json:"fallback_action,omitempty"`
    // enum: `block`, `permit`
    FileAction           *AamwProfileActionEnum `json:"file_action,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    VerdictThreshold     *int                   `json:"verdict_threshold,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AamwProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AamwProfile) String() string {
    return fmt.Sprintf(
    	"AamwProfile[Categories=%v, CreatedTime=%v, FallbackAction=%v, FileAction=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SiteId=%v, VerdictThreshold=%v, AdditionalProperties=%v]",
    	a.Categories, a.CreatedTime, a.FallbackAction, a.FileAction, a.Id, a.ModifiedTime, a.Name, a.OrgId, a.SiteId, a.VerdictThreshold, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AamwProfile.
// It customizes the JSON marshaling process for AamwProfile objects.
func (a AamwProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "categories", "created_time", "fallback_action", "file_action", "id", "modified_time", "name", "org_id", "site_id", "verdict_threshold"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AamwProfile object to a map representation for JSON marshaling.
func (a AamwProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Categories != nil {
        structMap["categories"] = a.Categories
    }
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    if a.FallbackAction != nil {
        structMap["fallback_action"] = a.FallbackAction
    }
    if a.FileAction != nil {
        structMap["file_action"] = a.FileAction
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.VerdictThreshold != nil {
        structMap["verdict_threshold"] = a.VerdictThreshold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AamwProfile.
// It customizes the JSON unmarshaling process for AamwProfile objects.
func (a *AamwProfile) UnmarshalJSON(input []byte) error {
    var temp tempAamwProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "categories", "created_time", "fallback_action", "file_action", "id", "modified_time", "name", "org_id", "site_id", "verdict_threshold")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Categories = temp.Categories
    a.CreatedTime = temp.CreatedTime
    a.FallbackAction = temp.FallbackAction
    a.FileAction = temp.FileAction
    a.Id = temp.Id
    a.ModifiedTime = temp.ModifiedTime
    a.Name = temp.Name
    a.OrgId = temp.OrgId
    a.SiteId = temp.SiteId
    a.VerdictThreshold = temp.VerdictThreshold
    return nil
}

// tempAamwProfile is a temporary struct used for validating the fields of AamwProfile.
type tempAamwProfile  struct {
    Categories       []AamwProfileCategory  `json:"categories,omitempty"`
    CreatedTime      *float64               `json:"created_time,omitempty"`
    FallbackAction   *AamwProfileActionEnum `json:"fallback_action,omitempty"`
    FileAction       *AamwProfileActionEnum `json:"file_action,omitempty"`
    Id               *uuid.UUID             `json:"id,omitempty"`
    ModifiedTime     *float64               `json:"modified_time,omitempty"`
    Name             *string                `json:"name,omitempty"`
    OrgId            *uuid.UUID             `json:"org_id,omitempty"`
    SiteId           *uuid.UUID             `json:"site_id,omitempty"`
    VerdictThreshold *int                   `json:"verdict_threshold,omitempty"`
}
