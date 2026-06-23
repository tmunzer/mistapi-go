// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientEventsSearch represents a MarvisClientEventsSearch struct.
// Paginated list of Marvis Client events
type MarvisClientEventsSearch struct {
	// Maximum number of results requested
	Limit *int `json:"limit,omitempty"`
	// List of Marvis Client events
	Results []MarvisClientEvent `json:"results,omitempty"`
	// Total number of matching results
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientEventsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientEventsSearch) String() string {
	return fmt.Sprintf(
		"MarvisClientEventsSearch[Limit=%v, Results=%v, Total=%v, AdditionalProperties=%v]",
		m.Limit, m.Results, m.Total, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientEventsSearch.
// It customizes the JSON marshaling process for MarvisClientEventsSearch objects.
func (m MarvisClientEventsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"limit", "results", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientEventsSearch object to a map representation for JSON marshaling.
func (m MarvisClientEventsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Limit != nil {
		structMap["limit"] = m.Limit
	}
	if m.Results != nil {
		structMap["results"] = m.Results
	}
	if m.Total != nil {
		structMap["total"] = m.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientEventsSearch.
// It customizes the JSON unmarshaling process for MarvisClientEventsSearch objects.
func (m *MarvisClientEventsSearch) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientEventsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "limit", "results", "total")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Limit = temp.Limit
	m.Results = temp.Results
	m.Total = temp.Total
	return nil
}

// tempMarvisClientEventsSearch is a temporary struct used for validating the fields of MarvisClientEventsSearch.
type tempMarvisClientEventsSearch struct {
	Limit   *int                `json:"limit,omitempty"`
	Results []MarvisClientEvent `json:"results,omitempty"`
	Total   *int                `json:"total,omitempty"`
}
