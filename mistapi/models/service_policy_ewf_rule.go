package models

import (
    "encoding/json"
)

// ServicePolicyEwfRule represents a ServicePolicyEwfRule struct.
type ServicePolicyEwfRule struct {
    AlertOnly            *bool                            `json:"alert_only,omitempty"`
    BlockMessage         *string                          `json:"block_message,omitempty"`
    Enabled              *bool                            `json:"enabled,omitempty"`
    // enum: `critical`, `standard`, `strict`
    Profile              *ServicePolicyEwfRuleProfileEnum `json:"profile,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyEwfRule.
// It customizes the JSON marshaling process for ServicePolicyEwfRule objects.
func (s ServicePolicyEwfRule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyEwfRule object to a map representation for JSON marshaling.
func (s ServicePolicyEwfRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AlertOnly != nil {
        structMap["alert_only"] = s.AlertOnly
    }
    if s.BlockMessage != nil {
        structMap["block_message"] = s.BlockMessage
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Profile != nil {
        structMap["profile"] = s.Profile
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicyEwfRule.
// It customizes the JSON unmarshaling process for ServicePolicyEwfRule objects.
func (s *ServicePolicyEwfRule) UnmarshalJSON(input []byte) error {
    var temp servicePolicyEwfRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "alert_only", "block_message", "enabled", "profile")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AlertOnly = temp.AlertOnly
    s.BlockMessage = temp.BlockMessage
    s.Enabled = temp.Enabled
    s.Profile = temp.Profile
    return nil
}

// servicePolicyEwfRule is a temporary struct used for validating the fields of ServicePolicyEwfRule.
type servicePolicyEwfRule  struct {
    AlertOnly    *bool                            `json:"alert_only,omitempty"`
    BlockMessage *string                          `json:"block_message,omitempty"`
    Enabled      *bool                            `json:"enabled,omitempty"`
    Profile      *ServicePolicyEwfRuleProfileEnum `json:"profile,omitempty"`
}
