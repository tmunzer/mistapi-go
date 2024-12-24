package models

import (
    "encoding/json"
    "fmt"
)

// GatewayComplianceVersion represents a GatewayComplianceVersion struct.
// version compliance score, major version for gateway, type
type GatewayComplianceVersion struct {
    MajorVersion         map[string]GatewayComplianceMajorVersionProperties `json:"major_version,omitempty"`
    Score                *float64                                           `json:"score,omitempty"`
    Type                 *string                                            `json:"type,omitempty"`
    AdditionalProperties map[string]interface{}                             `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayComplianceVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayComplianceVersion) String() string {
    return fmt.Sprintf(
    	"GatewayComplianceVersion[MajorVersion=%v, Score=%v, Type=%v, AdditionalProperties=%v]",
    	g.MajorVersion, g.Score, g.Type, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayComplianceVersion.
// It customizes the JSON marshaling process for GatewayComplianceVersion objects.
func (g GatewayComplianceVersion) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "major_version", "score", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayComplianceVersion object to a map representation for JSON marshaling.
func (g GatewayComplianceVersion) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "major_version", "score", "type")
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
type tempGatewayComplianceVersion  struct {
    MajorVersion map[string]GatewayComplianceMajorVersionProperties `json:"major_version,omitempty"`
    Score        *float64                                           `json:"score,omitempty"`
    Type         *string                                            `json:"type,omitempty"`
}
