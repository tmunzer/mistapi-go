// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsZeroizeFips represents a UtilsZeroizeFips struct.
type UtilsZeroizeFips struct {
    // FIPS zeroize password
    Password             string                 `json:"password"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsZeroizeFips,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsZeroizeFips) String() string {
    return fmt.Sprintf(
    	"UtilsZeroizeFips[Password=%v, AdditionalProperties=%v]",
    	u.Password, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsZeroizeFips.
// It customizes the JSON marshaling process for UtilsZeroizeFips objects.
func (u UtilsZeroizeFips) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "password"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsZeroizeFips object to a map representation for JSON marshaling.
func (u UtilsZeroizeFips) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["password"] = u.Password
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsZeroizeFips.
// It customizes the JSON unmarshaling process for UtilsZeroizeFips objects.
func (u *UtilsZeroizeFips) UnmarshalJSON(input []byte) error {
    var temp tempUtilsZeroizeFips
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Password = *temp.Password
    return nil
}

// tempUtilsZeroizeFips is a temporary struct used for validating the fields of UtilsZeroizeFips.
type tempUtilsZeroizeFips  struct {
    Password *string `json:"password"`
}

func (u *tempUtilsZeroizeFips) validate() error {
    var errs []string
    if u.Password == nil {
        errs = append(errs, "required field `password` is missing for type `utils_zeroize_fips`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
