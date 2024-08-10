package models

import (
    "encoding/json"
)

// SwitchMatchingRule represents a SwitchMatchingRule struct.
// property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`
type SwitchMatchingRule struct {
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    // role to match
    MatchRole            *string                                `json:"match_role,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    // Propery key is the interface name or interface range
    PortConfig           map[string]JunosPortConfig             `json:"port_config,omitempty"`
    // Property key is the port mirroring instance name
    // port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    // Switch settings
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatchingRule.
// It customizes the JSON marshaling process for SwitchMatchingRule objects.
func (s SwitchMatchingRule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatchingRule object to a map representation for JSON marshaling.
func (s SwitchMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = s.AdditionalConfigCmds
    }
    if s.MatchRole != nil {
        structMap["match_role"] = s.MatchRole
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.PortConfig != nil {
        structMap["port_config"] = s.PortConfig
    }
    if s.PortMirroring != nil {
        structMap["port_mirroring"] = s.PortMirroring
    }
    if s.SwitchMgmt != nil {
        structMap["switch_mgmt"] = s.SwitchMgmt.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMatchingRule.
// It customizes the JSON unmarshaling process for SwitchMatchingRule objects.
func (s *SwitchMatchingRule) UnmarshalJSON(input []byte) error {
    var temp tempSwitchMatchingRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_config_cmds", "match_role", "name", "port_config", "port_mirroring", "switch_mgmt")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AdditionalConfigCmds = temp.AdditionalConfigCmds
    s.MatchRole = temp.MatchRole
    s.Name = temp.Name
    s.PortConfig = temp.PortConfig
    s.PortMirroring = temp.PortMirroring
    s.SwitchMgmt = temp.SwitchMgmt
    return nil
}

// tempSwitchMatchingRule is a temporary struct used for validating the fields of SwitchMatchingRule.
type tempSwitchMatchingRule  struct {
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    MatchRole            *string                                `json:"match_role,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    PortConfig           map[string]JunosPortConfig             `json:"port_config,omitempty"`
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
}
