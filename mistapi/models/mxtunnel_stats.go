package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// MxtunnelStats represents a MxtunnelStats struct.
// MxTunnels statistics
type MxtunnelStats struct {
    Ap                   *string                 `json:"ap,omitempty"`
    ForSite              *bool                   `json:"for_site,omitempty"`
    LastSeen             *float64                `json:"last_seen,omitempty"`
    MxclusterId          *uuid.UUID              `json:"mxcluster_id,omitempty"`
    MxedgeId             *uuid.UUID              `json:"mxedge_id,omitempty"`
    MxtunnelId           *uuid.UUID              `json:"mxtunnel_id,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // MxEdge ID of the peer(mist edge to mist edge tunnel)
    PeerMxedgeId         *uuid.UUID              `json:"peer_mxedge_id,omitempty"`
    RemoteIp             *string                 `json:"remote_ip,omitempty"`
    RemotePort           *int                    `json:"remote_port,omitempty"`
    RxControlPkts        *int                    `json:"rx_control_pkts,omitempty"`
    // list of sessions
    Sessions             []MxtunnelStatsSession  `json:"sessions,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    // enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *MxtunnelStatsStateEnum `json:"state,omitempty"`
    TxControlPkts        *int                    `json:"tx_control_pkts,omitempty"`
    Uptime               *int                    `json:"uptime,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxtunnelStats.
// It customizes the JSON marshaling process for MxtunnelStats objects.
func (m MxtunnelStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxtunnelStats object to a map representation for JSON marshaling.
func (m MxtunnelStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Ap != nil {
        structMap["ap"] = m.Ap
    }
    if m.ForSite != nil {
        structMap["for_site"] = m.ForSite
    }
    if m.LastSeen != nil {
        structMap["last_seen"] = m.LastSeen
    }
    if m.MxclusterId != nil {
        structMap["mxcluster_id"] = m.MxclusterId
    }
    if m.MxedgeId != nil {
        structMap["mxedge_id"] = m.MxedgeId
    }
    if m.MxtunnelId != nil {
        structMap["mxtunnel_id"] = m.MxtunnelId
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.PeerMxedgeId != nil {
        structMap["peer_mxedge_id"] = m.PeerMxedgeId
    }
    if m.RemoteIp != nil {
        structMap["remote_ip"] = m.RemoteIp
    }
    if m.RemotePort != nil {
        structMap["remote_port"] = m.RemotePort
    }
    if m.RxControlPkts != nil {
        structMap["rx_control_pkts"] = m.RxControlPkts
    }
    if m.Sessions != nil {
        structMap["sessions"] = m.Sessions
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.State != nil {
        structMap["state"] = m.State
    }
    if m.TxControlPkts != nil {
        structMap["tx_control_pkts"] = m.TxControlPkts
    }
    if m.Uptime != nil {
        structMap["uptime"] = m.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxtunnelStats.
// It customizes the JSON unmarshaling process for MxtunnelStats objects.
func (m *MxtunnelStats) UnmarshalJSON(input []byte) error {
    var temp tempMxtunnelStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "for_site", "last_seen", "mxcluster_id", "mxedge_id", "mxtunnel_id", "org_id", "peer_mxedge_id", "remote_ip", "remote_port", "rx_control_pkts", "sessions", "site_id", "state", "tx_control_pkts", "uptime")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Ap = temp.Ap
    m.ForSite = temp.ForSite
    m.LastSeen = temp.LastSeen
    m.MxclusterId = temp.MxclusterId
    m.MxedgeId = temp.MxedgeId
    m.MxtunnelId = temp.MxtunnelId
    m.OrgId = temp.OrgId
    m.PeerMxedgeId = temp.PeerMxedgeId
    m.RemoteIp = temp.RemoteIp
    m.RemotePort = temp.RemotePort
    m.RxControlPkts = temp.RxControlPkts
    m.Sessions = temp.Sessions
    m.SiteId = temp.SiteId
    m.State = temp.State
    m.TxControlPkts = temp.TxControlPkts
    m.Uptime = temp.Uptime
    return nil
}

// tempMxtunnelStats is a temporary struct used for validating the fields of MxtunnelStats.
type tempMxtunnelStats  struct {
    Ap            *string                 `json:"ap,omitempty"`
    ForSite       *bool                   `json:"for_site,omitempty"`
    LastSeen      *float64                `json:"last_seen,omitempty"`
    MxclusterId   *uuid.UUID              `json:"mxcluster_id,omitempty"`
    MxedgeId      *uuid.UUID              `json:"mxedge_id,omitempty"`
    MxtunnelId    *uuid.UUID              `json:"mxtunnel_id,omitempty"`
    OrgId         *uuid.UUID              `json:"org_id,omitempty"`
    PeerMxedgeId  *uuid.UUID              `json:"peer_mxedge_id,omitempty"`
    RemoteIp      *string                 `json:"remote_ip,omitempty"`
    RemotePort    *int                    `json:"remote_port,omitempty"`
    RxControlPkts *int                    `json:"rx_control_pkts,omitempty"`
    Sessions      []MxtunnelStatsSession  `json:"sessions,omitempty"`
    SiteId        *uuid.UUID              `json:"site_id,omitempty"`
    State         *MxtunnelStatsStateEnum `json:"state,omitempty"`
    TxControlPkts *int                    `json:"tx_control_pkts,omitempty"`
    Uptime        *int                    `json:"uptime,omitempty"`
}
