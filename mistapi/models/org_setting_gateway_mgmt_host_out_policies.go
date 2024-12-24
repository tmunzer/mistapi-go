package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingGatewayMgmtHostOutPolicies represents a OrgSettingGatewayMgmtHostOutPolicies struct.
// optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim
type OrgSettingGatewayMgmtHostOutPolicies struct {
    Dns                  *OrgSettingGatewayMgmtHostOutPoliciesDns  `json:"dns,omitempty"`
    Mist                 *OrgSettingGatewayMgmtHostOutPoliciesMist `json:"mist,omitempty"`
    Ntp                  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"ntp,omitempty"`
    Pim                  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"pim,omitempty"`
    AdditionalProperties map[string]interface{}                    `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostOutPolicies,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostOutPolicies) String() string {
    return fmt.Sprintf(
    	"OrgSettingGatewayMgmtHostOutPolicies[Dns=%v, Mist=%v, Ntp=%v, Pim=%v, AdditionalProperties=%v]",
    	o.Dns, o.Mist, o.Ntp, o.Pim, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o OrgSettingGatewayMgmtHostOutPolicies) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "dns", "mist", "ntp", "pim"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPolicies object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPolicies) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Dns != nil {
        structMap["dns"] = o.Dns.toMap()
    }
    if o.Mist != nil {
        structMap["mist"] = o.Mist.toMap()
    }
    if o.Ntp != nil {
        structMap["ntp"] = o.Ntp.toMap()
    }
    if o.Pim != nil {
        structMap["pim"] = o.Pim.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPolicies.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPolicies objects.
func (o *OrgSettingGatewayMgmtHostOutPolicies) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtHostOutPolicies
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "mist", "ntp", "pim")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Dns = temp.Dns
    o.Mist = temp.Mist
    o.Ntp = temp.Ntp
    o.Pim = temp.Pim
    return nil
}

// tempOrgSettingGatewayMgmtHostOutPolicies is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPolicies.
type tempOrgSettingGatewayMgmtHostOutPolicies  struct {
    Dns  *OrgSettingGatewayMgmtHostOutPoliciesDns  `json:"dns,omitempty"`
    Mist *OrgSettingGatewayMgmtHostOutPoliciesMist `json:"mist,omitempty"`
    Ntp  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"ntp,omitempty"`
    Pim  *OrgSettingGatewayMgmtHostOutPoliciesNtp  `json:"pim,omitempty"`
}
