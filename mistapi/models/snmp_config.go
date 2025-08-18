// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SnmpConfig represents a SnmpConfig struct.
type SnmpConfig struct {
	ClientList  []SnmpConfigClientList `json:"client_list,omitempty"`
	Contact     *string                `json:"contact,omitempty"`
	Description *string                `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	EngineId    *string                `json:"engine_id,omitempty"`
	// enum: `local`, `use_mac_address`
	EngineIdType         *SnmpConfigEngineIdTypeEnum `json:"engine_id_type,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	Network              *string                     `json:"network,omitempty"`
	TrapGroups           []SnmpConfigTrapGroup       `json:"trap_groups,omitempty"`
	V2cConfig            []SnmpConfigV2cConfig       `json:"v2c_config,omitempty"`
	V3Config             *Snmpv3Config               `json:"v3_config,omitempty"`
	Views                []SnmpConfigView            `json:"views,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpConfig) String() string {
	return fmt.Sprintf(
		"SnmpConfig[ClientList=%v, Contact=%v, Description=%v, Enabled=%v, EngineId=%v, EngineIdType=%v, Location=%v, Name=%v, Network=%v, TrapGroups=%v, V2cConfig=%v, V3Config=%v, Views=%v, AdditionalProperties=%v]",
		s.ClientList, s.Contact, s.Description, s.Enabled, s.EngineId, s.EngineIdType, s.Location, s.Name, s.Network, s.TrapGroups, s.V2cConfig, s.V3Config, s.Views, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfig.
// It customizes the JSON marshaling process for SnmpConfig objects.
func (s SnmpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"client_list", "contact", "description", "enabled", "engine_id", "engine_id_type", "location", "name", "network", "trap_groups", "v2c_config", "v3_config", "views"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfig object to a map representation for JSON marshaling.
func (s SnmpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
	if s.EngineIdType != nil {
		structMap["engine_id_type"] = s.EngineIdType
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
	var temp tempSnmpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_list", "contact", "description", "enabled", "engine_id", "engine_id_type", "location", "name", "network", "trap_groups", "v2c_config", "v3_config", "views")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ClientList = temp.ClientList
	s.Contact = temp.Contact
	s.Description = temp.Description
	s.Enabled = temp.Enabled
	s.EngineId = temp.EngineId
	s.EngineIdType = temp.EngineIdType
	s.Location = temp.Location
	s.Name = temp.Name
	s.Network = temp.Network
	s.TrapGroups = temp.TrapGroups
	s.V2cConfig = temp.V2cConfig
	s.V3Config = temp.V3Config
	s.Views = temp.Views
	return nil
}

// tempSnmpConfig is a temporary struct used for validating the fields of SnmpConfig.
type tempSnmpConfig struct {
	ClientList   []SnmpConfigClientList      `json:"client_list,omitempty"`
	Contact      *string                     `json:"contact,omitempty"`
	Description  *string                     `json:"description,omitempty"`
	Enabled      *bool                       `json:"enabled,omitempty"`
	EngineId     *string                     `json:"engine_id,omitempty"`
	EngineIdType *SnmpConfigEngineIdTypeEnum `json:"engine_id_type,omitempty"`
	Location     *string                     `json:"location,omitempty"`
	Name         *string                     `json:"name,omitempty"`
	Network      *string                     `json:"network,omitempty"`
	TrapGroups   []SnmpConfigTrapGroup       `json:"trap_groups,omitempty"`
	V2cConfig    []SnmpConfigV2cConfig       `json:"v2c_config,omitempty"`
	V3Config     *Snmpv3Config               `json:"v3_config,omitempty"`
	Views        []SnmpConfigView            `json:"views,omitempty"`
}
