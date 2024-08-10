package models

import (
    "encoding/json"
)

// ResponseSwitchMetricsVersionComplianceDetails represents a ResponseSwitchMetricsVersionComplianceDetails struct.
type ResponseSwitchMetricsVersionComplianceDetails struct {
    MajorVersions        []SwitchMetricsComplianceMajorVersion `json:"major_versions,omitempty"`
    AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsVersionComplianceDetails.
// It customizes the JSON marshaling process for ResponseSwitchMetricsVersionComplianceDetails objects.
func (r ResponseSwitchMetricsVersionComplianceDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsVersionComplianceDetails object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsVersionComplianceDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.MajorVersions != nil {
        structMap["major_versions"] = r.MajorVersions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsVersionComplianceDetails.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsVersionComplianceDetails objects.
func (r *ResponseSwitchMetricsVersionComplianceDetails) UnmarshalJSON(input []byte) error {
    var temp tempResponseSwitchMetricsVersionComplianceDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "major_versions")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.MajorVersions = temp.MajorVersions
    return nil
}

// tempResponseSwitchMetricsVersionComplianceDetails is a temporary struct used for validating the fields of ResponseSwitchMetricsVersionComplianceDetails.
type tempResponseSwitchMetricsVersionComplianceDetails  struct {
    MajorVersions []SwitchMetricsComplianceMajorVersion `json:"major_versions,omitempty"`
}
