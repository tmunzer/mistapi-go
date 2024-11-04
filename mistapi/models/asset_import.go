package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AssetImport represents a AssetImport struct.
type AssetImport struct {
    Mac                  string         `json:"mac"`
    Name                 string         `json:"name"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AssetImport.
// It customizes the JSON marshaling process for AssetImport objects.
func (a AssetImport) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AssetImport object to a map representation for JSON marshaling.
func (a AssetImport) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["mac"] = a.Mac
    structMap["name"] = a.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetImport.
// It customizes the JSON unmarshaling process for AssetImport objects.
func (a *AssetImport) UnmarshalJSON(input []byte) error {
    var temp tempAssetImport
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "name")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Mac = *temp.Mac
    a.Name = *temp.Name
    return nil
}

// tempAssetImport is a temporary struct used for validating the fields of AssetImport.
type tempAssetImport  struct {
    Mac  *string `json:"mac"`
    Name *string `json:"name"`
}

func (a *tempAssetImport) validate() error {
    var errs []string
    if a.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `asset_import`")
    }
    if a.Name == nil {
        errs = append(errs, "required field `name` is missing for type `asset_import`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
