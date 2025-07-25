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

// ResponseMobileVerifySecret represents a ResponseMobileVerifySecret struct.
type ResponseMobileVerifySecret struct {
    Name                 string                 `json:"name"`
    OrgId                uuid.UUID              `json:"org_id"`
    Secret               string                 `json:"secret"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMobileVerifySecret,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMobileVerifySecret) String() string {
    return fmt.Sprintf(
    	"ResponseMobileVerifySecret[Name=%v, OrgId=%v, Secret=%v, AdditionalProperties=%v]",
    	r.Name, r.OrgId, r.Secret, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMobileVerifySecret.
// It customizes the JSON marshaling process for ResponseMobileVerifySecret objects.
func (r ResponseMobileVerifySecret) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "name", "org_id", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMobileVerifySecret object to a map representation for JSON marshaling.
func (r ResponseMobileVerifySecret) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["name"] = r.Name
    structMap["org_id"] = r.OrgId
    structMap["secret"] = r.Secret
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMobileVerifySecret.
// It customizes the JSON unmarshaling process for ResponseMobileVerifySecret objects.
func (r *ResponseMobileVerifySecret) UnmarshalJSON(input []byte) error {
    var temp tempResponseMobileVerifySecret
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "org_id", "secret")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Name = *temp.Name
    r.OrgId = *temp.OrgId
    r.Secret = *temp.Secret
    return nil
}

// tempResponseMobileVerifySecret is a temporary struct used for validating the fields of ResponseMobileVerifySecret.
type tempResponseMobileVerifySecret  struct {
    Name   *string    `json:"name"`
    OrgId  *uuid.UUID `json:"org_id"`
    Secret *string    `json:"secret"`
}

func (r *tempResponseMobileVerifySecret) validate() error {
    var errs []string
    if r.Name == nil {
        errs = append(errs, "required field `name` is missing for type `response_mobile_verify_secret`")
    }
    if r.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `response_mobile_verify_secret`")
    }
    if r.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `response_mobile_verify_secret`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
