package models

import (
    "encoding/json"
)

// SynthetictestInfo represents a SynthetictestInfo struct.
type SynthetictestInfo struct {
    By                   *string                          `json:"by,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    DeviceType           *SynthetictestInfoDeviceTypeEnum `json:"device_type,omitempty"`
    Failed               *bool                            `json:"failed,omitempty"`
    Latency              *int                             `json:"latency,omitempty"`
    Mac                  *string                          `json:"mac,omitempty"`
    PortId               *string                          `json:"port_id,omitempty"`
    Reason               *string                          `json:"reason,omitempty"`
    RxMbps               *int                             `json:"rx_mbps,omitempty"`
    StartTime            *int                             `json:"start_time,omitempty"`
    Status               *string                          `json:"status,omitempty"`
    Timestamp            *float64                         `json:"timestamp,omitempty"`
    TxMbps               *int                             `json:"tx_mbps,omitempty"`
    // enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest`
    Type                 *SynthetictestTypeEnum           `json:"type,omitempty"`
    VlanId               *int                             `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestInfo.
// It customizes the JSON marshaling process for SynthetictestInfo objects.
func (s SynthetictestInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestInfo object to a map representation for JSON marshaling.
func (s SynthetictestInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.By != nil {
        structMap["by"] = s.By
    }
    if s.DeviceType != nil {
        structMap["device_type"] = s.DeviceType
    }
    if s.Failed != nil {
        structMap["failed"] = s.Failed
    }
    if s.Latency != nil {
        structMap["latency"] = s.Latency
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
    }
    if s.Reason != nil {
        structMap["reason"] = s.Reason
    }
    if s.RxMbps != nil {
        structMap["rx_mbps"] = s.RxMbps
    }
    if s.StartTime != nil {
        structMap["start_time"] = s.StartTime
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.Timestamp != nil {
        structMap["timestamp"] = s.Timestamp
    }
    if s.TxMbps != nil {
        structMap["tx_mbps"] = s.TxMbps
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestInfo.
// It customizes the JSON unmarshaling process for SynthetictestInfo objects.
func (s *SynthetictestInfo) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "by", "device_type", "failed", "latency", "mac", "port_id", "reason", "rx_mbps", "start_time", "status", "timestamp", "tx_mbps", "type", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.By = temp.By
    s.DeviceType = temp.DeviceType
    s.Failed = temp.Failed
    s.Latency = temp.Latency
    s.Mac = temp.Mac
    s.PortId = temp.PortId
    s.Reason = temp.Reason
    s.RxMbps = temp.RxMbps
    s.StartTime = temp.StartTime
    s.Status = temp.Status
    s.Timestamp = temp.Timestamp
    s.TxMbps = temp.TxMbps
    s.Type = temp.Type
    s.VlanId = temp.VlanId
    return nil
}

// tempSynthetictestInfo is a temporary struct used for validating the fields of SynthetictestInfo.
type tempSynthetictestInfo  struct {
    By         *string                          `json:"by,omitempty"`
    DeviceType *SynthetictestInfoDeviceTypeEnum `json:"device_type,omitempty"`
    Failed     *bool                            `json:"failed,omitempty"`
    Latency    *int                             `json:"latency,omitempty"`
    Mac        *string                          `json:"mac,omitempty"`
    PortId     *string                          `json:"port_id,omitempty"`
    Reason     *string                          `json:"reason,omitempty"`
    RxMbps     *int                             `json:"rx_mbps,omitempty"`
    StartTime  *int                             `json:"start_time,omitempty"`
    Status     *string                          `json:"status,omitempty"`
    Timestamp  *float64                         `json:"timestamp,omitempty"`
    TxMbps     *int                             `json:"tx_mbps,omitempty"`
    Type       *SynthetictestTypeEnum           `json:"type,omitempty"`
    VlanId     *int                             `json:"vlan_id,omitempty"`
}
