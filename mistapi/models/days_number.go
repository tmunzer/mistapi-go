// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// DaysNumber represents a DaysNumber struct.
type DaysNumber struct {
    Days                 *int                   `json:"days,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DaysNumber,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DaysNumber) String() string {
    return fmt.Sprintf(
    	"DaysNumber[Days=%v, AdditionalProperties=%v]",
    	d.Days, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DaysNumber.
// It customizes the JSON marshaling process for DaysNumber objects.
func (d DaysNumber) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "days"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DaysNumber object to a map representation for JSON marshaling.
func (d DaysNumber) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Days != nil {
        structMap["days"] = d.Days
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DaysNumber.
// It customizes the JSON unmarshaling process for DaysNumber objects.
func (d *DaysNumber) UnmarshalJSON(input []byte) error {
    var temp tempDaysNumber
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "days")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Days = temp.Days
    return nil
}

// tempDaysNumber is a temporary struct used for validating the fields of DaysNumber.
type tempDaysNumber  struct {
    Days *int `json:"days,omitempty"`
}
