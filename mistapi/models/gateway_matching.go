package models

import (
	"encoding/json"
)

// GatewayMatching represents a GatewayMatching struct.
// Gateway matching
type GatewayMatching struct {
	Enable               *bool                 `json:"enable,omitempty"`
	Rules                []GatewayMatchingRule `json:"rules,omitempty"`
	AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayMatching.
// It customizes the JSON marshaling process for GatewayMatching objects.
func (g GatewayMatching) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMatching object to a map representation for JSON marshaling.
func (g GatewayMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Enable != nil {
		structMap["enable"] = g.Enable
	}
	if g.Rules != nil {
		structMap["rules"] = g.Rules
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMatching.
// It customizes the JSON unmarshaling process for GatewayMatching objects.
func (g *GatewayMatching) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMatching
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enable", "rules")
	if err != nil {
		return err
	}

	g.AdditionalProperties = additionalProperties
	g.Enable = temp.Enable
	g.Rules = temp.Rules
	return nil
}

// tempGatewayMatching is a temporary struct used for validating the fields of GatewayMatching.
type tempGatewayMatching struct {
	Enable *bool                 `json:"enable,omitempty"`
	Rules  []GatewayMatchingRule `json:"rules,omitempty"`
}
