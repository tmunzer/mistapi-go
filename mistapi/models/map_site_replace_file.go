// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MapSiteReplaceFile represents a MapSiteReplaceFile struct.
type MapSiteReplaceFile struct {
    File                 []byte                  `json:"file"`
    Json                 *MapSiteReplaceFileJson `json:"json,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for MapSiteReplaceFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapSiteReplaceFile) String() string {
    return fmt.Sprintf(
    	"MapSiteReplaceFile[File=%v, Json=%v, AdditionalProperties=%v]",
    	m.File, m.Json, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapSiteReplaceFile.
// It customizes the JSON marshaling process for MapSiteReplaceFile objects.
func (m MapSiteReplaceFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "file", "json"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteReplaceFile object to a map representation for JSON marshaling.
func (m MapSiteReplaceFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["file"] = m.File
    if m.Json != nil {
        structMap["json"] = m.Json.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapSiteReplaceFile.
// It customizes the JSON unmarshaling process for MapSiteReplaceFile objects.
func (m *MapSiteReplaceFile) UnmarshalJSON(input []byte) error {
    var temp tempMapSiteReplaceFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file", "json")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.File = *temp.File
    m.Json = temp.Json
    return nil
}

// tempMapSiteReplaceFile is a temporary struct used for validating the fields of MapSiteReplaceFile.
type tempMapSiteReplaceFile  struct {
    File *[]byte                 `json:"file"`
    Json *MapSiteReplaceFileJson `json:"json,omitempty"`
}

func (m *tempMapSiteReplaceFile) validate() error {
    var errs []string
    if m.File == nil {
        errs = append(errs, "required field `file` is missing for type `map_site_replace_file`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
