// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// AssetZone represents a AssetZone struct.
type AssetZone struct {
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Since                *float64               `json:"since,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AssetZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AssetZone) String() string {
    return fmt.Sprintf(
    	"AssetZone[Id=%v, Since=%v, AdditionalProperties=%v]",
    	a.Id, a.Since, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AssetZone.
// It customizes the JSON marshaling process for AssetZone objects.
func (a AssetZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AssetZone object to a map representation for JSON marshaling.
func (a AssetZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.Since != nil {
        structMap["since"] = a.Since
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetZone.
// It customizes the JSON unmarshaling process for AssetZone objects.
func (a *AssetZone) UnmarshalJSON(input []byte) error {
    var temp tempAssetZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "since")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Id = temp.Id
    a.Since = temp.Since
    return nil
}

// tempAssetZone is a temporary struct used for validating the fields of AssetZone.
type tempAssetZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *float64   `json:"since,omitempty"`
}
