// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SearchWirelessClientSession represents a SearchWirelessClientSession struct.
type SearchWirelessClientSession struct {
	End                  float64                 `json:"end"`
	Limit                int                     `json:"limit"`
	Next                 *string                 `json:"next,omitempty"`
	Results              []WirelessClientSession `json:"results"`
	Start                float64                 `json:"start"`
	Total                int                     `json:"total"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWirelessClientSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWirelessClientSession) String() string {
	return fmt.Sprintf(
		"SearchWirelessClientSession[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		s.End, s.Limit, s.Next, s.Results, s.Start, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWirelessClientSession.
// It customizes the JSON marshaling process for SearchWirelessClientSession objects.
func (s SearchWirelessClientSession) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SearchWirelessClientSession object to a map representation for JSON marshaling.
func (s SearchWirelessClientSession) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWirelessClientSession.
// It customizes the JSON unmarshaling process for SearchWirelessClientSession objects.
func (s *SearchWirelessClientSession) UnmarshalJSON(input []byte) error {
	var temp tempSearchWirelessClientSession
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

// tempSearchWirelessClientSession is a temporary struct used for validating the fields of SearchWirelessClientSession.
type tempSearchWirelessClientSession struct {
	End     *float64                 `json:"end"`
	Limit   *int                     `json:"limit"`
	Next    *string                  `json:"next,omitempty"`
	Results *[]WirelessClientSession `json:"results"`
	Start   *float64                 `json:"start"`
	Total   *int                     `json:"total"`
}

func (s *tempSearchWirelessClientSession) validate() error {
	var errs []string
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `search_wireless_client_session`")
	}
	if s.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `search_wireless_client_session`")
	}
	if s.Results == nil {
		errs = append(errs, "required field `results` is missing for type `search_wireless_client_session`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `search_wireless_client_session`")
	}
	if s.Total == nil {
		errs = append(errs, "required field `total` is missing for type `search_wireless_client_session`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
