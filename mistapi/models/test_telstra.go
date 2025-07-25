// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TestTelstra represents a TestTelstra struct.
type TestTelstra struct {
    // Telstra client id
    TelstraClientId      string                 `json:"telstra_client_id"`
    // Telstra client secret
    TelstraClientSecret  string                 `json:"telstra_client_secret"`
    // Phone number of the recipient of SMS with country code
    To                   string                 `json:"to"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TestTelstra,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TestTelstra) String() string {
    return fmt.Sprintf(
    	"TestTelstra[TelstraClientId=%v, TelstraClientSecret=%v, To=%v, AdditionalProperties=%v]",
    	t.TelstraClientId, t.TelstraClientSecret, t.To, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TestTelstra.
// It customizes the JSON marshaling process for TestTelstra objects.
func (t TestTelstra) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "telstra_client_id", "telstra_client_secret", "to"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TestTelstra object to a map representation for JSON marshaling.
func (t TestTelstra) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    structMap["telstra_client_id"] = t.TelstraClientId
    structMap["telstra_client_secret"] = t.TelstraClientSecret
    structMap["to"] = t.To
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TestTelstra.
// It customizes the JSON unmarshaling process for TestTelstra objects.
func (t *TestTelstra) UnmarshalJSON(input []byte) error {
    var temp tempTestTelstra
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "telstra_client_id", "telstra_client_secret", "to")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.TelstraClientId = *temp.TelstraClientId
    t.TelstraClientSecret = *temp.TelstraClientSecret
    t.To = *temp.To
    return nil
}

// tempTestTelstra is a temporary struct used for validating the fields of TestTelstra.
type tempTestTelstra  struct {
    TelstraClientId     *string `json:"telstra_client_id"`
    TelstraClientSecret *string `json:"telstra_client_secret"`
    To                  *string `json:"to"`
}

func (t *tempTestTelstra) validate() error {
    var errs []string
    if t.TelstraClientId == nil {
        errs = append(errs, "required field `telstra_client_id` is missing for type `test_telstra`")
    }
    if t.TelstraClientSecret == nil {
        errs = append(errs, "required field `telstra_client_secret` is missing for type `test_telstra`")
    }
    if t.To == nil {
        errs = append(errs, "required field `to` is missing for type `test_telstra`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
