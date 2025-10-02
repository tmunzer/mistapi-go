// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SsoMxedgeProxyAuthServer represents a SsoMxedgeProxyAuthServer struct.
type SsoMxedgeProxyAuthServer struct {
	Host *string `json:"host,omitempty"`
	Port *int    `json:"port,omitempty"`
	// Whether to require Message-Authenticator in requests
	RequireMessageAuthenticator *bool `json:"require_message_authenticator,omitempty"`
	// Authentication request retry
	Retry  *int    `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	// Authentication request timeout, in seconds
	Timeout              *int                   `json:"timeout,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoMxedgeProxyAuthServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoMxedgeProxyAuthServer) String() string {
	return fmt.Sprintf(
		"SsoMxedgeProxyAuthServer[Host=%v, Port=%v, RequireMessageAuthenticator=%v, Retry=%v, Secret=%v, Timeout=%v, AdditionalProperties=%v]",
		s.Host, s.Port, s.RequireMessageAuthenticator, s.Retry, s.Secret, s.Timeout, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoMxedgeProxyAuthServer.
// It customizes the JSON marshaling process for SsoMxedgeProxyAuthServer objects.
func (s SsoMxedgeProxyAuthServer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"host", "port", "require_message_authenticator", "retry", "secret", "timeout"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsoMxedgeProxyAuthServer object to a map representation for JSON marshaling.
func (s SsoMxedgeProxyAuthServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Host != nil {
		structMap["host"] = s.Host
	}
	if s.Port != nil {
		structMap["port"] = s.Port
	}
	if s.RequireMessageAuthenticator != nil {
		structMap["require_message_authenticator"] = s.RequireMessageAuthenticator
	}
	if s.Retry != nil {
		structMap["retry"] = s.Retry
	}
	if s.Secret != nil {
		structMap["secret"] = s.Secret
	}
	if s.Timeout != nil {
		structMap["timeout"] = s.Timeout
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoMxedgeProxyAuthServer.
// It customizes the JSON unmarshaling process for SsoMxedgeProxyAuthServer objects.
func (s *SsoMxedgeProxyAuthServer) UnmarshalJSON(input []byte) error {
	var temp tempSsoMxedgeProxyAuthServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "port", "require_message_authenticator", "retry", "secret", "timeout")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Host = temp.Host
	s.Port = temp.Port
	s.RequireMessageAuthenticator = temp.RequireMessageAuthenticator
	s.Retry = temp.Retry
	s.Secret = temp.Secret
	s.Timeout = temp.Timeout
	return nil
}

// tempSsoMxedgeProxyAuthServer is a temporary struct used for validating the fields of SsoMxedgeProxyAuthServer.
type tempSsoMxedgeProxyAuthServer struct {
	Host                        *string `json:"host,omitempty"`
	Port                        *int    `json:"port,omitempty"`
	RequireMessageAuthenticator *bool   `json:"require_message_authenticator,omitempty"`
	Retry                       *int    `json:"retry,omitempty"`
	Secret                      *string `json:"secret,omitempty"`
	Timeout                     *int    `json:"timeout,omitempty"`
}
