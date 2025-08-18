// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseSwitchMetricsVersionComplianceDetails represents a ResponseSwitchMetricsVersionComplianceDetails struct.
type ResponseSwitchMetricsVersionComplianceDetails struct {
	MajorVersions        []SwitchMetricsComplianceMajorVersion `json:"major_versions,omitempty"`
	AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchMetricsVersionComplianceDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchMetricsVersionComplianceDetails) String() string {
	return fmt.Sprintf(
		"ResponseSwitchMetricsVersionComplianceDetails[MajorVersions=%v, AdditionalProperties=%v]",
		r.MajorVersions, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsVersionComplianceDetails.
// It customizes the JSON marshaling process for ResponseSwitchMetricsVersionComplianceDetails objects.
func (r ResponseSwitchMetricsVersionComplianceDetails) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"major_versions"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsVersionComplianceDetails object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsVersionComplianceDetails) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "major_versions")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.MajorVersions = temp.MajorVersions
	return nil
}

// tempResponseSwitchMetricsVersionComplianceDetails is a temporary struct used for validating the fields of ResponseSwitchMetricsVersionComplianceDetails.
type tempResponseSwitchMetricsVersionComplianceDetails struct {
	MajorVersions []SwitchMetricsComplianceMajorVersion `json:"major_versions,omitempty"`
}
