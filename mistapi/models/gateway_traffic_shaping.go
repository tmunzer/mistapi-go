package models

import (
    "encoding/json"
    "fmt"
)

// GatewayTrafficShaping represents a GatewayTrafficShaping struct.
type GatewayTrafficShaping struct {
    // percentages for differet class of traffic: high / medium / low / best-effort
    // sum must be equal to 100
    ClassPercentages     []int                  `json:"class_percentages,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayTrafficShaping,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayTrafficShaping) String() string {
    return fmt.Sprintf(
    	"GatewayTrafficShaping[ClassPercentages=%v, Enabled=%v, AdditionalProperties=%v]",
    	g.ClassPercentages, g.Enabled, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayTrafficShaping.
// It customizes the JSON marshaling process for GatewayTrafficShaping objects.
func (g GatewayTrafficShaping) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "class_percentages", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayTrafficShaping object to a map representation for JSON marshaling.
func (g GatewayTrafficShaping) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.ClassPercentages != nil {
        structMap["class_percentages"] = g.ClassPercentages
    }
    if g.Enabled != nil {
        structMap["enabled"] = g.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayTrafficShaping.
// It customizes the JSON unmarshaling process for GatewayTrafficShaping objects.
func (g *GatewayTrafficShaping) UnmarshalJSON(input []byte) error {
    var temp tempGatewayTrafficShaping
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "class_percentages", "enabled")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.ClassPercentages = temp.ClassPercentages
    g.Enabled = temp.Enabled
    return nil
}

// tempGatewayTrafficShaping is a temporary struct used for validating the fields of GatewayTrafficShaping.
type tempGatewayTrafficShaping  struct {
    ClassPercentages []int `json:"class_percentages,omitempty"`
    Enabled          *bool `json:"enabled,omitempty"`
}
