// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseSelfOauthUrl represents a ResponseSelfOauthUrl struct.
type ResponseSelfOauthUrl struct {
	AuthorizationUrl     string                 `json:"authorization_url"`
	Linked               bool                   `json:"linked"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSelfOauthUrl,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSelfOauthUrl) String() string {
	return fmt.Sprintf(
		"ResponseSelfOauthUrl[AuthorizationUrl=%v, Linked=%v, AdditionalProperties=%v]",
		r.AuthorizationUrl, r.Linked, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSelfOauthUrl.
// It customizes the JSON marshaling process for ResponseSelfOauthUrl objects.
func (r ResponseSelfOauthUrl) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"authorization_url", "linked"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSelfOauthUrl object to a map representation for JSON marshaling.
func (r ResponseSelfOauthUrl) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["authorization_url"] = r.AuthorizationUrl
	structMap["linked"] = r.Linked
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSelfOauthUrl.
// It customizes the JSON unmarshaling process for ResponseSelfOauthUrl objects.
func (r *ResponseSelfOauthUrl) UnmarshalJSON(input []byte) error {
	var temp tempResponseSelfOauthUrl
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authorization_url", "linked")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.AuthorizationUrl = *temp.AuthorizationUrl
	r.Linked = *temp.Linked
	return nil
}

// tempResponseSelfOauthUrl is a temporary struct used for validating the fields of ResponseSelfOauthUrl.
type tempResponseSelfOauthUrl struct {
	AuthorizationUrl *string `json:"authorization_url"`
	Linked           *bool   `json:"linked"`
}

func (r *tempResponseSelfOauthUrl) validate() error {
	var errs []string
	if r.AuthorizationUrl == nil {
		errs = append(errs, "required field `authorization_url` is missing for type `response_self_oauth_url`")
	}
	if r.Linked == nil {
		errs = append(errs, "required field `linked` is missing for type `response_self_oauth_url`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
