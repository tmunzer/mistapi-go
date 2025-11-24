// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OspfPeerStatsSearchResult represents a OspfPeerStatsSearchResult struct.
type OspfPeerStatsSearchResult struct {
	End                  *int                              `json:"end,omitempty"`
	Limit                *int                              `json:"limit,omitempty"`
	Next                 *string                           `json:"next,omitempty"`
	Results              []OspfPeerStatsSearchResultsItems `json:"results,omitempty"`
	Start                *int                              `json:"start,omitempty"`
	Total                *int                              `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for OspfPeerStatsSearchResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OspfPeerStatsSearchResult) String() string {
	return fmt.Sprintf(
		"OspfPeerStatsSearchResult[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		o.End, o.Limit, o.Next, o.Results, o.Start, o.Total, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OspfPeerStatsSearchResult.
// It customizes the JSON marshaling process for OspfPeerStatsSearchResult objects.
func (o OspfPeerStatsSearchResult) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OspfPeerStatsSearchResult object to a map representation for JSON marshaling.
func (o OspfPeerStatsSearchResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.End != nil {
		structMap["end"] = o.End
	}
	if o.Limit != nil {
		structMap["limit"] = o.Limit
	}
	if o.Next != nil {
		structMap["next"] = o.Next
	}
	if o.Results != nil {
		structMap["results"] = o.Results
	}
	if o.Start != nil {
		structMap["start"] = o.Start
	}
	if o.Total != nil {
		structMap["total"] = o.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfPeerStatsSearchResult.
// It customizes the JSON unmarshaling process for OspfPeerStatsSearchResult objects.
func (o *OspfPeerStatsSearchResult) UnmarshalJSON(input []byte) error {
	var temp tempOspfPeerStatsSearchResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.End = temp.End
	o.Limit = temp.Limit
	o.Next = temp.Next
	o.Results = temp.Results
	o.Start = temp.Start
	o.Total = temp.Total
	return nil
}

// tempOspfPeerStatsSearchResult is a temporary struct used for validating the fields of OspfPeerStatsSearchResult.
type tempOspfPeerStatsSearchResult struct {
	End     *int                              `json:"end,omitempty"`
	Limit   *int                              `json:"limit,omitempty"`
	Next    *string                           `json:"next,omitempty"`
	Results []OspfPeerStatsSearchResultsItems `json:"results,omitempty"`
	Start   *int                              `json:"start,omitempty"`
	Total   *int                              `json:"total,omitempty"`
}
