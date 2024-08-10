package models

import (
    "encoding/json"
)

// AssetsImportFile represents a AssetsImportFile struct.
type AssetsImportFile struct {
    // CSV file
    File                 *[]byte        `json:"file,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AssetsImportFile.
// It customizes the JSON marshaling process for AssetsImportFile objects.
func (a AssetsImportFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AssetsImportFile object to a map representation for JSON marshaling.
func (a AssetsImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.File != nil {
        structMap["file"] = a.File
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetsImportFile.
// It customizes the JSON unmarshaling process for AssetsImportFile objects.
func (a *AssetsImportFile) UnmarshalJSON(input []byte) error {
    var temp tempAssetsImportFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.File = temp.File
    return nil
}

// tempAssetsImportFile is a temporary struct used for validating the fields of AssetsImportFile.
type tempAssetsImportFile  struct {
    File *[]byte `json:"file,omitempty"`
}
