package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RootPasswordString represents a RootPasswordString struct.
type RootPasswordString struct {
    RootPassword         string         `json:"root_password"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RootPasswordString.
// It customizes the JSON marshaling process for RootPasswordString objects.
func (r RootPasswordString) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RootPasswordString object to a map representation for JSON marshaling.
func (r RootPasswordString) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["root_password"] = r.RootPassword
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RootPasswordString.
// It customizes the JSON unmarshaling process for RootPasswordString objects.
func (r *RootPasswordString) UnmarshalJSON(input []byte) error {
    var temp tempRootPasswordString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "root_password")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.RootPassword = *temp.RootPassword
    return nil
}

// tempRootPasswordString is a temporary struct used for validating the fields of RootPasswordString.
type tempRootPasswordString  struct {
    RootPassword *string `json:"root_password"`
}

func (r *tempRootPasswordString) validate() error {
    var errs []string
    if r.RootPassword == nil {
        errs = append(errs, "required field `root_password` is missing for type `root_password_string`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
