package models

import (
    "encoding/json"
)

// OrgSettingGatewayMgmtHostOutPolicies represents a OrgSettingGatewayMgmtHostOutPolicies struct.
// optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp
type OrgSettingGatewayMgmtHostOutPolicies struct {
    Dns                  *OrgSettingGatewayMgmtHostOutPoliciesDns  `json:"dns,omitempty"`
    Mist                 *OrgSettingGatewayMgmtHostOutPoliciesMist `json:"mist,omitempty"`
    Ntp                  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"ntp,omitempty"`
    AdditionalProperties map[string]any                            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o OrgSettingGatewayMgmtHostOutPolicies) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPolicies object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPolicies) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Dns != nil {
        structMap["dns"] = o.Dns.toMap()
    }
    if o.Mist != nil {
        structMap["mist"] = o.Mist.toMap()
    }
    if o.Ntp != nil {
        structMap["ntp"] = o.Ntp.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o *OrgSettingGatewayMgmtHostOutPolicies) UnmarshalJSON(input []byte) error {
    var temp orgSettingGatewayMgmtHostOutPolicies
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dns", "mist", "ntp")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Dns = temp.Dns
    o.Mist = temp.Mist
    o.Ntp = temp.Ntp
    return nil
}

// orgSettingGatewayMgmtHostOutPolicies is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPolicies.
type orgSettingGatewayMgmtHostOutPolicies  struct {
    Dns  *OrgSettingGatewayMgmtHostOutPoliciesDns  `json:"dns,omitempty"`
    Mist *OrgSettingGatewayMgmtHostOutPoliciesMist `json:"mist,omitempty"`
    Ntp  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"ntp,omitempty"`
}
