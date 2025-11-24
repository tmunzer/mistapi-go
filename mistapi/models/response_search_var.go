// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseSearchVar represents a ResponseSearchVar struct.
type ResponseSearchVar struct {
	End                  *int                    `json:"end,omitempty"`
	Limit                *int                    `json:"limit,omitempty"`
	Next                 *string                 `json:"next,omitempty"`
	Results              []ResponseSearchVarItem `json:"results,omitempty"`
	Start                *int                    `json:"start,omitempty"`
	Total                *int                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSearchVar,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSearchVar) String() string {
	return fmt.Sprintf(
		"ResponseSearchVar[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchVar.
// It customizes the JSON marshaling process for ResponseSearchVar objects.
func (r ResponseSearchVar) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchVar object to a map representation for JSON marshaling.
func (r ResponseSearchVar) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	if r.Limit != nil {
		structMap["limit"] = r.Limit
	}
	if r.Next != nil {
		structMap["next"] = r.Next
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearchVar.
// It customizes the JSON unmarshaling process for ResponseSearchVar objects.
func (r *ResponseSearchVar) UnmarshalJSON(input []byte) error {
	var temp tempResponseSearchVar
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = temp.End
	r.Limit = temp.Limit
	r.Next = temp.Next
	r.Results = temp.Results
	r.Start = temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponseSearchVar is a temporary struct used for validating the fields of ResponseSearchVar.
type tempResponseSearchVar struct {
	End     *int                    `json:"end,omitempty"`
	Limit   *int                    `json:"limit,omitempty"`
	Next    *string                 `json:"next,omitempty"`
	Results []ResponseSearchVarItem `json:"results,omitempty"`
	Start   *int                    `json:"start,omitempty"`
	Total   *int                    `json:"total,omitempty"`
}
