// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// Marvis represents a Marvis struct.
type Marvis struct {
    AutoOperations       *MarvisAutoOperations  `json:"auto_operations,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Marvis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Marvis) String() string {
    return fmt.Sprintf(
    	"Marvis[AutoOperations=%v, AdditionalProperties=%v]",
    	m.AutoOperations, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Marvis.
// It customizes the JSON marshaling process for Marvis objects.
func (m Marvis) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "auto_operations"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Marvis object to a map representation for JSON marshaling.
func (m Marvis) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AutoOperations != nil {
        structMap["auto_operations"] = m.AutoOperations.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Marvis.
// It customizes the JSON unmarshaling process for Marvis objects.
func (m *Marvis) UnmarshalJSON(input []byte) error {
    var temp tempMarvis
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_operations")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AutoOperations = temp.AutoOperations
    return nil
}

// tempMarvis is a temporary struct used for validating the fields of Marvis.
type tempMarvis  struct {
    AutoOperations *MarvisAutoOperations `json:"auto_operations,omitempty"`
}
