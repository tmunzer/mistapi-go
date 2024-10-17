package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SiteSleMetrics represents a SiteSleMetrics struct.
type SiteSleMetrics struct {
	Enabled              []string       `json:"enabled"`
	Supported            []string       `json:"supported"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSleMetrics.
// It customizes the JSON marshaling process for SiteSleMetrics objects.
func (s SiteSleMetrics) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSleMetrics object to a map representation for JSON marshaling.
func (s SiteSleMetrics) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "supported")
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
