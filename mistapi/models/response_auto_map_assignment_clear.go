// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// ResponseAutoMapAssignmentClear represents a ResponseAutoMapAssignmentClear struct.
// Result returned after clearing auto map assignment candidates
type ResponseAutoMapAssignmentClear struct {
	// Human-readable description of the operation result
	Message string `json:"message"`
	// List of map IDs that were successfully rejected
	RejectedMaps         []uuid.UUID            `json:"rejected_maps"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoMapAssignmentClear,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoMapAssignmentClear) String() string {
	return fmt.Sprintf(
		"ResponseAutoMapAssignmentClear[Message=%v, RejectedMaps=%v, AdditionalProperties=%v]",
		r.Message, r.RejectedMaps, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoMapAssignmentClear.
// It customizes the JSON marshaling process for ResponseAutoMapAssignmentClear objects.
func (r ResponseAutoMapAssignmentClear) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"message", "rejected_maps"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoMapAssignmentClear object to a map representation for JSON marshaling.
func (r ResponseAutoMapAssignmentClear) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["message"] = r.Message
	structMap["rejected_maps"] = r.RejectedMaps
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoMapAssignmentClear.
// It customizes the JSON unmarshaling process for ResponseAutoMapAssignmentClear objects.
func (r *ResponseAutoMapAssignmentClear) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoMapAssignmentClear
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "message", "rejected_maps")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Message = *temp.Message
	r.RejectedMaps = *temp.RejectedMaps
	return nil
}

// tempResponseAutoMapAssignmentClear is a temporary struct used for validating the fields of ResponseAutoMapAssignmentClear.
type tempResponseAutoMapAssignmentClear struct {
	Message      *string      `json:"message"`
	RejectedMaps *[]uuid.UUID `json:"rejected_maps"`
}

func (r *tempResponseAutoMapAssignmentClear) validate() error {
	var errs []string
	if r.Message == nil {
		errs = append(errs, "required field `message` is missing for type `response_auto_map_assignment_clear`")
	}
	if r.RejectedMaps == nil {
		errs = append(errs, "required field `rejected_maps` is missing for type `response_auto_map_assignment_clear`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
