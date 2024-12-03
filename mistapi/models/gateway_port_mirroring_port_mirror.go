package models

import (
    "encoding/json"
)

// GatewayPortMirroringPortMirror represents a GatewayPortMirroringPortMirror struct.
type GatewayPortMirroringPortMirror struct {
    FamilyType           *string                `json:"family_type,omitempty"`
    IngressPortIds       []string               `json:"ingress_port_ids,omitempty"`
    OutputPortId         *string                `json:"output_port_id,omitempty"`
    Rate                 *int                   `json:"rate,omitempty"`
    RunLength            *int                   `json:"run_length,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortMirroringPortMirror.
// It customizes the JSON marshaling process for GatewayPortMirroringPortMirror objects.
func (g GatewayPortMirroringPortMirror) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "family_type", "ingress_port_ids", "output_port_id", "rate", "run_length"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortMirroringPortMirror object to a map representation for JSON marshaling.
func (g GatewayPortMirroringPortMirror) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.FamilyType != nil {
        structMap["family_type"] = g.FamilyType
    }
    if g.IngressPortIds != nil {
        structMap["ingress_port_ids"] = g.IngressPortIds
    }
    if g.OutputPortId != nil {
        structMap["output_port_id"] = g.OutputPortId
    }
    if g.Rate != nil {
        structMap["rate"] = g.Rate
    }
    if g.RunLength != nil {
        structMap["run_length"] = g.RunLength
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortMirroringPortMirror.
// It customizes the JSON unmarshaling process for GatewayPortMirroringPortMirror objects.
func (g *GatewayPortMirroringPortMirror) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortMirroringPortMirror
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "family_type", "ingress_port_ids", "output_port_id", "rate", "run_length")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.FamilyType = temp.FamilyType
    g.IngressPortIds = temp.IngressPortIds
    g.OutputPortId = temp.OutputPortId
    g.Rate = temp.Rate
    g.RunLength = temp.RunLength
    return nil
}

// tempGatewayPortMirroringPortMirror is a temporary struct used for validating the fields of GatewayPortMirroringPortMirror.
type tempGatewayPortMirroringPortMirror  struct {
    FamilyType     *string  `json:"family_type,omitempty"`
    IngressPortIds []string `json:"ingress_port_ids,omitempty"`
    OutputPortId   *string  `json:"output_port_id,omitempty"`
    Rate           *int     `json:"rate,omitempty"`
    RunLength      *int     `json:"run_length,omitempty"`
}
