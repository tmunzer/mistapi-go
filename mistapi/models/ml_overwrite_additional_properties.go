package models

import (
    "encoding/json"
)

// MlOverwriteAdditionalProperties represents a MlOverwriteAdditionalProperties struct.
type MlOverwriteAdditionalProperties struct {
    Int                  *int           `json:"int,omitempty"`
    Ple                  *int           `json:"ple,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MlOverwriteAdditionalProperties.
// It customizes the JSON marshaling process for MlOverwriteAdditionalProperties objects.
func (m MlOverwriteAdditionalProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MlOverwriteAdditionalProperties object to a map representation for JSON marshaling.
func (m MlOverwriteAdditionalProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Int != nil {
        structMap["int"] = m.Int
    }
    if m.Ple != nil {
        structMap["ple"] = m.Ple
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MlOverwriteAdditionalProperties.
// It customizes the JSON unmarshaling process for MlOverwriteAdditionalProperties objects.
func (m *MlOverwriteAdditionalProperties) UnmarshalJSON(input []byte) error {
    var temp tempMlOverwriteAdditionalProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "int", "ple")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Int = temp.Int
    m.Ple = temp.Ple
    return nil
}

// tempMlOverwriteAdditionalProperties is a temporary struct used for validating the fields of MlOverwriteAdditionalProperties.
type tempMlOverwriteAdditionalProperties  struct {
    Int *int `json:"int,omitempty"`
    Ple *int `json:"ple,omitempty"`
}
