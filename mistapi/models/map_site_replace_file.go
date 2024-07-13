package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MapSiteReplaceFile represents a MapSiteReplaceFile struct.
type MapSiteReplaceFile struct {
    File                 []byte                  `json:"file"`
    Json                 *MapSiteReplaceFileJson `json:"json,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapSiteReplaceFile.
// It customizes the JSON marshaling process for MapSiteReplaceFile objects.
func (m MapSiteReplaceFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteReplaceFile object to a map representation for JSON marshaling.
func (m MapSiteReplaceFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["file"] = m.File
    if m.Json != nil {
        structMap["json"] = m.Json.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapSiteReplaceFile.
// It customizes the JSON unmarshaling process for MapSiteReplaceFile objects.
func (m *MapSiteReplaceFile) UnmarshalJSON(input []byte) error {
    var temp mapSiteReplaceFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file", "json")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.File = *temp.File
    m.Json = temp.Json
    return nil
}

// mapSiteReplaceFile is a temporary struct used for validating the fields of MapSiteReplaceFile.
type mapSiteReplaceFile  struct {
    File *[]byte                 `json:"file"`
    Json *MapSiteReplaceFileJson `json:"json,omitempty"`
}

func (m *mapSiteReplaceFile) validate() error {
    var errs []string
    if m.File == nil {
        errs = append(errs, "required field `file` is missing for type `Map_Site_Replace_File`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
