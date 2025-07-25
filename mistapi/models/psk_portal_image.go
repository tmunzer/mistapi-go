// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// PskPortalImage represents a PskPortalImage struct.
type PskPortalImage struct {
    // Binary file
    File                 *[]byte                `json:"file,omitempty"`
    // JSON string describing the upload
    Json                 *string                `json:"json,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalImage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalImage) String() string {
    return fmt.Sprintf(
    	"PskPortalImage[File=%v, Json=%v, AdditionalProperties=%v]",
    	p.File, p.Json, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalImage.
// It customizes the JSON marshaling process for PskPortalImage objects.
func (p PskPortalImage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "file", "json"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalImage object to a map representation for JSON marshaling.
func (p PskPortalImage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.File != nil {
        structMap["file"] = p.File
    }
    if p.Json != nil {
        structMap["json"] = p.Json
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalImage.
// It customizes the JSON unmarshaling process for PskPortalImage objects.
func (p *PskPortalImage) UnmarshalJSON(input []byte) error {
    var temp tempPskPortalImage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file", "json")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.File = temp.File
    p.Json = temp.Json
    return nil
}

// tempPskPortalImage is a temporary struct used for validating the fields of PskPortalImage.
type tempPskPortalImage  struct {
    File *[]byte `json:"file,omitempty"`
    Json *string `json:"json,omitempty"`
}
