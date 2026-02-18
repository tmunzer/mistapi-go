// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JsiPbnSearch represents a JsiPbnSearch struct.
// PBN search response
type JsiPbnSearch struct {
	// End timestamp
	End *int `json:"end,omitempty"`
	// Number of results to return
	Limit *int `json:"limit,omitempty"`
	// Next page URL
	Next *string `json:"next,omitempty"`
	// List of PBN advisories
	Results []JsiPbnItem `json:"results,omitempty"`
	// Start timestamp
	Start *int `json:"start,omitempty"`
	// Total number of results
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsiPbnSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsiPbnSearch) String() string {
	return fmt.Sprintf(
		"JsiPbnSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		j.End, j.Limit, j.Next, j.Results, j.Start, j.Total, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsiPbnSearch.
// It customizes the JSON marshaling process for JsiPbnSearch objects.
func (j JsiPbnSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JsiPbnSearch object to a map representation for JSON marshaling.
func (j JsiPbnSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.End != nil {
		structMap["end"] = j.End
	}
	if j.Limit != nil {
		structMap["limit"] = j.Limit
	}
	if j.Next != nil {
		structMap["next"] = j.Next
	}
	if j.Results != nil {
		structMap["results"] = j.Results
	}
	if j.Start != nil {
		structMap["start"] = j.Start
	}
	if j.Total != nil {
		structMap["total"] = j.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsiPbnSearch.
// It customizes the JSON unmarshaling process for JsiPbnSearch objects.
func (j *JsiPbnSearch) UnmarshalJSON(input []byte) error {
	var temp tempJsiPbnSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.End = temp.End
	j.Limit = temp.Limit
	j.Next = temp.Next
	j.Results = temp.Results
	j.Start = temp.Start
	j.Total = temp.Total
	return nil
}

// tempJsiPbnSearch is a temporary struct used for validating the fields of JsiPbnSearch.
type tempJsiPbnSearch struct {
	End     *int         `json:"end,omitempty"`
	Limit   *int         `json:"limit,omitempty"`
	Next    *string      `json:"next,omitempty"`
	Results []JsiPbnItem `json:"results,omitempty"`
	Start   *int         `json:"start,omitempty"`
	Total   *int         `json:"total,omitempty"`
}
