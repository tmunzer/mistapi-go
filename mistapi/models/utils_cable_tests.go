package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsCableTests represents a UtilsCableTests struct.
type UtilsCableTests struct {
    // the port to run the cable test
    Port                 string         `json:"port"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsCableTests.
// It customizes the JSON marshaling process for UtilsCableTests objects.
func (u UtilsCableTests) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsCableTests object to a map representation for JSON marshaling.
func (u UtilsCableTests) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["port"] = u.Port
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsCableTests.
// It customizes the JSON unmarshaling process for UtilsCableTests objects.
func (u *UtilsCableTests) UnmarshalJSON(input []byte) error {
    var temp utilsCableTests
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Port = *temp.Port
    return nil
}

// utilsCableTests is a temporary struct used for validating the fields of UtilsCableTests.
type utilsCableTests  struct {
    Port *string `json:"port"`
}

func (u *utilsCableTests) validate() error {
    var errs []string
    if u.Port == nil {
        errs = append(errs, "required field `port` is missing for type `Utils_Cable_Tests`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
