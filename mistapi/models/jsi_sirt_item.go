// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JsiSirtItem represents a JsiSirtItem struct.
// SIRT advisory item
type JsiSirtItem struct {
	// CVSS score
	CvssScore *float64 `json:"cvss_score,omitempty"`
	// ID of the SIRT
	Id *string `json:"id,omitempty"`
	// OS models affected
	Models []string `json:"models,omitempty"`
	// Problem description
	Problem *string `json:"problem,omitempty"`
	// Release date of the SIRT issue
	PublishedDate *int `json:"published_date,omitempty"`
	// Release notes if any
	ReleaseNotes *string `json:"release_notes,omitempty"`
	// Severity of the issue
	Severity *string `json:"severity,omitempty"`
	// Solution for the security issue
	Solution *string `json:"solution,omitempty"`
	// Title of the SIRT
	Title *string `json:"title,omitempty"`
	// JSA updated timestamp
	UpdatedDate *int `json:"updated_date,omitempty"`
	// OS versions affected
	Versions []string `json:"versions,omitempty"`
	// Workaround provided
	Workaround           *string                `json:"workaround,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsiSirtItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsiSirtItem) String() string {
	return fmt.Sprintf(
		"JsiSirtItem[CvssScore=%v, Id=%v, Models=%v, Problem=%v, PublishedDate=%v, ReleaseNotes=%v, Severity=%v, Solution=%v, Title=%v, UpdatedDate=%v, Versions=%v, Workaround=%v, AdditionalProperties=%v]",
		j.CvssScore, j.Id, j.Models, j.Problem, j.PublishedDate, j.ReleaseNotes, j.Severity, j.Solution, j.Title, j.UpdatedDate, j.Versions, j.Workaround, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsiSirtItem.
// It customizes the JSON marshaling process for JsiSirtItem objects.
func (j JsiSirtItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"cvss_score", "id", "models", "problem", "published_date", "release_notes", "severity", "solution", "title", "updated_date", "versions", "workaround"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JsiSirtItem object to a map representation for JSON marshaling.
func (j JsiSirtItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.CvssScore != nil {
		structMap["cvss_score"] = j.CvssScore
	}
	if j.Id != nil {
		structMap["id"] = j.Id
	}
	if j.Models != nil {
		structMap["models"] = j.Models
	}
	if j.Problem != nil {
		structMap["problem"] = j.Problem
	}
	if j.PublishedDate != nil {
		structMap["published_date"] = j.PublishedDate
	}
	if j.ReleaseNotes != nil {
		structMap["release_notes"] = j.ReleaseNotes
	}
	if j.Severity != nil {
		structMap["severity"] = j.Severity
	}
	if j.Solution != nil {
		structMap["solution"] = j.Solution
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsiSirtItem.
// It customizes the JSON unmarshaling process for JsiSirtItem objects.
func (j *JsiSirtItem) UnmarshalJSON(input []byte) error {
	var temp tempJsiSirtItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cvss_score", "id", "models", "problem", "published_date", "release_notes", "severity", "solution", "title", "updated_date", "versions", "workaround")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.CvssScore = temp.CvssScore
	j.Id = temp.Id
	j.Models = temp.Models
	j.Problem = temp.Problem
	j.PublishedDate = temp.PublishedDate
	j.ReleaseNotes = temp.ReleaseNotes
	j.Severity = temp.Severity
	j.Solution = temp.Solution
	j.Title = temp.Title
	j.UpdatedDate = temp.UpdatedDate
	j.Versions = temp.Versions
	j.Workaround = temp.Workaround
	return nil
}

// tempJsiSirtItem is a temporary struct used for validating the fields of JsiSirtItem.
type tempJsiSirtItem struct {
	CvssScore     *float64 `json:"cvss_score,omitempty"`
	Id            *string  `json:"id,omitempty"`
	Models        []string `json:"models,omitempty"`
	Problem       *string  `json:"problem,omitempty"`
	PublishedDate *int     `json:"published_date,omitempty"`
	ReleaseNotes  *string  `json:"release_notes,omitempty"`
	Severity      *string  `json:"severity,omitempty"`
	Solution      *string  `json:"solution,omitempty"`
	Title         *string  `json:"title,omitempty"`
	UpdatedDate   *int     `json:"updated_date,omitempty"`
	Versions      []string `json:"versions,omitempty"`
	Workaround    *string  `json:"workaround,omitempty"`
}
