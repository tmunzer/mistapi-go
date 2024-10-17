package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// TwoFactorCode represents a TwoFactorCode struct.
type TwoFactorCode struct {
	TwoFactor            string         `json:"two_factor"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TwoFactorCode.
// It customizes the JSON marshaling process for TwoFactorCode objects.
func (t TwoFactorCode) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TwoFactorCode object to a map representation for JSON marshaling.
func (t TwoFactorCode) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, t.AdditionalProperties)
	structMap["two_factor"] = t.TwoFactor
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TwoFactorCode.
// It customizes the JSON unmarshaling process for TwoFactorCode objects.
func (t *TwoFactorCode) UnmarshalJSON(input []byte) error {
	var temp tempTwoFactorCode
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

// tempTwoFactorCode is a temporary struct used for validating the fields of TwoFactorCode.
type tempTwoFactorCode struct {
	TwoFactor *string `json:"two_factor"`
}

func (t *tempTwoFactorCode) validate() error {
	var errs []string
	if t.TwoFactor == nil {
		errs = append(errs, "required field `two_factor` is missing for type `two_factor_code`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
