// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisConfigActionsSearch represents a MarvisConfigActionsSearch struct.
// Paginated list of Marvis config actions
type MarvisConfigActionsSearch struct {
	// Search window end timestamp, in epoch seconds
	End *int `json:"end,omitempty"`
	// Maximum number of results requested
	Limit *int `json:"limit,omitempty"`
	// List of Marvis config actions
	Results []MarvisConfigAction `json:"results,omitempty"`
	// Search window start timestamp, in epoch seconds
	Start *int `json:"start,omitempty"`
	// Total number of matching results
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisConfigActionsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisConfigActionsSearch) String() string {
	return fmt.Sprintf(
		"MarvisConfigActionsSearch[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		m.End, m.Limit, m.Results, m.Start, m.Total, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisConfigActionsSearch.
// It customizes the JSON marshaling process for MarvisConfigActionsSearch objects.
func (m MarvisConfigActionsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"end", "limit", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisConfigActionsSearch object to a map representation for JSON marshaling.
func (m MarvisConfigActionsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.End != nil {
		structMap["end"] = m.End
	}
	if m.Limit != nil {
		structMap["limit"] = m.Limit
	}
	if m.Results != nil {
		structMap["results"] = m.Results
	}
	if m.Start != nil {
		structMap["start"] = m.Start
	}
	if m.Total != nil {
		structMap["total"] = m.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisConfigActionsSearch.
// It customizes the JSON unmarshaling process for MarvisConfigActionsSearch objects.
func (m *MarvisConfigActionsSearch) UnmarshalJSON(input []byte) error {
	var temp tempMarvisConfigActionsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.End = temp.End
	m.Limit = temp.Limit
	m.Results = temp.Results
	m.Start = temp.Start
	m.Total = temp.Total
	return nil
}

// tempMarvisConfigActionsSearch is a temporary struct used for validating the fields of MarvisConfigActionsSearch.
type tempMarvisConfigActionsSearch struct {
	End     *int                 `json:"end,omitempty"`
	Limit   *int                 `json:"limit,omitempty"`
	Results []MarvisConfigAction `json:"results,omitempty"`
	Start   *int                 `json:"start,omitempty"`
	Total   *int                 `json:"total,omitempty"`
}
