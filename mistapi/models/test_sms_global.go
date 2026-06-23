// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// TestSmsGlobal represents a TestSmsGlobal struct.
// Request body for validating SMSGlobal SMS gateway credentials
type TestSmsGlobal struct {
	// SMSGlobal API key used to send the test SMS
	SmsglobalApiKey string `json:"smsglobal_api_key"`
	// SMSGlobal API secret used to send the test SMS
	SmsglobalApiSecret string `json:"smsglobal_api_secret"`
	// Optional sender's number or sender ID. If not provided, uses the default number associated with the account
	SmsglobalSender *string `json:"smsglobal_sender,omitempty"`
	// Phone number of the recipient of SMS with country code
	To                   string                 `json:"to"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TestSmsGlobal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TestSmsGlobal) String() string {
	return fmt.Sprintf(
		"TestSmsGlobal[SmsglobalApiKey=%v, SmsglobalApiSecret=%v, SmsglobalSender=%v, To=%v, AdditionalProperties=%v]",
		t.SmsglobalApiKey, t.SmsglobalApiSecret, t.SmsglobalSender, t.To, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TestSmsGlobal.
// It customizes the JSON marshaling process for TestSmsGlobal objects.
func (t TestSmsGlobal) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"smsglobal_api_key", "smsglobal_api_secret", "smsglobal_sender", "to"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TestSmsGlobal object to a map representation for JSON marshaling.
func (t TestSmsGlobal) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	structMap["smsglobal_api_key"] = t.SmsglobalApiKey
	structMap["smsglobal_api_secret"] = t.SmsglobalApiSecret
	if t.SmsglobalSender != nil {
		structMap["smsglobal_sender"] = t.SmsglobalSender
	}
	structMap["to"] = t.To
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TestSmsGlobal.
// It customizes the JSON unmarshaling process for TestSmsGlobal objects.
func (t *TestSmsGlobal) UnmarshalJSON(input []byte) error {
	var temp tempTestSmsGlobal
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "smsglobal_api_key", "smsglobal_api_secret", "smsglobal_sender", "to")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.SmsglobalApiKey = *temp.SmsglobalApiKey
	t.SmsglobalApiSecret = *temp.SmsglobalApiSecret
	t.SmsglobalSender = temp.SmsglobalSender
	t.To = *temp.To
	return nil
}

// tempTestSmsGlobal is a temporary struct used for validating the fields of TestSmsGlobal.
type tempTestSmsGlobal struct {
	SmsglobalApiKey    *string `json:"smsglobal_api_key"`
	SmsglobalApiSecret *string `json:"smsglobal_api_secret"`
	SmsglobalSender    *string `json:"smsglobal_sender,omitempty"`
	To                 *string `json:"to"`
}

func (t *tempTestSmsGlobal) validate() error {
	var errs []string
	if t.SmsglobalApiKey == nil {
		errs = append(errs, "required field `smsglobal_api_key` is missing for type `test_sms_global`")
	}
	if t.SmsglobalApiSecret == nil {
		errs = append(errs, "required field `smsglobal_api_secret` is missing for type `test_sms_global`")
	}
	if t.To == nil {
		errs = append(errs, "required field `to` is missing for type `test_sms_global`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
