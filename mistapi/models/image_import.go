package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ImageImport represents a ImageImport struct.
type ImageImport struct {
    // binary file
    File                 []byte         `json:"file"`
    // JSON string describing your upload
    Json                 *string        `json:"json,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ImageImport.
// It customizes the JSON marshaling process for ImageImport objects.
func (i ImageImport) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the ImageImport object to a map representation for JSON marshaling.
func (i ImageImport) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["file"] = i.File
    if i.Json != nil {
        structMap["json"] = i.Json
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ImageImport.
// It customizes the JSON unmarshaling process for ImageImport objects.
func (i *ImageImport) UnmarshalJSON(input []byte) error {
    var temp tempImageImport
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
    
    i.AdditionalProperties = additionalProperties
    i.File = *temp.File
    i.Json = temp.Json
    return nil
}

// tempImageImport is a temporary struct used for validating the fields of ImageImport.
type tempImageImport  struct {
    File *[]byte `json:"file"`
    Json *string `json:"json,omitempty"`
}

func (i *tempImageImport) validate() error {
    var errs []string
    if i.File == nil {
        errs = append(errs, "required field `file` is missing for type `image_import`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
