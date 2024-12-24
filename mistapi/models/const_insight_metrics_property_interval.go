package models

import (
    "encoding/json"
    "fmt"
)

// ConstInsightMetricsPropertyInterval represents a ConstInsightMetricsPropertyInterval struct.
type ConstInsightMetricsPropertyInterval struct {
    Interval             *int                   `json:"interval,omitempty"`
    MaxAge               *int                   `json:"max_age,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyInterval,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyInterval) String() string {
    return fmt.Sprintf(
    	"ConstInsightMetricsPropertyInterval[Interval=%v, MaxAge=%v, AdditionalProperties=%v]",
    	c.Interval, c.MaxAge, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyInterval.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyInterval objects.
func (c ConstInsightMetricsPropertyInterval) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "interval", "max_age"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyInterval object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyInterval) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Interval != nil {
        structMap["interval"] = c.Interval
    }
    if c.MaxAge != nil {
        structMap["max_age"] = c.MaxAge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyInterval.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyInterval objects.
func (c *ConstInsightMetricsPropertyInterval) UnmarshalJSON(input []byte) error {
    var temp tempConstInsightMetricsPropertyInterval
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interval", "max_age")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Interval = temp.Interval
    c.MaxAge = temp.MaxAge
    return nil
}

// tempConstInsightMetricsPropertyInterval is a temporary struct used for validating the fields of ConstInsightMetricsPropertyInterval.
type tempConstInsightMetricsPropertyInterval  struct {
    Interval *int `json:"interval,omitempty"`
    MaxAge   *int `json:"max_age,omitempty"`
}
