// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// AutoMapAssignmentRequest represents a AutoMapAssignmentRequest struct.
type AutoMapAssignmentRequest struct {
	// Optional list of specific map IDs to apply/clear. If not provided or empty, all pending map assignments are accepted/rejected.
	MapIds               []uuid.UUID            `json:"map_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoMapAssignmentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoMapAssignmentRequest) String() string {
	return fmt.Sprintf(
		"AutoMapAssignmentRequest[MapIds=%v, AdditionalProperties=%v]",
		a.MapIds, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoMapAssignmentRequest.
// It customizes the JSON marshaling process for AutoMapAssignmentRequest objects.
func (a AutoMapAssignmentRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"map_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AutoMapAssignmentRequest object to a map representation for JSON marshaling.
func (a AutoMapAssignmentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.MapIds != nil {
		structMap["map_ids"] = a.MapIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoMapAssignmentRequest.
// It customizes the JSON unmarshaling process for AutoMapAssignmentRequest objects.
func (a *AutoMapAssignmentRequest) UnmarshalJSON(input []byte) error {
	var temp tempAutoMapAssignmentRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "map_ids")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.MapIds = temp.MapIds
	return nil
}

// tempAutoMapAssignmentRequest is a temporary struct used for validating the fields of AutoMapAssignmentRequest.
type tempAutoMapAssignmentRequest struct {
	MapIds []uuid.UUID `json:"map_ids,omitempty"`
}
