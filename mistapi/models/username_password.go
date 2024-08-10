package models

import (
    "encoding/json"
)

// UsernamePassword represents a UsernamePassword struct.
type UsernamePassword struct {
    Password             *string        `json:"password,omitempty"`
    Username             *string        `json:"username,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UsernamePassword.
// It customizes the JSON marshaling process for UsernamePassword objects.
func (u UsernamePassword) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UsernamePassword object to a map representation for JSON marshaling.
func (u UsernamePassword) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "password", "username")
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
