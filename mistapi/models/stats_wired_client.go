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

// StatsWiredClient represents a StatsWiredClient struct.
type StatsWiredClient struct {
    // Client authorization status
    AuthState            *string                `json:"auth_state,omitempty"`
    // Device ID the client is connected to
    DeviceId             *string                `json:"device_id,omitempty"`
    // Port on AP where the wired client is connected
    EthPort              *string                `json:"eth_port,omitempty"`
    // Time when last Tx/Rx observed
    LastSeen             *float64               `json:"last_seen,omitempty"`
    // Client mac
    Mac                  string                 `json:"mac"`
    // Amount of traffic received since connection
    RxBytes              Optional[int64]        `json:"rx_bytes"`
    // Amount of packets received since connection
    RxPkts               Optional[int64]        `json:"rx_pkts"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Amount of traffic sent since connection
    TxBytes              Optional[int64]        `json:"tx_bytes"`
    // Amount of packets sent since connection
    TxPkts               Optional[int64]        `json:"tx_pkts"`
    // How long, in seconds, has the client been connected
    Uptime               *float64               `json:"uptime,omitempty"`
    // VLAN id, could be empty
    VlanId               *float64               `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWiredClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWiredClient) String() string {
    return fmt.Sprintf(
    	"StatsWiredClient[AuthState=%v, DeviceId=%v, EthPort=%v, LastSeen=%v, Mac=%v, RxBytes=%v, RxPkts=%v, SiteId=%v, TxBytes=%v, TxPkts=%v, Uptime=%v, VlanId=%v, AdditionalProperties=%v]",
    	s.AuthState, s.DeviceId, s.EthPort, s.LastSeen, s.Mac, s.RxBytes, s.RxPkts, s.SiteId, s.TxBytes, s.TxPkts, s.Uptime, s.VlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWiredClient.
// It customizes the JSON marshaling process for StatsWiredClient objects.
func (s StatsWiredClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "auth_state", "device_id", "eth_port", "last_seen", "mac", "rx_bytes", "rx_pkts", "site_id", "tx_bytes", "tx_pkts", "uptime", "vlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWiredClient object to a map representation for JSON marshaling.
func (s StatsWiredClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AuthState != nil {
        structMap["auth_state"] = s.AuthState
    }
    if s.DeviceId != nil {
        structMap["device_id"] = s.DeviceId
    }
    if s.EthPort != nil {
        structMap["eth_port"] = s.EthPort
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    structMap["mac"] = s.Mac
    if s.RxBytes.IsValueSet() {
        if s.RxBytes.Value() != nil {
            structMap["rx_bytes"] = s.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if s.RxPkts.IsValueSet() {
        if s.RxPkts.Value() != nil {
            structMap["rx_pkts"] = s.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.TxBytes.IsValueSet() {
        if s.TxBytes.Value() != nil {
            structMap["tx_bytes"] = s.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if s.TxPkts.IsValueSet() {
        if s.TxPkts.Value() != nil {
            structMap["tx_pkts"] = s.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWiredClient.
// It customizes the JSON unmarshaling process for StatsWiredClient objects.
func (s *StatsWiredClient) UnmarshalJSON(input []byte) error {
    var temp tempStatsWiredClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_state", "device_id", "eth_port", "last_seen", "mac", "rx_bytes", "rx_pkts", "site_id", "tx_bytes", "tx_pkts", "uptime", "vlan_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AuthState = temp.AuthState
    s.DeviceId = temp.DeviceId
    s.EthPort = temp.EthPort
    s.LastSeen = temp.LastSeen
    s.Mac = *temp.Mac
    s.RxBytes = temp.RxBytes
    s.RxPkts = temp.RxPkts
    s.SiteId = temp.SiteId
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Uptime = temp.Uptime
    s.VlanId = temp.VlanId
    return nil
}

// tempStatsWiredClient is a temporary struct used for validating the fields of StatsWiredClient.
type tempStatsWiredClient  struct {
    AuthState *string         `json:"auth_state,omitempty"`
    DeviceId  *string         `json:"device_id,omitempty"`
    EthPort   *string         `json:"eth_port,omitempty"`
    LastSeen  *float64        `json:"last_seen,omitempty"`
    Mac       *string         `json:"mac"`
    RxBytes   Optional[int64] `json:"rx_bytes"`
    RxPkts    Optional[int64] `json:"rx_pkts"`
    SiteId    *uuid.UUID      `json:"site_id,omitempty"`
    TxBytes   Optional[int64] `json:"tx_bytes"`
    TxPkts    Optional[int64] `json:"tx_pkts"`
    Uptime    *float64        `json:"uptime,omitempty"`
    VlanId    *float64        `json:"vlan_id,omitempty"`
}

func (s *tempStatsWiredClient) validate() error {
    var errs []string
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `stats_wired_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
