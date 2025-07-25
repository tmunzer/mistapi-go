// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstAlarmDefinition represents a ConstAlarmDefinition struct.
type ConstAlarmDefinition struct {
    // Description of the alarm type
    Display                  string                 `json:"display"`
    Example                  *interface{}           `json:"example,omitempty"`
    // List of fields available in an alarm details payload (in REST APIs & Webhooks); e.g. `aps`, `switches`, `gateways`, `hostnames`, `ssids`, `bssids`
    Fields                   []string               `json:"fields"`
    // Group to which the alarm belongs
    Group                    string                 `json:"group"`
    // Key name of the alarm type
    Key                      string                 `json:"key"`
    // Marvis defined category to which the alarm belongs
    MarvisSuggestionCategory *string                `json:"marvis_suggestion_category,omitempty"`
    // Severity of the alarm
    Severity                 string                 `json:"severity"`
    AdditionalProperties     map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstAlarmDefinition,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstAlarmDefinition) String() string {
    return fmt.Sprintf(
    	"ConstAlarmDefinition[Display=%v, Example=%v, Fields=%v, Group=%v, Key=%v, MarvisSuggestionCategory=%v, Severity=%v, AdditionalProperties=%v]",
    	c.Display, c.Example, c.Fields, c.Group, c.Key, c.MarvisSuggestionCategory, c.Severity, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstAlarmDefinition.
// It customizes the JSON marshaling process for ConstAlarmDefinition objects.
func (c ConstAlarmDefinition) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "example", "fields", "group", "key", "marvis_suggestion_category", "severity"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstAlarmDefinition object to a map representation for JSON marshaling.
func (c ConstAlarmDefinition) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["display"] = c.Display
    if c.Example != nil {
        structMap["example"] = c.Example
    }
    structMap["fields"] = c.Fields
    structMap["group"] = c.Group
    structMap["key"] = c.Key
    if c.MarvisSuggestionCategory != nil {
        structMap["marvis_suggestion_category"] = c.MarvisSuggestionCategory
    }
    structMap["severity"] = c.Severity
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstAlarmDefinition.
// It customizes the JSON unmarshaling process for ConstAlarmDefinition objects.
func (c *ConstAlarmDefinition) UnmarshalJSON(input []byte) error {
    var temp tempConstAlarmDefinition
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "example", "fields", "group", "key", "marvis_suggestion_category", "severity")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = *temp.Display
    c.Example = temp.Example
    c.Fields = *temp.Fields
    c.Group = *temp.Group
    c.Key = *temp.Key
    c.MarvisSuggestionCategory = temp.MarvisSuggestionCategory
    c.Severity = *temp.Severity
    return nil
}

// tempConstAlarmDefinition is a temporary struct used for validating the fields of ConstAlarmDefinition.
type tempConstAlarmDefinition  struct {
    Display                  *string      `json:"display"`
    Example                  *interface{} `json:"example,omitempty"`
    Fields                   *[]string    `json:"fields"`
    Group                    *string      `json:"group"`
    Key                      *string      `json:"key"`
    MarvisSuggestionCategory *string      `json:"marvis_suggestion_category,omitempty"`
    Severity                 *string      `json:"severity"`
}

func (c *tempConstAlarmDefinition) validate() error {
    var errs []string
    if c.Display == nil {
        errs = append(errs, "required field `display` is missing for type `const_alarm_definition`")
    }
    if c.Fields == nil {
        errs = append(errs, "required field `fields` is missing for type `const_alarm_definition`")
    }
    if c.Group == nil {
        errs = append(errs, "required field `group` is missing for type `const_alarm_definition`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `const_alarm_definition`")
    }
    if c.Severity == nil {
        errs = append(errs, "required field `severity` is missing for type `const_alarm_definition`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
