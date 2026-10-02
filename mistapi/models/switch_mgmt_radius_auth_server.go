// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// SwitchMgmtRadiusAuthServer represents a SwitchMgmtRadiusAuthServer struct.
// RADIUS authentication server used for switch management authentication
type SwitchMgmtRadiusAuthServer struct {
	// Address or hostname of the RADIUS authentication server
	Host string `json:"host"`
	// Unique identifier for this RADIUS authentication server entry
	Id *uuid.UUID `json:"id,omitempty"`
	// RADIUS Auth Port, value from 1 to 65535, default is 1812
	Port *RadiusAuthPort `json:"port,omitempty"`
	// Shared secret used with this RADIUS authentication server
	Secret               string                 `json:"secret"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMgmtRadiusAuthServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMgmtRadiusAuthServer) String() string {
	return fmt.Sprintf(
		"SwitchMgmtRadiusAuthServer[Host=%v, Id=%v, Port=%v, Secret=%v, AdditionalProperties=%v]",
		s.Host, s.Id, s.Port, s.Secret, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMgmtRadiusAuthServer.
// It customizes the JSON marshaling process for SwitchMgmtRadiusAuthServer objects.
func (s SwitchMgmtRadiusAuthServer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"host", "id", "port", "secret"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchMgmtRadiusAuthServer object to a map representation for JSON marshaling.
func (s SwitchMgmtRadiusAuthServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["host"] = s.Host
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.Port != nil {
		structMap["port"] = s.Port.toMap()
	}
	structMap["secret"] = s.Secret
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMgmtRadiusAuthServer.
// It customizes the JSON unmarshaling process for SwitchMgmtRadiusAuthServer objects.
func (s *SwitchMgmtRadiusAuthServer) UnmarshalJSON(input []byte) error {
	var temp tempSwitchMgmtRadiusAuthServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "id", "port", "secret")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Host = *temp.Host
	s.Id = temp.Id
	s.Port = temp.Port
	s.Secret = *temp.Secret
	return nil
}

// tempSwitchMgmtRadiusAuthServer is a temporary struct used for validating the fields of SwitchMgmtRadiusAuthServer.
type tempSwitchMgmtRadiusAuthServer struct {
	Host   *string         `json:"host"`
	Id     *uuid.UUID      `json:"id,omitempty"`
	Port   *RadiusAuthPort `json:"port,omitempty"`
	Secret *string         `json:"secret"`
}

func (s *tempSwitchMgmtRadiusAuthServer) validate() error {
	var errs []string
	if s.Host == nil {
		errs = append(errs, "required field `host` is missing for type `switch_mgmt_radius_auth_server`")
	}
	if s.Secret == nil {
		errs = append(errs, "required field `secret` is missing for type `switch_mgmt_radius_auth_server`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
