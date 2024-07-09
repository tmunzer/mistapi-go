package models

import (
    "encoding/json"
)

// ExtraRoutePropertiesNextQualifiedProperties represents a ExtraRoutePropertiesNextQualifiedProperties struct.
type ExtraRoutePropertiesNextQualifiedProperties struct {
    Metric               Optional[int]  `json:"metric"`
    Preference           Optional[int]  `json:"preference"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ExtraRoutePropertiesNextQualifiedProperties.
// It customizes the JSON marshaling process for ExtraRoutePropertiesNextQualifiedProperties objects.
func (e ExtraRoutePropertiesNextQualifiedProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the ExtraRoutePropertiesNextQualifiedProperties object to a map representation for JSON marshaling.
func (e ExtraRoutePropertiesNextQualifiedProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Metric.IsValueSet() {
        if e.Metric.Value() != nil {
            structMap["metric"] = e.Metric.Value()
        } else {
            structMap["metric"] = nil
        }
    }
    if e.Preference.IsValueSet() {
        if e.Preference.Value() != nil {
            structMap["preference"] = e.Preference.Value()
        } else {
            structMap["preference"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRoutePropertiesNextQualifiedProperties.
// It customizes the JSON unmarshaling process for ExtraRoutePropertiesNextQualifiedProperties objects.
func (e *ExtraRoutePropertiesNextQualifiedProperties) UnmarshalJSON(input []byte) error {
    var temp extraRoutePropertiesNextQualifiedProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metric", "preference")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Metric = temp.Metric
    e.Preference = temp.Preference
    return nil
}

// extraRoutePropertiesNextQualifiedProperties is a temporary struct used for validating the fields of ExtraRoutePropertiesNextQualifiedProperties.
type extraRoutePropertiesNextQualifiedProperties  struct {
    Metric     Optional[int] `json:"metric"`
    Preference Optional[int] `json:"preference"`
}
