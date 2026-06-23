// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseEventsRogueSearch represents a ResponseEventsRogueSearch struct.
// Paginated response for rogue AP event search results
type ResponseEventsRogueSearch struct {
	// Epoch timestamp for the end of the rogue AP event search window
	End int `json:"end"`
	// Maximum number of rogue AP event records returned in this page
	Limit int `json:"limit"`
	// Pagination cursor or URL for retrieving the next page of rogue AP event records
	Next *string `json:"next,omitempty"`
	// Rogue AP event records returned by a search response
	Results []EventsRogue `json:"results"`
	// Epoch timestamp for the start of the rogue AP event search window
	Start int `json:"start"`
	// Number of rogue AP event records matching the search filters across all pages
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseEventsRogueSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseEventsRogueSearch) String() string {
	return fmt.Sprintf(
		"ResponseEventsRogueSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsRogueSearch.
// It customizes the JSON marshaling process for ResponseEventsRogueSearch objects.
func (r ResponseEventsRogueSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsRogueSearch object to a map representation for JSON marshaling.
func (r ResponseEventsRogueSearch) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsRogueSearch.
// It customizes the JSON unmarshaling process for ResponseEventsRogueSearch objects.
func (r *ResponseEventsRogueSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseEventsRogueSearch
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

// tempResponseEventsRogueSearch is a temporary struct used for validating the fields of ResponseEventsRogueSearch.
type tempResponseEventsRogueSearch struct {
	End     *int           `json:"end"`
	Limit   *int           `json:"limit"`
	Next    *string        `json:"next,omitempty"`
	Results *[]EventsRogue `json:"results"`
	Start   *int           `json:"start"`
	Total   *int           `json:"total"`
}

func (r *tempResponseEventsRogueSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_events_rogue_search`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_events_rogue_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_events_rogue_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_events_rogue_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_events_rogue_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
