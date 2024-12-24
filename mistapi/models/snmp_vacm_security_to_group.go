package models

import (
    "encoding/json"
    "fmt"
)

// SnmpVacmSecurityToGroup represents a SnmpVacmSecurityToGroup struct.
type SnmpVacmSecurityToGroup struct {
    Content              []SnmpVacmSecurityToGroupContentItem `json:"content,omitempty"`
    // enum: `usm`, `v1`, `v2c`
    SecurityModel        *SnmpVacmSecurityModelEnum           `json:"security_model,omitempty"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpVacmSecurityToGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpVacmSecurityToGroup) String() string {
    return fmt.Sprintf(
    	"SnmpVacmSecurityToGroup[Content=%v, SecurityModel=%v, AdditionalProperties=%v]",
    	s.Content, s.SecurityModel, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacmSecurityToGroup.
// It customizes the JSON marshaling process for SnmpVacmSecurityToGroup objects.
func (s SnmpVacmSecurityToGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "content", "security_model"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacmSecurityToGroup object to a map representation for JSON marshaling.
func (s SnmpVacmSecurityToGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Content != nil {
        structMap["content"] = s.Content
    }
    if s.SecurityModel != nil {
        structMap["security_model"] = s.SecurityModel
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpVacmSecurityToGroup.
// It customizes the JSON unmarshaling process for SnmpVacmSecurityToGroup objects.
func (s *SnmpVacmSecurityToGroup) UnmarshalJSON(input []byte) error {
    var temp tempSnmpVacmSecurityToGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "content", "security_model")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Content = temp.Content
    s.SecurityModel = temp.SecurityModel
    return nil
}

// tempSnmpVacmSecurityToGroup is a temporary struct used for validating the fields of SnmpVacmSecurityToGroup.
type tempSnmpVacmSecurityToGroup  struct {
    Content       []SnmpVacmSecurityToGroupContentItem `json:"content,omitempty"`
    SecurityModel *SnmpVacmSecurityModelEnum           `json:"security_model,omitempty"`
}
