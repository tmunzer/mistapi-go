package models

import (
    "encoding/json"
)

// StatsDeviceOtherVendorSpecific represents a StatsDeviceOtherVendorSpecific struct.
// when `vendor`==`cradlepoint`
type StatsDeviceOtherVendorSpecific struct {
    Ports                map[string]StatsDeviceOtherVendorSpecificPort `json:"ports,omitempty"`
    TargetVersion        *string                                       `json:"target_version,omitempty"`
    AdditionalProperties map[string]interface{}                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOtherVendorSpecific.
// It customizes the JSON marshaling process for StatsDeviceOtherVendorSpecific objects.
func (s StatsDeviceOtherVendorSpecific) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ports", "target_version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOtherVendorSpecific object to a map representation for JSON marshaling.
func (s StatsDeviceOtherVendorSpecific) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Ports != nil {
        structMap["ports"] = s.Ports
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ports", "target_version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Ports = temp.Ports
    s.TargetVersion = temp.TargetVersion
    return nil
}

// tempStatsDeviceOtherVendorSpecific is a temporary struct used for validating the fields of StatsDeviceOtherVendorSpecific.
type tempStatsDeviceOtherVendorSpecific  struct {
    Ports         map[string]StatsDeviceOtherVendorSpecificPort `json:"ports,omitempty"`
    TargetVersion *string                                       `json:"target_version,omitempty"`
}
