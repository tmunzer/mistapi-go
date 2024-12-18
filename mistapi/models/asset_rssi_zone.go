package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// AssetRssiZone represents a AssetRssiZone struct.
type AssetRssiZone struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Since                *float64               `json:"since,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AssetRssiZone.
// It customizes the JSON marshaling process for AssetRssiZone objects.
func (a AssetRssiZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AssetRssiZone object to a map representation for JSON marshaling.
func (a AssetRssiZone) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for AssetRssiZone.
// It customizes the JSON unmarshaling process for AssetRssiZone objects.
func (a *AssetRssiZone) UnmarshalJSON(input []byte) error {
    var temp tempAssetRssiZone
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

// tempAssetRssiZone is a temporary struct used for validating the fields of AssetRssiZone.
type tempAssetRssiZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *float64   `json:"since,omitempty"`
}
