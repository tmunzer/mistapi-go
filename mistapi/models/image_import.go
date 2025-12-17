// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ImageImport represents a ImageImport struct.
type ImageImport struct {
	// Binary file
	File                 []byte                 `json:"file"`
	Json                 *string                `json:"json,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ImageImport,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i ImageImport) String() string {
	return fmt.Sprintf(
		"ImageImport[File=%v, Json=%v, AdditionalProperties=%v]",
		i.File, i.Json, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ImageImport.
// It customizes the JSON marshaling process for ImageImport objects.
func (i ImageImport) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"file", "json"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the ImageImport object to a map representation for JSON marshaling.
func (i ImageImport) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file", "json")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.File = *temp.File
	i.Json = temp.Json
	return nil
}

// tempImageImport is a temporary struct used for validating the fields of ImageImport.
type tempImageImport struct {
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
	return errors.New(strings.Join(errs, "\n"))
}
