package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstApLed represents a ConstApLed struct.
type ConstApLed struct {
    Code                 string         `json:"code"`
    Description          string         `json:"description"`
    Key                  string         `json:"key"`
    Name                 string         `json:"name"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstApLed.
// It customizes the JSON marshaling process for ConstApLed objects.
func (c ConstApLed) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstApLed object to a map representation for JSON marshaling.
func (c ConstApLed) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["code"] = c.Code
    structMap["description"] = c.Description
    structMap["key"] = c.Key
    structMap["name"] = c.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstApLed.
// It customizes the JSON unmarshaling process for ConstApLed objects.
func (c *ConstApLed) UnmarshalJSON(input []byte) error {
    var temp tempConstApLed
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "code", "description", "key", "name")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Code = *temp.Code
    c.Description = *temp.Description
    c.Key = *temp.Key
    c.Name = *temp.Name
    return nil
}

// tempConstApLed is a temporary struct used for validating the fields of ConstApLed.
type tempConstApLed  struct {
    Code        *string `json:"code"`
    Description *string `json:"description"`
    Key         *string `json:"key"`
    Name        *string `json:"name"`
}

func (c *tempConstApLed) validate() error {
    var errs []string
    if c.Code == nil {
        errs = append(errs, "required field `code` is missing for type `const_ap_led`")
    }
    if c.Description == nil {
        errs = append(errs, "required field `description` is missing for type `const_ap_led`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `const_ap_led`")
    }
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `const_ap_led`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
