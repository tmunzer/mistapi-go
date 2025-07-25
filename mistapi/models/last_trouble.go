// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// LastTrouble represents a LastTrouble struct.
// Last trouble code of switch
type LastTrouble struct {
    // Code definitions list at [List Ap Led Definition]($e/Constants%20Definitions/listApLedDefinition)
    Code                 *string                `json:"code,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LastTrouble,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LastTrouble) String() string {
    return fmt.Sprintf(
    	"LastTrouble[Code=%v, Timestamp=%v, AdditionalProperties=%v]",
    	l.Code, l.Timestamp, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LastTrouble.
// It customizes the JSON marshaling process for LastTrouble objects.
func (l LastTrouble) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "code", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LastTrouble object to a map representation for JSON marshaling.
func (l LastTrouble) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Code != nil {
        structMap["code"] = l.Code
    }
    if l.Timestamp != nil {
        structMap["timestamp"] = l.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LastTrouble.
// It customizes the JSON unmarshaling process for LastTrouble objects.
func (l *LastTrouble) UnmarshalJSON(input []byte) error {
    var temp tempLastTrouble
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "timestamp")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Code = temp.Code
    l.Timestamp = temp.Timestamp
    return nil
}

// tempLastTrouble is a temporary struct used for validating the fields of LastTrouble.
type tempLastTrouble  struct {
    Code      *string  `json:"code,omitempty"`
    Timestamp *float64 `json:"timestamp,omitempty"`
}
