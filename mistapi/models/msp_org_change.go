// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// MspOrgChange represents a MspOrgChange struct.
type MspOrgChange struct {
	// enum: `assign`, `unassign`
	Op MspOrgChangeOperationEnum `json:"op"`
	// List of org_id
	OrgIds               []string               `json:"org_ids"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MspOrgChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MspOrgChange) String() string {
	return fmt.Sprintf(
		"MspOrgChange[Op=%v, OrgIds=%v, AdditionalProperties=%v]",
		m.Op, m.OrgIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MspOrgChange.
// It customizes the JSON marshaling process for MspOrgChange objects.
func (m MspOrgChange) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"op", "org_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MspOrgChange object to a map representation for JSON marshaling.
func (m MspOrgChange) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	structMap["op"] = m.Op
	structMap["org_ids"] = m.OrgIds
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MspOrgChange.
// It customizes the JSON unmarshaling process for MspOrgChange objects.
func (m *MspOrgChange) UnmarshalJSON(input []byte) error {
	var temp tempMspOrgChange
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "op", "org_ids")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Op = *temp.Op
	m.OrgIds = *temp.OrgIds
	return nil
}

// tempMspOrgChange is a temporary struct used for validating the fields of MspOrgChange.
type tempMspOrgChange struct {
	Op     *MspOrgChangeOperationEnum `json:"op"`
	OrgIds *[]string                  `json:"org_ids"`
}

func (m *tempMspOrgChange) validate() error {
	var errs []string
	if m.Op == nil {
		errs = append(errs, "required field `op` is missing for type `msp_org_change`")
	}
	if m.OrgIds == nil {
		errs = append(errs, "required field `org_ids` is missing for type `msp_org_change`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
