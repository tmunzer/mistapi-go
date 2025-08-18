// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseOrgSystemEventsSearch represents a ResponseOrgSystemEventsSearch struct.
type ResponseOrgSystemEventsSearch struct {
	End                  int                    `json:"end"`
	Limit                int                    `json:"limit"`
	Next                 *string                `json:"next,omitempty"`
	Results              []OrgSystemEvent       `json:"results"`
	Start                int                    `json:"start"`
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgSystemEventsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgSystemEventsSearch) String() string {
	return fmt.Sprintf(
		"ResponseOrgSystemEventsSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSystemEventsSearch.
// It customizes the JSON marshaling process for ResponseOrgSystemEventsSearch objects.
func (r ResponseOrgSystemEventsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSystemEventsSearch object to a map representation for JSON marshaling.
func (r ResponseOrgSystemEventsSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSystemEventsSearch.
// It customizes the JSON unmarshaling process for ResponseOrgSystemEventsSearch objects.
func (r *ResponseOrgSystemEventsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseOrgSystemEventsSearch
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

// tempResponseOrgSystemEventsSearch is a temporary struct used for validating the fields of ResponseOrgSystemEventsSearch.
type tempResponseOrgSystemEventsSearch struct {
	End     *int              `json:"end"`
	Limit   *int              `json:"limit"`
	Next    *string           `json:"next,omitempty"`
	Results *[]OrgSystemEvent `json:"results"`
	Start   *int              `json:"start"`
	Total   *int              `json:"total"`
}

func (r *tempResponseOrgSystemEventsSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_org_system_events_search`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_org_system_events_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_org_system_events_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_org_system_events_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_org_system_events_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
