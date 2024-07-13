package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayClusterSwap represents a GatewayClusterSwap struct.
type GatewayClusterSwap struct {
    // when `op` ==`replacement_nodeX`, new node1
    // 's MAC, the device has to be standalone and assigned to the same site
    Mac                  *string                  `json:"mac,omitempty"`
    Op                   GatewayClusterSwapOpEnum `json:"op"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayClusterSwap.
// It customizes the JSON marshaling process for GatewayClusterSwap objects.
func (g GatewayClusterSwap) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayClusterSwap object to a map representation for JSON marshaling.
func (g GatewayClusterSwap) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Mac != nil {
        structMap["mac"] = g.Mac
    }
    structMap["op"] = g.Op
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayClusterSwap.
// It customizes the JSON unmarshaling process for GatewayClusterSwap objects.
func (g *GatewayClusterSwap) UnmarshalJSON(input []byte) error {
    var temp gatewayClusterSwap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "op")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Mac = temp.Mac
    g.Op = *temp.Op
    return nil
}

// gatewayClusterSwap is a temporary struct used for validating the fields of GatewayClusterSwap.
type gatewayClusterSwap  struct {
    Mac *string                   `json:"mac,omitempty"`
    Op  *GatewayClusterSwapOpEnum `json:"op"`
}

func (g *gatewayClusterSwap) validate() error {
    var errs []string
    if g.Op == nil {
        errs = append(errs, "required field `op` is missing for type `Gateway_Cluster_Swap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
