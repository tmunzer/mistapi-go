package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// GatewayClusterNode represents a GatewayClusterNode struct.
type GatewayClusterNode struct {
	// Gateway MAC Address. Format is `[0-9a-f]{12}` (e.g "5684dae9ac8b")
	Mac                  string         `json:"mac"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayClusterNode.
// It customizes the JSON marshaling process for GatewayClusterNode objects.
func (g GatewayClusterNode) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayClusterNode object to a map representation for JSON marshaling.
func (g GatewayClusterNode) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, g.AdditionalProperties)
	structMap["mac"] = g.Mac
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayClusterNode.
// It customizes the JSON unmarshaling process for GatewayClusterNode objects.
func (g *GatewayClusterNode) UnmarshalJSON(input []byte) error {
	var temp tempGatewayClusterNode
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac")
	if err != nil {
		return err
	}

	g.AdditionalProperties = additionalProperties
	g.Mac = *temp.Mac
	return nil
}

// tempGatewayClusterNode is a temporary struct used for validating the fields of GatewayClusterNode.
type tempGatewayClusterNode struct {
	Mac *string `json:"mac"`
}

func (g *tempGatewayClusterNode) validate() error {
	var errs []string
	if g.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `gateway_cluster_node`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
