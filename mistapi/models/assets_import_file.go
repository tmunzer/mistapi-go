// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AssetsImportFile represents a AssetsImportFile struct.
type AssetsImportFile struct {
	// CSV file
	File                 *string                `json:"file,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AssetsImportFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AssetsImportFile) String() string {
	return fmt.Sprintf(
		"AssetsImportFile[File=%v, AdditionalProperties=%v]",
		a.File, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AssetsImportFile.
// It customizes the JSON marshaling process for AssetsImportFile objects.
func (a AssetsImportFile) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"file"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AssetsImportFile object to a map representation for JSON marshaling.
func (a AssetsImportFile) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.File = temp.File
	return nil
}

// tempAssetsImportFile is a temporary struct used for validating the fields of AssetsImportFile.
type tempAssetsImportFile struct {
	File *string `json:"file,omitempty"`
}
