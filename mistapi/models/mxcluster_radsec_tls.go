// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxclusterRadsecTls represents a MxclusterRadsecTls struct.
type MxclusterRadsecTls struct {
    Keypair              *string                `json:"keypair,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterRadsecTls,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterRadsecTls) String() string {
    return fmt.Sprintf(
    	"MxclusterRadsecTls[Keypair=%v, AdditionalProperties=%v]",
    	m.Keypair, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsecTls.
// It customizes the JSON marshaling process for MxclusterRadsecTls objects.
func (m MxclusterRadsecTls) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "keypair"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsecTls object to a map representation for JSON marshaling.
func (m MxclusterRadsecTls) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Keypair != nil {
        structMap["keypair"] = m.Keypair
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsecTls.
// It customizes the JSON unmarshaling process for MxclusterRadsecTls objects.
func (m *MxclusterRadsecTls) UnmarshalJSON(input []byte) error {
    var temp tempMxclusterRadsecTls
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "keypair")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Keypair = temp.Keypair
    return nil
}

// tempMxclusterRadsecTls is a temporary struct used for validating the fields of MxclusterRadsecTls.
type tempMxclusterRadsecTls  struct {
    Keypair *string `json:"keypair,omitempty"`
}
