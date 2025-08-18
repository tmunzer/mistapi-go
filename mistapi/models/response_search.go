// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseSearch represents a ResponseSearch struct.
type ResponseSearch struct {
	Limit                int                    `json:"limit"`
	Next                 *string                `json:"next,omitempty"`
	Page                 int                    `json:"page"`
	Results              []ResponseSearchItem   `json:"results"`
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSearch) String() string {
	return fmt.Sprintf(
		"ResponseSearch[Limit=%v, Next=%v, Page=%v, Results=%v, Total=%v, AdditionalProperties=%v]",
		r.Limit, r.Next, r.Page, r.Results, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearch.
// It customizes the JSON marshaling process for ResponseSearch objects.
func (r ResponseSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"limit", "next", "page", "results", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearch object to a map representation for JSON marshaling.
func (r ResponseSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["limit"] = r.Limit
	if r.Next != nil {
		structMap["next"] = r.Next
	}
	structMap["page"] = r.Page
	structMap["results"] = r.Results
	structMap["total"] = r.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearch.
// It customizes the JSON unmarshaling process for ResponseSearch objects.
func (r *ResponseSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "limit", "next", "page", "results", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Limit = *temp.Limit
	r.Next = temp.Next
	r.Page = *temp.Page
	r.Results = *temp.Results
	r.Total = *temp.Total
	return nil
}

// tempResponseSearch is a temporary struct used for validating the fields of ResponseSearch.
type tempResponseSearch struct {
	Limit   *int                  `json:"limit"`
	Next    *string               `json:"next,omitempty"`
	Page    *int                  `json:"page"`
	Results *[]ResponseSearchItem `json:"results"`
	Total   *int                  `json:"total"`
}

func (r *tempResponseSearch) validate() error {
	var errs []string
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_search`")
	}
	if r.Page == nil {
		errs = append(errs, "required field `page` is missing for type `response_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
