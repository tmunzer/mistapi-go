package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayCluster represents a GatewayCluster struct.
type GatewayCluster struct {
    // when replacing a node, either mac has to remain the same as existing cluster
    Nodes                []GatewayClusterNode `json:"nodes"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayCluster.
// It customizes the JSON marshaling process for GatewayCluster objects.
func (g GatewayCluster) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayCluster object to a map representation for JSON marshaling.
func (g GatewayCluster) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["nodes"] = g.Nodes
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayCluster.
// It customizes the JSON unmarshaling process for GatewayCluster objects.
func (g *GatewayCluster) UnmarshalJSON(input []byte) error {
    var temp tempGatewayCluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "nodes")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Nodes = *temp.Nodes
    return nil
}

// tempGatewayCluster is a temporary struct used for validating the fields of GatewayCluster.
type tempGatewayCluster  struct {
    Nodes *[]GatewayClusterNode `json:"nodes"`
}

func (g *tempGatewayCluster) validate() error {
    var errs []string
    if g.Nodes == nil {
        errs = append(errs, "required field `nodes` is missing for type `gateway_cluster`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
