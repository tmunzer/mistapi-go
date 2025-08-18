// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SynthetictestRadiusServer represents a SynthetictestRadiusServer struct.
type SynthetictestRadiusServer struct {
	// Specify the password associated with the username
	Password string `json:"password"`
	// Specify the access profile associated with the subscriber
	Profile *string `json:"profile,omitempty"`
	// Specify the subscriber username to test
	User                 string                 `json:"user"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestRadiusServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestRadiusServer) String() string {
	return fmt.Sprintf(
		"SynthetictestRadiusServer[Password=%v, Profile=%v, User=%v, AdditionalProperties=%v]",
		s.Password, s.Profile, s.User, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestRadiusServer.
// It customizes the JSON marshaling process for SynthetictestRadiusServer objects.
func (s SynthetictestRadiusServer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"password", "profile", "user"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestRadiusServer object to a map representation for JSON marshaling.
func (s SynthetictestRadiusServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["password"] = s.Password
	if s.Profile != nil {
		structMap["profile"] = s.Profile
	}
	structMap["user"] = s.User
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestRadiusServer.
// It customizes the JSON unmarshaling process for SynthetictestRadiusServer objects.
func (s *SynthetictestRadiusServer) UnmarshalJSON(input []byte) error {
	var temp tempSynthetictestRadiusServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password", "profile", "user")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Password = *temp.Password
	s.Profile = temp.Profile
	s.User = *temp.User
	return nil
}

// tempSynthetictestRadiusServer is a temporary struct used for validating the fields of SynthetictestRadiusServer.
type tempSynthetictestRadiusServer struct {
	Password *string `json:"password"`
	Profile  *string `json:"profile,omitempty"`
	User     *string `json:"user"`
}

func (s *tempSynthetictestRadiusServer) validate() error {
	var errs []string
	if s.Password == nil {
		errs = append(errs, "required field `password` is missing for type `synthetictest_radius_server`")
	}
	if s.User == nil {
		errs = append(errs, "required field `user` is missing for type `synthetictest_radius_server`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
