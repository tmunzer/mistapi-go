// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OrgUiSettingsTile represents a OrgUiSettingsTile struct.
type OrgUiSettingsTile struct {
    // Description of the tile
    Description          *string                    `json:"description,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                 `json:"id,omitempty"`
    // Whether the tile title is auto generated or not
    IsAutoTitle          *bool                      `json:"isAutoTitle,omitempty"`
    // Name of the tile
    Name                 *string                    `json:"name,omitempty"`
    // Natural Language query for the tile
    NlQuery              *string                    `json:"nl_query,omitempty"`
    Position             *OrgUiSettingsTilePosition `json:"position,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for OrgUiSettingsTile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgUiSettingsTile) String() string {
    return fmt.Sprintf(
    	"OrgUiSettingsTile[Description=%v, Id=%v, IsAutoTitle=%v, Name=%v, NlQuery=%v, Position=%v, AdditionalProperties=%v]",
    	o.Description, o.Id, o.IsAutoTitle, o.Name, o.NlQuery, o.Position, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgUiSettingsTile.
// It customizes the JSON marshaling process for OrgUiSettingsTile objects.
func (o OrgUiSettingsTile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "description", "id", "isAutoTitle", "name", "nl_query", "position"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgUiSettingsTile object to a map representation for JSON marshaling.
func (o OrgUiSettingsTile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Description != nil {
        structMap["description"] = o.Description
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.IsAutoTitle != nil {
        structMap["isAutoTitle"] = o.IsAutoTitle
    }
    if o.Name != nil {
        structMap["name"] = o.Name
    }
    if o.NlQuery != nil {
        structMap["nl_query"] = o.NlQuery
    }
    if o.Position != nil {
        structMap["position"] = o.Position.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgUiSettingsTile.
// It customizes the JSON unmarshaling process for OrgUiSettingsTile objects.
func (o *OrgUiSettingsTile) UnmarshalJSON(input []byte) error {
    var temp tempOrgUiSettingsTile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "description", "id", "isAutoTitle", "name", "nl_query", "position")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Description = temp.Description
    o.Id = temp.Id
    o.IsAutoTitle = temp.IsAutoTitle
    o.Name = temp.Name
    o.NlQuery = temp.NlQuery
    o.Position = temp.Position
    return nil
}

// tempOrgUiSettingsTile is a temporary struct used for validating the fields of OrgUiSettingsTile.
type tempOrgUiSettingsTile  struct {
    Description *string                    `json:"description,omitempty"`
    Id          *uuid.UUID                 `json:"id,omitempty"`
    IsAutoTitle *bool                      `json:"isAutoTitle,omitempty"`
    Name        *string                    `json:"name,omitempty"`
    NlQuery     *string                    `json:"nl_query,omitempty"`
    Position    *OrgUiSettingsTilePosition `json:"position,omitempty"`
}
