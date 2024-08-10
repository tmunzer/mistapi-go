package models

import (
    "encoding/json"
)

// OspfConfig represents a OspfConfig struct.
// Junos OSPF config
type OspfConfig struct {
    // OSPF areas to run on this device and the corresponding per-area-specific configs. Property key is the area
    Areas                map[string]OspfConfigArea `json:"areas,omitempty"`
    // whether to rung OSPF on this device
    Enabled              *bool                     `json:"enabled,omitempty"`
    // Bandwidth for calculating metric defaults (9600..4000000000000)
    ReferenceBandwidth   *string                   `json:"reference_bandwidth,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OspfConfig.
// It customizes the JSON marshaling process for OspfConfig objects.
func (o OspfConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OspfConfig object to a map representation for JSON marshaling.
func (o OspfConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Areas != nil {
        structMap["areas"] = o.Areas
    }
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    if o.ReferenceBandwidth != nil {
        structMap["reference_bandwidth"] = o.ReferenceBandwidth
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfConfig.
// It customizes the JSON unmarshaling process for OspfConfig objects.
func (o *OspfConfig) UnmarshalJSON(input []byte) error {
    var temp tempOspfConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "areas", "enabled", "reference_bandwidth")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Areas = temp.Areas
    o.Enabled = temp.Enabled
    o.ReferenceBandwidth = temp.ReferenceBandwidth
    return nil
}

// tempOspfConfig is a temporary struct used for validating the fields of OspfConfig.
type tempOspfConfig  struct {
    Areas              map[string]OspfConfigArea `json:"areas,omitempty"`
    Enabled            *bool                     `json:"enabled,omitempty"`
    ReferenceBandwidth *string                   `json:"reference_bandwidth,omitempty"`
}
