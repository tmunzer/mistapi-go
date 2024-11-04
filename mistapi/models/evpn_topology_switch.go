package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// EvpnTopologySwitch represents a EvpnTopologySwitch struct.
type EvpnTopologySwitch struct {
    // Switch Configuration.
    // You can configure `port_usages` and `networks` settings at the device level, but most of the time it's better use the Site Setting to achieve better consistency and be able to re-use the same settings across switches entries defined here will "replace" those defined in Site Setting/Network Template
    Config               *DeviceSwitch               `json:"config,omitempty"`
    DeviceprofileId      *uuid.UUID                  `json:"deviceprofile_id,omitempty"`
    Downlinks            []string                    `json:"downlinks,omitempty"`
    Esilaglinks          []string                    `json:"esilaglinks,omitempty"`
    EvpnId               *int                        `json:"evpn_id,omitempty"`
    Mac                  *string                     `json:"mac,omitempty"`
    Model                *string                     `json:"model,omitempty"`
    // optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.
    // * for CLOS, to group dist / access switches into pods
    // * for ERB/CRB, to group dist / esilag-access into pods
    Pod                  *int                        `json:"pod,omitempty"`
    // by default, core switches are assumed to be connecting all pods.
    // if you want to limit the pods, you can specify pods.
    Pods                 []int                       `json:"pods,omitempty"`
    // use `role`==`none` to remove a switch from the topology. enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`
    Role                 *EvpnTopologySwitchRoleEnum `json:"role,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    SuggestedDownlinks   []string                    `json:"suggested_downlinks,omitempty"`
    SuggestedEsilaglinks []string                    `json:"suggested_esilaglinks,omitempty"`
    SuggestedUplinks     []string                    `json:"suggested_uplinks,omitempty"`
    // if not specified in the request, suggested ones will be used
    Uplinks              []string                    `json:"uplinks,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitch.
// It customizes the JSON marshaling process for EvpnTopologySwitch objects.
func (e EvpnTopologySwitch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitch object to a map representation for JSON marshaling.
func (e EvpnTopologySwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Config != nil {
        structMap["config"] = e.Config.toMap()
    }
    if e.DeviceprofileId != nil {
        structMap["deviceprofile_id"] = e.DeviceprofileId
    }
    if e.Downlinks != nil {
        structMap["downlinks"] = e.Downlinks
    }
    if e.Esilaglinks != nil {
        structMap["esilaglinks"] = e.Esilaglinks
    }
    if e.EvpnId != nil {
        structMap["evpn_id"] = e.EvpnId
    }
    if e.Mac != nil {
        structMap["mac"] = e.Mac
    }
    if e.Model != nil {
        structMap["model"] = e.Model
    }
    if e.Pod != nil {
        structMap["pod"] = e.Pod
    }
    if e.Pods != nil {
        structMap["pods"] = e.Pods
    }
    if e.Role != nil {
        structMap["role"] = e.Role
    }
    if e.SiteId != nil {
        structMap["site_id"] = e.SiteId
    }
    if e.SuggestedDownlinks != nil {
        structMap["suggested_downlinks"] = e.SuggestedDownlinks
    }
    if e.SuggestedEsilaglinks != nil {
        structMap["suggested_esilaglinks"] = e.SuggestedEsilaglinks
    }
    if e.SuggestedUplinks != nil {
        structMap["suggested_uplinks"] = e.SuggestedUplinks
    }
    if e.Uplinks != nil {
        structMap["uplinks"] = e.Uplinks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitch.
// It customizes the JSON unmarshaling process for EvpnTopologySwitch objects.
func (e *EvpnTopologySwitch) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "config", "deviceprofile_id", "downlinks", "esilaglinks", "evpn_id", "mac", "model", "pod", "pods", "role", "site_id", "suggested_downlinks", "suggested_esilaglinks", "suggested_uplinks", "uplinks")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Config = temp.Config
    e.DeviceprofileId = temp.DeviceprofileId
    e.Downlinks = temp.Downlinks
    e.Esilaglinks = temp.Esilaglinks
    e.EvpnId = temp.EvpnId
    e.Mac = temp.Mac
    e.Model = temp.Model
    e.Pod = temp.Pod
    e.Pods = temp.Pods
    e.Role = temp.Role
    e.SiteId = temp.SiteId
    e.SuggestedDownlinks = temp.SuggestedDownlinks
    e.SuggestedEsilaglinks = temp.SuggestedEsilaglinks
    e.SuggestedUplinks = temp.SuggestedUplinks
    e.Uplinks = temp.Uplinks
    return nil
}

// tempEvpnTopologySwitch is a temporary struct used for validating the fields of EvpnTopologySwitch.
type tempEvpnTopologySwitch  struct {
    Config               *DeviceSwitch               `json:"config,omitempty"`
    DeviceprofileId      *uuid.UUID                  `json:"deviceprofile_id,omitempty"`
    Downlinks            []string                    `json:"downlinks,omitempty"`
    Esilaglinks          []string                    `json:"esilaglinks,omitempty"`
    EvpnId               *int                        `json:"evpn_id,omitempty"`
    Mac                  *string                     `json:"mac,omitempty"`
    Model                *string                     `json:"model,omitempty"`
    Pod                  *int                        `json:"pod,omitempty"`
    Pods                 []int                       `json:"pods,omitempty"`
    Role                 *EvpnTopologySwitchRoleEnum `json:"role,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    SuggestedDownlinks   []string                    `json:"suggested_downlinks,omitempty"`
    SuggestedEsilaglinks []string                    `json:"suggested_esilaglinks,omitempty"`
    SuggestedUplinks     []string                    `json:"suggested_uplinks,omitempty"`
    Uplinks              []string                    `json:"uplinks,omitempty"`
}
