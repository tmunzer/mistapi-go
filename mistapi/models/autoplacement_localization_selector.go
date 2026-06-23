// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AutoplacementLocalizationSelector represents a AutoplacementLocalizationSelector struct.
// Request body to apply or clear cached autoplacement or auto-orientation values for a map or subset of APs
type AutoplacementLocalizationSelector struct {
	// The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`
	For *UseAutoApValuesForEnum `json:"for,omitempty"`
	// List of AP MAC addresses to apply the action to. If omitted, the action applies to all APs on the map
	Macs                 []string               `json:"macs,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoplacementLocalizationSelector,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoplacementLocalizationSelector) String() string {
	return fmt.Sprintf(
		"AutoplacementLocalizationSelector[For=%v, Macs=%v, AdditionalProperties=%v]",
		a.For, a.Macs, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoplacementLocalizationSelector.
// It customizes the JSON marshaling process for AutoplacementLocalizationSelector objects.
func (a AutoplacementLocalizationSelector) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"for", "macs"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AutoplacementLocalizationSelector object to a map representation for JSON marshaling.
func (a AutoplacementLocalizationSelector) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.For != nil {
		structMap["for"] = a.For
	}
	if a.Macs != nil {
		structMap["macs"] = a.Macs
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoplacementLocalizationSelector.
// It customizes the JSON unmarshaling process for AutoplacementLocalizationSelector objects.
func (a *AutoplacementLocalizationSelector) UnmarshalJSON(input []byte) error {
	var temp tempAutoplacementLocalizationSelector
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "for", "macs")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.For = temp.For
	a.Macs = temp.Macs
	return nil
}

// tempAutoplacementLocalizationSelector is a temporary struct used for validating the fields of AutoplacementLocalizationSelector.
type tempAutoplacementLocalizationSelector struct {
	For  *UseAutoApValuesForEnum `json:"for,omitempty"`
	Macs []string                `json:"macs,omitempty"`
}
