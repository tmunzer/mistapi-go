// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseClientSessionsSearch represents a ResponseClientSessionsSearch struct.
type ResponseClientSessionsSearch struct {
	End                  int                                `json:"end"`
	Limit                int                                `json:"limit"`
	Next                 *string                            `json:"next,omitempty"`
	Results              []ResponseClientSessionsSearchItem `json:"results"`
	Start                int                                `json:"start"`
	Total                int                                `json:"total"`
	AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClientSessionsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClientSessionsSearch) String() string {
	return fmt.Sprintf(
		"ResponseClientSessionsSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClientSessionsSearch.
// It customizes the JSON marshaling process for ResponseClientSessionsSearch objects.
func (r ResponseClientSessionsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseClientSessionsSearch object to a map representation for JSON marshaling.
func (r ResponseClientSessionsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["limit"] = r.Limit
	if r.Next != nil {
		structMap["next"] = r.Next
	}
	structMap["results"] = r.Results
	structMap["start"] = r.Start
	structMap["total"] = r.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClientSessionsSearch.
// It customizes the JSON unmarshaling process for ResponseClientSessionsSearch objects.
func (r *ResponseClientSessionsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseClientSessionsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = *temp.End
	r.Limit = *temp.Limit
	r.Next = temp.Next
	r.Results = *temp.Results
	r.Start = *temp.Start
	r.Total = *temp.Total
	return nil
}

// tempResponseClientSessionsSearch is a temporary struct used for validating the fields of ResponseClientSessionsSearch.
type tempResponseClientSessionsSearch struct {
	End     *int                                `json:"end"`
	Limit   *int                                `json:"limit"`
	Next    *string                             `json:"next,omitempty"`
	Results *[]ResponseClientSessionsSearchItem `json:"results"`
	Start   *int                                `json:"start"`
	Total   *int                                `json:"total"`
}

func (r *tempResponseClientSessionsSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_client_sessions_search`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_client_sessions_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_client_sessions_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_client_sessions_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_client_sessions_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
