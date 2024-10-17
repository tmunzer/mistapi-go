package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ConstLanguage represents a ConstLanguage struct.
type ConstLanguage struct {
	Display              string         `json:"display"`
	DisplayNative        string         `json:"display_native"`
	Key                  string         `json:"key"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstLanguage.
// It customizes the JSON marshaling process for ConstLanguage objects.
func (c ConstLanguage) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstLanguage object to a map representation for JSON marshaling.
func (c ConstLanguage) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["display"] = c.Display
	structMap["display_native"] = c.DisplayNative
	structMap["key"] = c.Key
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstLanguage.
// It customizes the JSON unmarshaling process for ConstLanguage objects.
func (c *ConstLanguage) UnmarshalJSON(input []byte) error {
	var temp tempConstLanguage
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "display", "display_native", "key")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Display = *temp.Display
	c.DisplayNative = *temp.DisplayNative
	c.Key = *temp.Key
	return nil
}

// tempConstLanguage is a temporary struct used for validating the fields of ConstLanguage.
type tempConstLanguage struct {
	Display       *string `json:"display"`
	DisplayNative *string `json:"display_native"`
	Key           *string `json:"key"`
}

func (c *tempConstLanguage) validate() error {
	var errs []string
	if c.Display == nil {
		errs = append(errs, "required field `display` is missing for type `const_language`")
	}
	if c.DisplayNative == nil {
		errs = append(errs, "required field `display_native` is missing for type `const_language`")
	}
	if c.Key == nil {
		errs = append(errs, "required field `key` is missing for type `const_language`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
