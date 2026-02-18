// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AosRegisterCmd represents a AosRegisterCmd struct.
// AOS Brownfield Registration Commands
type AosRegisterCmd struct {
	// AOS-specific CLI commands that can be copied and pasted directly into an AOS device to register it with Mist. Includes registration code and configuration commands.
	CliCommands          *string                `json:"cli_commands,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AosRegisterCmd,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AosRegisterCmd) String() string {
	return fmt.Sprintf(
		"AosRegisterCmd[CliCommands=%v, AdditionalProperties=%v]",
		a.CliCommands, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AosRegisterCmd.
// It customizes the JSON marshaling process for AosRegisterCmd objects.
func (a AosRegisterCmd) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"cli_commands"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AosRegisterCmd object to a map representation for JSON marshaling.
func (a AosRegisterCmd) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.CliCommands != nil {
		structMap["cli_commands"] = a.CliCommands
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AosRegisterCmd.
// It customizes the JSON unmarshaling process for AosRegisterCmd objects.
func (a *AosRegisterCmd) UnmarshalJSON(input []byte) error {
	var temp tempAosRegisterCmd
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

// tempAosRegisterCmd is a temporary struct used for validating the fields of AosRegisterCmd.
type tempAosRegisterCmd struct {
	CliCommands *string `json:"cli_commands,omitempty"`
}
