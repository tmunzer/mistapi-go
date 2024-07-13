package models

import (
    "encoding/json"
)

// ResponseSwitchMetrics represents a ResponseSwitchMetrics struct.
type ResponseSwitchMetrics struct {
    ActivePortsSummary   *ResponseSwitchMetricsActivePortsSummary `json:"active_ports_summary,omitempty"`
    ConfigSuccess        *ResponseSwitchMetricsConfigSuccess      `json:"config_success,omitempty"`
    VersionCompliance    *ResponseSwitchMetricsVersionCompliance  `json:"version_compliance,omitempty"`
    AdditionalProperties map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetrics.
// It customizes the JSON marshaling process for ResponseSwitchMetrics objects.
func (r ResponseSwitchMetrics) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetrics object to a map representation for JSON marshaling.
func (r ResponseSwitchMetrics) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp responseSwitchMetrics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "active_ports_summary", "config_success", "version_compliance")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ActivePortsSummary = temp.ActivePortsSummary
    r.ConfigSuccess = temp.ConfigSuccess
    r.VersionCompliance = temp.VersionCompliance
    return nil
}

// responseSwitchMetrics is a temporary struct used for validating the fields of ResponseSwitchMetrics.
type responseSwitchMetrics  struct {
    ActivePortsSummary *ResponseSwitchMetricsActivePortsSummary `json:"active_ports_summary,omitempty"`
    ConfigSuccess      *ResponseSwitchMetricsConfigSuccess      `json:"config_success,omitempty"`
    VersionCompliance  *ResponseSwitchMetricsVersionCompliance  `json:"version_compliance,omitempty"`
}
