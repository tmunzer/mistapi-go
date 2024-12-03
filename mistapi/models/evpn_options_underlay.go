package models

import (
    "encoding/json"
)

// EvpnOptionsUnderlay represents a EvpnOptionsUnderlay struct.
type EvpnOptionsUnderlay struct {
    // Underlay BGP Base AS Number
    AsBase               *int                   `json:"as_base,omitempty"`
    RoutedIdPrefix       *string                `json:"routed_id_prefix,omitempty"`
    // underlay subnet, by default, `10.255.240.0/20`, or `fd31:5700::/64` for ipv6
    Subnet               *string                `json:"subnet,omitempty"`
    // if v6 is desired for underlay
    UseIpv6              *bool                  `json:"use_ipv6,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptionsUnderlay.
// It customizes the JSON marshaling process for EvpnOptionsUnderlay objects.
func (e EvpnOptionsUnderlay) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "as_base", "routed_id_prefix", "subnet", "use_ipv6"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptionsUnderlay object to a map representation for JSON marshaling.
func (e EvpnOptionsUnderlay) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.AsBase != nil {
        structMap["as_base"] = e.AsBase
    }
    if e.RoutedIdPrefix != nil {
        structMap["routed_id_prefix"] = e.RoutedIdPrefix
    }
    if e.Subnet != nil {
        structMap["subnet"] = e.Subnet
    }
    if e.UseIpv6 != nil {
        structMap["use_ipv6"] = e.UseIpv6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptionsUnderlay.
// It customizes the JSON unmarshaling process for EvpnOptionsUnderlay objects.
func (e *EvpnOptionsUnderlay) UnmarshalJSON(input []byte) error {
    var temp tempEvpnOptionsUnderlay
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as_base", "routed_id_prefix", "subnet", "use_ipv6")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.AsBase = temp.AsBase
    e.RoutedIdPrefix = temp.RoutedIdPrefix
    e.Subnet = temp.Subnet
    e.UseIpv6 = temp.UseIpv6
    return nil
}

// tempEvpnOptionsUnderlay is a temporary struct used for validating the fields of EvpnOptionsUnderlay.
type tempEvpnOptionsUnderlay  struct {
    AsBase         *int    `json:"as_base,omitempty"`
    RoutedIdPrefix *string `json:"routed_id_prefix,omitempty"`
    Subnet         *string `json:"subnet,omitempty"`
    UseIpv6        *bool   `json:"use_ipv6,omitempty"`
}
