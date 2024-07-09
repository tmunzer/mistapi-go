package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayClusterForm represents a GatewayClusterForm struct.
type GatewayClusterForm struct {
    Nodes                []GatewayClusterFormNode `json:"nodes"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayClusterForm.
// It customizes the JSON marshaling process for GatewayClusterForm objects.
func (g GatewayClusterForm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayClusterForm object to a map representation for JSON marshaling.
func (g GatewayClusterForm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["nodes"] = g.Nodes
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayClusterForm.
// It customizes the JSON unmarshaling process for GatewayClusterForm objects.
func (g *GatewayClusterForm) UnmarshalJSON(input []byte) error {
    var temp gatewayClusterForm
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

// gatewayClusterForm is a temporary struct used for validating the fields of GatewayClusterForm.
type gatewayClusterForm  struct {
    Nodes *[]GatewayClusterFormNode `json:"nodes"`
}

func (g *gatewayClusterForm) validate() error {
    var errs []string
    if g.Nodes == nil {
        errs = append(errs, "required field `nodes` is missing for type `Gateway_Cluster_Form`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
