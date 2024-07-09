package models

import (
    "encoding/json"
)

// WanUsages represents a WanUsages struct.
type WanUsages struct {
    Mac                  *string        `json:"mac,omitempty"`
    PathType             *string        `json:"path_type,omitempty"`
    PathWeight           *int           `json:"path_weight,omitempty"`
    PeerMac              *string        `json:"peer_mac,omitempty"`
    PeerPortId           *string        `json:"peer_port_id,omitempty"`
    Policy               *string        `json:"policy,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    Tenant               *string        `json:"tenant,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WanUsages.
// It customizes the JSON marshaling process for WanUsages objects.
func (w WanUsages) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WanUsages object to a map representation for JSON marshaling.
func (w WanUsages) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.PathType != nil {
        structMap["path_type"] = w.PathType
    }
    if w.PathWeight != nil {
        structMap["path_weight"] = w.PathWeight
    }
    if w.PeerMac != nil {
        structMap["peer_mac"] = w.PeerMac
    }
    if w.PeerPortId != nil {
        structMap["peer_port_id"] = w.PeerPortId
    }
    if w.Policy != nil {
        structMap["policy"] = w.Policy
    }
    if w.PortId != nil {
        structMap["port_id"] = w.PortId
    }
    if w.Tenant != nil {
        structMap["tenant"] = w.Tenant
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WanUsages.
// It customizes the JSON unmarshaling process for WanUsages objects.
func (w *WanUsages) UnmarshalJSON(input []byte) error {
    var temp wanUsages
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "path_type", "path_weight", "peer_mac", "peer_port_id", "policy", "port_id", "tenant")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Mac = temp.Mac
    w.PathType = temp.PathType
    w.PathWeight = temp.PathWeight
    w.PeerMac = temp.PeerMac
    w.PeerPortId = temp.PeerPortId
    w.Policy = temp.Policy
    w.PortId = temp.PortId
    w.Tenant = temp.Tenant
    return nil
}

// wanUsages is a temporary struct used for validating the fields of WanUsages.
type wanUsages  struct {
    Mac        *string `json:"mac,omitempty"`
    PathType   *string `json:"path_type,omitempty"`
    PathWeight *int    `json:"path_weight,omitempty"`
    PeerMac    *string `json:"peer_mac,omitempty"`
    PeerPortId *string `json:"peer_port_id,omitempty"`
    Policy     *string `json:"policy,omitempty"`
    PortId     *string `json:"port_id,omitempty"`
    Tenant     *string `json:"tenant,omitempty"`
}
