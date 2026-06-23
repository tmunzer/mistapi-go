// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// IssuedClientCertificatesResults represents a IssuedClientCertificatesResults struct.
// Issued client certificate search results wrapper
type IssuedClientCertificatesResults struct {
	// Maximum number of results requested
	Limit *int `json:"limit,omitempty"`
	// Current page number
	Page *int `json:"page,omitempty"`
	// Issued client certificates returned by a SCEP certificate query
	Results              []IssuedClientCertificate `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for IssuedClientCertificatesResults,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssuedClientCertificatesResults) String() string {
	return fmt.Sprintf(
		"IssuedClientCertificatesResults[Limit=%v, Page=%v, Results=%v, AdditionalProperties=%v]",
		i.Limit, i.Page, i.Results, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssuedClientCertificatesResults.
// It customizes the JSON marshaling process for IssuedClientCertificatesResults objects.
func (i IssuedClientCertificatesResults) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"limit", "page", "results"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the IssuedClientCertificatesResults object to a map representation for JSON marshaling.
func (i IssuedClientCertificatesResults) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Limit != nil {
		structMap["limit"] = i.Limit
	}
	if i.Page != nil {
		structMap["page"] = i.Page
	}
	if i.Results != nil {
		structMap["results"] = i.Results
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssuedClientCertificatesResults.
// It customizes the JSON unmarshaling process for IssuedClientCertificatesResults objects.
func (i *IssuedClientCertificatesResults) UnmarshalJSON(input []byte) error {
	var temp tempIssuedClientCertificatesResults
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "limit", "page", "results")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.Limit = temp.Limit
	i.Page = temp.Page
	i.Results = temp.Results
	return nil
}

// tempIssuedClientCertificatesResults is a temporary struct used for validating the fields of IssuedClientCertificatesResults.
type tempIssuedClientCertificatesResults struct {
	Limit   *int                      `json:"limit,omitempty"`
	Page    *int                      `json:"page,omitempty"`
	Results []IssuedClientCertificate `json:"results,omitempty"`
}
