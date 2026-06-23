// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseSwitchPortSearch represents a ResponseSwitchPortSearch struct.
// Paginated response for switch and gateway port statistics search results
type ResponseSwitchPortSearch struct {
	// Epoch timestamp, in seconds, for the end of the port statistics reporting window
	End *int `json:"end,omitempty"`
	// Maximum number of port statistics records returned in this page
	Limit int `json:"limit"`
	// URL for retrieving the next page of switch and gateway port statistics results
	Next *string `json:"next,omitempty"`
	// Switch and gateway port statistics records returned by search
	Results []StatsSwitchPort `json:"results"`
	// Epoch timestamp, in seconds, for the start of the port statistics reporting window
	Start *int `json:"start,omitempty"`
	// Number of port statistics records matching the search filters across all pages
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchPortSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchPortSearch) String() string {
	return fmt.Sprintf(
		"ResponseSwitchPortSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchPortSearch.
// It customizes the JSON marshaling process for ResponseSwitchPortSearch objects.
func (r ResponseSwitchPortSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchPortSearch object to a map representation for JSON marshaling.
func (r ResponseSwitchPortSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	structMap["limit"] = r.Limit
	if r.Next != nil {
		structMap["next"] = r.Next
	}
	structMap["results"] = r.Results
	if r.Start != nil {
		structMap["start"] = r.Start
	}
	structMap["total"] = r.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchPortSearch.
// It customizes the JSON unmarshaling process for ResponseSwitchPortSearch objects.
func (r *ResponseSwitchPortSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseSwitchPortSearch
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

	r.End = temp.End
	r.Limit = *temp.Limit
	r.Next = temp.Next
	r.Results = *temp.Results
	r.Start = temp.Start
	r.Total = *temp.Total
	return nil
}

// tempResponseSwitchPortSearch is a temporary struct used for validating the fields of ResponseSwitchPortSearch.
type tempResponseSwitchPortSearch struct {
	End     *int               `json:"end,omitempty"`
	Limit   *int               `json:"limit"`
	Next    *string            `json:"next,omitempty"`
	Results *[]StatsSwitchPort `json:"results"`
	Start   *int               `json:"start,omitempty"`
	Total   *int               `json:"total"`
}

func (r *tempResponseSwitchPortSearch) validate() error {
	var errs []string
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_switch_port_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_switch_port_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_switch_port_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
