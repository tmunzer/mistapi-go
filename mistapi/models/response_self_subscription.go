package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// ResponseSelfSubscription represents a ResponseSelfSubscription struct.
type ResponseSelfSubscription struct {
	OrgId                uuid.UUID      `json:"org_id"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSelfSubscription.
// It customizes the JSON marshaling process for ResponseSelfSubscription objects.
func (r ResponseSelfSubscription) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSelfSubscription object to a map representation for JSON marshaling.
func (r ResponseSelfSubscription) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "org_id")
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
