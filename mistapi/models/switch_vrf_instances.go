package models

import (
    "encoding/json"
)

// SwitchVrfInstances represents a SwitchVrfInstances struct.
// Property key is the network name
type SwitchVrfInstances struct {
    Networks             []string       `json:"networks,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchVrfInstances.
// It customizes the JSON marshaling process for SwitchVrfInstances objects.
func (s SwitchVrfInstances) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchVrfInstances object to a map representation for JSON marshaling.
func (s SwitchVrfInstances) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Networks != nil {
        structMap["networks"] = s.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVrfInstances.
// It customizes the JSON unmarshaling process for SwitchVrfInstances objects.
func (s *SwitchVrfInstances) UnmarshalJSON(input []byte) error {
    var temp tempSwitchVrfInstances
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "networks")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Networks = temp.Networks
    return nil
}

// tempSwitchVrfInstances is a temporary struct used for validating the fields of SwitchVrfInstances.
type tempSwitchVrfInstances  struct {
    Networks []string `json:"networks,omitempty"`
}
