package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsZeroiseFips represents a UtilsZeroiseFips struct.
type UtilsZeroiseFips struct {
    // FIPS zeroize password
    Password             string                 `json:"password"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsZeroiseFips,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsZeroiseFips) String() string {
    return fmt.Sprintf(
    	"UtilsZeroiseFips[Password=%v, AdditionalProperties=%v]",
    	u.Password, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsZeroiseFips.
// It customizes the JSON marshaling process for UtilsZeroiseFips objects.
func (u UtilsZeroiseFips) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "password"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsZeroiseFips object to a map representation for JSON marshaling.
func (u UtilsZeroiseFips) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["password"] = u.Password
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsZeroiseFips.
// It customizes the JSON unmarshaling process for UtilsZeroiseFips objects.
func (u *UtilsZeroiseFips) UnmarshalJSON(input []byte) error {
    var temp tempUtilsZeroiseFips
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Password = *temp.Password
    return nil
}

// tempUtilsZeroiseFips is a temporary struct used for validating the fields of UtilsZeroiseFips.
type tempUtilsZeroiseFips  struct {
    Password *string `json:"password"`
}

func (u *tempUtilsZeroiseFips) validate() error {
    var errs []string
    if u.Password == nil {
        errs = append(errs, "required field `password` is missing for type `utils_zeroise_fips`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
