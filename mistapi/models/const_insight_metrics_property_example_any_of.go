// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstInsightMetricsPropertyExampleAnyOf represents a ConstInsightMetricsPropertyExampleAnyOf struct.
type ConstInsightMetricsPropertyExampleAnyOf struct {
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyExampleAnyOf,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyExampleAnyOf) String() string {
	return fmt.Sprintf(
		"ConstInsightMetricsPropertyExampleAnyOf[AdditionalProperties=%v]",
		c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyExampleAnyOf.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyExampleAnyOf objects.
func (c ConstInsightMetricsPropertyExampleAnyOf) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyExampleAnyOf object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyExampleAnyOf) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyExampleAnyOf.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyExampleAnyOf objects.
func (c *ConstInsightMetricsPropertyExampleAnyOf) UnmarshalJSON(input []byte) error {
	var temp tempConstInsightMetricsPropertyExampleAnyOf
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input)
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	return nil
}

// tempConstInsightMetricsPropertyExampleAnyOf is a temporary struct used for validating the fields of ConstInsightMetricsPropertyExampleAnyOf.
type tempConstInsightMetricsPropertyExampleAnyOf struct {
}
