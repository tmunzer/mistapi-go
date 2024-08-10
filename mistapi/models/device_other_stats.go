package models

import (
    "encoding/json"
)

// DeviceOtherStats represents a DeviceOtherStats struct.
type DeviceOtherStats struct {
    ConfigStatus         *string                         `json:"config_status,omitempty"`
    LastConfig           *int                            `json:"last_config,omitempty"`
    LastSeen             *int                            `json:"last_seen,omitempty"`
    Mac                  *string                         `json:"mac,omitempty"`
    Status               *string                         `json:"status,omitempty"`
    Uptime               *int                            `json:"uptime,omitempty"`
    Vendor               *string                         `json:"vendor,omitempty"`
    // when `vendor`==`cradlepoint`
    VendorSpecific       *DeviceOtherStatsVendorSpecific `json:"vendor_specific,omitempty"`
    Version              *string                         `json:"version,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceOtherStats.
// It customizes the JSON marshaling process for DeviceOtherStats objects.
func (d DeviceOtherStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceOtherStats object to a map representation for JSON marshaling.
func (d DeviceOtherStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.ConfigStatus != nil {
        structMap["config_status"] = d.ConfigStatus
    }
    if d.LastConfig != nil {
        structMap["last_config"] = d.LastConfig
    }
    if d.LastSeen != nil {
        structMap["last_seen"] = d.LastSeen
    }
    if d.Mac != nil {
        structMap["mac"] = d.Mac
    }
    if d.Status != nil {
        structMap["status"] = d.Status
    }
    if d.Uptime != nil {
        structMap["uptime"] = d.Uptime
    }
    if d.Vendor != nil {
        structMap["vendor"] = d.Vendor
    }
    if d.VendorSpecific != nil {
        structMap["vendor_specific"] = d.VendorSpecific.toMap()
    }
    if d.Version != nil {
        structMap["version"] = d.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceOtherStats.
// It customizes the JSON unmarshaling process for DeviceOtherStats objects.
func (d *DeviceOtherStats) UnmarshalJSON(input []byte) error {
    var temp tempDeviceOtherStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "config_status", "last_config", "last_seen", "mac", "status", "uptime", "vendor", "vendor_specific", "version")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.ConfigStatus = temp.ConfigStatus
    d.LastConfig = temp.LastConfig
    d.LastSeen = temp.LastSeen
    d.Mac = temp.Mac
    d.Status = temp.Status
    d.Uptime = temp.Uptime
    d.Vendor = temp.Vendor
    d.VendorSpecific = temp.VendorSpecific
    d.Version = temp.Version
    return nil
}

// tempDeviceOtherStats is a temporary struct used for validating the fields of DeviceOtherStats.
type tempDeviceOtherStats  struct {
    ConfigStatus   *string                         `json:"config_status,omitempty"`
    LastConfig     *int                            `json:"last_config,omitempty"`
    LastSeen       *int                            `json:"last_seen,omitempty"`
    Mac            *string                         `json:"mac,omitempty"`
    Status         *string                         `json:"status,omitempty"`
    Uptime         *int                            `json:"uptime,omitempty"`
    Vendor         *string                         `json:"vendor,omitempty"`
    VendorSpecific *DeviceOtherStatsVendorSpecific `json:"vendor_specific,omitempty"`
    Version        *string                         `json:"version,omitempty"`
}
