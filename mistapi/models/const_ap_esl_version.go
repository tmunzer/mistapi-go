// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstApEslVersion represents a ConstApEslVersion struct.
type ConstApEslVersion struct {
	EslVersion           *string                `json:"esl_version,omitempty"`
	Model                *string                `json:"model,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstApEslVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstApEslVersion) String() string {
	return fmt.Sprintf(
		"ConstApEslVersion[EslVersion=%v, Model=%v, AdditionalProperties=%v]",
		c.EslVersion, c.Model, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstApEslVersion.
// It customizes the JSON marshaling process for ConstApEslVersion objects.
func (c ConstApEslVersion) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"esl_version", "model"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstApEslVersion object to a map representation for JSON marshaling.
func (c ConstApEslVersion) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.EslVersion != nil {
		structMap["esl_version"] = c.EslVersion
	}
	if c.Model != nil {
		structMap["model"] = c.Model
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstApEslVersion.
// It customizes the JSON unmarshaling process for ConstApEslVersion objects.
func (c *ConstApEslVersion) UnmarshalJSON(input []byte) error {
	var temp tempConstApEslVersion
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "esl_version", "model")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.EslVersion = temp.EslVersion
	c.Model = temp.Model
	return nil
}

// tempConstApEslVersion is a temporary struct used for validating the fields of ConstApEslVersion.
type tempConstApEslVersion struct {
	EslVersion *string `json:"esl_version,omitempty"`
	Model      *string `json:"model,omitempty"`
}
