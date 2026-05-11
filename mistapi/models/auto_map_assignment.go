// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AutoMapAssignment represents a AutoMapAssignment struct.
type AutoMapAssignment struct {
	// If `true`, validates the site's APs without starting the map assignment process. Returns device validity and estimated runtime.
	Dryrun *bool `json:"dryrun,omitempty"`
	// If `true`, forces data collection via orchestration. If `false`, attempts to use existing BLE data first.
	ForceCollection      *bool                  `json:"force_collection,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoMapAssignment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoMapAssignment) String() string {
	return fmt.Sprintf(
		"AutoMapAssignment[Dryrun=%v, ForceCollection=%v, AdditionalProperties=%v]",
		a.Dryrun, a.ForceCollection, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoMapAssignment.
// It customizes the JSON marshaling process for AutoMapAssignment objects.
func (a AutoMapAssignment) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"dryrun", "force_collection"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AutoMapAssignment object to a map representation for JSON marshaling.
func (a AutoMapAssignment) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Dryrun != nil {
		structMap["dryrun"] = a.Dryrun
	}
	if a.ForceCollection != nil {
		structMap["force_collection"] = a.ForceCollection
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoMapAssignment.
// It customizes the JSON unmarshaling process for AutoMapAssignment objects.
func (a *AutoMapAssignment) UnmarshalJSON(input []byte) error {
	var temp tempAutoMapAssignment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dryrun", "force_collection")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Dryrun = temp.Dryrun
	a.ForceCollection = temp.ForceCollection
	return nil
}

// tempAutoMapAssignment is a temporary struct used for validating the fields of AutoMapAssignment.
type tempAutoMapAssignment struct {
	Dryrun          *bool `json:"dryrun,omitempty"`
	ForceCollection *bool `json:"force_collection,omitempty"`
}
