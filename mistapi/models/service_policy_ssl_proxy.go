// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ServicePolicySslProxy represents a ServicePolicySslProxy struct.
// For SRX-only
type ServicePolicySslProxy struct {
    // enum: `medium`, `strong`, `weak`
    CiphersCategory      *SslProxyCiphersCategoryEnum `json:"ciphers_category,omitempty"`
    Enabled              *bool                        `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySslProxy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySslProxy) String() string {
    return fmt.Sprintf(
    	"ServicePolicySslProxy[CiphersCategory=%v, Enabled=%v, AdditionalProperties=%v]",
    	s.CiphersCategory, s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySslProxy.
// It customizes the JSON marshaling process for ServicePolicySslProxy objects.
func (s ServicePolicySslProxy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ciphers_category", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySslProxy object to a map representation for JSON marshaling.
func (s ServicePolicySslProxy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CiphersCategory != nil {
        structMap["ciphers_category"] = s.CiphersCategory
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySslProxy.
// It customizes the JSON unmarshaling process for ServicePolicySslProxy objects.
func (s *ServicePolicySslProxy) UnmarshalJSON(input []byte) error {
    var temp tempServicePolicySslProxy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ciphers_category", "enabled")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CiphersCategory = temp.CiphersCategory
    s.Enabled = temp.Enabled
    return nil
}

// tempServicePolicySslProxy is a temporary struct used for validating the fields of ServicePolicySslProxy.
type tempServicePolicySslProxy  struct {
    CiphersCategory *SslProxyCiphersCategoryEnum `json:"ciphers_category,omitempty"`
    Enabled         *bool                        `json:"enabled,omitempty"`
}
