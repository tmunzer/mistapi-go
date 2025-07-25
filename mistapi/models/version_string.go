// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// VersionString represents a VersionString struct.
type VersionString struct {
    Version              *string                `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VersionString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VersionString) String() string {
    return fmt.Sprintf(
    	"VersionString[Version=%v, AdditionalProperties=%v]",
    	v.Version, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VersionString.
// It customizes the JSON marshaling process for VersionString objects.
func (v VersionString) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VersionString object to a map representation for JSON marshaling.
func (v VersionString) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Version != nil {
        structMap["version"] = v.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VersionString.
// It customizes the JSON unmarshaling process for VersionString objects.
func (v *VersionString) UnmarshalJSON(input []byte) error {
    var temp tempVersionString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "version")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Version = temp.Version
    return nil
}

// tempVersionString is a temporary struct used for validating the fields of VersionString.
type tempVersionString  struct {
    Version *string `json:"version,omitempty"`
}
