package models

import (
    "encoding/json"
)

// WlanAppQosAppsProperties represents a WlanAppQosAppsProperties struct.
type WlanAppQosAppsProperties struct {
    Dscp                 *int           `json:"dscp,omitempty"`
    // subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
    DstSubnet            *string        `json:"dst_subnet,omitempty"`
    // subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load)
    SrcSubnet            *string        `json:"src_subnet,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQosAppsProperties.
// It customizes the JSON marshaling process for WlanAppQosAppsProperties objects.
func (w WlanAppQosAppsProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQosAppsProperties object to a map representation for JSON marshaling.
func (w WlanAppQosAppsProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Dscp != nil {
        structMap["dscp"] = w.Dscp
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dscp", "dst_subnet", "src_subnet")
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
    Dscp      *int    `json:"dscp,omitempty"`
    DstSubnet *string `json:"dst_subnet,omitempty"`
    SrcSubnet *string `json:"src_subnet,omitempty"`
}
