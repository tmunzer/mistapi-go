package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsWiredClient represents a StatsWiredClient struct.
type StatsWiredClient struct {
    Id                   *string        `json:"_id,omitempty"`
    // TTL of the validity of the stat
    Ttl                  *float64       `json:"_ttl,omitempty"`
    // client authorization status
    AuthState            *string        `json:"auth_state,omitempty"`
    // Device ID the client is connected to
    DeviceId             *string        `json:"device_id,omitempty"`
    // port on AP where the wired client is connected
    EthPort              *string        `json:"eth_port,omitempty"`
    // time when last Tx/Rx observed
    LastSeen             *float64       `json:"last_seen,omitempty"`
    // client mac
    Mac                  string         `json:"mac"`
    // amount of traffic sent to client since client connects
    RxBytes              *float64       `json:"rx_bytes,omitempty"`
    // amount of traffic sent to client since client connects
    RxPkts               *float64       `json:"rx_pkts,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    // amount of traffic received from client since client connects
    TxBytes              *float64       `json:"tx_bytes,omitempty"`
    // amount of traffic received from client since client connects
    TxPkts               *float64       `json:"tx_pkts,omitempty"`
    // how long, in seconds, has the client been connected
    Uptime               *float64       `json:"uptime,omitempty"`
    // vlan id, could be empty
    VlanId               *float64       `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWiredClient.
// It customizes the JSON marshaling process for StatsWiredClient objects.
func (s StatsWiredClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWiredClient object to a map representation for JSON marshaling.
func (s StatsWiredClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["_id"] = s.Id
    }
    if s.Ttl != nil {
        structMap["_ttl"] = s.Ttl
    }
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
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    }
    if s.RxPkts != nil {
        structMap["rx_pkts"] = s.RxPkts
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    }
    if s.TxPkts != nil {
        structMap["tx_pkts"] = s.TxPkts
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
    var temp statsWiredClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "_id", "_ttl", "auth_state", "device_id", "eth_port", "last_seen", "mac", "rx_bytes", "rx_pkts", "site_id", "tx_bytes", "tx_pkts", "uptime", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.Ttl = temp.Ttl
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

// statsWiredClient is a temporary struct used for validating the fields of StatsWiredClient.
type statsWiredClient  struct {
    Id        *string    `json:"_id,omitempty"`
    Ttl       *float64   `json:"_ttl,omitempty"`
    AuthState *string    `json:"auth_state,omitempty"`
    DeviceId  *string    `json:"device_id,omitempty"`
    EthPort   *string    `json:"eth_port,omitempty"`
    LastSeen  *float64   `json:"last_seen,omitempty"`
    Mac       *string    `json:"mac"`
    RxBytes   *float64   `json:"rx_bytes,omitempty"`
    RxPkts    *float64   `json:"rx_pkts,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    TxBytes   *float64   `json:"tx_bytes,omitempty"`
    TxPkts    *float64   `json:"tx_pkts,omitempty"`
    Uptime    *float64   `json:"uptime,omitempty"`
    VlanId    *float64   `json:"vlan_id,omitempty"`
}

func (s *statsWiredClient) validate() error {
    var errs []string
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Stats_Wired_Client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
