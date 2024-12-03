package models

import (
    "encoding/json"
)

// StatsDeviceOther represents a StatsDeviceOther struct.
type StatsDeviceOther struct {
    ConfigStatus         *string                         `json:"config_status,omitempty"`
    LastConfig           *int                            `json:"last_config,omitempty"`
    LastSeen             *int                            `json:"last_seen,omitempty"`
    Mac                  *string                         `json:"mac,omitempty"`
    Status               *string                         `json:"status,omitempty"`
    Uptime               *int                            `json:"uptime,omitempty"`
    Vendor               *string                         `json:"vendor,omitempty"`
    // when `vendor`==`cradlepoint`
    VendorSpecific       *StatsDeviceOtherVendorSpecific `json:"vendor_specific,omitempty"`
    Version              *string                         `json:"version,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsDeviceOther.
// It customizes the JSON marshaling process for StatsDeviceOther objects.
func (s StatsDeviceOther) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "config_status", "last_config", "last_seen", "mac", "status", "uptime", "vendor", "vendor_specific", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDeviceOther object to a map representation for JSON marshaling.
func (s StatsDeviceOther) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ConfigStatus != nil {
        structMap["config_status"] = s.ConfigStatus
    }
    if s.LastConfig != nil {
        structMap["last_config"] = s.LastConfig
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_status", "last_config", "last_seen", "mac", "status", "uptime", "vendor", "vendor_specific", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ConfigStatus = temp.ConfigStatus
    s.LastConfig = temp.LastConfig
    s.LastSeen = temp.LastSeen
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
    ConfigStatus   *string                         `json:"config_status,omitempty"`
    LastConfig     *int                            `json:"last_config,omitempty"`
    LastSeen       *int                            `json:"last_seen,omitempty"`
    Mac            *string                         `json:"mac,omitempty"`
    Status         *string                         `json:"status,omitempty"`
    Uptime         *int                            `json:"uptime,omitempty"`
    Vendor         *string                         `json:"vendor,omitempty"`
    VendorSpecific *StatsDeviceOtherVendorSpecific `json:"vendor_specific,omitempty"`
    Version        *string                         `json:"version,omitempty"`
}
