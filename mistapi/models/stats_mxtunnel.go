// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsMxtunnel represents a StatsMxtunnel struct.
// MxTunnels statistics
type StatsMxtunnel struct {
    Ap                   *string                 `json:"ap,omitempty"`
    ForSite              *bool                   `json:"for_site,omitempty"`
    Fwupdate             *FwupdateStat           `json:"fwupdate,omitempty"`
    // Last seen timestamp
    LastSeen             Optional[float64]       `json:"last_seen"`
    Mtu                  *int                    `json:"mtu,omitempty"`
    MxclusterId          *uuid.UUID              `json:"mxcluster_id,omitempty"`
    MxedgeId             *uuid.UUID              `json:"mxedge_id,omitempty"`
    MxtunnelId           *uuid.UUID              `json:"mxtunnel_id,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // MxEdge ID of the peer(mist edge to mist edge tunnel)
    PeerMxedgeId         *uuid.UUID              `json:"peer_mxedge_id,omitempty"`
    RemoteIp             string                  `json:"remote_ip"`
    RemotePort           *int                    `json:"remote_port,omitempty"`
    RxControlPkts        *int                    `json:"rx_control_pkts,omitempty"`
    // List of sessions
    Sessions             []StatsMxtunnelSession  `json:"sessions,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    // enum: `established`, `established_with_sessions`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *StatsMxtunnelStateEnum `json:"state,omitempty"`
    TxControlPkts        *int                    `json:"tx_control_pkts,omitempty"`
    Uptime               *int                    `json:"uptime,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxtunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxtunnel) String() string {
    return fmt.Sprintf(
    	"StatsMxtunnel[Ap=%v, ForSite=%v, Fwupdate=%v, LastSeen=%v, Mtu=%v, MxclusterId=%v, MxedgeId=%v, MxtunnelId=%v, OrgId=%v, PeerMxedgeId=%v, RemoteIp=%v, RemotePort=%v, RxControlPkts=%v, Sessions=%v, SiteId=%v, State=%v, TxControlPkts=%v, Uptime=%v, AdditionalProperties=%v]",
    	s.Ap, s.ForSite, s.Fwupdate, s.LastSeen, s.Mtu, s.MxclusterId, s.MxedgeId, s.MxtunnelId, s.OrgId, s.PeerMxedgeId, s.RemoteIp, s.RemotePort, s.RxControlPkts, s.Sessions, s.SiteId, s.State, s.TxControlPkts, s.Uptime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxtunnel.
// It customizes the JSON marshaling process for StatsMxtunnel objects.
func (s StatsMxtunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap", "for_site", "fwupdate", "last_seen", "mtu", "mxcluster_id", "mxedge_id", "mxtunnel_id", "org_id", "peer_mxedge_id", "remote_ip", "remote_port", "rx_control_pkts", "sessions", "site_id", "state", "tx_control_pkts", "uptime"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxtunnel object to a map representation for JSON marshaling.
func (s StatsMxtunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Ap != nil {
        structMap["ap"] = s.Ap
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.Fwupdate != nil {
        structMap["fwupdate"] = s.Fwupdate.toMap()
    }
    if s.LastSeen.IsValueSet() {
        if s.LastSeen.Value() != nil {
            structMap["last_seen"] = s.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if s.Mtu != nil {
        structMap["mtu"] = s.Mtu
    }
    if s.MxclusterId != nil {
        structMap["mxcluster_id"] = s.MxclusterId
    }
    if s.MxedgeId != nil {
        structMap["mxedge_id"] = s.MxedgeId
    }
    if s.MxtunnelId != nil {
        structMap["mxtunnel_id"] = s.MxtunnelId
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.PeerMxedgeId != nil {
        structMap["peer_mxedge_id"] = s.PeerMxedgeId
    }
    structMap["remote_ip"] = s.RemoteIp
    if s.RemotePort != nil {
        structMap["remote_port"] = s.RemotePort
    }
    if s.RxControlPkts != nil {
        structMap["rx_control_pkts"] = s.RxControlPkts
    }
    if s.Sessions != nil {
        structMap["sessions"] = s.Sessions
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.State != nil {
        structMap["state"] = s.State
    }
    if s.TxControlPkts != nil {
        structMap["tx_control_pkts"] = s.TxControlPkts
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxtunnel.
// It customizes the JSON unmarshaling process for StatsMxtunnel objects.
func (s *StatsMxtunnel) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxtunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "for_site", "fwupdate", "last_seen", "mtu", "mxcluster_id", "mxedge_id", "mxtunnel_id", "org_id", "peer_mxedge_id", "remote_ip", "remote_port", "rx_control_pkts", "sessions", "site_id", "state", "tx_control_pkts", "uptime")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Ap = temp.Ap
    s.ForSite = temp.ForSite
    s.Fwupdate = temp.Fwupdate
    s.LastSeen = temp.LastSeen
    s.Mtu = temp.Mtu
    s.MxclusterId = temp.MxclusterId
    s.MxedgeId = temp.MxedgeId
    s.MxtunnelId = temp.MxtunnelId
    s.OrgId = temp.OrgId
    s.PeerMxedgeId = temp.PeerMxedgeId
    s.RemoteIp = *temp.RemoteIp
    s.RemotePort = temp.RemotePort
    s.RxControlPkts = temp.RxControlPkts
    s.Sessions = temp.Sessions
    s.SiteId = temp.SiteId
    s.State = temp.State
    s.TxControlPkts = temp.TxControlPkts
    s.Uptime = temp.Uptime
    return nil
}

// tempStatsMxtunnel is a temporary struct used for validating the fields of StatsMxtunnel.
type tempStatsMxtunnel  struct {
    Ap            *string                 `json:"ap,omitempty"`
    ForSite       *bool                   `json:"for_site,omitempty"`
    Fwupdate      *FwupdateStat           `json:"fwupdate,omitempty"`
    LastSeen      Optional[float64]       `json:"last_seen"`
    Mtu           *int                    `json:"mtu,omitempty"`
    MxclusterId   *uuid.UUID              `json:"mxcluster_id,omitempty"`
    MxedgeId      *uuid.UUID              `json:"mxedge_id,omitempty"`
    MxtunnelId    *uuid.UUID              `json:"mxtunnel_id,omitempty"`
    OrgId         *uuid.UUID              `json:"org_id,omitempty"`
    PeerMxedgeId  *uuid.UUID              `json:"peer_mxedge_id,omitempty"`
    RemoteIp      *string                 `json:"remote_ip"`
    RemotePort    *int                    `json:"remote_port,omitempty"`
    RxControlPkts *int                    `json:"rx_control_pkts,omitempty"`
    Sessions      []StatsMxtunnelSession  `json:"sessions,omitempty"`
    SiteId        *uuid.UUID              `json:"site_id,omitempty"`
    State         *StatsMxtunnelStateEnum `json:"state,omitempty"`
    TxControlPkts *int                    `json:"tx_control_pkts,omitempty"`
    Uptime        *int                    `json:"uptime,omitempty"`
}

func (s *tempStatsMxtunnel) validate() error {
    var errs []string
    if s.RemoteIp == nil {
        errs = append(errs, "required field `remote_ip` is missing for type `stats_mxtunnel`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
