package models

import (
    "encoding/json"
)

// ServicePolicyAppqoe represents a ServicePolicyAppqoe struct.
// For SRX Only
type ServicePolicyAppqoe struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyAppqoe.
// It customizes the JSON marshaling process for ServicePolicyAppqoe objects.
func (s ServicePolicyAppqoe) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyAppqoe object to a map representation for JSON marshaling.
func (s ServicePolicyAppqoe) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicyAppqoe.
// It customizes the JSON unmarshaling process for ServicePolicyAppqoe objects.
func (s *ServicePolicyAppqoe) UnmarshalJSON(input []byte) error {
    var temp servicePolicyAppqoe
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    return nil
}

// servicePolicyAppqoe is a temporary struct used for validating the fields of ServicePolicyAppqoe.
type servicePolicyAppqoe  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
