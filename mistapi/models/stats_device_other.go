package models

import (
    "encoding/json"
    "fmt"
)

// StatsDeviceOther represents a StatsDeviceOther struct.
type StatsDeviceOther struct {
    CachedStats          *bool                                      `json:"cached_stats,omitempty"`
    ConfigStatus         *string                                    `json:"config_status,omitempty"`
    // Property key is the connected device MAC Address
    ConnectedDevices     map[string]StatsDeviceOtherConnectedDevice `json:"connected_devices,omitempty"`
    // Property key is the interface name
    Interfaces           map[string]StatsDeviceOtherInterface       `json:"interfaces,omitempty"`
    LastConfig           *int                                       `json:"last_config,omitempty"`
    // Last seen timestamp
    LastSeen             Optional[float64]                          `json:"last_seen"`
    LldpEnabled          *bool                                      `json:"lldp_enabled,omitempty"`
    Mac                  *string                                    `json:"mac,omitempty"`
    Status               *string                                    `json:"status,omitempty"`
    Uptime               *int                                       `json:"uptime,omitempty"`
    Vendor               *string                                    `json:"vendor,omitempty"`
    // When `vendor`==`cradlepoint`
    VendorSpecific       *StatsDeviceOtherVendorSpecific            `json:"vendor_specific,omitempty"`
    Version              *string                                    `json:"version,omitempty"`
    AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for StatsDeviceOther,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsDeviceOther) String() string {
    return fmt.Sprintf(
    	"StatsDeviceOther[CachedStats=%v, ConfigStatus=%v, ConnectedDevices=%v, Interfaces=%v, LastConfig=%v, LastSeen=%v, LldpEnabled=%v, Mac=%v, Status=%v, Uptime=%v, Vendor=%v, VendorSpecific=%v, Version=%v, AdditionalProperties=%v]",
    	s.CachedStats, s.ConfigStatus, s.ConnectedDevices, s.Interfaces, s.LastConfig, s.LastSeen, s.LldpEnabled, s.Mac, s.Status, s.Uptime, s.Vendor, s.VendorSpecific, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOther.
// It customizes the JSON marshaling process for StatsDeviceOther objects.
func (s StatsDeviceOther) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "cached_stats", "config_status", "connected_devices", "interfaces", "last_config", "last_seen", "lldp_enabled", "mac", "status", "uptime", "vendor", "vendor_specific", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOther object to a map representation for JSON marshaling.
func (s StatsDeviceOther) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CachedStats != nil {
        structMap["cached_stats"] = s.CachedStats
    }
    if s.ConfigStatus != nil {
        structMap["config_status"] = s.ConfigStatus
    }
    if s.ConnectedDevices != nil {
        structMap["connected_devices"] = s.ConnectedDevices
    }
    if s.Interfaces != nil {
        structMap["interfaces"] = s.Interfaces
    }
    if s.LastConfig != nil {
        structMap["last_config"] = s.LastConfig
    }
    if s.LastSeen.IsValueSet() {
        if s.LastSeen.Value() != nil {
            structMap["last_seen"] = s.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if s.LldpEnabled != nil {
        structMap["lldp_enabled"] = s.LldpEnabled
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.Vendor != nil {
        structMap["vendor"] = s.Vendor
    }
    if s.VendorSpecific != nil {
        structMap["vendor_specific"] = s.VendorSpecific.toMap()
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDeviceOther.
// It customizes the JSON unmarshaling process for StatsDeviceOther objects.
func (s *StatsDeviceOther) UnmarshalJSON(input []byte) error {
    var temp tempStatsDeviceOther
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cached_stats", "config_status", "connected_devices", "interfaces", "last_config", "last_seen", "lldp_enabled", "mac", "status", "uptime", "vendor", "vendor_specific", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CachedStats = temp.CachedStats
    s.ConfigStatus = temp.ConfigStatus
    s.ConnectedDevices = temp.ConnectedDevices
    s.Interfaces = temp.Interfaces
    s.LastConfig = temp.LastConfig
    s.LastSeen = temp.LastSeen
    s.LldpEnabled = temp.LldpEnabled
    s.Mac = temp.Mac
    s.Status = temp.Status
    s.Uptime = temp.Uptime
    s.Vendor = temp.Vendor
    s.VendorSpecific = temp.VendorSpecific
    s.Version = temp.Version
    return nil
}

// tempStatsDeviceOther is a temporary struct used for validating the fields of StatsDeviceOther.
type tempStatsDeviceOther  struct {
    CachedStats      *bool                                      `json:"cached_stats,omitempty"`
    ConfigStatus     *string                                    `json:"config_status,omitempty"`
    ConnectedDevices map[string]StatsDeviceOtherConnectedDevice `json:"connected_devices,omitempty"`
    Interfaces       map[string]StatsDeviceOtherInterface       `json:"interfaces,omitempty"`
    LastConfig       *int                                       `json:"last_config,omitempty"`
    LastSeen         Optional[float64]                          `json:"last_seen"`
    LldpEnabled      *bool                                      `json:"lldp_enabled,omitempty"`
    Mac              *string                                    `json:"mac,omitempty"`
    Status           *string                                    `json:"status,omitempty"`
    Uptime           *int                                       `json:"uptime,omitempty"`
    Vendor           *string                                    `json:"vendor,omitempty"`
    VendorSpecific   *StatsDeviceOtherVendorSpecific            `json:"vendor_specific,omitempty"`
    Version          *string                                    `json:"version,omitempty"`
}
