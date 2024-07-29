package models

import (
    "encoding/json"
)

// SwitchStpConfig represents a SwitchStpConfig struct.
type SwitchStpConfig struct {
    // enum: `rstp`, `vstp`
    Type                 *SwitchStpConfigTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStpConfig.
// It customizes the JSON marshaling process for SwitchStpConfig objects.
func (s SwitchStpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStpConfig object to a map representation for JSON marshaling.
func (s SwitchStpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStpConfig.
// It customizes the JSON unmarshaling process for SwitchStpConfig objects.
func (s *SwitchStpConfig) UnmarshalJSON(input []byte) error {
    var temp switchStpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Type = temp.Type
    return nil
}

// switchStpConfig is a temporary struct used for validating the fields of SwitchStpConfig.
type switchStpConfig  struct {
    Type *SwitchStpConfigTypeEnum `json:"type,omitempty"`
}
