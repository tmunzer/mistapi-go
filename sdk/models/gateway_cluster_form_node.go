package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayClusterFormNode represents a GatewayClusterFormNode struct.
type GatewayClusterFormNode struct {
    // when replacing a noce, either mac has to remain the same as existing cluster
    Mac                  string         `json:"mac"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayClusterFormNode.
// It customizes the JSON marshaling process for GatewayClusterFormNode objects.
func (g GatewayClusterFormNode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayClusterFormNode object to a map representation for JSON marshaling.
func (g GatewayClusterFormNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["mac"] = g.Mac
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayClusterFormNode.
// It customizes the JSON unmarshaling process for GatewayClusterFormNode objects.
func (g *GatewayClusterFormNode) UnmarshalJSON(input []byte) error {
    var temp gatewayClusterFormNode
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

// gatewayClusterFormNode is a temporary struct used for validating the fields of GatewayClusterFormNode.
type gatewayClusterFormNode  struct {
    Mac *string `json:"mac"`
}

func (g *gatewayClusterFormNode) validate() error {
    var errs []string
    if g.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Gateway_Cluster_Form_Node`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
