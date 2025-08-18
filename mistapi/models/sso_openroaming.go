// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SsoOpenroaming represents a SsoOpenroaming struct.
// if `idp_type`==`openroaming`
type SsoOpenroaming struct {
	// SSIDs that support OpenRoaming
	Ssids []string `json:"ssids"`
	// Optional WBA-issued certificate. If not provided, the default WBA-issued certificate for Juniper will be used.
	WbaCert              *string                `json:"wba_cert,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoOpenroaming,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoOpenroaming) String() string {
	return fmt.Sprintf(
		"SsoOpenroaming[Ssids=%v, WbaCert=%v, AdditionalProperties=%v]",
		s.Ssids, s.WbaCert, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoOpenroaming.
// It customizes the JSON marshaling process for SsoOpenroaming objects.
func (s SsoOpenroaming) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"ssids", "wba_cert"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SsoOpenroaming object to a map representation for JSON marshaling.
func (s SsoOpenroaming) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["ssids"] = s.Ssids
	if s.WbaCert != nil {
		structMap["wba_cert"] = s.WbaCert
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoOpenroaming.
// It customizes the JSON unmarshaling process for SsoOpenroaming objects.
func (s *SsoOpenroaming) UnmarshalJSON(input []byte) error {
	var temp tempSsoOpenroaming
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ssids", "wba_cert")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Ssids = *temp.Ssids
	s.WbaCert = temp.WbaCert
	return nil
}

// tempSsoOpenroaming is a temporary struct used for validating the fields of SsoOpenroaming.
type tempSsoOpenroaming struct {
	Ssids   *[]string `json:"ssids"`
	WbaCert *string   `json:"wba_cert,omitempty"`
}

func (s *tempSsoOpenroaming) validate() error {
	var errs []string
	if s.Ssids == nil {
		errs = append(errs, "required field `ssids` is missing for type `sso_openroaming`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
