// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMarvisClientsSearch represents a StatsMarvisClientsSearch struct.
// Paginated list of Marvis Client stats records
type StatsMarvisClientsSearch struct {
	// Maximum number of results requested
	Limit *int `json:"limit,omitempty"`
	// List of Marvis Client stats records
	Results []StatsMarvisClient `json:"results,omitempty"`
	// Total number of matching results
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMarvisClientsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMarvisClientsSearch) String() string {
	return fmt.Sprintf(
		"StatsMarvisClientsSearch[Limit=%v, Results=%v, Total=%v, AdditionalProperties=%v]",
		s.Limit, s.Results, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMarvisClientsSearch.
// It customizes the JSON marshaling process for StatsMarvisClientsSearch objects.
func (s StatsMarvisClientsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"limit", "results", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMarvisClientsSearch object to a map representation for JSON marshaling.
func (s StatsMarvisClientsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Limit != nil {
		structMap["limit"] = s.Limit
	}
	if s.Results != nil {
		structMap["results"] = s.Results
	}
	if s.Total != nil {
		structMap["total"] = s.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMarvisClientsSearch.
// It customizes the JSON unmarshaling process for StatsMarvisClientsSearch objects.
func (s *StatsMarvisClientsSearch) UnmarshalJSON(input []byte) error {
	var temp tempStatsMarvisClientsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "limit", "results", "total")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Limit = temp.Limit
	s.Results = temp.Results
	s.Total = temp.Total
	return nil
}

// tempStatsMarvisClientsSearch is a temporary struct used for validating the fields of StatsMarvisClientsSearch.
type tempStatsMarvisClientsSearch struct {
	Limit   *int                `json:"limit,omitempty"`
	Results []StatsMarvisClient `json:"results,omitempty"`
	Total   *int                `json:"total,omitempty"`
}
