package models

import (
    "encoding/json"
)

// ApStatsIotStatAdditionalProperties represents a ApStatsIotStatAdditionalProperties struct.
type ApStatsIotStatAdditionalProperties struct {
    Value                *int           `json:"value,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsIotStatAdditionalProperties.
// It customizes the JSON marshaling process for ApStatsIotStatAdditionalProperties objects.
func (a ApStatsIotStatAdditionalProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsIotStatAdditionalProperties object to a map representation for JSON marshaling.
func (a ApStatsIotStatAdditionalProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Value != nil {
        structMap["value"] = a.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsIotStatAdditionalProperties.
// It customizes the JSON unmarshaling process for ApStatsIotStatAdditionalProperties objects.
func (a *ApStatsIotStatAdditionalProperties) UnmarshalJSON(input []byte) error {
    var temp apStatsIotStatAdditionalProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "value")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Value = temp.Value
    return nil
}

// apStatsIotStatAdditionalProperties is a temporary struct used for validating the fields of ApStatsIotStatAdditionalProperties.
type apStatsIotStatAdditionalProperties  struct {
    Value *int `json:"value,omitempty"`
}
