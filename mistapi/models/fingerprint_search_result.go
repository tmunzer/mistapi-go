// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// FingerprintSearchResult represents a FingerprintSearchResult struct.
type FingerprintSearchResult struct {
	End                  int                    `json:"end"`
	Limit                int                    `json:"limit"`
	Next                 *string                `json:"next,omitempty"`
	Results              []Fingerprint          `json:"results"`
	Start                int                    `json:"start"`
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FingerprintSearchResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FingerprintSearchResult) String() string {
	return fmt.Sprintf(
		"FingerprintSearchResult[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		f.End, f.Limit, f.Next, f.Results, f.Start, f.Total, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FingerprintSearchResult.
// It customizes the JSON marshaling process for FingerprintSearchResult objects.
func (f FingerprintSearchResult) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FingerprintSearchResult object to a map representation for JSON marshaling.
func (f FingerprintSearchResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	structMap["end"] = f.End
	structMap["limit"] = f.Limit
	if f.Next != nil {
		structMap["next"] = f.Next
	}
	structMap["results"] = f.Results
	structMap["start"] = f.Start
	structMap["total"] = f.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FingerprintSearchResult.
// It customizes the JSON unmarshaling process for FingerprintSearchResult objects.
func (f *FingerprintSearchResult) UnmarshalJSON(input []byte) error {
	var temp tempFingerprintSearchResult
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
	f.AdditionalProperties = additionalProperties

	f.End = *temp.End
	f.Limit = *temp.Limit
	f.Next = temp.Next
	f.Results = *temp.Results
	f.Start = *temp.Start
	f.Total = *temp.Total
	return nil
}

// tempFingerprintSearchResult is a temporary struct used for validating the fields of FingerprintSearchResult.
type tempFingerprintSearchResult struct {
	End     *int           `json:"end"`
	Limit   *int           `json:"limit"`
	Next    *string        `json:"next,omitempty"`
	Results *[]Fingerprint `json:"results"`
	Start   *int           `json:"start"`
	Total   *int           `json:"total"`
}

func (f *tempFingerprintSearchResult) validate() error {
	var errs []string
	if f.End == nil {
		errs = append(errs, "required field `end` is missing for type `fingerprint_search_result`")
	}
	if f.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `fingerprint_search_result`")
	}
	if f.Results == nil {
		errs = append(errs, "required field `results` is missing for type `fingerprint_search_result`")
	}
	if f.Start == nil {
		errs = append(errs, "required field `start` is missing for type `fingerprint_search_result`")
	}
	if f.Total == nil {
		errs = append(errs, "required field `total` is missing for type `fingerprint_search_result`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
