// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ConstMarvisClientVersion represents a ConstMarvisClientVersion struct.
type ConstMarvisClientVersion struct {
    Label                *string                `json:"label,omitempty"`
    Notes                *string                `json:"notes,omitempty"`
    // Client OS
    Os                   *string                `json:"os,omitempty"`
    // Client download url
    Url                  *string                `json:"url,omitempty"`
    // Client version
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstMarvisClientVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstMarvisClientVersion) String() string {
    return fmt.Sprintf(
    	"ConstMarvisClientVersion[Label=%v, Notes=%v, Os=%v, Url=%v, Version=%v, AdditionalProperties=%v]",
    	c.Label, c.Notes, c.Os, c.Url, c.Version, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstMarvisClientVersion.
// It customizes the JSON marshaling process for ConstMarvisClientVersion objects.
func (c ConstMarvisClientVersion) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "label", "notes", "os", "url", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstMarvisClientVersion object to a map representation for JSON marshaling.
func (c ConstMarvisClientVersion) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Label != nil {
        structMap["label"] = c.Label
    }
    if c.Notes != nil {
        structMap["notes"] = c.Notes
    }
    if c.Os != nil {
        structMap["os"] = c.Os
    }
    if c.Url != nil {
        structMap["url"] = c.Url
    }
    if c.Version != nil {
        structMap["version"] = c.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstMarvisClientVersion.
// It customizes the JSON unmarshaling process for ConstMarvisClientVersion objects.
func (c *ConstMarvisClientVersion) UnmarshalJSON(input []byte) error {
    var temp tempConstMarvisClientVersion
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "label", "notes", "os", "url", "version")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Label = temp.Label
    c.Notes = temp.Notes
    c.Os = temp.Os
    c.Url = temp.Url
    c.Version = temp.Version
    return nil
}

// tempConstMarvisClientVersion is a temporary struct used for validating the fields of ConstMarvisClientVersion.
type tempConstMarvisClientVersion  struct {
    Label   *string `json:"label,omitempty"`
    Notes   *string `json:"notes,omitempty"`
    Os      *string `json:"os,omitempty"`
    Url     *string `json:"url,omitempty"`
    Version *string `json:"version,omitempty"`
}
