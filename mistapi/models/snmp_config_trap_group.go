package models

import (
    "encoding/json"
)

// SnmpConfigTrapGroup represents a SnmpConfigTrapGroup struct.
type SnmpConfigTrapGroup struct {
    Categories           []string                  `json:"categories,omitempty"`
    // Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html
    GroupName            *string                   `json:"group_name,omitempty"`
    Targets              []string                  `json:"targets,omitempty"`
    // enum: `all`, `v1`, `v2`
    Version              *SnmpConfigTrapVerionEnum `json:"version,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigTrapGroup.
// It customizes the JSON marshaling process for SnmpConfigTrapGroup objects.
func (s SnmpConfigTrapGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigTrapGroup object to a map representation for JSON marshaling.
func (s SnmpConfigTrapGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Categories != nil {
        structMap["categories"] = s.Categories
    }
    if s.GroupName != nil {
        structMap["group_name"] = s.GroupName
    }
    if s.Targets != nil {
        structMap["targets"] = s.Targets
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfigTrapGroup.
// It customizes the JSON unmarshaling process for SnmpConfigTrapGroup objects.
func (s *SnmpConfigTrapGroup) UnmarshalJSON(input []byte) error {
    var temp snmpConfigTrapGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "categories", "group_name", "targets", "version")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Categories = temp.Categories
    s.GroupName = temp.GroupName
    s.Targets = temp.Targets
    s.Version = temp.Version
    return nil
}

// snmpConfigTrapGroup is a temporary struct used for validating the fields of SnmpConfigTrapGroup.
type snmpConfigTrapGroup  struct {
    Categories []string                  `json:"categories,omitempty"`
    GroupName  *string                   `json:"group_name,omitempty"`
    Targets    []string                  `json:"targets,omitempty"`
    Version    *SnmpConfigTrapVerionEnum `json:"version,omitempty"`
}
