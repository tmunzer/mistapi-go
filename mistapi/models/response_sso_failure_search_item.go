// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseSsoFailureSearchItem represents a ResponseSsoFailureSearchItem struct.
type ResponseSsoFailureSearchItem struct {
	Detail           string `json:"detail"`
	SamlAssertionXml string `json:"saml_assertion_xml"`
	// Epoch (seconds)
	Timestamp            float64                `json:"timestamp"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsoFailureSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsoFailureSearchItem) String() string {
	return fmt.Sprintf(
		"ResponseSsoFailureSearchItem[Detail=%v, SamlAssertionXml=%v, Timestamp=%v, AdditionalProperties=%v]",
		r.Detail, r.SamlAssertionXml, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsoFailureSearchItem.
// It customizes the JSON marshaling process for ResponseSsoFailureSearchItem objects.
func (r ResponseSsoFailureSearchItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"detail", "saml_assertion_xml", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsoFailureSearchItem object to a map representation for JSON marshaling.
func (r ResponseSsoFailureSearchItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["detail"] = r.Detail
	structMap["saml_assertion_xml"] = r.SamlAssertionXml
	structMap["timestamp"] = r.Timestamp
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsoFailureSearchItem.
// It customizes the JSON unmarshaling process for ResponseSsoFailureSearchItem objects.
func (r *ResponseSsoFailureSearchItem) UnmarshalJSON(input []byte) error {
	var temp tempResponseSsoFailureSearchItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "detail", "saml_assertion_xml", "timestamp")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Detail = *temp.Detail
	r.SamlAssertionXml = *temp.SamlAssertionXml
	r.Timestamp = *temp.Timestamp
	return nil
}

// tempResponseSsoFailureSearchItem is a temporary struct used for validating the fields of ResponseSsoFailureSearchItem.
type tempResponseSsoFailureSearchItem struct {
	Detail           *string  `json:"detail"`
	SamlAssertionXml *string  `json:"saml_assertion_xml"`
	Timestamp        *float64 `json:"timestamp"`
}

func (r *tempResponseSsoFailureSearchItem) validate() error {
	var errs []string
	if r.Detail == nil {
		errs = append(errs, "required field `detail` is missing for type `response_sso_failure_search_item`")
	}
	if r.SamlAssertionXml == nil {
		errs = append(errs, "required field `saml_assertion_xml` is missing for type `response_sso_failure_search_item`")
	}
	if r.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `response_sso_failure_search_item`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
