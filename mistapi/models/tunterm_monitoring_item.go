package models

import (
    "encoding/json"
)

// TuntermMonitoringItem represents a TuntermMonitoringItem struct.
type TuntermMonitoringItem struct {
    // can be ip, ipv6, hostname
    Host                 *string                        `json:"host,omitempty"`
    // when `protocol`==`tcp`
    Port                 *int                           `json:"port,omitempty"`
    // enum: `arp`, `ping`, `tcp`
    Protocol             *TunternMonitoringProtocolEnum `json:"protocol,omitempty"`
    Timeout              *int                           `json:"timeout,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TuntermMonitoringItem.
// It customizes the JSON marshaling process for TuntermMonitoringItem objects.
func (t TuntermMonitoringItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TuntermMonitoringItem object to a map representation for JSON marshaling.
func (t TuntermMonitoringItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Host != nil {
        structMap["host"] = t.Host
    }
    if t.Port != nil {
        structMap["port"] = t.Port
    }
    if t.Protocol != nil {
        structMap["protocol"] = t.Protocol
    }
    if t.Timeout != nil {
        structMap["timeout"] = t.Timeout
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermMonitoringItem.
// It customizes the JSON unmarshaling process for TuntermMonitoringItem objects.
func (t *TuntermMonitoringItem) UnmarshalJSON(input []byte) error {
    var temp tuntermMonitoringItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "port", "protocol", "timeout")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Host = temp.Host
    t.Port = temp.Port
    t.Protocol = temp.Protocol
    t.Timeout = temp.Timeout
    return nil
}

// tuntermMonitoringItem is a temporary struct used for validating the fields of TuntermMonitoringItem.
type tuntermMonitoringItem  struct {
    Host     *string                        `json:"host,omitempty"`
    Port     *int                           `json:"port,omitempty"`
    Protocol *TunternMonitoringProtocolEnum `json:"protocol,omitempty"`
    Timeout  *int                           `json:"timeout,omitempty"`
}
