// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// ResponseSelfSubscription represents a ResponseSelfSubscription struct.
type ResponseSelfSubscription struct {
	OrgId                uuid.UUID              `json:"org_id"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSelfSubscription,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSelfSubscription) String() string {
	return fmt.Sprintf(
		"ResponseSelfSubscription[OrgId=%v, AdditionalProperties=%v]",
		r.OrgId, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSelfSubscription.
// It customizes the JSON marshaling process for ResponseSelfSubscription objects.
func (r ResponseSelfSubscription) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"org_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSelfSubscription object to a map representation for JSON marshaling.
func (r ResponseSelfSubscription) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["org_id"] = r.OrgId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSelfSubscription.
// It customizes the JSON unmarshaling process for ResponseSelfSubscription objects.
func (r *ResponseSelfSubscription) UnmarshalJSON(input []byte) error {
	var temp tempResponseSelfSubscription
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.OrgId = *temp.OrgId
	return nil
}

// tempResponseSelfSubscription is a temporary struct used for validating the fields of ResponseSelfSubscription.
type tempResponseSelfSubscription struct {
	OrgId *uuid.UUID `json:"org_id"`
}

func (r *tempResponseSelfSubscription) validate() error {
	var errs []string
	if r.OrgId == nil {
		errs = append(errs, "required field `org_id` is missing for type `response_self_subscription`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
