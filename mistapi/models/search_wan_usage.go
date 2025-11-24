// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SearchWanUsage represents a SearchWanUsage struct.
type SearchWanUsage struct {
	End                  *float64               `json:"end,omitempty"`
	Limit                *int                   `json:"limit,omitempty"`
	Results              []WanUsages            `json:"results,omitempty"`
	Start                *float64               `json:"start,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWanUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWanUsage) String() string {
	return fmt.Sprintf(
		"SearchWanUsage[End=%v, Limit=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
		s.End, s.Limit, s.Results, s.Start, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWanUsage.
// It customizes the JSON marshaling process for SearchWanUsage objects.
func (s SearchWanUsage) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"end", "limit", "results", "start"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SearchWanUsage object to a map representation for JSON marshaling.
func (s SearchWanUsage) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.End != nil {
		structMap["end"] = s.End
	}
	if s.Limit != nil {
		structMap["limit"] = s.Limit
	}
	if s.Results != nil {
		structMap["results"] = s.Results
	}
	if s.Start != nil {
		structMap["start"] = s.Start
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWanUsage.
// It customizes the JSON unmarshaling process for SearchWanUsage objects.
func (s *SearchWanUsage) UnmarshalJSON(input []byte) error {
	var temp tempSearchWanUsage
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.End = temp.End
	s.Limit = temp.Limit
	s.Results = temp.Results
	s.Start = temp.Start
	return nil
}

// tempSearchWanUsage is a temporary struct used for validating the fields of SearchWanUsage.
type tempSearchWanUsage struct {
	End     *float64    `json:"end,omitempty"`
	Limit   *int        `json:"limit,omitempty"`
	Results []WanUsages `json:"results,omitempty"`
	Start   *float64    `json:"start,omitempty"`
}
