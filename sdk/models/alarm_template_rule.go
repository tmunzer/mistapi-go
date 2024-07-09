package models

import (
    "encoding/json"
)

// AlarmTemplateRule represents a AlarmTemplateRule struct.
type AlarmTemplateRule struct {
    // Delivery object to configure the alarm delivery
    Delivery             *Delivery      `json:"delivery,omitempty"`
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AlarmTemplateRule.
// It customizes the JSON marshaling process for AlarmTemplateRule objects.
func (a AlarmTemplateRule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AlarmTemplateRule object to a map representation for JSON marshaling.
func (a AlarmTemplateRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Delivery != nil {
        structMap["delivery"] = a.Delivery.toMap()
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AlarmTemplateRule.
// It customizes the JSON unmarshaling process for AlarmTemplateRule objects.
func (a *AlarmTemplateRule) UnmarshalJSON(input []byte) error {
    var temp alarmTemplateRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "delivery", "enabled")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Delivery = temp.Delivery
    a.Enabled = temp.Enabled
    return nil
}

// alarmTemplateRule is a temporary struct used for validating the fields of AlarmTemplateRule.
type alarmTemplateRule  struct {
    Delivery *Delivery `json:"delivery,omitempty"`
    Enabled  *bool     `json:"enabled,omitempty"`
}
