// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// EmailString represents a EmailString struct.
type EmailString struct {
    Email                string                 `json:"email"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EmailString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EmailString) String() string {
    return fmt.Sprintf(
    	"EmailString[Email=%v, AdditionalProperties=%v]",
    	e.Email, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EmailString.
// It customizes the JSON marshaling process for EmailString objects.
func (e EmailString) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "email"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EmailString object to a map representation for JSON marshaling.
func (e EmailString) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["email"] = e.Email
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EmailString.
// It customizes the JSON unmarshaling process for EmailString objects.
func (e *EmailString) UnmarshalJSON(input []byte) error {
    var temp tempEmailString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "email")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Email = *temp.Email
    return nil
}

// tempEmailString is a temporary struct used for validating the fields of EmailString.
type tempEmailString  struct {
    Email *string `json:"email"`
}

func (e *tempEmailString) validate() error {
    var errs []string
    if e.Email == nil {
        errs = append(errs, "required field `email` is missing for type `email_string`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
