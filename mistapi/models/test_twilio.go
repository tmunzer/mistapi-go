// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TestTwilio represents a TestTwilio struct.
type TestTwilio struct {
    // One of the numbers you have in your Twilio account
    From                 string                 `json:"from"`
    // Phone number of the recipient of SMS
    To                   string                 `json:"to"`
    // Auth Token associated with twilio account
    TwilioAuthToken      string                 `json:"twilio_auth_token"`
    // Twilio Account SID
    TwilioSid            string                 `json:"twilio_sid"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TestTwilio,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TestTwilio) String() string {
    return fmt.Sprintf(
    	"TestTwilio[From=%v, To=%v, TwilioAuthToken=%v, TwilioSid=%v, AdditionalProperties=%v]",
    	t.From, t.To, t.TwilioAuthToken, t.TwilioSid, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TestTwilio.
// It customizes the JSON marshaling process for TestTwilio objects.
func (t TestTwilio) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "from", "to", "twilio_auth_token", "twilio_sid"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TestTwilio object to a map representation for JSON marshaling.
func (t TestTwilio) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    structMap["from"] = t.From
    structMap["to"] = t.To
    structMap["twilio_auth_token"] = t.TwilioAuthToken
    structMap["twilio_sid"] = t.TwilioSid
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TestTwilio.
// It customizes the JSON unmarshaling process for TestTwilio objects.
func (t *TestTwilio) UnmarshalJSON(input []byte) error {
    var temp tempTestTwilio
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "from", "to", "twilio_auth_token", "twilio_sid")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.From = *temp.From
    t.To = *temp.To
    t.TwilioAuthToken = *temp.TwilioAuthToken
    t.TwilioSid = *temp.TwilioSid
    return nil
}

// tempTestTwilio is a temporary struct used for validating the fields of TestTwilio.
type tempTestTwilio  struct {
    From            *string `json:"from"`
    To              *string `json:"to"`
    TwilioAuthToken *string `json:"twilio_auth_token"`
    TwilioSid       *string `json:"twilio_sid"`
}

func (t *tempTestTwilio) validate() error {
    var errs []string
    if t.From == nil {
        errs = append(errs, "required field `from` is missing for type `test_twilio`")
    }
    if t.To == nil {
        errs = append(errs, "required field `to` is missing for type `test_twilio`")
    }
    if t.TwilioAuthToken == nil {
        errs = append(errs, "required field `twilio_auth_token` is missing for type `test_twilio`")
    }
    if t.TwilioSid == nil {
        errs = append(errs, "required field `twilio_sid` is missing for type `test_twilio`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
