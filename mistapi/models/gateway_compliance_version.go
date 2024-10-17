package models

import (
	"encoding/json"
)

// GatewayComplianceVersion represents a GatewayComplianceVersion struct.
// version compliance score, major version for gateway, type
type GatewayComplianceVersion struct {
	MajorVersion         map[string]GatewayComplianceMajorVersionProperties `json:"major_version,omitempty"`
	Score                *float64                                           `json:"score,omitempty"`
	Type                 *string                                            `json:"type,omitempty"`
	AdditionalProperties map[string]any                                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayComplianceVersion.
// It customizes the JSON marshaling process for GatewayComplianceVersion objects.
func (g GatewayComplianceVersion) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayComplianceVersion object to a map representation for JSON marshaling.
func (g GatewayComplianceVersion) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, g.AdditionalProperties)
	if g.MajorVersion != nil {
		structMap["major_version"] = g.MajorVersion
	}
	if g.Score != nil {
		structMap["score"] = g.Score
	}
	if g.Type != nil {
		structMap["type"] = g.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayComplianceVersion.
// It customizes the JSON unmarshaling process for GatewayComplianceVersion objects.
func (g *GatewayComplianceVersion) UnmarshalJSON(input []byte) error {
	var temp tempGatewayComplianceVersion
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "major_version", "score", "type")
	if err != nil {
		return err
	}

	g.AdditionalProperties = additionalProperties
	g.MajorVersion = temp.MajorVersion
	g.Score = temp.Score
	g.Type = temp.Type
	return nil
}

// tempGatewayComplianceVersion is a temporary struct used for validating the fields of GatewayComplianceVersion.
type tempGatewayComplianceVersion struct {
	MajorVersion map[string]GatewayComplianceMajorVersionProperties `json:"major_version,omitempty"`
	Score        *float64                                           `json:"score,omitempty"`
	Type         *string                                            `json:"type,omitempty"`
}
