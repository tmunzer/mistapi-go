// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// JuniperAccount represents a JuniperAccount struct.
type JuniperAccount struct {
    LinkedBy             *string                `json:"linked_by,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JuniperAccount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JuniperAccount) String() string {
    return fmt.Sprintf(
    	"JuniperAccount[LinkedBy=%v, Name=%v, AdditionalProperties=%v]",
    	j.LinkedBy, j.Name, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JuniperAccount.
// It customizes the JSON marshaling process for JuniperAccount objects.
func (j JuniperAccount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "linked_by", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JuniperAccount object to a map representation for JSON marshaling.
func (j JuniperAccount) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.LinkedBy != nil {
        structMap["linked_by"] = j.LinkedBy
    }
    if j.Name != nil {
        structMap["name"] = j.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JuniperAccount.
// It customizes the JSON unmarshaling process for JuniperAccount objects.
func (j *JuniperAccount) UnmarshalJSON(input []byte) error {
    var temp tempJuniperAccount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "linked_by", "name")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.LinkedBy = temp.LinkedBy
    j.Name = temp.Name
    return nil
}

// tempJuniperAccount is a temporary struct used for validating the fields of JuniperAccount.
type tempJuniperAccount  struct {
    LinkedBy *string `json:"linked_by,omitempty"`
    Name     *string `json:"name,omitempty"`
}
