// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAsyncClaimsList represents a ResponseAsyncClaimsList struct.
// List of async inventory claim jobs for the organization
type ResponseAsyncClaimsList struct {
	// Async claim job status records
	Claims               []ResponseAsyncClaimStatus `json:"claims,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncClaimsList,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncClaimsList) String() string {
	return fmt.Sprintf(
		"ResponseAsyncClaimsList[Claims=%v, AdditionalProperties=%v]",
		r.Claims, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncClaimsList.
// It customizes the JSON marshaling process for ResponseAsyncClaimsList objects.
func (r ResponseAsyncClaimsList) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"claims"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncClaimsList object to a map representation for JSON marshaling.
func (r ResponseAsyncClaimsList) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Claims != nil {
		structMap["claims"] = r.Claims
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncClaimsList.
// It customizes the JSON unmarshaling process for ResponseAsyncClaimsList objects.
func (r *ResponseAsyncClaimsList) UnmarshalJSON(input []byte) error {
	var temp tempResponseAsyncClaimsList
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "claims")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Claims = temp.Claims
	return nil
}

// tempResponseAsyncClaimsList is a temporary struct used for validating the fields of ResponseAsyncClaimsList.
type tempResponseAsyncClaimsList struct {
	Claims []ResponseAsyncClaimStatus `json:"claims,omitempty"`
}
