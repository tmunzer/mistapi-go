package models

import (
    "encoding/json"
    "fmt"
)

// SleThreshold represents a SleThreshold struct.
type SleThreshold struct {
    Default              *float64               `json:"default,omitempty"`
    Direction            *string                `json:"direction,omitempty"`
    Maximum              *float64               `json:"maximum,omitempty"`
    Metric               *string                `json:"metric,omitempty"`
    Minimum              *float64               `json:"minimum,omitempty"`
    Threshold            *string                `json:"threshold,omitempty"`
    Units                *string                `json:"units,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleThreshold,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleThreshold) String() string {
    return fmt.Sprintf(
    	"SleThreshold[Default=%v, Direction=%v, Maximum=%v, Metric=%v, Minimum=%v, Threshold=%v, Units=%v, AdditionalProperties=%v]",
    	s.Default, s.Direction, s.Maximum, s.Metric, s.Minimum, s.Threshold, s.Units, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleThreshold.
// It customizes the JSON marshaling process for SleThreshold objects.
func (s SleThreshold) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "default", "direction", "maximum", "metric", "minimum", "threshold", "units"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleThreshold object to a map representation for JSON marshaling.
func (s SleThreshold) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Default != nil {
        structMap["default"] = s.Default
    }
    if s.Direction != nil {
        structMap["direction"] = s.Direction
    }
    if s.Maximum != nil {
        structMap["maximum"] = s.Maximum
    }
    if s.Metric != nil {
        structMap["metric"] = s.Metric
    }
    if s.Minimum != nil {
        structMap["minimum"] = s.Minimum
    }
    if s.Threshold != nil {
        structMap["threshold"] = s.Threshold
    }
    if s.Units != nil {
        structMap["units"] = s.Units
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleThreshold.
// It customizes the JSON unmarshaling process for SleThreshold objects.
func (s *SleThreshold) UnmarshalJSON(input []byte) error {
    var temp tempSleThreshold
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "default", "direction", "maximum", "metric", "minimum", "threshold", "units")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Default = temp.Default
    s.Direction = temp.Direction
    s.Maximum = temp.Maximum
    s.Metric = temp.Metric
    s.Minimum = temp.Minimum
    s.Threshold = temp.Threshold
    s.Units = temp.Units
    return nil
}

// tempSleThreshold is a temporary struct used for validating the fields of SleThreshold.
type tempSleThreshold  struct {
    Default   *float64 `json:"default,omitempty"`
    Direction *string  `json:"direction,omitempty"`
    Maximum   *float64 `json:"maximum,omitempty"`
    Metric    *string  `json:"metric,omitempty"`
    Minimum   *float64 `json:"minimum,omitempty"`
    Threshold *string  `json:"threshold,omitempty"`
    Units     *string  `json:"units,omitempty"`
}
