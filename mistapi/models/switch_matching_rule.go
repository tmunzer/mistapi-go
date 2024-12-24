package models

import (
    "encoding/json"
    "fmt"
)

// SwitchMatchingRule represents a SwitchMatchingRule struct.
// property key defines the type of matching, value is the string to match. e.g:
// * `match_name[0:3]`: switch name must match the first 3 letters of the property value
// * `match_name[2:6]`: switch name must match the property value from the 2nd to the 6th letter
// * `match_model[0-8]`: switch model must match the first 8 letters of the property value
// * `match_role`: switch role must match the property value
type SwitchMatchingRule struct {
    // additional CLI commands to append to the generated Junos config. **Note**: no check is done
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    // In-Band Management interface configuration
    IpConfig             *SwitchMatchingRuleIpConfig            `json:"ip_config,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    // Out-of-Band Management interface configuration
    OobIpConfig          *SwitchMatchingRuleOobIpConfig         `json:"oob_ip_config,omitempty"`
    // Propery key is the interface name or interface range
    PortConfig           map[string]JunosPortConfig             `json:"port_config,omitempty"`
    // Property key is the port mirroring instance name
    // port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 port mirrorings is allowed
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    // Switch settings
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    AdditionalProperties map[string]string                      `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMatchingRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMatchingRule) String() string {
    return fmt.Sprintf(
    	"SwitchMatchingRule[AdditionalConfigCmds=%v, IpConfig=%v, Name=%v, OobIpConfig=%v, PortConfig=%v, PortMirroring=%v, SwitchMgmt=%v, AdditionalProperties=%v]",
    	s.AdditionalConfigCmds, s.IpConfig, s.Name, s.OobIpConfig, s.PortConfig, s.PortMirroring, s.SwitchMgmt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMatchingRule.
// It customizes the JSON marshaling process for SwitchMatchingRule objects.
func (s SwitchMatchingRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "additional_config_cmds", "ip_config", "name", "oob_ip_config", "port_config", "port_mirroring", "switch_mgmt"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMatchingRule object to a map representation for JSON marshaling.
func (s SwitchMatchingRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = s.AdditionalConfigCmds
    }
    if s.IpConfig != nil {
        structMap["ip_config"] = s.IpConfig.toMap()
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.OobIpConfig != nil {
        structMap["oob_ip_config"] = s.OobIpConfig.toMap()
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
    additionalProperties, err := ExtractAdditionalProperties[string](input, "additional_config_cmds", "ip_config", "name", "oob_ip_config", "port_config", "port_mirroring", "switch_mgmt")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AdditionalConfigCmds = temp.AdditionalConfigCmds
    s.IpConfig = temp.IpConfig
    s.Name = temp.Name
    s.OobIpConfig = temp.OobIpConfig
    s.PortConfig = temp.PortConfig
    s.PortMirroring = temp.PortMirroring
    s.SwitchMgmt = temp.SwitchMgmt
    return nil
}

// tempSwitchMatchingRule is a temporary struct used for validating the fields of SwitchMatchingRule.
type tempSwitchMatchingRule  struct {
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    IpConfig             *SwitchMatchingRuleIpConfig            `json:"ip_config,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    OobIpConfig          *SwitchMatchingRuleOobIpConfig         `json:"oob_ip_config,omitempty"`
    PortConfig           map[string]JunosPortConfig             `json:"port_config,omitempty"`
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
}
