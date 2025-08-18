// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SkyatpListIp represents a SkyatpListIp struct.
type SkyatpListIp struct {
	Comment              *string                `json:"comment,omitempty"`
	Value                string                 `json:"value"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SkyatpListIp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SkyatpListIp) String() string {
	return fmt.Sprintf(
		"SkyatpListIp[Comment=%v, Value=%v, AdditionalProperties=%v]",
		s.Comment, s.Value, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SkyatpListIp.
// It customizes the JSON marshaling process for SkyatpListIp objects.
func (s SkyatpListIp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"comment", "value"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SkyatpListIp object to a map representation for JSON marshaling.
func (s SkyatpListIp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Comment != nil {
		structMap["comment"] = s.Comment
	}
	structMap["value"] = s.Value
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SkyatpListIp.
// It customizes the JSON unmarshaling process for SkyatpListIp objects.
func (s *SkyatpListIp) UnmarshalJSON(input []byte) error {
	var temp tempSkyatpListIp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "comment", "value")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Comment = temp.Comment
	s.Value = *temp.Value
	return nil
}

// tempSkyatpListIp is a temporary struct used for validating the fields of SkyatpListIp.
type tempSkyatpListIp struct {
	Comment *string `json:"comment,omitempty"`
	Value   *string `json:"value"`
}

func (s *tempSkyatpListIp) validate() error {
	var errs []string
	if s.Value == nil {
		errs = append(errs, "required field `value` is missing for type `skyatp_list_ip`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
