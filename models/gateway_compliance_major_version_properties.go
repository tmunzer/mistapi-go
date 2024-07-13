package models

import (
    "encoding/json"
)

// GatewayComplianceMajorVersionProperties represents a GatewayComplianceMajorVersionProperties struct.
type GatewayComplianceMajorVersionProperties struct {
    MajorCount           *int           `json:"major_count,omitempty"`
    MajorVersion         *string        `json:"major_version,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayComplianceMajorVersionProperties.
// It customizes the JSON marshaling process for GatewayComplianceMajorVersionProperties objects.
func (g GatewayComplianceMajorVersionProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayComplianceMajorVersionProperties object to a map representation for JSON marshaling.
func (g GatewayComplianceMajorVersionProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
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
    var temp gatewayComplianceMajorVersionProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "major_count", "major_version")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.MajorCount = temp.MajorCount
    g.MajorVersion = temp.MajorVersion
    return nil
}

// gatewayComplianceMajorVersionProperties is a temporary struct used for validating the fields of GatewayComplianceMajorVersionProperties.
type gatewayComplianceMajorVersionProperties  struct {
    MajorCount   *int    `json:"major_count,omitempty"`
    MajorVersion *string `json:"major_version,omitempty"`
}
