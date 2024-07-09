package models

import (
    "encoding/json"
)

// GatewayMatchingRule represents a GatewayMatchingRule struct.
type GatewayMatchingRule struct {
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds []string                   `json:"additional_config_cmds,omitempty"`
    Name                 *string                    `json:"name,omitempty"`
    PortConfig           map[string]JunosPortConfig `json:"port_config,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayMatchingRule.
// It customizes the JSON marshaling process for GatewayMatchingRule objects.
func (g GatewayMatchingRule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayMatchingRule object to a map representation for JSON marshaling.
func (g GatewayMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = g.AdditionalConfigCmds
    }
    if g.Name != nil {
        structMap["name"] = g.Name
    }
    if g.PortConfig != nil {
        structMap["port_config"] = g.PortConfig
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMatchingRule.
// It customizes the JSON unmarshaling process for GatewayMatchingRule objects.
func (g *GatewayMatchingRule) UnmarshalJSON(input []byte) error {
    var temp gatewayMatchingRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_config_cmds", "name", "port_config")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.AdditionalConfigCmds = temp.AdditionalConfigCmds
    g.Name = temp.Name
    g.PortConfig = temp.PortConfig
    return nil
}

// gatewayMatchingRule is a temporary struct used for validating the fields of GatewayMatchingRule.
type gatewayMatchingRule  struct {
    AdditionalConfigCmds []string                   `json:"additional_config_cmds,omitempty"`
    Name                 *string                    `json:"name,omitempty"`
    PortConfig           map[string]JunosPortConfig `json:"port_config,omitempty"`
}