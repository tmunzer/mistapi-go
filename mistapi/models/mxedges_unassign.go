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

// MxedgesUnassign represents a MxedgesUnassign struct.
type MxedgesUnassign struct {
	MxedgeIds            []uuid.UUID            `json:"mxedge_ids"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgesUnassign,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgesUnassign) String() string {
	return fmt.Sprintf(
		"MxedgesUnassign[MxedgeIds=%v, AdditionalProperties=%v]",
		m.MxedgeIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgesUnassign.
// It customizes the JSON marshaling process for MxedgesUnassign objects.
func (m MxedgesUnassign) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"mxedge_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgesUnassign object to a map representation for JSON marshaling.
func (m MxedgesUnassign) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	structMap["mxedge_ids"] = m.MxedgeIds
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgesUnassign.
// It customizes the JSON unmarshaling process for MxedgesUnassign objects.
func (m *MxedgesUnassign) UnmarshalJSON(input []byte) error {
	var temp tempMxedgesUnassign
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxedge_ids")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.MxedgeIds = *temp.MxedgeIds
	return nil
}

// tempMxedgesUnassign is a temporary struct used for validating the fields of MxedgesUnassign.
type tempMxedgesUnassign struct {
	MxedgeIds *[]uuid.UUID `json:"mxedge_ids"`
}

func (m *tempMxedgesUnassign) validate() error {
	var errs []string
	if m.MxedgeIds == nil {
		errs = append(errs, "required field `mxedge_ids` is missing for type `mxedges_unassign`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
