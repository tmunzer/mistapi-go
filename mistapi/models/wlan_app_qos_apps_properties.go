// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanAppQosAppsProperties represents a WlanAppQosAppsProperties struct.
type WlanAppQosAppsProperties struct {
    // DSCP value range between 0 and 63
    Dscp                 *Dscp                  `json:"dscp,omitempty"`
    // Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
    DstSubnet            *string                `json:"dst_subnet,omitempty"`
    // Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
    SrcSubnet            *string                `json:"src_subnet,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanAppQosAppsProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanAppQosAppsProperties) String() string {
    return fmt.Sprintf(
    	"WlanAppQosAppsProperties[Dscp=%v, DstSubnet=%v, SrcSubnet=%v, AdditionalProperties=%v]",
    	w.Dscp, w.DstSubnet, w.SrcSubnet, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQosAppsProperties.
// It customizes the JSON marshaling process for WlanAppQosAppsProperties objects.
func (w WlanAppQosAppsProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "dscp", "dst_subnet", "src_subnet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQosAppsProperties object to a map representation for JSON marshaling.
func (w WlanAppQosAppsProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Dscp != nil {
        structMap["dscp"] = w.Dscp.toMap()
    }
    if w.DstSubnet != nil {
        structMap["dst_subnet"] = w.DstSubnet
    }
    if w.SrcSubnet != nil {
        structMap["src_subnet"] = w.SrcSubnet
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAppQosAppsProperties.
// It customizes the JSON unmarshaling process for WlanAppQosAppsProperties objects.
func (w *WlanAppQosAppsProperties) UnmarshalJSON(input []byte) error {
    var temp tempWlanAppQosAppsProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dscp", "dst_subnet", "src_subnet")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Dscp = temp.Dscp
    w.DstSubnet = temp.DstSubnet
    w.SrcSubnet = temp.SrcSubnet
    return nil
}

// tempWlanAppQosAppsProperties is a temporary struct used for validating the fields of WlanAppQosAppsProperties.
type tempWlanAppQosAppsProperties  struct {
    Dscp      *Dscp   `json:"dscp,omitempty"`
    DstSubnet *string `json:"dst_subnet,omitempty"`
    SrcSubnet *string `json:"src_subnet,omitempty"`
}
