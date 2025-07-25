// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseAssignSuccess represents a ResponseAssignSuccess struct.
type ResponseAssignSuccess struct {
    Success              []string               `json:"success"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAssignSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAssignSuccess) String() string {
    return fmt.Sprintf(
    	"ResponseAssignSuccess[Success=%v, AdditionalProperties=%v]",
    	r.Success, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAssignSuccess.
// It customizes the JSON marshaling process for ResponseAssignSuccess objects.
func (r ResponseAssignSuccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "success"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAssignSuccess object to a map representation for JSON marshaling.
func (r ResponseAssignSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["success"] = r.Success
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAssignSuccess.
// It customizes the JSON unmarshaling process for ResponseAssignSuccess objects.
func (r *ResponseAssignSuccess) UnmarshalJSON(input []byte) error {
    var temp tempResponseAssignSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "success")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Success = *temp.Success
    return nil
}

// tempResponseAssignSuccess is a temporary struct used for validating the fields of ResponseAssignSuccess.
type tempResponseAssignSuccess  struct {
    Success *[]string `json:"success"`
}

func (r *tempResponseAssignSuccess) validate() error {
    var errs []string
    if r.Success == nil {
        errs = append(errs, "required field `success` is missing for type `response_assign_success`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
