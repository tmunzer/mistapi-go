// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AoscxRegisterCmd represents a AoscxRegisterCmd struct.
// AOSCX Brownfield Registration Commands
type AoscxRegisterCmd struct {
	// AOSCX-specific CLI commands that can be copied and pasted directly into an AOSCX device to register it with Mist
	CliCommands          *string                `json:"cli_commands,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AoscxRegisterCmd,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AoscxRegisterCmd) String() string {
	return fmt.Sprintf(
		"AoscxRegisterCmd[CliCommands=%v, AdditionalProperties=%v]",
		a.CliCommands, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AoscxRegisterCmd.
// It customizes the JSON marshaling process for AoscxRegisterCmd objects.
func (a AoscxRegisterCmd) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"cli_commands"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AoscxRegisterCmd object to a map representation for JSON marshaling.
func (a AoscxRegisterCmd) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.CliCommands != nil {
		structMap["cli_commands"] = a.CliCommands
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AoscxRegisterCmd.
// It customizes the JSON unmarshaling process for AoscxRegisterCmd objects.
func (a *AoscxRegisterCmd) UnmarshalJSON(input []byte) error {
	var temp tempAoscxRegisterCmd
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cli_commands")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.CliCommands = temp.CliCommands
	return nil
}

// tempAoscxRegisterCmd is a temporary struct used for validating the fields of AoscxRegisterCmd.
type tempAoscxRegisterCmd struct {
	CliCommands *string `json:"cli_commands,omitempty"`
}
