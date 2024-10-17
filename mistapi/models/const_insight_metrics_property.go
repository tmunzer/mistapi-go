package models

import (
	"encoding/json"
)

// ConstInsightMetricsProperty represents a ConstInsightMetricsProperty struct.
type ConstInsightMetricsProperty struct {
	Ctype       []string                             `json:"ctype,omitempty"`
	Description *string                              `json:"description,omitempty"`
	Example     []ConstInsightMetricsPropertyExample `json:"example,omitempty"`
	// Property key is the interval (e.g. 10m, 1h, ...)
	Intervals map[string]ConstInsightMetricsPropertyInterval `json:"intervals,omitempty"`
	Keys      *interface{}                                   `json:"keys,omitempty"`
	// Property key is the parameter name
	Params map[string]ConstInsightMetricsPropertyParam `json:"params,omitempty"`
	// Property key is the duration (e.g. 1d, 1w, ...)
	ReportDuration       map[string]ConstInsightMetricsPropertyReportDuration `json:"report_duration,omitempty"`
	ReportScopes         []string                                             `json:"report_scopes,omitempty"`
	Scopes               []ConstInsightMetricsPropertyScopeEnum               `json:"scopes,omitempty"`
	SleBaselined         *bool                                                `json:"sle_baselined,omitempty"`
	SleClassifiers       []string                                             `json:"sle_classifiers,omitempty"`
	Type                 *string                                              `json:"type,omitempty"`
	Unit                 *string                                              `json:"unit,omitempty"`
	Values               *interface{}                                         `json:"values,omitempty"`
	AdditionalProperties map[string]any                                       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsProperty.
// It customizes the JSON marshaling process for ConstInsightMetricsProperty objects.
func (c ConstInsightMetricsProperty) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsProperty object to a map representation for JSON marshaling.
func (c ConstInsightMetricsProperty) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Ctype != nil {
		structMap["ctype"] = c.Ctype
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	if c.Example != nil {
		structMap["example"] = c.Example
	}
	if c.Intervals != nil {
		structMap["intervals"] = c.Intervals
	}
	if c.Keys != nil {
		structMap["keys"] = c.Keys
	}
	if c.Params != nil {
		structMap["params"] = c.Params
	}
	if c.ReportDuration != nil {
		structMap["report_duration"] = c.ReportDuration
	}
	if c.ReportScopes != nil {
		structMap["report_scopes"] = c.ReportScopes
	}
	if c.Scopes != nil {
		structMap["scopes"] = c.Scopes
	}
	if c.SleBaselined != nil {
		structMap["sle_baselined"] = c.SleBaselined
	}
	if c.SleClassifiers != nil {
		structMap["sle_classifiers"] = c.SleClassifiers
	}
	if c.Type != nil {
		structMap["type"] = c.Type
	}
	if c.Unit != nil {
		structMap["unit"] = c.Unit
	}
	if c.Values != nil {
		structMap["values"] = c.Values
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsProperty.
// It customizes the JSON unmarshaling process for ConstInsightMetricsProperty objects.
func (c *ConstInsightMetricsProperty) UnmarshalJSON(input []byte) error {
	var temp tempConstInsightMetricsProperty
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ctype", "description", "example", "intervals", "keys", "params", "report_duration", "report_scopes", "scopes", "sle_baselined", "sle_classifiers", "type", "unit", "values")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Ctype = temp.Ctype
	c.Description = temp.Description
	c.Example = temp.Example
	c.Intervals = temp.Intervals
	c.Keys = temp.Keys
	c.Params = temp.Params
	c.ReportDuration = temp.ReportDuration
	c.ReportScopes = temp.ReportScopes
	c.Scopes = temp.Scopes
	c.SleBaselined = temp.SleBaselined
	c.SleClassifiers = temp.SleClassifiers
	c.Type = temp.Type
	c.Unit = temp.Unit
	c.Values = temp.Values
	return nil
}

// tempConstInsightMetricsProperty is a temporary struct used for validating the fields of ConstInsightMetricsProperty.
type tempConstInsightMetricsProperty struct {
	Ctype          []string                                             `json:"ctype,omitempty"`
	Description    *string                                              `json:"description,omitempty"`
	Example        []ConstInsightMetricsPropertyExample                 `json:"example,omitempty"`
	Intervals      map[string]ConstInsightMetricsPropertyInterval       `json:"intervals,omitempty"`
	Keys           *interface{}                                         `json:"keys,omitempty"`
	Params         map[string]ConstInsightMetricsPropertyParam          `json:"params,omitempty"`
	ReportDuration map[string]ConstInsightMetricsPropertyReportDuration `json:"report_duration,omitempty"`
	ReportScopes   []string                                             `json:"report_scopes,omitempty"`
	Scopes         []ConstInsightMetricsPropertyScopeEnum               `json:"scopes,omitempty"`
	SleBaselined   *bool                                                `json:"sle_baselined,omitempty"`
	SleClassifiers []string                                             `json:"sle_classifiers,omitempty"`
	Type           *string                                              `json:"type,omitempty"`
	Unit           *string                                              `json:"unit,omitempty"`
	Values         *interface{}                                         `json:"values,omitempty"`
}
