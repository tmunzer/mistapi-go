package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// TwoFactorString represents a TwoFactorString struct.
type TwoFactorString struct {
	TwoFactor            string         `json:"two_factor"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TwoFactorString.
// It customizes the JSON marshaling process for TwoFactorString objects.
func (t TwoFactorString) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TwoFactorString object to a map representation for JSON marshaling.
func (t TwoFactorString) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, t.AdditionalProperties)
	structMap["two_factor"] = t.TwoFactor
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TwoFactorString.
// It customizes the JSON unmarshaling process for TwoFactorString objects.
func (t *TwoFactorString) UnmarshalJSON(input []byte) error {
	var temp tempTwoFactorString
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "two_factor")
	if err != nil {
		return err
	}

	t.AdditionalProperties = additionalProperties
	t.TwoFactor = *temp.TwoFactor
	return nil
}

// tempTwoFactorString is a temporary struct used for validating the fields of TwoFactorString.
type tempTwoFactorString struct {
	TwoFactor *string `json:"two_factor"`
}

func (t *tempTwoFactorString) validate() error {
	var errs []string
	if t.TwoFactor == nil {
		errs = append(errs, "required field `two_factor` is missing for type `two_factor_string`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
