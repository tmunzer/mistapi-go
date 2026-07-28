// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseDeviceFlowRecordsSearch represents a ResponseDeviceFlowRecordsSearch struct.
// Paginated response for device flow record search results
type ResponseDeviceFlowRecordsSearch struct {
	// Epoch timestamp, in seconds, for the end of the flow record search window
	End *int `json:"end,omitempty"`
	// Maximum number of flow records returned in this page
	Limit *int `json:"limit,omitempty"`
	// Flow records matching the search filters
	Results []FlowRecord `json:"results,omitempty"`
	// Cursor token for retrieving the next page of flow records
	SearchAfter *string `json:"search_after,omitempty"`
	// Epoch timestamp, in seconds, for the start of the flow record search window
	Start *int `json:"start,omitempty"`
	// Number of flow records matching the search filters
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceFlowRecordsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceFlowRecordsSearch) String() string {
	return fmt.Sprintf(
		"ResponseDeviceFlowRecordsSearch[End=%v, Limit=%v, Results=%v, SearchAfter=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Results, r.SearchAfter, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceFlowRecordsSearch.
// It customizes the JSON marshaling process for ResponseDeviceFlowRecordsSearch objects.
func (r ResponseDeviceFlowRecordsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "results", "search_after", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceFlowRecordsSearch object to a map representation for JSON marshaling.
func (r ResponseDeviceFlowRecordsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	if r.Limit != nil {
		structMap["limit"] = r.Limit
	}
	if r.Results != nil {
		structMap["results"] = r.Results
	}
	if r.SearchAfter != nil {
		structMap["search_after"] = r.SearchAfter
	}
	if r.Start != nil {
		structMap["start"] = r.Start
	}
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceFlowRecordsSearch.
// It customizes the JSON unmarshaling process for ResponseDeviceFlowRecordsSearch objects.
func (r *ResponseDeviceFlowRecordsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseDeviceFlowRecordsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "search_after", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = temp.End
	r.Limit = temp.Limit
	r.Results = temp.Results
	r.SearchAfter = temp.SearchAfter
	r.Start = temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponseDeviceFlowRecordsSearch is a temporary struct used for validating the fields of ResponseDeviceFlowRecordsSearch.
type tempResponseDeviceFlowRecordsSearch struct {
	End         *int         `json:"end,omitempty"`
	Limit       *int         `json:"limit,omitempty"`
	Results     []FlowRecord `json:"results,omitempty"`
	SearchAfter *string      `json:"search_after,omitempty"`
	Start       *int         `json:"start,omitempty"`
	Total       *int         `json:"total,omitempty"`
}
