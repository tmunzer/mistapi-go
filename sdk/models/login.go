package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// Login represents a Login struct.
type Login struct {
    Email                string         `json:"email"`
    Password             string         `json:"password"`
    TwoFactor            *string        `json:"two_factor,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Login.
// It customizes the JSON marshaling process for Login objects.
func (l Login) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the Login object to a map representation for JSON marshaling.
func (l Login) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["email"] = l.Email
    structMap["password"] = l.Password
    if l.TwoFactor != nil {
        structMap["two_factor"] = l.TwoFactor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Login.
// It customizes the JSON unmarshaling process for Login objects.
func (l *Login) UnmarshalJSON(input []byte) error {
    var temp login
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "email", "password", "two_factor")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Email = *temp.Email
    l.Password = *temp.Password
    l.TwoFactor = temp.TwoFactor
    return nil
}

// login is a temporary struct used for validating the fields of Login.
type login  struct {
    Email     *string `json:"email"`
    Password  *string `json:"password"`
    TwoFactor *string `json:"two_factor,omitempty"`
}

func (l *login) validate() error {
    var errs []string
    if l.Email == nil {
        errs = append(errs, "required field `email` is missing for type `Login`")
    }
    if l.Password == nil {
        errs = append(errs, "required field `password` is missing for type `Login`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
