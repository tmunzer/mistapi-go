// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SearchWiredClient represents a SearchWiredClient struct.
type SearchWiredClient struct {
	End                  float64                `json:"end"`
	Limit                int                    `json:"limit"`
	Next                 *string                `json:"next,omitempty"`
	Results              []WiredClientResponse  `json:"results"`
	Start                float64                `json:"start"`
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWiredClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWiredClient) String() string {
	return fmt.Sprintf(
		"SearchWiredClient[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		s.End, s.Limit, s.Next, s.Results, s.Start, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWiredClient.
// It customizes the JSON marshaling process for SearchWiredClient objects.
func (s SearchWiredClient) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SearchWiredClient object to a map representation for JSON marshaling.
func (s SearchWiredClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["end"] = s.End
	structMap["limit"] = s.Limit
	if s.Next != nil {
		structMap["next"] = s.Next
	}
	structMap["results"] = s.Results
	structMap["start"] = s.Start
	structMap["total"] = s.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWiredClient.
// It customizes the JSON unmarshaling process for SearchWiredClient objects.
func (s *SearchWiredClient) UnmarshalJSON(input []byte) error {
	var temp tempSearchWiredClient
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
	s.AdditionalProperties = additionalProperties

	s.End = *temp.End
	s.Limit = *temp.Limit
	s.Next = temp.Next
	s.Results = *temp.Results
	s.Start = *temp.Start
	s.Total = *temp.Total
	return nil
}

// tempSearchWiredClient is a temporary struct used for validating the fields of SearchWiredClient.
type tempSearchWiredClient struct {
	End     *float64               `json:"end"`
	Limit   *int                   `json:"limit"`
	Next    *string                `json:"next,omitempty"`
	Results *[]WiredClientResponse `json:"results"`
	Start   *float64               `json:"start"`
	Total   *int                   `json:"total"`
}

func (s *tempSearchWiredClient) validate() error {
	var errs []string
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `search_wired_client`")
	}
	if s.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `search_wired_client`")
	}
	if s.Results == nil {
		errs = append(errs, "required field `results` is missing for type `search_wired_client`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `search_wired_client`")
	}
	if s.Total == nil {
		errs = append(errs, "required field `total` is missing for type `search_wired_client`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
