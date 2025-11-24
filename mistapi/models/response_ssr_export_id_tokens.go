// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseSsrExportIdTokens represents a ResponseSsrExportIdTokens struct.
type ResponseSsrExportIdTokens struct {
	Results              []ResponseSsrExportIdTokensResultsItem `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrExportIdTokens,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrExportIdTokens) String() string {
	return fmt.Sprintf(
		"ResponseSsrExportIdTokens[Results=%v, AdditionalProperties=%v]",
		r.Results, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrExportIdTokens.
// It customizes the JSON marshaling process for ResponseSsrExportIdTokens objects.
func (r ResponseSsrExportIdTokens) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"results"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrExportIdTokens object to a map representation for JSON marshaling.
func (r ResponseSsrExportIdTokens) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Results != nil {
		structMap["results"] = r.Results
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrExportIdTokens.
// It customizes the JSON unmarshaling process for ResponseSsrExportIdTokens objects.
func (r *ResponseSsrExportIdTokens) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsrExportIdTokens
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Results = temp.Results
	return nil
}

// tempResponseSsrExportIdTokens is a temporary struct used for validating the fields of ResponseSsrExportIdTokens.
type tempResponseSsrExportIdTokens struct {
	Results []ResponseSsrExportIdTokensResultsItem `json:"results,omitempty"`
}
