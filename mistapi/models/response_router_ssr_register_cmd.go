// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseRouterSsrRegisterCmd represents a ResponseRouterSsrRegisterCmd struct.
type ResponseRouterSsrRegisterCmd struct {
	ConductorCmd         *string                `json:"conductor_cmd,omitempty"`
	RegistrationCode     *string                `json:"registration_code,omitempty"`
	RouterShellCmd       *string                `json:"router_shell_cmd,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseRouterSsrRegisterCmd,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseRouterSsrRegisterCmd) String() string {
	return fmt.Sprintf(
		"ResponseRouterSsrRegisterCmd[ConductorCmd=%v, RegistrationCode=%v, RouterShellCmd=%v, AdditionalProperties=%v]",
		r.ConductorCmd, r.RegistrationCode, r.RouterShellCmd, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseRouterSsrRegisterCmd.
// It customizes the JSON marshaling process for ResponseRouterSsrRegisterCmd objects.
func (r ResponseRouterSsrRegisterCmd) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"conductor_cmd", "registration_code", "router_shell_cmd"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseRouterSsrRegisterCmd object to a map representation for JSON marshaling.
func (r ResponseRouterSsrRegisterCmd) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ConductorCmd != nil {
		structMap["conductor_cmd"] = r.ConductorCmd
	}
	if r.RegistrationCode != nil {
		structMap["registration_code"] = r.RegistrationCode
	}
	if r.RouterShellCmd != nil {
		structMap["router_shell_cmd"] = r.RouterShellCmd
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseRouterSsrRegisterCmd.
// It customizes the JSON unmarshaling process for ResponseRouterSsrRegisterCmd objects.
func (r *ResponseRouterSsrRegisterCmd) UnmarshalJSON(input []byte) error {
	var temp tempResponseRouterSsrRegisterCmd
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "conductor_cmd", "registration_code", "router_shell_cmd")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.ConductorCmd = temp.ConductorCmd
	r.RegistrationCode = temp.RegistrationCode
	r.RouterShellCmd = temp.RouterShellCmd
	return nil
}

// tempResponseRouterSsrRegisterCmd is a temporary struct used for validating the fields of ResponseRouterSsrRegisterCmd.
type tempResponseRouterSsrRegisterCmd struct {
	ConductorCmd     *string `json:"conductor_cmd,omitempty"`
	RegistrationCode *string `json:"registration_code,omitempty"`
	RouterShellCmd   *string `json:"router_shell_cmd,omitempty"`
}
