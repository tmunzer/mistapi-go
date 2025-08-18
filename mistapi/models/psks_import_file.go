// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// PsksImportFile represents a PsksImportFile struct.
type PsksImportFile struct {
	File                 *[]byte                `json:"file,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PsksImportFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PsksImportFile) String() string {
	return fmt.Sprintf(
		"PsksImportFile[File=%v, AdditionalProperties=%v]",
		p.File, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PsksImportFile.
// It customizes the JSON marshaling process for PsksImportFile objects.
func (p PsksImportFile) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"file"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PsksImportFile object to a map representation for JSON marshaling.
func (p PsksImportFile) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.File != nil {
		structMap["file"] = p.File
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PsksImportFile.
// It customizes the JSON unmarshaling process for PsksImportFile objects.
func (p *PsksImportFile) UnmarshalJSON(input []byte) error {
	var temp tempPsksImportFile
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.File = temp.File
	return nil
}

// tempPsksImportFile is a temporary struct used for validating the fields of PsksImportFile.
type tempPsksImportFile struct {
	File *[]byte `json:"file,omitempty"`
}
