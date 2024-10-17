package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CodeString represents a CodeString struct.
type CodeString struct {
	Code                 string         `json:"code"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CodeString.
// It customizes the JSON marshaling process for CodeString objects.
func (c CodeString) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CodeString object to a map representation for JSON marshaling.
func (c CodeString) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["code"] = c.Code
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CodeString.
// It customizes the JSON unmarshaling process for CodeString objects.
func (c *CodeString) UnmarshalJSON(input []byte) error {
	var temp tempCodeString
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "code")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Code = *temp.Code
	return nil
}

// tempCodeString is a temporary struct used for validating the fields of CodeString.
type tempCodeString struct {
	Code *string `json:"code"`
}

func (c *tempCodeString) validate() error {
	var errs []string
	if c.Code == nil {
		errs = append(errs, "required field `code` is missing for type `code_string`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
