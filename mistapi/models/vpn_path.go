package models

import (
    "encoding/json"
    "fmt"
)

// VpnPath represents a VpnPath struct.
type VpnPath struct {
    // enum: `broadband`, `lte`
    BfdProfile           *VpnPathBfdProfileEnum          `json:"bfd_profile,omitempty"`
    // If `type`==`mesh` and for SSR only, whether to use tunnel mode
    BfdUseTunnelMode     *bool                           `json:"bfd_use_tunnel_mode,omitempty"`
    // If different from the wan port
    Ip                   *string                         `json:"ip,omitempty"`
    // If `type`==`mesh`, Property key is the Peer Interface name
    PeerPaths            map[string]VpnPathPeerPathsPeer `json:"peer_paths,omitempty"`
    Pod                  *int                            `json:"pod,omitempty"`
    TrafficShaping       *VpnPathTrafficShaping          `json:"traffic_shaping,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPath) String() string {
    return fmt.Sprintf(
    	"VpnPath[BfdProfile=%v, BfdUseTunnelMode=%v, Ip=%v, PeerPaths=%v, Pod=%v, TrafficShaping=%v, AdditionalProperties=%v]",
    	v.BfdProfile, v.BfdUseTunnelMode, v.Ip, v.PeerPaths, v.Pod, v.TrafficShaping, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPath.
// It customizes the JSON marshaling process for VpnPath objects.
func (v VpnPath) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "bfd_profile", "bfd_use_tunnel_mode", "ip", "peer_paths", "pod", "traffic_shaping"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPath object to a map representation for JSON marshaling.
func (v VpnPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.BfdProfile != nil {
        structMap["bfd_profile"] = v.BfdProfile
    }
    if v.BfdUseTunnelMode != nil {
        structMap["bfd_use_tunnel_mode"] = v.BfdUseTunnelMode
    }
    if v.Ip != nil {
        structMap["ip"] = v.Ip
    }
    if v.PeerPaths != nil {
        structMap["peer_paths"] = v.PeerPaths
    }
    if v.Pod != nil {
        structMap["pod"] = v.Pod
    }
    if v.TrafficShaping != nil {
        structMap["traffic_shaping"] = v.TrafficShaping.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPath.
// It customizes the JSON unmarshaling process for VpnPath objects.
func (v *VpnPath) UnmarshalJSON(input []byte) error {
    var temp tempVpnPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bfd_profile", "bfd_use_tunnel_mode", "ip", "peer_paths", "pod", "traffic_shaping")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.BfdProfile = temp.BfdProfile
    v.BfdUseTunnelMode = temp.BfdUseTunnelMode
    v.Ip = temp.Ip
    v.PeerPaths = temp.PeerPaths
    v.Pod = temp.Pod
    v.TrafficShaping = temp.TrafficShaping
    return nil
}

// tempVpnPath is a temporary struct used for validating the fields of VpnPath.
type tempVpnPath  struct {
    BfdProfile       *VpnPathBfdProfileEnum          `json:"bfd_profile,omitempty"`
    BfdUseTunnelMode *bool                           `json:"bfd_use_tunnel_mode,omitempty"`
    Ip               *string                         `json:"ip,omitempty"`
    PeerPaths        map[string]VpnPathPeerPathsPeer `json:"peer_paths,omitempty"`
    Pod              *int                            `json:"pod,omitempty"`
    TrafficShaping   *VpnPathTrafficShaping          `json:"traffic_shaping,omitempty"`
}
