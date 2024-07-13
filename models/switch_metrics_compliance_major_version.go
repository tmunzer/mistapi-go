package models

import (
    "encoding/json"
)

// SwitchMetricsComplianceMajorVersion represents a SwitchMetricsComplianceMajorVersion struct.
type SwitchMetricsComplianceMajorVersion struct {
    MajorCount           *int           `json:"major_count,omitempty"`
    MajorVersion         *string        `json:"major_version,omitempty"`
    Model                *string        `json:"model,omitempty"`
    SystemNames          []string       `json:"system_names,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMetricsComplianceMajorVersion.
// It customizes the JSON marshaling process for SwitchMetricsComplianceMajorVersion objects.
func (s SwitchMetricsComplianceMajorVersion) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMetricsComplianceMajorVersion object to a map representation for JSON marshaling.
func (s SwitchMetricsComplianceMajorVersion) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.MajorCount != nil {
        structMap["major_count"] = s.MajorCount
    }
    if s.MajorVersion != nil {
        structMap["major_version"] = s.MajorVersion
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.SystemNames != nil {
        structMap["system_names"] = s.SystemNames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMetricsComplianceMajorVersion.
// It customizes the JSON unmarshaling process for SwitchMetricsComplianceMajorVersion objects.
func (s *SwitchMetricsComplianceMajorVersion) UnmarshalJSON(input []byte) error {
    var temp switchMetricsComplianceMajorVersion
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "major_count", "major_version", "model", "system_names")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.MajorCount = temp.MajorCount
    s.MajorVersion = temp.MajorVersion
    s.Model = temp.Model
    s.SystemNames = temp.SystemNames
    return nil
}

// switchMetricsComplianceMajorVersion is a temporary struct used for validating the fields of SwitchMetricsComplianceMajorVersion.
type switchMetricsComplianceMajorVersion  struct {
    MajorCount   *int     `json:"major_count,omitempty"`
    MajorVersion *string  `json:"major_version,omitempty"`
    Model        *string  `json:"model,omitempty"`
    SystemNames  []string `json:"system_names,omitempty"`
}