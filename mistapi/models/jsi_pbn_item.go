// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JsiPbnItem represents a JsiPbnItem struct.
// PBN (Problem Bug Notification) advisory item
type JsiPbnItem struct {
	// Type of the bug (Day-1, Regression)
	BugType *string `json:"bug_type,omitempty"`
	// Risk level
	CustomerRisk *string `json:"customer_risk,omitempty"`
	// Release in which the issue was fixed
	FixedIn *string `json:"fixed_in,omitempty"`
	// ID of the PBN
	Id *string `json:"id,omitempty"`
	// Release introduced in
	IntroducedIn *string `json:"introduced_in,omitempty"`
	// OS models affected
	Models []string `json:"models,omitempty"`
	// Product family affected
	ProductFamily *string `json:"product_family,omitempty"`
	// Release notes for this PBN
	ReleaseNotes *string `json:"release_notes,omitempty"`
	// Restoration steps
	Restoration *string `json:"restoration,omitempty"`
	// Title of the issue
	Title *string `json:"title,omitempty"`
	// PBN updated timestamp
	UpdatedDate *int `json:"updated_date,omitempty"`
	// OS versions affected
	Versions []string `json:"versions,omitempty"`
	// Workaround for this issue
	Workaround *string `json:"workaround,omitempty"`
	// Any workaround available
	WorkaroundProvided   *string                `json:"workaround_provided,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsiPbnItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsiPbnItem) String() string {
	return fmt.Sprintf(
		"JsiPbnItem[BugType=%v, CustomerRisk=%v, FixedIn=%v, Id=%v, IntroducedIn=%v, Models=%v, ProductFamily=%v, ReleaseNotes=%v, Restoration=%v, Title=%v, UpdatedDate=%v, Versions=%v, Workaround=%v, WorkaroundProvided=%v, AdditionalProperties=%v]",
		j.BugType, j.CustomerRisk, j.FixedIn, j.Id, j.IntroducedIn, j.Models, j.ProductFamily, j.ReleaseNotes, j.Restoration, j.Title, j.UpdatedDate, j.Versions, j.Workaround, j.WorkaroundProvided, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsiPbnItem.
// It customizes the JSON marshaling process for JsiPbnItem objects.
func (j JsiPbnItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"bug_type", "customer_risk", "fixed_in", "id", "introduced_in", "models", "product_family", "release_notes", "restoration", "title", "updated_date", "versions", "workaround", "workaround_provided"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JsiPbnItem object to a map representation for JSON marshaling.
func (j JsiPbnItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.BugType != nil {
		structMap["bug_type"] = j.BugType
	}
	if j.CustomerRisk != nil {
		structMap["customer_risk"] = j.CustomerRisk
	}
	if j.FixedIn != nil {
		structMap["fixed_in"] = j.FixedIn
	}
	if j.Id != nil {
		structMap["id"] = j.Id
	}
	if j.IntroducedIn != nil {
		structMap["introduced_in"] = j.IntroducedIn
	}
	if j.Models != nil {
		structMap["models"] = j.Models
	}
	if j.ProductFamily != nil {
		structMap["product_family"] = j.ProductFamily
	}
	if j.ReleaseNotes != nil {
		structMap["release_notes"] = j.ReleaseNotes
	}
	if j.Restoration != nil {
		structMap["restoration"] = j.Restoration
	}
	if j.Title != nil {
		structMap["title"] = j.Title
	}
	if j.UpdatedDate != nil {
		structMap["updated_date"] = j.UpdatedDate
	}
	if j.Versions != nil {
		structMap["versions"] = j.Versions
	}
	if j.Workaround != nil {
		structMap["workaround"] = j.Workaround
	}
	if j.WorkaroundProvided != nil {
		structMap["workaround_provided"] = j.WorkaroundProvided
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsiPbnItem.
// It customizes the JSON unmarshaling process for JsiPbnItem objects.
func (j *JsiPbnItem) UnmarshalJSON(input []byte) error {
	var temp tempJsiPbnItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bug_type", "customer_risk", "fixed_in", "id", "introduced_in", "models", "product_family", "release_notes", "restoration", "title", "updated_date", "versions", "workaround", "workaround_provided")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.BugType = temp.BugType
	j.CustomerRisk = temp.CustomerRisk
	j.FixedIn = temp.FixedIn
	j.Id = temp.Id
	j.IntroducedIn = temp.IntroducedIn
	j.Models = temp.Models
	j.ProductFamily = temp.ProductFamily
	j.ReleaseNotes = temp.ReleaseNotes
	j.Restoration = temp.Restoration
	j.Title = temp.Title
	j.UpdatedDate = temp.UpdatedDate
	j.Versions = temp.Versions
	j.Workaround = temp.Workaround
	j.WorkaroundProvided = temp.WorkaroundProvided
	return nil
}

// tempJsiPbnItem is a temporary struct used for validating the fields of JsiPbnItem.
type tempJsiPbnItem struct {
	BugType            *string  `json:"bug_type,omitempty"`
	CustomerRisk       *string  `json:"customer_risk,omitempty"`
	FixedIn            *string  `json:"fixed_in,omitempty"`
	Id                 *string  `json:"id,omitempty"`
	IntroducedIn       *string  `json:"introduced_in,omitempty"`
	Models             []string `json:"models,omitempty"`
	ProductFamily      *string  `json:"product_family,omitempty"`
	ReleaseNotes       *string  `json:"release_notes,omitempty"`
	Restoration        *string  `json:"restoration,omitempty"`
	Title              *string  `json:"title,omitempty"`
	UpdatedDate        *int     `json:"updated_date,omitempty"`
	Versions           []string `json:"versions,omitempty"`
	Workaround         *string  `json:"workaround,omitempty"`
	WorkaroundProvided *string  `json:"workaround_provided,omitempty"`
}
