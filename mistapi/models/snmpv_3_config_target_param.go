package models

import (
    "encoding/json"
)

// Snmpv3ConfigTargetParam represents a Snmpv3ConfigTargetParam struct.
type Snmpv3ConfigTargetParam struct {
    // enum: `v1`, `v2c`, `v3`
    MessageProcessingModel *Snmpv3ConfigTargetParamMessProcessModelEnum `json:"message_processing_model,omitempty"`
    Name                   *string                                      `json:"name,omitempty"`
    // refer to profile-name in notify_filter
    NotifyFilter           *string                                      `json:"notify_filter,omitempty"`
    // enum: `authentication`, `none`, `privacy`
    SecurityLevel          *Snmpv3ConfigTargetParamSecurityLevelEnum    `json:"security_level,omitempty"`
    // enum: `usm`, `v1`, `v2c`
    SecurityModel          *Snmpv3ConfigTargetParamSecurityModelEnum    `json:"security_model,omitempty"`
    // refer to security_name in usm
    SecurityName           *string                                      `json:"security_name,omitempty"`
    AdditionalProperties   map[string]any                               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigTargetParam.
// It customizes the JSON marshaling process for Snmpv3ConfigTargetParam objects.
func (s Snmpv3ConfigTargetParam) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigTargetParam object to a map representation for JSON marshaling.
func (s Snmpv3ConfigTargetParam) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.MessageProcessingModel != nil {
        structMap["message_processing_model"] = s.MessageProcessingModel
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NotifyFilter != nil {
        structMap["notify_filter"] = s.NotifyFilter
    }
    if s.SecurityLevel != nil {
        structMap["security_level"] = s.SecurityLevel
    }
    if s.SecurityModel != nil {
        structMap["security_model"] = s.SecurityModel
    }
    if s.SecurityName != nil {
        structMap["security_name"] = s.SecurityName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigTargetParam.
// It customizes the JSON unmarshaling process for Snmpv3ConfigTargetParam objects.
func (s *Snmpv3ConfigTargetParam) UnmarshalJSON(input []byte) error {
    var temp tempSnmpv3ConfigTargetParam
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "message_processing_model", "name", "notify_filter", "security_level", "security_model", "security_name")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.MessageProcessingModel = temp.MessageProcessingModel
    s.Name = temp.Name
    s.NotifyFilter = temp.NotifyFilter
    s.SecurityLevel = temp.SecurityLevel
    s.SecurityModel = temp.SecurityModel
    s.SecurityName = temp.SecurityName
    return nil
}

// tempSnmpv3ConfigTargetParam is a temporary struct used for validating the fields of Snmpv3ConfigTargetParam.
type tempSnmpv3ConfigTargetParam  struct {
    MessageProcessingModel *Snmpv3ConfigTargetParamMessProcessModelEnum `json:"message_processing_model,omitempty"`
    Name                   *string                                      `json:"name,omitempty"`
    NotifyFilter           *string                                      `json:"notify_filter,omitempty"`
    SecurityLevel          *Snmpv3ConfigTargetParamSecurityLevelEnum    `json:"security_level,omitempty"`
    SecurityModel          *Snmpv3ConfigTargetParamSecurityModelEnum    `json:"security_model,omitempty"`
    SecurityName           *string                                      `json:"security_name,omitempty"`
}
