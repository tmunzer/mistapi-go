package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// WiredClientResponse represents a WiredClientResponse struct.
type WiredClientResponse struct {
    DeviceMac            []string                               `json:"device_mac,omitempty"`
    DeviceMacPort        []WiredClientResponseDeviceMacPortItem `json:"device_mac_port,omitempty"`
    Ip                   []string                               `json:"ip,omitempty"`
    Mac                  *string                                `json:"mac,omitempty"`
    OrgId                *uuid.UUID                             `json:"org_id,omitempty"`
    PortId               []string                               `json:"port_id,omitempty"`
    SiteId               *uuid.UUID                             `json:"site_id,omitempty"`
    Timestamp            *float64                               `json:"timestamp,omitempty"`
    Vlan                 []int                                  `json:"vlan,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WiredClientResponse.
// It customizes the JSON marshaling process for WiredClientResponse objects.
func (w WiredClientResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WiredClientResponse object to a map representation for JSON marshaling.
func (w WiredClientResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.DeviceMac != nil {
        structMap["device_mac"] = w.DeviceMac
    }
    if w.DeviceMacPort != nil {
        structMap["device_mac_port"] = w.DeviceMacPort
    }
    if w.Ip != nil {
        structMap["ip"] = w.Ip
    }
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.PortId != nil {
        structMap["port_id"] = w.PortId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.Vlan != nil {
        structMap["vlan"] = w.Vlan
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WiredClientResponse.
// It customizes the JSON unmarshaling process for WiredClientResponse objects.
func (w *WiredClientResponse) UnmarshalJSON(input []byte) error {
    var temp wiredClientResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "device_mac_port", "ip", "mac", "org_id", "port_id", "site_id", "timestamp", "vlan")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.DeviceMac = temp.DeviceMac
    w.DeviceMacPort = temp.DeviceMacPort
    w.Ip = temp.Ip
    w.Mac = temp.Mac
    w.OrgId = temp.OrgId
    w.PortId = temp.PortId
    w.SiteId = temp.SiteId
    w.Timestamp = temp.Timestamp
    w.Vlan = temp.Vlan
    return nil
}

// wiredClientResponse is a temporary struct used for validating the fields of WiredClientResponse.
type wiredClientResponse  struct {
    DeviceMac     []string                               `json:"device_mac,omitempty"`
    DeviceMacPort []WiredClientResponseDeviceMacPortItem `json:"device_mac_port,omitempty"`
    Ip            []string                               `json:"ip,omitempty"`
    Mac           *string                                `json:"mac,omitempty"`
    OrgId         *uuid.UUID                             `json:"org_id,omitempty"`
    PortId        []string                               `json:"port_id,omitempty"`
    SiteId        *uuid.UUID                             `json:"site_id,omitempty"`
    Timestamp     *float64                               `json:"timestamp,omitempty"`
    Vlan          []int                                  `json:"vlan,omitempty"`
}
