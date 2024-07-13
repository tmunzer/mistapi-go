package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayClusterNode represents a GatewayClusterNode struct.
type GatewayClusterNode struct {
    // when replacing a noce, either mac has to remain the same as existing cluster
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
    var temp gatewayClusterNode
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

// gatewayClusterNode is a temporary struct used for validating the fields of GatewayClusterNode.
type gatewayClusterNode  struct {
    Mac *string `json:"mac"`
}

func (g *gatewayClusterNode) validate() error {
    var errs []string
    if g.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Gateway_Cluster_Node`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
