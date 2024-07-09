package models

import (
    "encoding/json"
)

// WlanAppQosOthersItem represents a WlanAppQosOthersItem struct.
type WlanAppQosOthersItem struct {
    Dscp                 *int           `json:"dscp,omitempty"`
    DstSubnet            *string        `json:"dst_subnet,omitempty"`
    PortRanges           *string        `json:"port_ranges,omitempty"`
    Protocol             *string        `json:"protocol,omitempty"`
    SrcSubnet            *string        `json:"src_subnet,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQosOthersItem.
// It customizes the JSON marshaling process for WlanAppQosOthersItem objects.
func (w WlanAppQosOthersItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQosOthersItem object to a map representation for JSON marshaling.
func (w WlanAppQosOthersItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Dscp != nil {
        structMap["dscp"] = w.Dscp
    }
    if w.DstSubnet != nil {
        structMap["dst_subnet"] = w.DstSubnet
    }
    if w.PortRanges != nil {
        structMap["port_ranges"] = w.PortRanges
    }
    if w.Protocol != nil {
        structMap["protocol"] = w.Protocol
    }
    if w.SrcSubnet != nil {
        structMap["src_subnet"] = w.SrcSubnet
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAppQosOthersItem.
// It customizes the JSON unmarshaling process for WlanAppQosOthersItem objects.
func (w *WlanAppQosOthersItem) UnmarshalJSON(input []byte) error {
    var temp wlanAppQosOthersItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dscp", "dst_subnet", "port_ranges", "protocol", "src_subnet")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Dscp = temp.Dscp
    w.DstSubnet = temp.DstSubnet
    w.PortRanges = temp.PortRanges
    w.Protocol = temp.Protocol
    w.SrcSubnet = temp.SrcSubnet
    return nil
}

// wlanAppQosOthersItem is a temporary struct used for validating the fields of WlanAppQosOthersItem.
type wlanAppQosOthersItem  struct {
    Dscp       *int    `json:"dscp,omitempty"`
    DstSubnet  *string `json:"dst_subnet,omitempty"`
    PortRanges *string `json:"port_ranges,omitempty"`
    Protocol   *string `json:"protocol,omitempty"`
    SrcSubnet  *string `json:"src_subnet,omitempty"`
}
