// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ClientCertSerialNumbers represents a ClientCertSerialNumbers struct.
type ClientCertSerialNumbers struct {
    SerialNumbers        []string               `json:"serial_numbers,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClientCertSerialNumbers,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClientCertSerialNumbers) String() string {
    return fmt.Sprintf(
    	"ClientCertSerialNumbers[SerialNumbers=%v, AdditionalProperties=%v]",
    	c.SerialNumbers, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClientCertSerialNumbers.
// It customizes the JSON marshaling process for ClientCertSerialNumbers objects.
func (c ClientCertSerialNumbers) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "serial_numbers"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ClientCertSerialNumbers object to a map representation for JSON marshaling.
func (c ClientCertSerialNumbers) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.SerialNumbers != nil {
        structMap["serial_numbers"] = c.SerialNumbers
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientCertSerialNumbers.
// It customizes the JSON unmarshaling process for ClientCertSerialNumbers objects.
func (c *ClientCertSerialNumbers) UnmarshalJSON(input []byte) error {
    var temp tempClientCertSerialNumbers
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "serial_numbers")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.SerialNumbers = temp.SerialNumbers
    return nil
}

// tempClientCertSerialNumbers is a temporary struct used for validating the fields of ClientCertSerialNumbers.
type tempClientCertSerialNumbers  struct {
    SerialNumbers []string `json:"serial_numbers,omitempty"`
}
