package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// BinaryStream represents a BinaryStream struct.
type BinaryStream struct {
    // file to updload
    File                 []byte         `json:"file"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BinaryStream.
// It customizes the JSON marshaling process for BinaryStream objects.
func (b BinaryStream) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BinaryStream object to a map representation for JSON marshaling.
func (b BinaryStream) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["file"] = b.File
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BinaryStream.
// It customizes the JSON unmarshaling process for BinaryStream objects.
func (b *BinaryStream) UnmarshalJSON(input []byte) error {
    var temp tempBinaryStream
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.File = *temp.File
    return nil
}

// tempBinaryStream is a temporary struct used for validating the fields of BinaryStream.
type tempBinaryStream  struct {
    File *[]byte `json:"file"`
}

func (b *tempBinaryStream) validate() error {
    var errs []string
    if b.File == nil {
        errs = append(errs, "required field `file` is missing for type `binary_stream`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
