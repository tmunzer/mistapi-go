// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// UsernamePassword represents a UsernamePassword struct.
type UsernamePassword struct {
    Password             *string                `json:"password,omitempty"`
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UsernamePassword,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UsernamePassword) String() string {
    return fmt.Sprintf(
    	"UsernamePassword[Password=%v, Username=%v, AdditionalProperties=%v]",
    	u.Password, u.Username, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UsernamePassword.
// It customizes the JSON marshaling process for UsernamePassword objects.
func (u UsernamePassword) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "password", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UsernamePassword object to a map representation for JSON marshaling.
func (u UsernamePassword) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Password != nil {
        structMap["password"] = u.Password
    }
    if u.Username != nil {
        structMap["username"] = u.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsernamePassword.
// It customizes the JSON unmarshaling process for UsernamePassword objects.
func (u *UsernamePassword) UnmarshalJSON(input []byte) error {
    var temp tempUsernamePassword
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password", "username")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Password = temp.Password
    u.Username = temp.Username
    return nil
}

// tempUsernamePassword is a temporary struct used for validating the fields of UsernamePassword.
type tempUsernamePassword  struct {
    Password *string `json:"password,omitempty"`
    Username *string `json:"username,omitempty"`
}
