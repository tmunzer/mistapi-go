package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstEvent represents a ConstEvent struct.
type ConstEvent struct {
    Description          *string        `json:"description,omitempty"`
    Display              string         `json:"display"`
    Example              *interface{}   `json:"example,omitempty"`
    Key                  string         `json:"key"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstEvent.
// It customizes the JSON marshaling process for ConstEvent objects.
func (c ConstEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstEvent object to a map representation for JSON marshaling.
func (c ConstEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    structMap["display"] = c.Display
    if c.Example != nil {
        structMap["example"] = c.Example
    }
    structMap["key"] = c.Key
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstEvent.
// It customizes the JSON unmarshaling process for ConstEvent objects.
func (c *ConstEvent) UnmarshalJSON(input []byte) error {
    var temp tempConstEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "description", "display", "example", "key")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Description = temp.Description
    c.Display = *temp.Display
    c.Example = temp.Example
    c.Key = *temp.Key
    return nil
}

// tempConstEvent is a temporary struct used for validating the fields of ConstEvent.
type tempConstEvent  struct {
    Description *string      `json:"description,omitempty"`
    Display     *string      `json:"display"`
    Example     *interface{} `json:"example,omitempty"`
    Key         *string      `json:"key"`
}

func (c *tempConstEvent) validate() error {
    var errs []string
    if c.Display == nil {
        errs = append(errs, "required field `display` is missing for type `const_event`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `const_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
