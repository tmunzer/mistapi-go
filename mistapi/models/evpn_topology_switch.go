package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// EvpnTopologySwitch represents a EvpnTopologySwitch struct.
type EvpnTopologySwitch struct {
    Config               *EvpnTopologySwitchConfig  `json:"config,omitempty"`
    DeviceprofileId      *uuid.UUID                 `json:"deviceprofile_id,omitempty"`
    DownlinkIps          []string                   `json:"downlink_ips,omitempty"`
    Downlinks            []string                   `json:"downlinks,omitempty"`
    Esilaglinks          []string                   `json:"esilaglinks,omitempty"`
    EvpnId               *int                       `json:"evpn_id,omitempty"`
    Mac                  string                     `json:"mac"`
    Model                *string                    `json:"model,omitempty"`
    // optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.
    // * for CLOS, to group dist / access switches into pods
    // * for ERB/CRB, to group dist / esilag-access into pods
    Pod                  *int                       `json:"pod,omitempty"`
    // by default, core switches are assumed to be connecting all pods.
    // if you want to limit the pods, you can specify pods.
    Pods                 []int                      `json:"pods,omitempty"`
    // use `role`==`none` to remove a switch from the topology. enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`
    Role                 EvpnTopologySwitchRoleEnum `json:"role"`
    RouterId             *string                    `json:"router_id,omitempty"`
    SiteId               *uuid.UUID                 `json:"site_id,omitempty"`
    SuggestedDownlinks   []string                   `json:"suggested_downlinks,omitempty"`
    SuggestedEsilaglinks []string                   `json:"suggested_esilaglinks,omitempty"`
    SuggestedUplinks     []string                   `json:"suggested_uplinks,omitempty"`
    Uplinks              []string                   `json:"uplinks,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologySwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologySwitch) String() string {
    return fmt.Sprintf(
    	"EvpnTopologySwitch[Config=%v, DeviceprofileId=%v, DownlinkIps=%v, Downlinks=%v, Esilaglinks=%v, EvpnId=%v, Mac=%v, Model=%v, Pod=%v, Pods=%v, Role=%v, RouterId=%v, SiteId=%v, SuggestedDownlinks=%v, SuggestedEsilaglinks=%v, SuggestedUplinks=%v, Uplinks=%v, AdditionalProperties=%v]",
    	e.Config, e.DeviceprofileId, e.DownlinkIps, e.Downlinks, e.Esilaglinks, e.EvpnId, e.Mac, e.Model, e.Pod, e.Pods, e.Role, e.RouterId, e.SiteId, e.SuggestedDownlinks, e.SuggestedEsilaglinks, e.SuggestedUplinks, e.Uplinks, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitch.
// It customizes the JSON marshaling process for EvpnTopologySwitch objects.
func (e EvpnTopologySwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "config", "deviceprofile_id", "downlink_ips", "downlinks", "esilaglinks", "evpn_id", "mac", "model", "pod", "pods", "role", "router_id", "site_id", "suggested_downlinks", "suggested_esilaglinks", "suggested_uplinks", "uplinks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitch object to a map representation for JSON marshaling.
func (e EvpnTopologySwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Config != nil {
        structMap["config"] = e.Config.toMap()
    }
    if e.DeviceprofileId != nil {
        structMap["deviceprofile_id"] = e.DeviceprofileId
    }
    if e.DownlinkIps != nil {
        structMap["downlink_ips"] = e.DownlinkIps
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
    structMap["mac"] = e.Mac
    if e.Model != nil {
        structMap["model"] = e.Model
    }
    if e.Pod != nil {
        structMap["pod"] = e.Pod
    }
    if e.Pods != nil {
        structMap["pods"] = e.Pods
    }
    structMap["role"] = e.Role
    if e.RouterId != nil {
        structMap["router_id"] = e.RouterId
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config", "deviceprofile_id", "downlink_ips", "downlinks", "esilaglinks", "evpn_id", "mac", "model", "pod", "pods", "role", "router_id", "site_id", "suggested_downlinks", "suggested_esilaglinks", "suggested_uplinks", "uplinks")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Config = temp.Config
    e.DeviceprofileId = temp.DeviceprofileId
    e.DownlinkIps = temp.DownlinkIps
    e.Downlinks = temp.Downlinks
    e.Esilaglinks = temp.Esilaglinks
    e.EvpnId = temp.EvpnId
    e.Mac = *temp.Mac
    e.Model = temp.Model
    e.Pod = temp.Pod
    e.Pods = temp.Pods
    e.Role = *temp.Role
    e.RouterId = temp.RouterId
    e.SiteId = temp.SiteId
    e.SuggestedDownlinks = temp.SuggestedDownlinks
    e.SuggestedEsilaglinks = temp.SuggestedEsilaglinks
    e.SuggestedUplinks = temp.SuggestedUplinks
    e.Uplinks = temp.Uplinks
    return nil
}

// tempEvpnTopologySwitch is a temporary struct used for validating the fields of EvpnTopologySwitch.
type tempEvpnTopologySwitch  struct {
    Config               *EvpnTopologySwitchConfig   `json:"config,omitempty"`
    DeviceprofileId      *uuid.UUID                  `json:"deviceprofile_id,omitempty"`
    DownlinkIps          []string                    `json:"downlink_ips,omitempty"`
    Downlinks            []string                    `json:"downlinks,omitempty"`
    Esilaglinks          []string                    `json:"esilaglinks,omitempty"`
    EvpnId               *int                        `json:"evpn_id,omitempty"`
    Mac                  *string                     `json:"mac"`
    Model                *string                     `json:"model,omitempty"`
    Pod                  *int                        `json:"pod,omitempty"`
    Pods                 []int                       `json:"pods,omitempty"`
    Role                 *EvpnTopologySwitchRoleEnum `json:"role"`
    RouterId             *string                     `json:"router_id,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    SuggestedDownlinks   []string                    `json:"suggested_downlinks,omitempty"`
    SuggestedEsilaglinks []string                    `json:"suggested_esilaglinks,omitempty"`
    SuggestedUplinks     []string                    `json:"suggested_uplinks,omitempty"`
    Uplinks              []string                    `json:"uplinks,omitempty"`
}

func (e *tempEvpnTopologySwitch) validate() error {
    var errs []string
    if e.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `evpn_topology_switch`")
    }
    if e.Role == nil {
        errs = append(errs, "required field `role` is missing for type `evpn_topology_switch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
