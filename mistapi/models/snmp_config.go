package models

import (
    "encoding/json"
)

// SnmpConfig represents a SnmpConfig struct.
type SnmpConfig struct {
    ClientList           []SnmpConfigClientList  `json:"client_list,omitempty"`
    Contact              *string                 `json:"contact,omitempty"`
    Description          *string                 `json:"description,omitempty"`
    Enabled              *bool                   `json:"enabled,omitempty"`
    // enum: `engine-id-suffix`, `local`, `use-default-ip-address`, `use_mac-address`
    EngineId             *SnmpConfigEngineIdEnum `json:"engine_id,omitempty"`
    Location             *string                 `json:"location,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    Network              *string                 `json:"network,omitempty"`
    TrapGroups           []SnmpConfigTrapGroup   `json:"trap_groups,omitempty"`
    V2cConfig            []SnmpConfigV2CConfig   `json:"v2c_config,omitempty"`
    V3Config             *Snmpv3Config           `json:"v3_config,omitempty"`
    Views                []SnmpConfigView        `json:"views,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfig.
// It customizes the JSON marshaling process for SnmpConfig objects.
func (s SnmpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfig object to a map representation for JSON marshaling.
func (s SnmpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ClientList != nil {
        structMap["client_list"] = s.ClientList
    }
    if s.Contact != nil {
        structMap["contact"] = s.Contact
    }
    if s.Description != nil {
        structMap["description"] = s.Description
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.EngineId != nil {
        structMap["engine_id"] = s.EngineId
    }
    if s.Location != nil {
        structMap["location"] = s.Location
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Network != nil {
        structMap["network"] = s.Network
    }
    if s.TrapGroups != nil {
        structMap["trap_groups"] = s.TrapGroups
    }
    if s.V2cConfig != nil {
        structMap["v2c_config"] = s.V2cConfig
    }
    if s.V3Config != nil {
        structMap["v3_config"] = s.V3Config.toMap()
    }
    if s.Views != nil {
        structMap["views"] = s.Views
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfig.
// It customizes the JSON unmarshaling process for SnmpConfig objects.
func (s *SnmpConfig) UnmarshalJSON(input []byte) error {
    var temp snmpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "client_list", "contact", "description", "enabled", "engine_id", "location", "name", "network", "trap_groups", "v2c_config", "v3_config", "views")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ClientList = temp.ClientList
    s.Contact = temp.Contact
    s.Description = temp.Description
    s.Enabled = temp.Enabled
    s.EngineId = temp.EngineId
    s.Location = temp.Location
    s.Name = temp.Name
    s.Network = temp.Network
    s.TrapGroups = temp.TrapGroups
    s.V2cConfig = temp.V2cConfig
    s.V3Config = temp.V3Config
    s.Views = temp.Views
    return nil
}

// snmpConfig is a temporary struct used for validating the fields of SnmpConfig.
type snmpConfig  struct {
    ClientList  []SnmpConfigClientList  `json:"client_list,omitempty"`
    Contact     *string                 `json:"contact,omitempty"`
    Description *string                 `json:"description,omitempty"`
    Enabled     *bool                   `json:"enabled,omitempty"`
    EngineId    *SnmpConfigEngineIdEnum `json:"engine_id,omitempty"`
    Location    *string                 `json:"location,omitempty"`
    Name        *string                 `json:"name,omitempty"`
    Network     *string                 `json:"network,omitempty"`
    TrapGroups  []SnmpConfigTrapGroup   `json:"trap_groups,omitempty"`
    V2cConfig   []SnmpConfigV2CConfig   `json:"v2c_config,omitempty"`
    V3Config    *Snmpv3Config           `json:"v3_config,omitempty"`
    Views       []SnmpConfigView        `json:"views,omitempty"`
}
