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

// ResponseAutoMapAssignmentApply represents a ResponseAutoMapAssignmentApply struct.
type ResponseAutoMapAssignmentApply struct {
	// List of map IDs that were successfully accepted
	AcceptedMaps []uuid.UUID `json:"accepted_maps"`
	// Human-readable description of the operation result
	Message              string                 `json:"message"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoMapAssignmentApply,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoMapAssignmentApply) String() string {
	return fmt.Sprintf(
		"ResponseAutoMapAssignmentApply[AcceptedMaps=%v, Message=%v, AdditionalProperties=%v]",
		r.AcceptedMaps, r.Message, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoMapAssignmentApply.
// It customizes the JSON marshaling process for ResponseAutoMapAssignmentApply objects.
func (r ResponseAutoMapAssignmentApply) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"accepted_maps", "message"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoMapAssignmentApply object to a map representation for JSON marshaling.
func (r ResponseAutoMapAssignmentApply) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["accepted_maps"] = r.AcceptedMaps
	structMap["message"] = r.Message
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoMapAssignmentApply.
// It customizes the JSON unmarshaling process for ResponseAutoMapAssignmentApply objects.
func (r *ResponseAutoMapAssignmentApply) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoMapAssignmentApply
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accepted_maps", "message")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.AcceptedMaps = *temp.AcceptedMaps
	r.Message = *temp.Message
	return nil
}

// tempResponseAutoMapAssignmentApply is a temporary struct used for validating the fields of ResponseAutoMapAssignmentApply.
type tempResponseAutoMapAssignmentApply struct {
	AcceptedMaps *[]uuid.UUID `json:"accepted_maps"`
	Message      *string      `json:"message"`
}

func (r *tempResponseAutoMapAssignmentApply) validate() error {
	var errs []string
	if r.AcceptedMaps == nil {
		errs = append(errs, "required field `accepted_maps` is missing for type `response_auto_map_assignment_apply`")
	}
	if r.Message == nil {
		errs = append(errs, "required field `message` is missing for type `response_auto_map_assignment_apply`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
