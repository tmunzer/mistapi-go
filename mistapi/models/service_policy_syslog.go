// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicySyslog represents a ServicePolicySyslog struct.
// Required for syslog logging
type ServicePolicySyslog struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	ServerNames          []string               `json:"server_names,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySyslog,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySyslog) String() string {
	return fmt.Sprintf(
		"ServicePolicySyslog[Enabled=%v, ServerNames=%v, AdditionalProperties=%v]",
		s.Enabled, s.ServerNames, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySyslog.
// It customizes the JSON marshaling process for ServicePolicySyslog objects.
func (s ServicePolicySyslog) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled", "server_names"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySyslog object to a map representation for JSON marshaling.
func (s ServicePolicySyslog) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.ServerNames != nil {
		structMap["server_names"] = s.ServerNames
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySyslog.
// It customizes the JSON unmarshaling process for ServicePolicySyslog objects.
func (s *ServicePolicySyslog) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicySyslog
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "server_names")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Enabled = temp.Enabled
	s.ServerNames = temp.ServerNames
	return nil
}

// tempServicePolicySyslog is a temporary struct used for validating the fields of ServicePolicySyslog.
type tempServicePolicySyslog struct {
	Enabled     *bool    `json:"enabled,omitempty"`
	ServerNames []string `json:"server_names,omitempty"`
}
