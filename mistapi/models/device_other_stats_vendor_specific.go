package models

import (
    "encoding/json"
)

// DeviceOtherStatsVendorSpecific represents a DeviceOtherStatsVendorSpecific struct.
// when `vendor`==`cradlepoint`
type DeviceOtherStatsVendorSpecific struct {
    Ports                map[string]DeviceOtherStatsVendorSpecificPort `json:"ports,omitempty"`
    TargetVersion        *string                                       `json:"target_version,omitempty"`
    AdditionalProperties map[string]any                                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceOtherStatsVendorSpecific.
// It customizes the JSON marshaling process for DeviceOtherStatsVendorSpecific objects.
func (d DeviceOtherStatsVendorSpecific) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceOtherStatsVendorSpecific object to a map representation for JSON marshaling.
func (d DeviceOtherStatsVendorSpecific) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Ports != nil {
        structMap["ports"] = d.Ports
    }
    if d.TargetVersion != nil {
        structMap["target_version"] = d.TargetVersion
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceOtherStatsVendorSpecific.
// It customizes the JSON unmarshaling process for DeviceOtherStatsVendorSpecific objects.
func (d *DeviceOtherStatsVendorSpecific) UnmarshalJSON(input []byte) error {
    var temp tempDeviceOtherStatsVendorSpecific
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ports", "target_version")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Ports = temp.Ports
    d.TargetVersion = temp.TargetVersion
    return nil
}

// tempDeviceOtherStatsVendorSpecific is a temporary struct used for validating the fields of DeviceOtherStatsVendorSpecific.
type tempDeviceOtherStatsVendorSpecific  struct {
    Ports         map[string]DeviceOtherStatsVendorSpecificPort `json:"ports,omitempty"`
    TargetVersion *string                                       `json:"target_version,omitempty"`
}
