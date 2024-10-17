package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SsrVersion represents a SsrVersion struct.
type SsrVersion struct {
	Default              *bool          `json:"default,omitempty"`
	Package              string         `json:"package"`
	Version              string         `json:"version"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsrVersion.
// It customizes the JSON marshaling process for SsrVersion objects.
func (s SsrVersion) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SsrVersion object to a map representation for JSON marshaling.
func (s SsrVersion) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Default != nil {
		structMap["default"] = s.Default
	}
	structMap["package"] = s.Package
	structMap["version"] = s.Version
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrVersion.
// It customizes the JSON unmarshaling process for SsrVersion objects.
func (s *SsrVersion) UnmarshalJSON(input []byte) error {
	var temp tempSsrVersion
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "default", "package", "version")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Default = temp.Default
	s.Package = *temp.Package
	s.Version = *temp.Version
	return nil
}

// tempSsrVersion is a temporary struct used for validating the fields of SsrVersion.
type tempSsrVersion struct {
	Default *bool   `json:"default,omitempty"`
	Package *string `json:"package"`
	Version *string `json:"version"`
}

func (s *tempSsrVersion) validate() error {
	var errs []string
	if s.Package == nil {
		errs = append(errs, "required field `package` is missing for type `ssr_version`")
	}
	if s.Version == nil {
		errs = append(errs, "required field `version` is missing for type `ssr_version`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
