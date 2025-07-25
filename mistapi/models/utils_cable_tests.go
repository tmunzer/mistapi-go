// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsCableTests represents a UtilsCableTests struct.
type UtilsCableTests struct {
    // The port to run the cable test
    Port                 string                 `json:"port"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsCableTests,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsCableTests) String() string {
    return fmt.Sprintf(
    	"UtilsCableTests[Port=%v, AdditionalProperties=%v]",
    	u.Port, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsCableTests.
// It customizes the JSON marshaling process for UtilsCableTests objects.
func (u UtilsCableTests) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsCableTests object to a map representation for JSON marshaling.
func (u UtilsCableTests) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["port"] = u.Port
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsCableTests.
// It customizes the JSON unmarshaling process for UtilsCableTests objects.
func (u *UtilsCableTests) UnmarshalJSON(input []byte) error {
    var temp tempUtilsCableTests
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Port = *temp.Port
    return nil
}

// tempUtilsCableTests is a temporary struct used for validating the fields of UtilsCableTests.
type tempUtilsCableTests  struct {
    Port *string `json:"port"`
}

func (u *tempUtilsCableTests) validate() error {
    var errs []string
    if u.Port == nil {
        errs = append(errs, "required field `port` is missing for type `utils_cable_tests`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
