// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayComplianceMajorVersionProperties represents a GatewayComplianceMajorVersionProperties struct.
type GatewayComplianceMajorVersionProperties struct {
    MajorCount           *int                   `json:"major_count,omitempty"`
    MajorVersion         *string                `json:"major_version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayComplianceMajorVersionProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayComplianceMajorVersionProperties) String() string {
    return fmt.Sprintf(
    	"GatewayComplianceMajorVersionProperties[MajorCount=%v, MajorVersion=%v, AdditionalProperties=%v]",
    	g.MajorCount, g.MajorVersion, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayComplianceMajorVersionProperties.
// It customizes the JSON marshaling process for GatewayComplianceMajorVersionProperties objects.
func (g GatewayComplianceMajorVersionProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "major_count", "major_version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayComplianceMajorVersionProperties object to a map representation for JSON marshaling.
func (g GatewayComplianceMajorVersionProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.MajorCount != nil {
        structMap["major_count"] = g.MajorCount
    }
    if g.MajorVersion != nil {
        structMap["major_version"] = g.MajorVersion
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayComplianceMajorVersionProperties.
// It customizes the JSON unmarshaling process for GatewayComplianceMajorVersionProperties objects.
func (g *GatewayComplianceMajorVersionProperties) UnmarshalJSON(input []byte) error {
    var temp tempGatewayComplianceMajorVersionProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "major_count", "major_version")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.MajorCount = temp.MajorCount
    g.MajorVersion = temp.MajorVersion
    return nil
}

// tempGatewayComplianceMajorVersionProperties is a temporary struct used for validating the fields of GatewayComplianceMajorVersionProperties.
type tempGatewayComplianceMajorVersionProperties  struct {
    MajorCount   *int    `json:"major_count,omitempty"`
    MajorVersion *string `json:"major_version,omitempty"`
}
