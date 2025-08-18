// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// TwoFactorCode represents a TwoFactorCode struct.
type TwoFactorCode struct {
	TwoFactor            string                 `json:"two_factor"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TwoFactorCode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TwoFactorCode) String() string {
	return fmt.Sprintf(
		"TwoFactorCode[TwoFactor=%v, AdditionalProperties=%v]",
		t.TwoFactor, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TwoFactorCode.
// It customizes the JSON marshaling process for TwoFactorCode objects.
func (t TwoFactorCode) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"two_factor"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TwoFactorCode object to a map representation for JSON marshaling.
func (t TwoFactorCode) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "two_factor")
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
