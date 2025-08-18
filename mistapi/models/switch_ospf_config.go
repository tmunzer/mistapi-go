// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchOspfConfig represents a SwitchOspfConfig struct.
type SwitchOspfConfig struct {
    // Property key is the area name. Defines the OSPF areas configured on the switch.
    Areas                map[string]SwitchOspfConfigArea     `json:"areas,omitempty"`
    // Enable OSPF on the switch
    Enabled              *bool                               `json:"enabled,omitempty"`
    // optional, for basic scenario, `import_policy` can be specified and can be applied to all networks in all areas if not explicitly specified
    ExportPolicy         *string                             `json:"export_policy,omitempty"`
    // optional, for basic scenario, `import_policy` can be specified and can be applied to all networks in all areas if not explicitly specified
    ImportPolicy         *string                             `json:"import_policy,omitempty"`
    // Reference bandwidth. Integer(100000) or String (10g)
    ReferenceBandwidth   *SwitchOspfConfigReferenceBandwidth `json:"reference_bandwidth,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchOspfConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchOspfConfig) String() string {
    return fmt.Sprintf(
    	"SwitchOspfConfig[Areas=%v, Enabled=%v, ExportPolicy=%v, ImportPolicy=%v, ReferenceBandwidth=%v, AdditionalProperties=%v]",
    	s.Areas, s.Enabled, s.ExportPolicy, s.ImportPolicy, s.ReferenceBandwidth, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchOspfConfig.
// It customizes the JSON marshaling process for SwitchOspfConfig objects.
func (s SwitchOspfConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "areas", "enabled", "export_policy", "import_policy", "reference_bandwidth"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchOspfConfig object to a map representation for JSON marshaling.
func (s SwitchOspfConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Areas != nil {
        structMap["areas"] = s.Areas
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.ExportPolicy != nil {
        structMap["export_policy"] = s.ExportPolicy
    }
    if s.ImportPolicy != nil {
        structMap["import_policy"] = s.ImportPolicy
    }
    if s.ReferenceBandwidth != nil {
        structMap["reference_bandwidth"] = s.ReferenceBandwidth.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchOspfConfig.
// It customizes the JSON unmarshaling process for SwitchOspfConfig objects.
func (s *SwitchOspfConfig) UnmarshalJSON(input []byte) error {
    var temp tempSwitchOspfConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "areas", "enabled", "export_policy", "import_policy", "reference_bandwidth")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Areas = temp.Areas
    s.Enabled = temp.Enabled
    s.ExportPolicy = temp.ExportPolicy
    s.ImportPolicy = temp.ImportPolicy
    s.ReferenceBandwidth = temp.ReferenceBandwidth
    return nil
}

// tempSwitchOspfConfig is a temporary struct used for validating the fields of SwitchOspfConfig.
type tempSwitchOspfConfig  struct {
    Areas              map[string]SwitchOspfConfigArea     `json:"areas,omitempty"`
    Enabled            *bool                               `json:"enabled,omitempty"`
    ExportPolicy       *string                             `json:"export_policy,omitempty"`
    ImportPolicy       *string                             `json:"import_policy,omitempty"`
    ReferenceBandwidth *SwitchOspfConfigReferenceBandwidth `json:"reference_bandwidth,omitempty"`
}
