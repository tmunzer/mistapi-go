// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MacAddresses represents a MacAddresses struct.
type MacAddresses struct {
    Macs                 []string               `json:"macs"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MacAddresses,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MacAddresses) String() string {
    return fmt.Sprintf(
    	"MacAddresses[Macs=%v, AdditionalProperties=%v]",
    	m.Macs, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MacAddresses.
// It customizes the JSON marshaling process for MacAddresses objects.
func (m MacAddresses) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "macs"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MacAddresses object to a map representation for JSON marshaling.
func (m MacAddresses) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["macs"] = m.Macs
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MacAddresses.
// It customizes the JSON unmarshaling process for MacAddresses objects.
func (m *MacAddresses) UnmarshalJSON(input []byte) error {
    var temp tempMacAddresses
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "macs")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Macs = *temp.Macs
    return nil
}

// tempMacAddresses is a temporary struct used for validating the fields of MacAddresses.
type tempMacAddresses  struct {
    Macs *[]string `json:"macs"`
}

func (m *tempMacAddresses) validate() error {
    var errs []string
    if m.Macs == nil {
        errs = append(errs, "required field `macs` is missing for type `mac_addresses`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
