package models

import (
    "encoding/json"
)

// OrgSettingVpnOptions represents a OrgSettingVpnOptions struct.
type OrgSettingVpnOptions struct {
    AsBase               *int           `json:"as_base,omitempty"`
    // equiring /12 or bigger to support 16 private IPs for 65535 gateways
    StSubnet             *string        `json:"st_subnet,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingVpnOptions.
// It customizes the JSON marshaling process for OrgSettingVpnOptions objects.
func (o OrgSettingVpnOptions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingVpnOptions object to a map representation for JSON marshaling.
func (o OrgSettingVpnOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AsBase != nil {
        structMap["as_base"] = o.AsBase
    }
    if o.StSubnet != nil {
        structMap["st_subnet"] = o.StSubnet
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingVpnOptions.
// It customizes the JSON unmarshaling process for OrgSettingVpnOptions objects.
func (o *OrgSettingVpnOptions) UnmarshalJSON(input []byte) error {
    var temp orgSettingVpnOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "as_base", "st_subnet")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.AsBase = temp.AsBase
    o.StSubnet = temp.StSubnet
    return nil
}

// orgSettingVpnOptions is a temporary struct used for validating the fields of OrgSettingVpnOptions.
type orgSettingVpnOptions  struct {
    AsBase   *int    `json:"as_base,omitempty"`
    StSubnet *string `json:"st_subnet,omitempty"`
}
