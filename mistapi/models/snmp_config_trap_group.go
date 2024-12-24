package models

import (
    "encoding/json"
    "fmt"
)

// SnmpConfigTrapGroup represents a SnmpConfigTrapGroup struct.
type SnmpConfigTrapGroup struct {
    Categories           []string                  `json:"categories,omitempty"`
    // Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html
    GroupName            *string                   `json:"group_name,omitempty"`
    Targets              []string                  `json:"targets,omitempty"`
    // enum: `all`, `v1`, `v2`
    Version              *SnmpConfigTrapVerionEnum `json:"version,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpConfigTrapGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpConfigTrapGroup) String() string {
    return fmt.Sprintf(
    	"SnmpConfigTrapGroup[Categories=%v, GroupName=%v, Targets=%v, Version=%v, AdditionalProperties=%v]",
    	s.Categories, s.GroupName, s.Targets, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigTrapGroup.
// It customizes the JSON marshaling process for SnmpConfigTrapGroup objects.
func (s SnmpConfigTrapGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "categories", "group_name", "targets", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigTrapGroup object to a map representation for JSON marshaling.
func (s SnmpConfigTrapGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSnmpConfigTrapGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "categories", "group_name", "targets", "version")
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

// tempSnmpConfigTrapGroup is a temporary struct used for validating the fields of SnmpConfigTrapGroup.
type tempSnmpConfigTrapGroup  struct {
    Categories []string                  `json:"categories,omitempty"`
    GroupName  *string                   `json:"group_name,omitempty"`
    Targets    []string                  `json:"targets,omitempty"`
    Version    *SnmpConfigTrapVerionEnum `json:"version,omitempty"`
}
