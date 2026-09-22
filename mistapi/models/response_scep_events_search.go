// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseScepEventsSearch represents a ResponseScepEventsSearch struct.
// Paginated Mist SCEP PKI operation event search response
type ResponseScepEventsSearch struct {
	// End of the SCEP event search window, in epoch seconds
	End *int `json:"end,omitempty"`
	// Maximum number of SCEP events returned per page
	Limit *int `json:"limit,omitempty"`
	// Current page of SCEP event search results
	Page *int `json:"page,omitempty"`
	// SCEP PKI operation events returned by a search
	Results []ScepEvent `json:"results,omitempty"`
	// Start of the SCEP event search window, in epoch seconds
	Start *int `json:"start,omitempty"`
	// Number of SCEP events matching the search
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseScepEventsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseScepEventsSearch) String() string {
	return fmt.Sprintf(
		"ResponseScepEventsSearch[End=%v, Limit=%v, Page=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Page, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseScepEventsSearch.
// It customizes the JSON marshaling process for ResponseScepEventsSearch objects.
func (r ResponseScepEventsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "page", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseScepEventsSearch object to a map representation for JSON marshaling.
func (r ResponseScepEventsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	if r.Limit != nil {
		structMap["limit"] = r.Limit
	}
	if r.Page != nil {
		structMap["page"] = r.Page
	}
	if r.Results != nil {
		structMap["results"] = r.Results
	}
	if r.Start != nil {
		structMap["start"] = r.Start
	}
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseScepEventsSearch.
// It customizes the JSON unmarshaling process for ResponseScepEventsSearch objects.
func (r *ResponseScepEventsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseScepEventsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "page", "results", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = temp.End
	r.Limit = temp.Limit
	r.Page = temp.Page
	r.Results = temp.Results
	r.Start = temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponseScepEventsSearch is a temporary struct used for validating the fields of ResponseScepEventsSearch.
type tempResponseScepEventsSearch struct {
	End     *int        `json:"end,omitempty"`
	Limit   *int        `json:"limit,omitempty"`
	Page    *int        `json:"page,omitempty"`
	Results []ScepEvent `json:"results,omitempty"`
	Start   *int        `json:"start,omitempty"`
	Total   *int        `json:"total,omitempty"`
}
