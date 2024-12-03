package models

import (
    "encoding/json"
)

// VrfConfig represents a VrfConfig struct.
type VrfConfig struct {
    // whether to enable VRF (when supported on the device)
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrfConfig.
// It customizes the JSON marshaling process for VrfConfig objects.
func (v VrfConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VrfConfig object to a map representation for JSON marshaling.
func (v VrfConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Enabled != nil {
        structMap["enabled"] = v.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrfConfig.
// It customizes the JSON unmarshaling process for VrfConfig objects.
func (v *VrfConfig) UnmarshalJSON(input []byte) error {
    var temp tempVrfConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Enabled = temp.Enabled
    return nil
}

// tempVrfConfig is a temporary struct used for validating the fields of VrfConfig.
type tempVrfConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
