package models

import (
    "encoding/json"
    "fmt"
)

// ServicePolicyEwfRule represents a ServicePolicyEwfRule struct.
type ServicePolicyEwfRule struct {
    AlertOnly            *bool                            `json:"alert_only,omitempty"`
    BlockMessage         *string                          `json:"block_message,omitempty"`
    Enabled              *bool                            `json:"enabled,omitempty"`
    // enum: `critical`, `standard`, `strict`
    Profile              *ServicePolicyEwfRuleProfileEnum `json:"profile,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicyEwfRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicyEwfRule) String() string {
    return fmt.Sprintf(
    	"ServicePolicyEwfRule[AlertOnly=%v, BlockMessage=%v, Enabled=%v, Profile=%v, AdditionalProperties=%v]",
    	s.AlertOnly, s.BlockMessage, s.Enabled, s.Profile, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyEwfRule.
// It customizes the JSON marshaling process for ServicePolicyEwfRule objects.
func (s ServicePolicyEwfRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "alert_only", "block_message", "enabled", "profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyEwfRule object to a map representation for JSON marshaling.
func (s ServicePolicyEwfRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempServicePolicyEwfRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alert_only", "block_message", "enabled", "profile")
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

// tempServicePolicyEwfRule is a temporary struct used for validating the fields of ServicePolicyEwfRule.
type tempServicePolicyEwfRule  struct {
    AlertOnly    *bool                            `json:"alert_only,omitempty"`
    BlockMessage *string                          `json:"block_message,omitempty"`
    Enabled      *bool                            `json:"enabled,omitempty"`
    Profile      *ServicePolicyEwfRuleProfileEnum `json:"profile,omitempty"`
}
