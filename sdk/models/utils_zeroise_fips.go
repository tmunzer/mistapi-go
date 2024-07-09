package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsZeroiseFips represents a UtilsZeroiseFips struct.
type UtilsZeroiseFips struct {
    // FIPS zeroize password
    Password             string         `json:"password"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsZeroiseFips.
// It customizes the JSON marshaling process for UtilsZeroiseFips objects.
func (u UtilsZeroiseFips) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsZeroiseFips object to a map representation for JSON marshaling.
func (u UtilsZeroiseFips) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["password"] = u.Password
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsZeroiseFips.
// It customizes the JSON unmarshaling process for UtilsZeroiseFips objects.
func (u *UtilsZeroiseFips) UnmarshalJSON(input []byte) error {
    var temp utilsZeroiseFips
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "password")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Password = *temp.Password
    return nil
}

// utilsZeroiseFips is a temporary struct used for validating the fields of UtilsZeroiseFips.
type utilsZeroiseFips  struct {
    Password *string `json:"password"`
}

func (u *utilsZeroiseFips) validate() error {
    var errs []string
    if u.Password == nil {
        errs = append(errs, "required field `password` is missing for type `Utils_Zeroise_Fips`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
