// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SiteSleMetrics represents a SiteSleMetrics struct.
type SiteSleMetrics struct {
	Enabled              []string               `json:"enabled"`
	Supported            []string               `json:"supported"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSleMetrics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSleMetrics) String() string {
	return fmt.Sprintf(
		"SiteSleMetrics[Enabled=%v, Supported=%v, AdditionalProperties=%v]",
		s.Enabled, s.Supported, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSleMetrics.
// It customizes the JSON marshaling process for SiteSleMetrics objects.
func (s SiteSleMetrics) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled", "supported"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSleMetrics object to a map representation for JSON marshaling.
func (s SiteSleMetrics) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["enabled"] = s.Enabled
	structMap["supported"] = s.Supported
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSleMetrics.
// It customizes the JSON unmarshaling process for SiteSleMetrics objects.
func (s *SiteSleMetrics) UnmarshalJSON(input []byte) error {
	var temp tempSiteSleMetrics
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "supported")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Enabled = *temp.Enabled
	s.Supported = *temp.Supported
	return nil
}

// tempSiteSleMetrics is a temporary struct used for validating the fields of SiteSleMetrics.
type tempSiteSleMetrics struct {
	Enabled   *[]string `json:"enabled"`
	Supported *[]string `json:"supported"`
}

func (s *tempSiteSleMetrics) validate() error {
	var errs []string
	if s.Enabled == nil {
		errs = append(errs, "required field `enabled` is missing for type `site_sle_metrics`")
	}
	if s.Supported == nil {
		errs = append(errs, "required field `supported` is missing for type `site_sle_metrics`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
