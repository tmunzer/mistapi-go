package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstApLed represents a ConstApLed struct.
type ConstApLed struct {
    Code                 string                 `json:"code"`
    Description          string                 `json:"description"`
    Key                  string                 `json:"key"`
    Name                 string                 `json:"name"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstApLed,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstApLed) String() string {
    return fmt.Sprintf(
    	"ConstApLed[Code=%v, Description=%v, Key=%v, Name=%v, AdditionalProperties=%v]",
    	c.Code, c.Description, c.Key, c.Name, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstApLed.
// It customizes the JSON marshaling process for ConstApLed objects.
func (c ConstApLed) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "code", "description", "key", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstApLed object to a map representation for JSON marshaling.
func (c ConstApLed) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "description", "key", "name")
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
    return errors.New(strings.Join (errs, "\n"))
}
