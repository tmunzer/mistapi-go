// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// Snmpv3Config represents a Snmpv3Config struct.
type Snmpv3Config struct {
    Notify               []Snmpv3ConfigNotifyItems       `json:"notify,omitempty"`
    NotifyFilter         []Snmpv3ConfigNotifyFilterItem  `json:"notify_filter,omitempty"`
    TargetAddress        []Snmpv3ConfigTargetAddressItem `json:"target_address,omitempty"`
    TargetParameters     []Snmpv3ConfigTargetParam       `json:"target_parameters,omitempty"`
    Usm                  []SnmpUsm                       `json:"usm,omitempty"`
    Vacm                 *SnmpVacm                       `json:"vacm,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for Snmpv3Config,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Snmpv3Config) String() string {
    return fmt.Sprintf(
    	"Snmpv3Config[Notify=%v, NotifyFilter=%v, TargetAddress=%v, TargetParameters=%v, Usm=%v, Vacm=%v, AdditionalProperties=%v]",
    	s.Notify, s.NotifyFilter, s.TargetAddress, s.TargetParameters, s.Usm, s.Vacm, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3Config.
// It customizes the JSON marshaling process for Snmpv3Config objects.
func (s Snmpv3Config) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "notify", "notify_filter", "target_address", "target_parameters", "usm", "vacm"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3Config object to a map representation for JSON marshaling.
func (s Snmpv3Config) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Notify != nil {
        structMap["notify"] = s.Notify
    }
    if s.NotifyFilter != nil {
        structMap["notify_filter"] = s.NotifyFilter
    }
    if s.TargetAddress != nil {
        structMap["target_address"] = s.TargetAddress
    }
    if s.TargetParameters != nil {
        structMap["target_parameters"] = s.TargetParameters
    }
    if s.Usm != nil {
        structMap["usm"] = s.Usm
    }
    if s.Vacm != nil {
        structMap["vacm"] = s.Vacm.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3Config.
// It customizes the JSON unmarshaling process for Snmpv3Config objects.
func (s *Snmpv3Config) UnmarshalJSON(input []byte) error {
    var temp tempSnmpv3Config
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "notify", "notify_filter", "target_address", "target_parameters", "usm", "vacm")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Notify = temp.Notify
    s.NotifyFilter = temp.NotifyFilter
    s.TargetAddress = temp.TargetAddress
    s.TargetParameters = temp.TargetParameters
    s.Usm = temp.Usm
    s.Vacm = temp.Vacm
    return nil
}

// tempSnmpv3Config is a temporary struct used for validating the fields of Snmpv3Config.
type tempSnmpv3Config  struct {
    Notify           []Snmpv3ConfigNotifyItems       `json:"notify,omitempty"`
    NotifyFilter     []Snmpv3ConfigNotifyFilterItem  `json:"notify_filter,omitempty"`
    TargetAddress    []Snmpv3ConfigTargetAddressItem `json:"target_address,omitempty"`
    TargetParameters []Snmpv3ConfigTargetParam       `json:"target_parameters,omitempty"`
    Usm              []SnmpUsm                       `json:"usm,omitempty"`
    Vacm             *SnmpVacm                       `json:"vacm,omitempty"`
}
