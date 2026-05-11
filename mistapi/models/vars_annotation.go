// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// VarsAnnotation represents a VarsAnnotation struct.
// Annotation for a single var, helping identify its purpose and enabling auto-complete/enumeration in UI
type VarsAnnotation struct {
	// User-provided note to describe what this var was created for
	Note *string `json:"note,omitempty"`
	// Used to identify where to enumerate / auto-complete the field from. Default is `generic` (plain string, no special handling).
	// enum: `generic`, `mxtunnel_id`
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VarsAnnotation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VarsAnnotation) String() string {
	return fmt.Sprintf(
		"VarsAnnotation[Note=%v, Type=%v, AdditionalProperties=%v]",
		v.Note, v.Type, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VarsAnnotation.
// It customizes the JSON marshaling process for VarsAnnotation objects.
func (v VarsAnnotation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"note", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VarsAnnotation object to a map representation for JSON marshaling.
func (v VarsAnnotation) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	if v.Note != nil {
		structMap["note"] = v.Note
	}
	if v.Type != nil {
		structMap["type"] = v.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VarsAnnotation.
// It customizes the JSON unmarshaling process for VarsAnnotation objects.
func (v *VarsAnnotation) UnmarshalJSON(input []byte) error {
	var temp tempVarsAnnotation
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "note", "type")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.Note = temp.Note
	v.Type = temp.Type
	return nil
}

// tempVarsAnnotation is a temporary struct used for validating the fields of VarsAnnotation.
type tempVarsAnnotation struct {
	Note *string `json:"note,omitempty"`
	Type *string `json:"type,omitempty"`
}
