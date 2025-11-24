// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// InventorySearch represents a InventorySearch struct.
type InventorySearch struct {
	End                  *int                    `json:"end,omitempty"`
	Limit                *int                    `json:"limit,omitempty"`
	Next                 *string                 `json:"next,omitempty"`
	Results              []InventorySearchResult `json:"results,omitempty"`
	Start                *int                    `json:"start,omitempty"`
	Total                *int                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for InventorySearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InventorySearch) String() string {
	return fmt.Sprintf(
		"InventorySearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		i.End, i.Limit, i.Next, i.Results, i.Start, i.Total, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InventorySearch.
// It customizes the JSON marshaling process for InventorySearch objects.
func (i InventorySearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InventorySearch object to a map representation for JSON marshaling.
func (i InventorySearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.End != nil {
		structMap["end"] = i.End
	}
	if i.Limit != nil {
		structMap["limit"] = i.Limit
	}
	if i.Next != nil {
		structMap["next"] = i.Next
	}
	if i.Results != nil {
		structMap["results"] = i.Results
	}
	if i.Start != nil {
		structMap["start"] = i.Start
	}
	if i.Total != nil {
		structMap["total"] = i.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InventorySearch.
// It customizes the JSON unmarshaling process for InventorySearch objects.
func (i *InventorySearch) UnmarshalJSON(input []byte) error {
	var temp tempInventorySearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.End = temp.End
	i.Limit = temp.Limit
	i.Next = temp.Next
	i.Results = temp.Results
	i.Start = temp.Start
	i.Total = temp.Total
	return nil
}

// tempInventorySearch is a temporary struct used for validating the fields of InventorySearch.
type tempInventorySearch struct {
	End     *int                    `json:"end,omitempty"`
	Limit   *int                    `json:"limit,omitempty"`
	Next    *string                 `json:"next,omitempty"`
	Results []InventorySearchResult `json:"results,omitempty"`
	Start   *int                    `json:"start,omitempty"`
	Total   *int                    `json:"total,omitempty"`
}
