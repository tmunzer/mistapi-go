package models

import (
    "encoding/json"
)

// OrgSettingGatewayMgmt represents a OrgSettingGatewayMgmt struct.
type OrgSettingGatewayMgmt struct {
    AppProbing           *OrgSettingGatewayMgmtAppProbing      `json:"app_probing,omitempty"`
    // consumes uplink bandwidth, requires WA license
    AppUsage             *bool                                 `json:"app_usage,omitempty"`
    // optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim
    HostOutPolicies      *OrgSettingGatewayMgmtHostOutPolicies `json:"host_out_policies,omitempty"`
    OverlayIp            *OrgSettingGatewayMgmtOverlayIp       `json:"overlay_ip,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmt.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmt objects.
func (o OrgSettingGatewayMgmt) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "app_probing", "app_usage", "host_out_policies", "overlay_ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmt object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AppProbing != nil {
        structMap["app_probing"] = o.AppProbing.toMap()
    }
    if o.AppUsage != nil {
        structMap["app_usage"] = o.AppUsage
    }
    if o.HostOutPolicies != nil {
        structMap["host_out_policies"] = o.HostOutPolicies.toMap()
    }
    if o.OverlayIp != nil {
        structMap["overlay_ip"] = o.OverlayIp.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmt.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmt objects.
func (o *OrgSettingGatewayMgmt) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "app_probing", "app_usage", "host_out_policies", "overlay_ip")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.AppProbing = temp.AppProbing
    o.AppUsage = temp.AppUsage
    o.HostOutPolicies = temp.HostOutPolicies
    o.OverlayIp = temp.OverlayIp
    return nil
}

// tempOrgSettingGatewayMgmt is a temporary struct used for validating the fields of OrgSettingGatewayMgmt.
type tempOrgSettingGatewayMgmt  struct {
    AppProbing      *OrgSettingGatewayMgmtAppProbing      `json:"app_probing,omitempty"`
    AppUsage        *bool                                 `json:"app_usage,omitempty"`
    HostOutPolicies *OrgSettingGatewayMgmtHostOutPolicies `json:"host_out_policies,omitempty"`
    OverlayIp       *OrgSettingGatewayMgmtOverlayIp       `json:"overlay_ip,omitempty"`
}
