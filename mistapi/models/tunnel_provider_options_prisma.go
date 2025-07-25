// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// TunnelProviderOptionsPrisma represents a TunnelProviderOptionsPrisma struct.
type TunnelProviderOptionsPrisma struct {
    // For prisma-ipsec, service account name to used for tunnel auto provisioning
    ServiceAccountName   *string                `json:"service_account_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelProviderOptionsPrisma,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelProviderOptionsPrisma) String() string {
    return fmt.Sprintf(
    	"TunnelProviderOptionsPrisma[ServiceAccountName=%v, AdditionalProperties=%v]",
    	t.ServiceAccountName, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsPrisma.
// It customizes the JSON marshaling process for TunnelProviderOptionsPrisma objects.
func (t TunnelProviderOptionsPrisma) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "service_account_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsPrisma object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsPrisma) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.ServiceAccountName != nil {
        structMap["service_account_name"] = t.ServiceAccountName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsPrisma.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsPrisma objects.
func (t *TunnelProviderOptionsPrisma) UnmarshalJSON(input []byte) error {
    var temp tempTunnelProviderOptionsPrisma
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_account_name")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.ServiceAccountName = temp.ServiceAccountName
    return nil
}

// tempTunnelProviderOptionsPrisma is a temporary struct used for validating the fields of TunnelProviderOptionsPrisma.
type tempTunnelProviderOptionsPrisma  struct {
    ServiceAccountName *string `json:"service_account_name,omitempty"`
}
