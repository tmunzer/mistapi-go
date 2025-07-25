// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseSwitchMetrics represents a ResponseSwitchMetrics struct.
type ResponseSwitchMetrics struct {
    ActivePortsSummary   *ResponseSwitchMetricsActivePortsSummary `json:"active_ports_summary,omitempty"`
    ConfigSuccess        *ResponseSwitchMetricsConfigSuccess      `json:"config_success,omitempty"`
    VersionCompliance    *ResponseSwitchMetricsVersionCompliance  `json:"version_compliance,omitempty"`
    AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchMetrics,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchMetrics) String() string {
    return fmt.Sprintf(
    	"ResponseSwitchMetrics[ActivePortsSummary=%v, ConfigSuccess=%v, VersionCompliance=%v, AdditionalProperties=%v]",
    	r.ActivePortsSummary, r.ConfigSuccess, r.VersionCompliance, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetrics.
// It customizes the JSON marshaling process for ResponseSwitchMetrics objects.
func (r ResponseSwitchMetrics) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "active_ports_summary", "config_success", "version_compliance"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetrics object to a map representation for JSON marshaling.
func (r ResponseSwitchMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ActivePortsSummary != nil {
        structMap["active_ports_summary"] = r.ActivePortsSummary.toMap()
    }
    if r.ConfigSuccess != nil {
        structMap["config_success"] = r.ConfigSuccess.toMap()
    }
    if r.VersionCompliance != nil {
        structMap["version_compliance"] = r.VersionCompliance.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetrics.
// It customizes the JSON unmarshaling process for ResponseSwitchMetrics objects.
func (r *ResponseSwitchMetrics) UnmarshalJSON(input []byte) error {
    var temp tempResponseSwitchMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active_ports_summary", "config_success", "version_compliance")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ActivePortsSummary = temp.ActivePortsSummary
    r.ConfigSuccess = temp.ConfigSuccess
    r.VersionCompliance = temp.VersionCompliance
    return nil
}

// tempResponseSwitchMetrics is a temporary struct used for validating the fields of ResponseSwitchMetrics.
type tempResponseSwitchMetrics  struct {
    ActivePortsSummary *ResponseSwitchMetricsActivePortsSummary `json:"active_ports_summary,omitempty"`
    ConfigSuccess      *ResponseSwitchMetricsConfigSuccess      `json:"config_success,omitempty"`
    VersionCompliance  *ResponseSwitchMetricsVersionCompliance  `json:"version_compliance,omitempty"`
}
