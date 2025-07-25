// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// LoginFailures represents a LoginFailures struct.
// Failed login attempts
type LoginFailures struct {
    // Email address of the user
    Email                *string                `json:"email,omitempty"`
    // Last failure time
    LastFailureAt        *int                   `json:"last_failure_at,omitempty"`
    // Number of failed login attempts
    NumAttempts          *int                   `json:"num_attempts,omitempty"`
    // List of up to 32 unique source IP addresses, ordered with the most recent first
    ScrIps               []string               `json:"scr_ips,omitempty"`
    // List of up to 32 unique User-Agent strings, ordered with the most recent first
    UserAgents           []string               `json:"user_agents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LoginFailures,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LoginFailures) String() string {
    return fmt.Sprintf(
    	"LoginFailures[Email=%v, LastFailureAt=%v, NumAttempts=%v, ScrIps=%v, UserAgents=%v, AdditionalProperties=%v]",
    	l.Email, l.LastFailureAt, l.NumAttempts, l.ScrIps, l.UserAgents, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LoginFailures.
// It customizes the JSON marshaling process for LoginFailures objects.
func (l LoginFailures) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "email", "last_failure_at", "num_attempts", "scr_ips", "user_agents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LoginFailures object to a map representation for JSON marshaling.
func (l LoginFailures) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Email != nil {
        structMap["email"] = l.Email
    }
    if l.LastFailureAt != nil {
        structMap["last_failure_at"] = l.LastFailureAt
    }
    if l.NumAttempts != nil {
        structMap["num_attempts"] = l.NumAttempts
    }
    if l.ScrIps != nil {
        structMap["scr_ips"] = l.ScrIps
    }
    if l.UserAgents != nil {
        structMap["user_agents"] = l.UserAgents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LoginFailures.
// It customizes the JSON unmarshaling process for LoginFailures objects.
func (l *LoginFailures) UnmarshalJSON(input []byte) error {
    var temp tempLoginFailures
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "email", "last_failure_at", "num_attempts", "scr_ips", "user_agents")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Email = temp.Email
    l.LastFailureAt = temp.LastFailureAt
    l.NumAttempts = temp.NumAttempts
    l.ScrIps = temp.ScrIps
    l.UserAgents = temp.UserAgents
    return nil
}

// tempLoginFailures is a temporary struct used for validating the fields of LoginFailures.
type tempLoginFailures  struct {
    Email         *string  `json:"email,omitempty"`
    LastFailureAt *int     `json:"last_failure_at,omitempty"`
    NumAttempts   *int     `json:"num_attempts,omitempty"`
    ScrIps        []string `json:"scr_ips,omitempty"`
    UserAgents    []string `json:"user_agents,omitempty"`
}
