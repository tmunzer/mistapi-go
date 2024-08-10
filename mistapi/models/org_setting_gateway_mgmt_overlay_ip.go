package models

import (
    "encoding/json"
)

// OrgSettingGatewayMgmtOverlayIp represents a OrgSettingGatewayMgmtOverlayIp struct.
type OrgSettingGatewayMgmtOverlayIp struct {
    // when it's going overlay, a routable IP to overlay will be required
    Ip                   *string        `json:"ip,omitempty"`
    // for SSR HA cluster, another IP for node1 will be required, too
    Node1Ip              *string        `json:"node1_ip,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtOverlayIp.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtOverlayIp objects.
func (o OrgSettingGatewayMgmtOverlayIp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtOverlayIp object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtOverlayIp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Ip != nil {
        structMap["ip"] = o.Ip
    }
    if o.Node1Ip != nil {
        structMap["node1_ip"] = o.Node1Ip
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtOverlayIp.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtOverlayIp objects.
func (o *OrgSettingGatewayMgmtOverlayIp) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtOverlayIp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "node1_ip")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Ip = temp.Ip
    o.Node1Ip = temp.Node1Ip
    return nil
}

// tempOrgSettingGatewayMgmtOverlayIp is a temporary struct used for validating the fields of OrgSettingGatewayMgmtOverlayIp.
type tempOrgSettingGatewayMgmtOverlayIp  struct {
    Ip      *string `json:"ip,omitempty"`
    Node1Ip *string `json:"node1_ip,omitempty"`
}
