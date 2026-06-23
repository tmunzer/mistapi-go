// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponsePcapSearch represents a ResponsePcapSearch struct.
// Paginated response for packet capture search results
type ResponsePcapSearch struct {
	// Epoch timestamp, in seconds, for the end of the packet capture search window
	End int `json:"end"`
	// Maximum number of packet capture records returned in this page
	Limit int `json:"limit"`
	// URL for retrieving the next page of packet capture search results
	Next *string `json:"next,omitempty"`
	// Packet capture records returned by search
	Results []ResponsePcapSearchItem `json:"results"`
	// Epoch timestamp, in seconds, for the start of the packet capture search window
	Start int `json:"start"`
	// Number of packet capture records matching the search filters across all pages
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapSearch) String() string {
	return fmt.Sprintf(
		"ResponsePcapSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapSearch.
// It customizes the JSON marshaling process for ResponsePcapSearch objects.
func (r ResponsePcapSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapSearch object to a map representation for JSON marshaling.
func (r ResponsePcapSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["limit"] = r.Limit
	if r.Next != nil {
		structMap["next"] = r.Next
	}
	structMap["results"] = r.Results
	structMap["start"] = r.Start
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapSearch.
// It customizes the JSON unmarshaling process for ResponsePcapSearch objects.
func (r *ResponsePcapSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponsePcapSearch
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
	r.Total = temp.Total
	return nil
}

// tempResponsePcapSearch is a temporary struct used for validating the fields of ResponsePcapSearch.
type tempResponsePcapSearch struct {
	End     *int                      `json:"end"`
	Limit   *int                      `json:"limit"`
	Next    *string                   `json:"next,omitempty"`
	Results *[]ResponsePcapSearchItem `json:"results"`
	Start   *int                      `json:"start"`
	Total   *int                      `json:"total,omitempty"`
}

func (r *tempResponsePcapSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_pcap_search`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_pcap_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_pcap_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_pcap_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
