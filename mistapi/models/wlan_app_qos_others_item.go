package models

import (
    "encoding/json"
    "fmt"
)

// WlanAppQosOthersItem represents a WlanAppQosOthersItem struct.
type WlanAppQosOthersItem struct {
    Dscp                 *int                   `json:"dscp,omitempty"`
    DstSubnet            *string                `json:"dst_subnet,omitempty"`
    PortRanges           *string                `json:"port_ranges,omitempty"`
    Protocol             *string                `json:"protocol,omitempty"`
    SrcSubnet            *string                `json:"src_subnet,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanAppQosOthersItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanAppQosOthersItem) String() string {
    return fmt.Sprintf(
    	"WlanAppQosOthersItem[Dscp=%v, DstSubnet=%v, PortRanges=%v, Protocol=%v, SrcSubnet=%v, AdditionalProperties=%v]",
    	w.Dscp, w.DstSubnet, w.PortRanges, w.Protocol, w.SrcSubnet, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQosOthersItem.
// It customizes the JSON marshaling process for WlanAppQosOthersItem objects.
func (w WlanAppQosOthersItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "dscp", "dst_subnet", "port_ranges", "protocol", "src_subnet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQosOthersItem object to a map representation for JSON marshaling.
func (w WlanAppQosOthersItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp tempWlanAppQosOthersItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dscp", "dst_subnet", "port_ranges", "protocol", "src_subnet")
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

// tempWlanAppQosOthersItem is a temporary struct used for validating the fields of WlanAppQosOthersItem.
type tempWlanAppQosOthersItem  struct {
    Dscp       *int    `json:"dscp,omitempty"`
    DstSubnet  *string `json:"dst_subnet,omitempty"`
    PortRanges *string `json:"port_ranges,omitempty"`
    Protocol   *string `json:"protocol,omitempty"`
    SrcSubnet  *string `json:"src_subnet,omitempty"`
}
