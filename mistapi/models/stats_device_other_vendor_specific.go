// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsDeviceOtherVendorSpecific represents a StatsDeviceOtherVendorSpecific struct.
// When `vendor`==`cradlepoint`
type StatsDeviceOtherVendorSpecific struct {
    Interfaces           map[string]StatsDeviceOtherVendorSpecificPort `json:"interfaces,omitempty"`
    TargetVersion        *string                                       `json:"target_version,omitempty"`
    AdditionalProperties map[string]interface{}                        `json:"_"`
}

// String implements the fmt.Stringer interface for StatsDeviceOtherVendorSpecific,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsDeviceOtherVendorSpecific) String() string {
    return fmt.Sprintf(
    	"StatsDeviceOtherVendorSpecific[Interfaces=%v, TargetVersion=%v, AdditionalProperties=%v]",
    	s.Interfaces, s.TargetVersion, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOtherVendorSpecific.
// It customizes the JSON marshaling process for StatsDeviceOtherVendorSpecific objects.
func (s StatsDeviceOtherVendorSpecific) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "interfaces", "target_version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOtherVendorSpecific object to a map representation for JSON marshaling.
func (s StatsDeviceOtherVendorSpecific) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Interfaces != nil {
        structMap["interfaces"] = s.Interfaces
    }
    if s.TargetVersion != nil {
        structMap["target_version"] = s.TargetVersion
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDeviceOtherVendorSpecific.
// It customizes the JSON unmarshaling process for StatsDeviceOtherVendorSpecific objects.
func (s *StatsDeviceOtherVendorSpecific) UnmarshalJSON(input []byte) error {
    var temp tempStatsDeviceOtherVendorSpecific
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interfaces", "target_version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Interfaces = temp.Interfaces
    s.TargetVersion = temp.TargetVersion
    return nil
}

// tempStatsDeviceOtherVendorSpecific is a temporary struct used for validating the fields of StatsDeviceOtherVendorSpecific.
type tempStatsDeviceOtherVendorSpecific  struct {
    Interfaces    map[string]StatsDeviceOtherVendorSpecificPort `json:"interfaces,omitempty"`
    TargetVersion *string                                       `json:"target_version,omitempty"`
}
