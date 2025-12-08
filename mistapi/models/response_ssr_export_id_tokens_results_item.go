// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseSsrExportIdTokensResultsItem represents a ResponseSsrExportIdTokensResultsItem struct.
type ResponseSsrExportIdTokensResultsItem struct {
	Mac                  *string                `json:"mac,omitempty"`
	Token                *string                `json:"token,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrExportIdTokensResultsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrExportIdTokensResultsItem) String() string {
	return fmt.Sprintf(
		"ResponseSsrExportIdTokensResultsItem[Mac=%v, Token=%v, AdditionalProperties=%v]",
		r.Mac, r.Token, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrExportIdTokensResultsItem.
// It customizes the JSON marshaling process for ResponseSsrExportIdTokensResultsItem objects.
func (r ResponseSsrExportIdTokensResultsItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"mac", "token"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrExportIdTokensResultsItem object to a map representation for JSON marshaling.
func (r ResponseSsrExportIdTokensResultsItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Mac != nil {
		structMap["mac"] = r.Mac
	}
	if r.Token != nil {
		structMap["token"] = r.Token
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrExportIdTokensResultsItem.
// It customizes the JSON unmarshaling process for ResponseSsrExportIdTokensResultsItem objects.
func (r *ResponseSsrExportIdTokensResultsItem) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsrExportIdTokensResultsItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "token")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Mac = temp.Mac
	r.Token = temp.Token
	return nil
}

// tempResponseSsrExportIdTokensResultsItem is a temporary struct used for validating the fields of ResponseSsrExportIdTokensResultsItem.
type tempResponseSsrExportIdTokensResultsItem struct {
	Mac   *string `json:"mac,omitempty"`
	Token *string `json:"token,omitempty"`
}
